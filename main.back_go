package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello from KBot!")
    token := os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        fmt.Println("No TELEGRAM_TOKEN provided")
    } else {
        fmt.Println("Bot would start here with token:", token)
    }
}
