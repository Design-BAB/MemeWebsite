# 📰 The Daily Memer

A delightfully simple web app written in Go that fetches and displays wholesome memes using the Meme API.

## 🛠 Features

- Pulls fresh memes from [r/wholesomememes](https://www.reddit.com/r/wholesomememes)
- Written in Go using `net/http` and `html/template`
- Uses [HTMX](https://htmx.org) for dynamic content loading

## 📁 Project Structure

```
.
├── main.go              // Go web server
├── static/
│   ├── index.html       // HTML template
│   └── sakura-vader.css // Minimal Vader-themed CSS
```

## ⚙️ How It Works

1. Homepage serves `index.html`
2. HTMX sends an async request to `/meme`
3. Backend fetches a wholesome meme and returns it as an `<img>` tag

## 🧱 Technologies Used

- **Go** – lightweight server with `net/http` and `html/template`
- **HTMX** – handles dynamic meme fetching
- **[Meme API](https://meme-api.com)** – source of wholesome goodness
- **Sakura.css (Vader edition)** – for minimalist vibes

## 📜 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT)
