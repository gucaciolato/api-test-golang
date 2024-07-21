package main

import (
    "log"
    "net/http"
    "github.com/gucaciolato/api-test-golang/handlers"
)

func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)

    log.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
