package router

import (
	"github.com/gin-gonic/gin"
	"github.com/weizhe0422/GolangPractice/Practice4/demo/handlers"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handlers.HelloPage)
	}
}
