// server.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/get_something", func(c *gin.Context) {
		key := c.Query("key")
		c.String(http.StatusOK, "you got %s", key)
		// c.JSON(http.StatusOK, gin.H{
		// 	"got": key,
		// })
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
