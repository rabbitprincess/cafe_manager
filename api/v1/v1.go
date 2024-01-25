package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/service"
)

var (
	AdminApi map[string]gin.HandlerFunc
	MenuApi  map[string]gin.HandlerFunc
)

func InitHandler(serv *service.Service) {
	AdminApi = map[string]gin.HandlerFunc{
		"/admin/register": Register(serv.Admin),
		"/admin/login":    Login(serv.Admin),
		"/admin/logout":   Logout(serv.Admin),
	}

	MenuApi = map[string]gin.HandlerFunc{
		"/menu/get":    GetMenu(serv.Menu),
		"/menu/search": SearchMenu(serv.Menu),
		"/menu/add":    AddMenu(serv.Menu),
		"/menu/update": UpdateMenu(serv.Menu),
		"/menu/delete": DeleteMenu(serv.Menu),
	}
}
