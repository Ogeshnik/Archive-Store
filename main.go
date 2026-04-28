package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ===== MODEL =====
type Item struct {
	ID          int
	Brand       string
	Year        int
	Model       string
	Price       float64
	Description string
	Image       string
}

// ===== DATA =====
var archive = []Item{
	{ID: 1, Brand: "PROTOTYPES", Year: 2024, Model: "HOODIE", Price: 80000, Image: "1.jpg"},
	{ID: 2, Brand: "PROTOTYPES", Year: 2023, Model: "T-SHIRT", Price: 20000, Image: "2.webp"},
	{ID: 3, Brand: "PROTOTYPES", Year: 2024, Model: "RED TEE", Price: 25000, Image: "3.webp"},
	{ID: 4, Brand: "PROTOTYPES", Year: 2024, Model: "BLACK TEE", Price: 18000, Image: "4.jpg"},
	{ID: 5, Brand: "PROTOTYPES", Year: 2024, Model: "Hoodie", Price: 56000, Image: "5.webp"},
	{ID: 6, Brand: "PROTOTYPES", Year: 2024, Model: "Hoщdie", Price: 65000, Image: "6.jpg"},
	{ID: 7, Brand: "PROTOTYPES", Year: 2024, Model: "BASIC HOODIE", Price: 40000, Image: "7.jpg"},
	{ID: 8, Brand: "PROTOTYPES", Year: 2024, Model: "WASHED HOODIE", Price: 45000, Image: "8.webp"},
	{ID: 9, Brand: "PROTOTYPES", Year: 2024, Model: "GRAPHIC HOODIE", Price: 50000, Image: "9.webp"},
}

var rickArchive = []Item{
	{ID: 10, Brand: "RICK OWENS", Year: 2023, Model: "T-shirt", Price: 65000, Image: "rick1.jpg"},
	{ID: 11, Brand: "RICK OWENS", Year: 2024, Model: "T-shirt", Price: 80000, Image: "rick2.jpg"},
	{ID: 12, Brand: "RICK OWENS", Year: 2024, Model: "T-shirt", Price: 95000, Image: "rick3.webp"},
	{ID: 13, Brand: "RICK OWENS", Year: 2024, Model: "MOUNTAIN HOODIE", Price: 55000, Image: "rick4.jpg"},
	{ID: 14, Brand: "RICK OWENS", Year: 2024, Model: "Mountain zip hoodie", Price: 60000, Image: "rick5.jpg"},
	{ID: 15, Brand: "RICK OWENS", Year: 2024, Model: "Hoodie", Price: 30000, Image: "rick6.jpg"},
	{ID: 16, Brand: "RICK OWENS", Year: 2024, Model: "LEVEL TEE", Price: 28000, Image: "rick7.jpg"},
	{ID: 17, Brand: "RICK OWENS", Year: 2024, Model: "T-shirt", Price: 70000, Image: "rick8.jpg"},
	{ID: 18, Brand: "RICK OWENS", Year: 2024, Model: "Hoodie", Price: 75000, Image: "rick9.jpg"},
}

var nineArchive = []Item{
	{ID: 20, Brand: "NUMBER (N)INE", Year: 2004, Model: "Hoodie", Price: 45000, Image: "num1.jpg"},
	{ID: 21, Brand: "NUMBER (N)INE", Year: 2005, Model: "T-shirt", Price: 35000, Image: "num2.jpg"},
	{ID: 22, Brand: "NUMBER (N)INE", Year: 2005, Model: "longsleeve", Price: 15000, Image: "num3.jpg"},
	{ID: 23, Brand: "NUMBER (N)INE", Year: 2004, Model: "longsleeve", Price: 55000, Image: "num4.jpg"},
	{ID: 24, Brand: "NUMBER (N)INE", Year: 2004, Model: "T-shirt", Price: 40000, Image: "num5.webp"},
	{ID: 25, Brand: "NUMBER (N)INE", Year: 2005, Model: "hoodie", Price: 38000, Image: "num6.webp"},
	{ID: 26, Brand: "NUMBER (N)INE", Year: 2004, Model: "Zip-hoodie", Price: 60000, Image: "num7.webp"},
	{ID: 27, Brand: "NUMBER (N)INE", Year: 2005, Model: "T-shirt", Price: 25000, Image: "num8.webp"},
	{ID: 28, Brand: "NUMBER (N)INE", Year: 2004, Model: "T-shirt", Price: 42000, Image: "num9.webp"},
}
var StarbucksArchive = []Item{
	{ID: 29, Brand: "Haunted Starbucks", Year: 2025, Model: "T-shirt", Price: 20000, Image: "haunt1.jpg"},
	{ID: 30, Brand: "Haunted Starbucks", Year: 2026, Model: "Hoodie", Price: 10000, Image: "haunt2.jpg"},
	{ID: 31, Brand: "Haunted Starbucks", Year: 2025, Model: "Longsleeve", Price: 15000, Image: "haunt3.jpg"},
}
var MiharaArchive = []Item{
	{ID: 32, Brand: "Mihara", Year: 2020, Model: "Hoodie", Price: 45000, Image: "mihara1.jpg"},
	{ID: 33, Brand: "Mihara", Year: 2015, Model: "T-shirt", Price: 35000, Image: "mihara2.webp"},
	{ID: 34, Brand: "Mihara", Year: 2016, Model: "T-shirt", Price: 15000, Image: "mihara3.jpg"},
	{ID: 35, Brand: "Mihara", Year: 2014, Model: "HOODIE", Price: 55000, Image: "mihara4.jpg"},
	{ID: 36, Brand: "Mihara", Year: 2012, Model: "T-shirt", Price: 40000, Image: "mihara5.jpg"},
}
var SaintMichaelArchive = []Item{
	{ID: 37, Brand: "SaintMichael", Year: 2004, Model: "T-shirt", Price: 45000, Image: "saint1.jpg"},
	{ID: 38, Brand: "SaintMichael", Year: 2005, Model: "Hoodie", Price: 35000, Image: "saint2.jpg"},
	{ID: 39, Brand: "SaintMichael", Year: 2005, Model: "T-shirt", Price: 15000, Image: "saint3.jpg"},
	{ID: 40, Brand: "SaintMichael", Year: 2004, Model: "T-shirt", Price: 55000, Image: "saint4.webp"},
	{ID: 41, Brand: "SaintMichael", Year: 2004, Model: "T-shirt", Price: 40000, Image: "saint5.jpg"},
	{ID: 42, Brand: "SaintMichael", Year: 2005, Model: "T-shirt", Price: 38000, Image: "saint6.jpg"},
	{ID: 43, Brand: "SaintMichael", Year: 2004, Model: "T-shirt", Price: 60000, Image: "saint7.jpg"},
	{ID: 44, Brand: "SaintMichael", Year: 2005, Model: "T-shirt", Price: 25000, Image: "saint8.avif"},
}
var IamsoloistArchive = []Item{
	{ID: 45, Brand: "Iamsoloist", Year: 2004, Model: "T-shirt", Price: 45000, Image: "solo1.webp"},
	{ID: 46, Brand: "Iamsoloist", Year: 2005, Model: "T-shirt", Price: 35000, Image: "solo2.jpg"},
	{ID: 47, Brand: "Iamsoloist", Year: 2005, Model: "Hoodie", Price: 15000, Image: "solo3.webp"},
	{ID: 48, Brand: "Iamsoloist", Year: 2004, Model: "T-shirt", Price: 55000, Image: "solo4.webp"},
	{ID: 49, Brand: "Iamsoloist", Year: 2004, Model: "T-shirt", Price: 40000, Image: "solo5.jpg"},
	{ID: 50, Brand: "Iamsoloist", Year: 2005, Model: "blazer", Price: 38000, Image: "solo6.jpg"},
}

