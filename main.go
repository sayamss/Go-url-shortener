package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	var urls = make(map[string]string)

	// urls
	urls["/ga"] = "http://www.google.com"
	urls["/fb"] = "http://www.facebook.com"
	urls["/gm"] = "http://www.gmail.com"
	urls["/yt"] = "http://www.youtube.com"
	urls["/az"] = "http://www.amazon.com"
	urls["/in"] = "http://www.instagram.com"
	urls["/wk"] = "http://www.wikipedia.com"
	urls["/rd"] = "http://www.reddit.com"
	urls["/ln"] = "http://www.linkedin.com"

	// store the current url
	url := r.URL.RequestURI()

	// loop through urls
	for key, value := range urls {
		// Redirect to the desired url if url matches the key
		if url == key {
			http.Redirect(w, r, value, http.StatusSeeOther)
		}
	}
}

func handleRequest() {

	http.HandleFunc("/", redirect)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	handleRequest()
}
