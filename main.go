package main

import (
    "fmt"
    "net/http"
)

func main() {
    hub := NewHub()
    go hub.Run()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        ServeWs(hub, w, r)
    })

    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}
