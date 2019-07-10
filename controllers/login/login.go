package login

import (
	"tat_gogogo/utls/login"

	"github.com/gin-gonic/gin"
)

/*
HandleLogin is a function for gin to handle login api
*/
func HandleLogin(c *gin.Context) {
	login.HandleRequest(c)
}
