package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(c *gin.Context, code int, err error) {
	c.JSON(code, Response{
		Meta: Meta{
			Code:    code,
			Message: err.Error(),
		},
		Data: nil,
	})
}

func HandleData(c *gin.Context, data interface{}) {
	var code int
	var message string
	if data == nil {
		code = http.StatusNoContent
		message = "no content"
	} else {
		code = http.StatusOK
		message = "ok"
	}
	c.JSON(code, Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	})
}
