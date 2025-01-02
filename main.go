package main

import (
    "net/http"
)

func redirectToPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../page1/index.html")
}


func main() {
	http.HandleFunc("/taxa", redirectToPage)
	http.ListenAndServe(":8080", nil)

}
