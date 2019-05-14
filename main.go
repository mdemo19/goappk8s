package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	fmt.Println("Hello, world.")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}
