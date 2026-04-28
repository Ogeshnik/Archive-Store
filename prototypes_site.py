<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MAXIM | DATA OPERATOR</title>
    <style>
        :root { --main-color: #00ff41; --bg-color: #050505; --accent-color: #555; }
        * { box-sizing: border-box; }
        body { 
            background-color: var(--bg-color); 
            color: #fff; 
            font-family: "Courier New", monospace; 
            padding: 5vw; 
            line-height: 1.6; 
            margin: 0;
            overflow-x: hidden;
        }
        
        /* Эффект шума на фоне */
        body::before {
            content: "";
            position: fixed; top: 0; left: 0; width: 100%; height: 100%;
            background: repeating-linear-gradient(0deg, rgba(0,0,0,0.05), rgba(0,0,0,0.05) 1px, transparent 1px, transparent 2px);
            pointer-events: none; z-index: 10;
        }

        .header { border-bottom: 1px solid #222; padding-bottom: 30px; margin-bottom: 50px; }
        .name { font-size: clamp(1.5rem, 5vw, 3rem); letter-spacing: -2px; text-transform: uppercase; font-weight: 900; }
        .tag { color: var(--accent-color); font-size: 0.8rem; margin-top: 5px; }

        .section { margin-bottom: 60px; animation: fadeIn 1s ease-out; }
        .case-title { color: var(--main-color); text-transform: uppercase; font-weight: bold; margin-bottom: 15px; display: flex; align-items: center; }
        .case-title::before { content: "> "; margin-right: 10px; }

        .case-item { background: #0d0d0d; border: 1px solid #1a1a1a; padding: 20px; margin-bottom: 15px; transition: 0.3s; }
        .case-item:hover { border-color: var(--main-color); background: #111; }
        .case-item p { margin: 0; font-size: 0.95rem; }
        .case-item span { color: var(--accent-color); font-size: 0.8rem; display: block; margin-top: 10px; }

        .stack { display: flex; flex-wrap: wrap; gap: 10px; }
        .stack-item { background: #1a1a1a; padding: 5px 12px; font-size: 0.85rem; border-radius: 2px; border: 1px solid #333; }

        a { color: #fff; text-decoration: none; border-bottom: 1px dashed var(--accent-color); transition: 0.3s; }
        a:hover { color: var(--main-color); border-bottom-style: solid; border-color: var(--main-color); }

        @keyframes fadeIn { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: translateY(0); } }

        @media (max-width: 600px) { body { padding: 20px; } }
    </style>
</head>
<body>

    <div class="header">
        <div class="name">Maxim_System.v1</div>
        <div class="tag">// STATUS: ACTIVE // ROLE: DATA_ARCHITECT // ST_PETERSBURG</div>
    </div>

    <div class="section">
        <div class="case-title">ВОРК / CASES</div>
        
        <div class="case-item">
            <p>Автоматизация сбора B2B баз (производства, образовательные центры).</p>
            <span>[PLAYWRIGHT / SELENIUM / MULTI-THREADING]</span>
        </div>

        <div class="case-item">
            <p>Разработка финансовых калькуляторов с анализом ROI и маржинальности.</p>
            <span>[PYTHON / GOOGLE_SHEETS_API / MATH_LOGIC]</span>
        </div>

        <div class="case-item">
            <p>Мониторинг и фильтрация Telegram/VK сообществ по охватам.</p>
            <span>[SOCIAL_GRAPH / DATA_MINING]</span>
        </div>
    </div>

    <div class="section">
        <div class="case-title">СТЕК / STACK</div>
        <div class="stack">
            <div class="stack-item">Python 3.14</div>
            <div class="stack-item">Playwright</div>
            <div class="stack-item">Pandas</div>
            <div class="stack-item">HTML5/CSS3</div>
            <div class="stack-item">Linux Admin</div>
        </div>
    </div>

    <div class="section">
        <div class="case-title">СВЯЗЬ / CONTACTS</div>
        <p>TG: <a href="https://t.me" target="_blank">@ТВОЙ_НИК</a></p>
        <p>Email: <a href="mailto:ТВОЯ_ПОЧТА@mail.com">ТВОЯ_ПОЧТА@mail.com</a></p>
        <p>Events: <span style="color: #fff;">In real life only. 🐍</span></p>
    </div>

</body>
</html>
