package v1

import "github.com/gin-gonic/gin"

var AdminApi map[string]gin.HandlerFunc = map[string]gin.HandlerFunc{
	"/admin/register": Register,
	"/admin/login":    Login,
	"/admin/logout":   Logout,
}

var MenuApi map[string]gin.HandlerFunc = map[string]gin.HandlerFunc{
	"/menu/get":    GetMenu,
	"/menu/search": SearchMenu,
	"/menu/add":    AddMenu,
	"/menu/update": UpdateMenu,
	"/menu/delete": DeleteMenu,
}
