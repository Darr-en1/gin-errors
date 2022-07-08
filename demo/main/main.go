package main

import (
	"github.com/Darr-en1/gin-errors/demo"
	"github.com/Darr-en1/gin-errors/errors"
	"github.com/gin-gonic/gin"
	"log"
)

func initRouter() *gin.Engine {
	var router = gin.Default()
	router.GET("/blog", errors.ErrWrapper(web.GetBlog))
	return router
}

func main() {

	router := initRouter()

	err := router.Run()
	if err != nil {
		log.Panic("service startup fails")
	}
}
