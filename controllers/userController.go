package controllers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/database"
	"main/models"
	"main/middleware"
"fmt"
"strconv"


)

func AllUsers(c *gin.Context) {
	if err:= middleware.IsAuthorized(c , "users"); err!= nil {
		return
	}

	// if we make request to localhost:8080/api/users?page=1 returns 1st 5 records
	page , _ := strconv.Atoi(c.Query("page"))

	// if user makes req to localhost:8080/api/users without passing page query param , we set page to 1
	if page < 1 {
		page = 1
	}
  
	// // return this many records
	// limit:= 5
	// // if page = 2 , we will return records 6 and onwards (skip first 5)
	// offset := (page-1) * limit

	// // total num of users
	// var total int64

	// // will hold the results we return through API 
	// var users []models.User 

	// // Preload("Role")  will load the foreign key as well (role, rolename)
	// database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	// // Assigning total number of records to total var
	// database.DB.Model(&models.User{}).Count(&total)

	// c.IndentedJSON(http.StatusOK , gin.H{"data":users,
	// "meta":gin.H{"total":total,"page":page,
	// "last_page":math.Ceil(float64(int(total)/limit)),},})

	c.IndentedJSON(http.StatusOK , gin.H(models.Paginate(database.DB , &models.User{} , page)))

}

func CreateUser(c *gin.Context) {
	if err:= middleware.IsAuthorized(c , "users"); err!= nil {
		return
	}

	var user models.User 

	// parsing data from request and assigning to user
	
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
}


	user.SetPassword("1234")

	database.DB.Create(&user)

	c.IndentedJSON(http.StatusOK , user)

}
	// get the user from DB that has the ID in the request 

func GetUser(c *gin.Context) {
	if err:= middleware.IsAuthorized(c , "users"); err!= nil {
		return
	}
	// reading the query param passed to API route 

	// app.GET("/api/users/:id" , controllers.GetUser)

	queryId := c.Param("id")

	// converting id to int 
	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)

	user := models.User{
		Id : uint(id),
	}

	// find user from DB that matches ID passed in query string

	database.DB.Find(&user)

	c.IndentedJSON(http.StatusOK , user)

}

func UpdateUser(c *gin.Context) {
	if err:= middleware.IsAuthorized(c , "users"); err!= nil {
		return
	}

	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)

	var user models.User 

	user = models.User{
		Id : uint(id),
	}

	// parsing data from request and assigning to user
	
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
}

	database.DB.Model(&user).Updates(user)

	c.IndentedJSON(http.StatusOK , user)


}

func DeleteUser(c *gin.Context) {
	if err:= middleware.IsAuthorized(c , "users"); err!= nil {
		return
	}

	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)


}