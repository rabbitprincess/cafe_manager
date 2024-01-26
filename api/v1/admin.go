package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	"github.com/gokch/cafe_manager/service"
)

// TODO : make swagger
// http://localhost:3000/admin/register?id=admin1&name=admin1&pw=1234&phone=010-1234-5678
func Register(admin *service.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		pw := c.Query("pw")
		phone := c.Query("phone")

		if err := admin.Register(id, name, pw, phone); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}

func Login(admin *service.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		pw := c.Query("pw")

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
