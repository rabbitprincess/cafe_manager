package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	"github.com/gokch/cafe_manager/service"
)

func Register(admin *service.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		name := c.Params.ByName("name")
		pw := c.Params.ByName("pw")
		phone := c.Params.ByName("phone")

		if err := admin.Register(id, name, pw, phone); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}

func Login(admin *service.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		pw := c.Params.ByName("pw")

		if err := admin.Login(id, pw); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}

func Logout(admin *service.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		// no need
	}
}
