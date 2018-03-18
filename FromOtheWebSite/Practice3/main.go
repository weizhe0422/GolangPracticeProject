package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", HelloPage)
		v1.GET("hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "hello, %s", name)
		})
		v1.GET("/helloWho", func(c *gin.Context) {
			firstName := c.DefaultQuery("firstname", "Guest")
			lastName := c.Query("lastname")
			c.String(http.StatusOK, "Hello, %s %s", firstName, lastName)

		})
	}

	r.LoadHTMLGlob("templates/*")
	v2 := r.Group("/v2")
	{
		v2.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title": "hello, Gin",
			})
		})

		v2.POST("/post", func(c *gin.Context) {
			id := c.Query("id")
			page := c.DefaultQuery("page", "0")
			name := c.PostForm("name")
			message := c.PostForm("message")

			fmt.Printf("id: %s, page: %s, name: %s, message: %s", id, page, name, message)

		})
	}

	r.Run(":8080")

}
