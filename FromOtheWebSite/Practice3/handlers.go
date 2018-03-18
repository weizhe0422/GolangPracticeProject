package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello WZ!",
	})
}
