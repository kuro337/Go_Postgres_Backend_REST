package controllers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/database"
	"main/models"
"fmt"
"strconv"

)

func AllRoles(c *gin.Context) {
	var roles []models.Role 

	database.DB.Find(&roles)

	c.IndentedJSON(http.StatusOK , roles)

}

// DTO : Data Transfer Objects 

// 1 role can have many permissions and 1 permission can also belong to multiple roles

// hence many to many

type RoleCreateDTO struct {
	name string
	permissions []string 

}

func CreateRole(c *gin.Context) {
	var roleDto map[string]interface{}

	// parsing data from request and assigning to roleDto
	
	if err := c.BindJSON(&roleDto); err != nil {
		fmt.Println(err)
}
list := roleDto["permissions"].([]interface{})

  permissions := make([]models.Permission , len(list))

	for i,permissionId := range list {
		id,_ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

  role:= models.Role{
		Name: roleDto["name"].(string),
		Permissions : permissions,
	}

	database.DB.Create(&role)

	c.IndentedJSON(http.StatusOK , role)

}
	// get the role from DB that has the ID in the request 

func GetRole(c *gin.Context) {

	// reading the query param passed to API route 

	// app.GET("/api/users/:id" , controllers.GetUser)

	queryId := c.Param("id")

	// converting id to int 
	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)

	role := models.Role{
		Id : uint(id),
	}

	// find user from DB that matches ID passed in query string

	database.DB.Preload("Permissions").Find(&role)

	c.IndentedJSON(http.StatusOK , role)

}

func UpdateRole(c *gin.Context) {

	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)


	var roleDto map[string]interface{}

	// parsing data from request and assigning to user
	
	if err := c.BindJSON(&roleDto); err != nil {
		fmt.Println(err)
}

list := roleDto["permissions"].([]interface{})

  permissions := make([]models.Permission , len(list))

	for i,permissionId := range list {
		id,_ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	// Below 2 lines will delete existing entries for role_permissions table of specified role_id 
	var result interface{}
	database.DB.Table("role_permissions").Where("role_id",id).Delete(&result)

	// Now we create the new entry for role 
  role:= models.Role{
		Id : uint(id),
		Name: roleDto["name"].(string),
		Permissions : permissions,
	}

	database.DB.Model(&role).Updates(role)

	c.IndentedJSON(http.StatusOK , role)


}

func DeleteRole(c *gin.Context) {
	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)


}