package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	getServerIP()
	router := gin.Default()
	fmt.Println("Hello, world.")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.Run(":8080")
}

func getServerIP() {
	var dbHome string
	dbHome = os.Getenv("SERVER_DB_HOST")
	fmt.Println("get SERVER_DB_HOST: ")
	fmt.Println(dbHome)
}
