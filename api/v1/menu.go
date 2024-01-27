package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
			HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}

		product, err := p.GetMenu(uint64(nSeq))
		if err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}

		HandleData(c, product)
	}
}

func ListMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq := c.Query("seq")
		if seq == "" {
			HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}

		products, err := p.ListMenu(uint64(nSeq))
		if err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, products)
	}
}

func SearchMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		products, err := p.SearchMenu(name)
		if err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, products)
	}
}

var Menu struct {
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Cost        int64  `json:"cost"`
	Expire      int64  `json:"expire"`
	Size        string `json:"size"`
}

func AddMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBind(&Menu); err != nil {
			HandleError(c, http.StatusBadRequest, err)
			return
		}

		if err := p.AddMenu(Menu.Category, Menu.Name, Menu.Description, Menu.Price, Menu.Cost, Menu.Expire, Menu.Size); err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, nil)
	}
}

func UpdateMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		var category, name, description *string
		var price, cost, expire *int64
		var size *string
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
			sizeVal := c.GetString("size")
			size = &sizeVal
		}
		if err := p.UpdateMenu(category, name, description, price, cost, expire, size); err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, nil)
	}
}

func DeleteMenu(p *service.Menu) gin.HandlerFunc {
	return func(c *gin.Context) {
		seq, exist := c.Params.Get("seq")
		if exist != true {
			HandleError(c, http.StatusBadRequest, ErrSeqNotExist)
			return
		}
		nSeq, ok := utilx.ParseInt64(seq)
		if ok != true {
			HandleError(c, http.StatusBadRequest, ErrSeqInvalidType)
			return
		}
		if err := p.DeleteMenu(uint64(nSeq)); err != nil {
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, nil)
	}
}