// ===== FIND ITEM =====
func findItemByID(id string) *Item {
	search := func(items []Item) *Item {
		for i := range items {
			if fmt.Sprintf("%d", items[i].ID) == id {
				return &items[i]
			}
		}
		return nil
	}

	if item := search(archive); item != nil {
		return item
	}
	if item := search(rickArchive); item != nil {
		return item
	}
	if item := search(nineArchive); item != nil {
		return item
	}
	if item := search(StarbucksArchive); item != nil {
		return item
	}
	if item := search(MiharaArchive); item != nil {
		return item
	}
	if item := search(SaintMichaelArchive); item != nil {
		return item
	}

	if item := search(IamsoloistArchive); item != nil {
		return item
	}

	return nil
}
func ГЛАВНАЯ(c *gin.Context) {
	c.HTML(http.StatusOK, "prototypes.html", archive)
}
func prototypes(c *gin.Context) {
	c.HTML(http.StatusOK, "prototypes_shop.html", archive)
}
func rick(c *gin.Context) {
	c.HTML(http.StatusOK, "rick_owens_shop.html", rickArchive)
}
func numberNine(c *gin.Context) {
	c.HTML(http.StatusOK, "number_nine_shop.html", nineArchive)
}
func hauntedstarbucks(c *gin.Context) {
	c.HTML(http.StatusOK, "hauntedstarbucks.html", StarbucksArchive)
}
func mihara(c *gin.Context) {
	c.HTML(http.StatusOK, "mihara.html", MiharaArchive)
}
func saintmichael(c *gin.Context) {
	c.HTML(http.StatusOK, "saintmichael.html", SaintMichaelArchive)
}
func soloist(c *gin.Context) {
	c.HTML(http.StatusOK, "soloist.html", IamsoloistArchive)
}
func product(c *gin.Context) {
	id := c.Param("id")
	item := findItemByID(id)
	if item == nil {
		c.String(http.StatusNotFound, "Товар не найден")
		return
	}
	c.HTML(http.StatusOK, "product.html", item)
}

func apiItems(c *gin.Context) {
	all := append([]Item{}, archive...)
	all = append(all, rickArchive...)
	all = append(all, nineArchive...)
	all = append(all, StarbucksArchive...)
	all = append(all, MiharaArchive...)
	all = append(all, SaintMichaelArchive...)
	all = append(all, IamsoloistArchive...)

	c.JSON(http.StatusOK, all)
}

// ===== MAIN =====
func main() {
	r := gin.Default()

	// static files
	r.Static("/img", "./img")
	r.Static("/static", "./static")

	// templates
	r.LoadHTMLGlob("templates/*")

	// routes
	r.GET("/", ГЛАВНАЯ)
	r.GET("/prototypes", prototypes)
	r.GET("/rick", rick)
	r.GET("/numbernine", numberNine)
	r.GET("/product/:id", product)
	r.GET("/hauntedstarbucks", hauntedstarbucks)
	r.GET("/mihara", mihara)
	r.GET("/saintmichael", saintmichael)
	r.GET("/soloist", soloist)

	// api
	r.GET("/api/items", apiItems)

	// run server
	r.Run(":8080")
}
