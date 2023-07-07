package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.POST("/params", func(c *gin.Context) {

		// id := c.Query("id")
		// page := c.DefaultQuery("page", "0")
		// name := c.PostForm("name")
		message := c.PostForm("params")
		fmt.Println(message)
		// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

	})
	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.POST("/upload", func(c *gin.Context) {
	// 	// Single file
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	c.SaveUploadedFile(file, dst)

	// 	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// })
	router.Run("0.0.0.0:8080")

	// or call server listen and serve directly
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }
	// server.ListenAndServe()
}
