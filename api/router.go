package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
)

func init() {
	r := gin.Default()
	r.Use(middleware.ResponseFormat)
	r.GET("/ping", func(c *gin.Context) {

	})
	// handler.
}
