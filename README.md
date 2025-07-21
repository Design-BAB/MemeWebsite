📰 The Daily Memer

A small web app written in Go that fetches and displays a wholesome meme using the Meme API.
🛠 Features

#Fetches a new meme from r/wholesomememes
#Written in Go with net/http and html/template
#Uses HTMX to load memes dynamically

📁 Structure

.
├── main.go              # Main Go server
├── static/
│   ├── index.html       # HTML template
│   └── sakura-vader.css # Minimal CSS theme

⚙️ How It Works

#The homepage serves index.html

#HTMX triggers a request to /meme

#The Go backend fetches a meme and returns an <img> tag

🧱 Technologies Used
#Go (net/http, html/template)

#HTMX (for async fetching)

#Meme API (https://meme-api.com)

#Sakura.css theme: Vader edition

📜 License

MIT
