package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	v1 "github.com/gokch/cafe_manager/api/v1"
)

func init() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// TODO - prefix
	r.Use(middleware.ResponseFormat)
	for k, v := range v1.AdminApi {
		r.GET(k, v)
	}
	// add auth
	for k, v := range v1.MenuApi {
		r.GET(k, v)
	}
}
