package main

import (
    "log"
    _ "log"
    "net/http"
    _ "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w,r)
        return

    }
    w.Write([]byte("шалом"))
}

func showSnippet(w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("Отображения заметки"))
}

func createSnippet(w http.ResponseWriter, r *http.Request)  {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", "http.MethodPost")
        http.Error(w, "Метод запрещен", 405)
        return
    }

    w.Write([]byte("Coздание новой заметки"))
}

func main()  {
    mux :=http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
    
}

