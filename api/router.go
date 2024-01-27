package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gokch/cafe_manager/api/v1"
	"github.com/gokch/cafe_manager/service"
)

func InitRouter(serv *service.Service, r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.NoRoute()
	// r.NoMethod()

	authRouter := r.Group("/auth")
	authRouter.GET("/register", v1.Register(serv.Admin))

	jwt := v1.NewJWTMiddleware(serv.Admin)
	authRouter.POST("/login", jwt.LoginHandler)
	authRouter.POST("/logout", jwt.LogoutHandler)

	menuRouter := r.Group("/menu")
	menuRouter.Use(jwt.MiddlewareFunc()) // check auth
	menuRouter.GET("/get", v1.GetMenu(serv.Menu))
	menuRouter.GET("/search", v1.SearchMenu(serv.Menu))
	menuRouter.GET("/list", v1.ListMenu(serv.Menu))
	menuRouter.POST("/add", v1.AddMenu(serv.Menu))
	menuRouter.POST("/update", v1.UpdateMenu(serv.Menu))
	menuRouter.POST("/delete", v1.DeleteMenu(serv.Menu))
}
