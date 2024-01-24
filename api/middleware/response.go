package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Meta Meta `json:"meta"`
	data map[string]interface{}
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseFormat(ctx *gin.Context) {
	ctx.Next()

	// 처리 중에 오류가 있는지 확인
	err := ctx.Errors.Last()
	if err != nil {
		// 오류가 있으면 처리하고 오류 응답을 전송
		code := http.StatusInternalServerError
		message := err.Error()

		ctx.JSON(code, Response{
			Meta: Meta{
				Code:    code,
				Message: message,
			},
			data: nil,
		})
	} else {
		// 오류가 없으면 성공 응답을 전송
		data := ctx.GetStringMap("data")
		var code int
		if len(data) >= 0 {
			code = http.StatusOK
		} else {
			code = http.StatusNoContent
		}

		ctx.JSON(code, Response{
			Meta: Meta{
				Code:    code,
				Message: "ok",
			},
			data: ctx.GetStringMap("data"),
		})
	}
	return
}
