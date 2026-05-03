package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// ===== MODELS =====

type Item struct {
	ID       uint `gorm:"primaryKey"`
	Brand    string
	Year     int
	Model    string
	Price    float64
	Category string
	Image    string
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

var db *gorm.DB

type ClientInfo struct {
	Attempts     int
	LastRequest  time.Time
	BlockedUntil time.Time
}

var clientInfo = make(map[string]*ClientInfo)
var clientMutex sync.Mutex

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminSecret := os.Getenv("ADMIN_SECRET")
		token, err := c.Cookie("admin_token")
		if err != nil || token != adminSecret {
			c.HTML(http.StatusForbidden, "login.html", gin.H{"Status": "ДОСТУП ЗАПРЕЩЕН: ТЫ КТО ТАКОЙ?"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getClientIP(c *gin.Context) string {
	if ip := c.GetHeader("X-Real-IP"); ip != "" {
		return ip
	}
	if ips := c.GetHeader("X-Forwarded-For"); ips != "" {
		return strings.TrimSpace(strings.Split(ips, ",")[0])
	}
	addr := c.Request.RemoteAddr
	if colon := strings.LastIndex(addr, ":"); colon != -1 {
		return addr[:colon]
	}
	return addr
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)
		clientMutex.Lock()
		info, ok := clientInfo[ip]
		if !ok {
			info = &ClientInfo{LastRequest: time.Now()}
			clientInfo[ip] = info
		}

		now := time.Now()
		if info.BlockedUntil.After(now) {
			clientMutex.Unlock()
			c.String(http.StatusTooManyRequests, "Слишком много запросов. Попробуйте позже.")
			c.Abort()
			return
		}

		if now.Sub(info.LastRequest) > time.Minute {
			info.Attempts = 0
		}
		info.Attempts++
		info.LastRequest = now

		if info.Attempts > 80 {
			info.BlockedUntil = now.Add(1 * time.Minute)
			clientMutex.Unlock()
			c.String(http.StatusTooManyRequests, "Слишком много запросов. Попробуйте через минуту.")
			c.Abort()
			return
		}
		clientMutex.Unlock()
		c.Next()
	}
}

func setSecureCookie(c *gin.Context, name, value string, maxAge int) {
	secure := c.Request.TLS != nil
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func recaptchaSiteKey() string {
	return os.Getenv("RECAPTCHA_SITEKEY")
}

func validateCaptcha(c *gin.Context) bool {
	secret := os.Getenv("RECAPTCHA_SECRET")
	if secret == "" {
		return true
	}

	response := c.PostForm("g-recaptcha-response")
	if response == "" {
		return false
	}

	return verifyCaptcha(response)
}

func verifyCaptcha(response string) bool {
	secret := os.Getenv("RECAPTCHA_SECRET")
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {secret}, "response": {response}})
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false
	}
	return result.Success
}

func authPageData() gin.H {
	return gin.H{"RecaptchaSiteKey": recaptchaSiteKey()}
}

// ===== UTILS =====

// ===== HANDLERS =====
func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("archive.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("[DATABASE ERROR]: %v\n", err)
		panic("DATABASE CONNECTION FAILED")
	}
	db.AutoMigrate(&User{}, &Item{})
	fmt.Println("[SYSTEM]: DATABASE INITIALIZED")
}

func register(c *gin.Context) {
	if !validateCaptcha(c) {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Подтвердите, что вы не робот", "RecaptchaSiteKey": recaptchaSiteKey()})
		return
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := User{Email: email, Password: string(hashedPassword)}

	if err := db.Create(&user).Error; err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "EMAIL ALREADY TAKEN", "RecaptchaSiteKey": recaptchaSiteKey()})
		return
	}

	setSecureCookie(c, "user_email", email, 3600)
	c.Redirect(http.StatusSeeOther, "/")
}

func signupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", authPageData())
}

func login(c *gin.Context) {
	if !validateCaptcha(c) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Подтвердите, что вы не робот", "RecaptchaSiteKey": recaptchaSiteKey()})
		return
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusNotFound, "login.html", gin.H{"error": "ОБЪЕКТ НЕ НАЙДЕН", "RecaptchaSiteKey": recaptchaSiteKey()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "НЕВЕРНЫЙ ПАРОЛЬ", "RecaptchaSiteKey": recaptchaSiteKey()})
		return
	}

	setSecureCookie(c, "user_email", user.Email, 3600*24)
	c.Redirect(http.StatusSeeOther, "/")
}

