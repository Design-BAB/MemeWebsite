package main

import (
  "embed"
  "fmt"
  "net/http"
  "html/template"
  "log"
  "os/exec"
  "encoding/json"
)

//go:embed static/*
var staticFiles embed.FS


func openBrowser(url string) {
	var err error

	err = exec.Command("xdg-open", url).Start()
	if err != nil {
		log.Println("Error opening browser:", err)
  }
}

func newMeme(meme string) *Meme{
  m := Meme{Image: meme}
  return &m
}

type Meme struct {
  Image string
}

type OnlineMeme struct {
  Image string `json:"url"`
}

func fetchMeme() *Meme {
  fmt.Println("I am now going to go get the meme!")
  apiURL := fmt.Sprintf("https://meme-api.com/gimme/wholesomememes")
  // Make an HTTP GET request to the  API
	resp, err := http.Get(apiURL)
	if err != nil {
		 // If the request fails, return the error
		log.Fatalf("Error getting meme: %v", err)
    return nil
	}
	defer resp.Body.Close() // Always close the response body when done
  if resp.StatusCode != http.StatusOK {
    log.Printf("Meme API returned status code: %d", resp.StatusCode)
    return nil
  }
  var theMeme OnlineMeme
	// Decode the JSON response into our theMeme struct
	err = json.NewDecoder(resp.Body).Decode(&theMeme)
if err != nil {
    log.Fatalf("Error getting meme: %v", err)
		return nil // If decoding fails, return the error
	}
  return newMeme(theMeme.Image)
}

func main() {
    fmt.Println("Hello! Going to start the web server!")
    //now this is the function keep in mind that h1 means handler #1
    h1 := func (w http.ResponseWriter, r *http.Request) {
      tmp1, err := template.ParseFS(staticFiles, "static/index.html")
		  if err != nil {
			  log.Fatalf("Error parsing index.html: %v", err)
			  http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			  return
		  }
      tmp1.Execute(w, nil)
      w.Header().Set("Content-Type", "text/html")
    }

    fs := http.FS(staticFiles)
	  http.Handle("/static/", http.FileServer(fs))
    //This will handle the function.
    //since this is "/" basicly when you go to the homepage it tells the person Hey! go run h1 when visitng the homepage
    http.HandleFunc("/", h1)


    http.HandleFunc("/meme", func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/html")
      theMeme := fetchMeme()
      w.Write([]byte(`<img src="` + theMeme.Image + `" width="700">`))
    })

    openBrowser("http://localhost:8000")
    //log.fatal will recond something if it failed to make a webserver
    //http listen and serve creates that server.
    log.Fatal(http.ListenAndServe(":8000", nil)) 
  }
