package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://educative.io",
	"https://educative.io/teach",
	"https://educative.io/assessments",
	"https://educative.io/projects",
	"https://educative.io/paths",
	"https://educative.io/learning-plans",
	"https://educative.io/learn",
	"https://educative.io/edpresso",
	"https://educative.io/explore",
	"https://educative.io/efer-a-friend",
	"https://google.com",
	"https://twitter.com",
}

func fetchUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(resp.Status)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	for _, url := range urls {
		go fetchUrl(url)
	}
	fmt.Fprintf(w, "All Response received successfully")
}

func handleRequests() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":7079", nil)
}

func main() {
	handleRequests()
}