func findItemByID(id string) *Item {
	var item Item
	if err := db.First(&item, id).Error; err != nil {
		return nil
	}
	return &item
}

// ===== ROUTES =====
func main() {

	godotenv.Load()
	initDB()

	r := gin.Default()
	r.Use(rateLimitMiddleware())
	r.Static("/img", "./img")
	r.Static("/static", "./static")
	r.StaticFile("/prototypes.css", "./prototypes.css")
	r.LoadHTMLGlob("templates/*")

	// 2. Создаем сервер

	// 3. Настраиваем статику и шаблоны

	// --- АДМИН ЛОГИН ---
	r.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin_login.html", nil)
	})

	r.POST("/admin/login", func(c *gin.Context) {
		password := c.PostForm("password")
		adminSecret := os.Getenv("ADMIN_SECRET")

		if password != adminSecret {
			c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{"error": "НЕВЕРНЫЙ ПАРОЛЬ"})
			return
		}

		setSecureCookie(c, "admin_token", adminSecret, 3600*24)
		c.Redirect(http.StatusSeeOther, "/admin/")
	})

	admin := r.Group("/admin")
	admin.Use(AuthRequired())
	{
		admin.GET("/", func(c *gin.Context) {
			c.HTML(200, "admin.html", nil)
		})

		admin.GET("/logout", func(c *gin.Context) {
			http.SetCookie(c.Writer, &http.Cookie{Name: "admin_token", Value: "", Path: "/", MaxAge: -1, HttpOnly: true, SameSite: http.SameSiteLaxMode})
			c.Redirect(http.StatusSeeOther, "/")
		})

		admin.POST("/add", func(c *gin.Context) {

			var newItem Item

			newItem.Brand = c.PostForm("brand")
			newItem.Model = c.PostForm("model")
			newItem.Image = c.PostForm("image")
			newItem.Category = c.PostForm("category")

			price, _ := strconv.Atoi(c.PostForm("price"))
			newItem.Price = float64(price)

			year, _ := strconv.Atoi(c.PostForm("year"))
			newItem.Year = year

			db.Create(&newItem)

			c.Redirect(303, "/")
		})
	}
	// --- ГЛАВНАЯ СТРАНИЦА (ОБЪЕДИНЕННАЯ) ---
	r.GET("/", func(c *gin.Context) {
		email, err := c.Cookie("user_email")

		isLoggedIn := err == nil && email != ""

		status := "ВХОД НЕ ВЫПОЛНЕН"
		if isLoggedIn {
			status = "ВЫ В СВОЕМ АККАУНТЕ"
		}

		var items []Item
		db.Find(&items)

		c.HTML(http.StatusOK, "prototypes.html", gin.H{
			"Category":   "all",
			"Items":      items,
			"IsLoggedIn": isLoggedIn,
			"Status":     status,
		})
	})
	// --- МАГАЗИНЫ ---
	r.GET("/shop/:cat", func(c *gin.Context) {
		category := c.Param("cat") // Сервер «ловит» название бренда из ссылки
		var items []Item

		// Идем в базу и берем только то, что помечено этой категорией
		db.Where("category = ?", category).Find(&items)

		// Показываем твой шаблон prototypes.html, но с нужным товаром
		c.HTML(http.StatusOK, "clothes.html", gin.H{
			"Category": category,
			"Items":    items,
		})
	})
	// --- ВХОД И РЕГИСТРАЦИЯ ---
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", authPageData())
	})
	r.POST("/login", login)

	r.GET("/signup", signupPage)
	r.POST("/register", register)

	// --- ОСТАЛЬНОЕ ---

	r.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", gin.H{"title": "Корзина"})
	})

	r.GET("/logout", func(c *gin.Context) {
		// Мы перезаписываем куку пустой строкой и ставим время жизни -1 (удаление)
		http.SetCookie(c.Writer, &http.Cookie{Name: "user_email", Value: "", Path: "/", MaxAge: -1, HttpOnly: true, SameSite: http.SameSiteLaxMode})

		// Кидаем обратно на главную
		c.Redirect(http.StatusSeeOther, "/")
	})
	// --- ЗАПУСК ---
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	fmt.Println("[SYSTEM]: ARCHIVE SERVER RUNNING ON :8080")
	srv.ListenAndServe()
}
