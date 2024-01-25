package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	"github.com/gokch/cafe_manager/service"
	"github.com/gokch/cafe_manager/utilx"
)

// FIXME : need structured error
var (
	ErrSeqNotExist    = errors.New("seq is not exist")
	ErrSeqInvalidType = errors.New("invalid seq type")
)

func GetMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq, exist := c.Get("seq")
		if exist != true {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}

		product, err := p.GetMenu(nSeq)
		if err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}

		middleware.HandleData(c, product)
	}
}

func SearchMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AddMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
