package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}

func main() {
    fmt.Println("Hello!")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })
    router.Run(":8080")
}