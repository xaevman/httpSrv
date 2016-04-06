package main

import (
    "log"
    "net/http"
)

func main() {
    // Simple static webserver:
    log.Fatal(http.ListenAndServe("127.0.0.1:8081", http.FileServer(http.Dir("."))))
}
