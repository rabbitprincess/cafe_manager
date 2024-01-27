package v1

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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
			HandleError(c, http.StatusInternalServerError, err)
			return
		}
		HandleData(c, nil)
	}
}

var identityKey = "id"

var Login struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

func NewJWTMiddleware(admin *service.Admin) *jwt.GinJWTMiddleware {
	// Set up the Gin JWT middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// PayloadFunc: func(data interface{}) jwt.MapClaims {
		// },
		// IdentityHandler: func(c *gin.Context) interface{} {
		// },
		Authenticator: func(c *gin.Context) (interface{}, error) {

			if err := c.ShouldBind(&Login); err != nil {
				HandleError(c, http.StatusUnauthorized, jwt.ErrMissingLoginValues)
				return "", err
			}
			id := Login.ID
			pw := Login.PW

			if err := admin.Login(id, pw); err != nil {
				HandleError(c, http.StatusUnauthorized, err)
				return nil, jwt.ErrFailedAuthentication
			}
			return &Response{
				Meta: Meta{
					Code:    200,
					Message: "ok",
				},
				Data: nil,
			}, nil
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			HandleError(c, http.StatusUnauthorized, errors.New(message))
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		panic("JWT Error:" + err.Error())
	}

	return authMiddleware
}
