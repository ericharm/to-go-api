package controllers

import (
    "fmt"
    "html"
    "net/http"
)



// this route can go in another controller
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

