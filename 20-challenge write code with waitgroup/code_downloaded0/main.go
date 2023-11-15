package main

import (
	"fmt"
	"net/http"
	"sync"
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

func fetchUrl(url string, wg *sync.WaitGroup) {
	// Write your code here
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write your code here
}

func handleRequests() {
	// Write your code here
}

func main() {
	handleRequests()
}
