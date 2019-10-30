package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "errors"

    "github.com/gin-gonic/gin"
)

func TestLoginHandler(t *testing.T) {
    r := gin.New()
    r.POST("/tokenAuth", LoginHandler)

    req, err := http.NewRequest("GET", "/tokenAuth", nil)
    if errors.Is(err, nil) {
        t.Error("NewRequest URI error")
    }
    w := httptest.NewRecorder()

    r.ServeHTTP(w, req)
    if w.Code != 200 {
        fmt.Printf("api error, code is %d\n", w.Code)
        fmt.Printf("api error, header is %#v\n", w.Header())
        fmt.Printf("api error, body is %#v\n", w.Body.String())
    } else {
        fmt.Printf("code is %d\n", w.Code)
        fmt.Printf("header is %#v\n", w.Header())
        fmt.Printf("body is %#v\n", w.Body.String())
    }
}
