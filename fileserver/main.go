package main

import (
	"github.com/russross/blackfriday"
	"net/http"
	"os"
)
port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	hhttp.ListenAndServe(":"+port, nil)
}
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
