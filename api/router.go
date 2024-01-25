package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	v1 "github.com/gokch/cafe_manager/api/v1"
	"github.com/gokch/cafe_manager/service"
)

func InitRouter(serv *service.Service, r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.NoRoute()
	// r.NoMethod()
	jwt := middleware.NewJwtMiddleWare()

	adminRouter := r.Group("/admin")
	adminRouter.GET("/register", v1.Register(serv.Admin))
	adminRouter.POST("/login", v1.Login(serv.Admin), jwt.LoginHandler)
	adminRouter.POST("/logout", v1.Logout(serv.Admin), jwt.LogoutHandler)

	menuRouter := r.Group("/menu")
	menuRouter.Use(jwt.MiddlewareFunc()) // check auth
	menuRouter.GET("/get", v1.GetMenu(serv.Menu))
	menuRouter.GET("/search", v1.SearchMenu(serv.Menu))
	menuRouter.GET("/list", v1.ListMenu(serv.Menu))
	menuRouter.GET("/add", v1.AddMenu(serv.Menu))
	menuRouter.GET("/update", v1.UpdateMenu(serv.Menu))
	menuRouter.GET("/delete", v1.DeleteMenu(serv.Menu))
}
