package main

import (
    "net/http"
    "os"

    "github.com/russross/blackfriday"
)

func main () {
    puerto: = os.getenv("PORT")
    if puerto == "" {
        puerto = "8080"
    }

    http.HandleFunc("/markdown" , GeneraDesdeMarkdown)
    http.Handle("/", http.FileServer(http.Dir("publico")))
    http.ListenAndServe(":" + puerto, nil)
}

func GeneraDesdeMarkdown (rw http.ResponseWriter, r *http.Request) {
    html := blackfriday.MarkdownCommon([]byte (r.FormValue("cuerpo")))
    rw.Write(html)
}
