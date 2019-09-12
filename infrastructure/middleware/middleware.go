package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	StudentID string `form:"studentID" json:"studentID" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

/*
NewAuthMiddleware handles jwt middleware
*/
func NewAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "jwt",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24 * 30,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*login); ok {
				return jwt.MapClaims{
					"studentID": v.StudentID,
					"password":  v.Password,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return &login{
				StudentID: claims["studentID"].(string),
				Password:  claims["password"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			studentID := loginVals.StudentID
			password := loginVals.Password

			return &login{StudentID: studentID, Password: password}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*login); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

}
