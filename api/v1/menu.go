package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokch/cafe_manager/api/middleware"
	"github.com/gokch/cafe_manager/service"
	"github.com/gokch/cafe_manager/utilx"
)

// TODO : need structured error
var (
	ErrSeqNotExist    = errors.New("seq is not exist")
	ErrSeqInvalidType = errors.New("invalid seq type")
)

func GetMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq := c.Query("seq")
		if seq == "" {
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

func ListMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq := c.Query("seq")
		if seq == "" {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}

		products, err := p.ListMenu(nSeq)
		if err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, products)
	}
}

func SearchMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		products, err := p.SearchMenu(name)
		if err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, products)
	}
}

func AddMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.GetString("category")
		name := c.GetString("name")
		description := c.GetString("description")
		price := c.GetInt64("price")
		cost := c.GetInt64("cost")
		expire := c.GetInt64("expire")
		size := c.GetBool("size")

		if err := p.AddMenu(category, name, description, price, cost, expire, size); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}

func UpdateMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category, name, description *string
		var price, cost, expire *int64
		var size *bool
		if _, ok := c.Get("category"); ok {
			categoryVal := c.GetString("category")
			category = &categoryVal
		}
		if _, ok := c.Get("name"); ok {
			nameVal := c.GetString("name")
			name = &nameVal
		}
		if _, ok := c.Get("description"); ok {
			descriptionVal := c.GetString("description")
			description = &descriptionVal
		}
		if _, ok := c.Get("price"); ok {
			priceVal := c.GetInt64("price")
			price = &priceVal
		}
		if _, ok := c.Get("cost"); ok {
			costVal := c.GetInt64("cost")
			cost = &costVal
		}
		if _, ok := c.Get("expire"); ok {
			expireVal := c.GetInt64("expire")
			expire = &expireVal
		}
		if _, ok := c.Get("size"); ok {
			sizeVal := c.GetBool("size")
			size = &sizeVal
		}
		if err := p.UpdateMenu(category, name, description, price, cost, expire, size); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}

func DeleteMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq, exist := c.Params.Get("seq")
		if exist != true {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			middleware.HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}
		if err := p.DeleteMenu(nSeq); err != nil {
			middleware.HandleError(c, http.StatusInternalServerError, err)
			return
		}
		middleware.HandleData(c, nil)
	}
}
