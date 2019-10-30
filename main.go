package main

import (
    "fmt"
    "html/template"
    "os"
    "sync"

    "github.com/gin-gonic/gin"
)

type templateHandler struct {
    once   sync.Once
    filename string
    temp     *template.Template
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    fmt.Println("running...")

    r := gin.Default()

    // Login
    r.POST("/tokenAuth", LoginHandler)
    // Tokenが有効か確認
    r.GET(
        "/tokenAuthenticate",
        RequiredTokenAuthenticationHandler,
    )
    
    r.Run(":" + port)
}
