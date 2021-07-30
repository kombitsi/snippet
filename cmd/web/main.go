package main

import (
    "log"
    "net/http"
)


func main()  {
    mux :=http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)
    fileServer := http.FileServer(http.Dir("./ui/static/"))


    log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
    
}

