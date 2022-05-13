package middleware

import (
	"github.com/gin-gonic/gin"
	"main/util"
	"net/http"
	"fmt"
	"main/models"
	"main/database"
	"strconv"
)

func IsAuthorized(c *gin.Context , page string) error{

		// pulls Cookie that has been set by Register Func

		cookie , err := c.Cookie("jwt")
		if err!=nil {
			fmt.Println(err)
		}
			 // if user unauthenticated :
			 Id, err := util.ParseJwt(cookie)
			 if  err!=nil  {
				 c.IndentedJSON(http.StatusUnauthorized, "Unauthenticated Request")
			 c.Next()
			 }

			 userId , _ := strconv.Atoi(Id)

			  user:= models.User{
				 Id: uint(userId),
			 }

			 database.DB.Preload("Role").Find(&user)

			 role := models.Role{
				 Id : user.RoleId,
			 }

			 database.DB.Preload("Permissions").Find(&role)

			 if c.Request.Method == "GET" {
				 for _,permissions:= range role.Permissions {
					 if permissions.Name == "view_" + page || permissions.Name == "edit_"+page {
						 return err
					 }

				 }
			 } else {
				for _,permissions:= range role.Permissions {
					if permissions.Name == "edit_" + page  {
						return err
					}

				}
			 }

			 c.IndentedJSON(http.StatusUnauthorized, "Unauthorized to perform action")

return nil
}