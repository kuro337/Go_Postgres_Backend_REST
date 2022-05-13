package middleware

import (
	"github.com/gin-gonic/gin"
	"main/util"
	"net/http"
	"fmt"
)

func IsAuthenticated(c *gin.Context) {

		// pulls Cookie that has been set by Register Func

	cookie , err := c.Cookie("jwt")
 if err!=nil {
	 fmt.Println(err)
 }
		// if user unauthenticated :
		if _, err := util.ParseJwt(cookie); err!=nil  {
			c.IndentedJSON(http.StatusUnauthorized, "Unauthenticated Request")
		c.Next()
		}


}