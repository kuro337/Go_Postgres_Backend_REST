package controllers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/database"
	"main/models"
"fmt"
"strconv"


)

func AllProducts(c *gin.Context) {

	// if we make request to localhost:8080/api/products?page=1 returns 1st 5 records
	page , _ := strconv.Atoi(c.Query("page"))

	// if user makes req to localhost:8080/api/products without passing page query param , we set page to 1
	if page < 1 {
		page = 1
	}

  // we created a function to Paginate results at models/paginate.go 

	c.IndentedJSON(http.StatusOK , gin.H(models.Paginate( database.DB ,&models.Product{}, page )))

}

func CreateProduct(c *gin.Context) {
	var product models.Product 

	// parsing data from request and assigning to product
	
	if err := c.BindJSON(&product); err != nil {
		fmt.Println(err)
}



	database.DB.Create(&product)

	c.IndentedJSON(http.StatusOK , product)

}
	// get the user from DB that has the ID in the request 

func GetProduct(c *gin.Context) {

	// reading the query param passed to API route 

	// app.GET("/api/users/:id" , controllers.GetUser)

	queryId := c.Param("id")

	// converting id to int 
	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)

	product := models.Product{
		Id : uint(id),
	}

	// find user from DB that matches ID passed in query string

	database.DB.Find(&product)

	c.IndentedJSON(http.StatusOK , product)

}

func UpdateProduct(c *gin.Context) {

	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	fmt.Println(id)

	var product models.Product 

	product = models.Product{
		Id : uint(id),
	}

	// parsing data from request and assigning to user
	
	if err := c.BindJSON(&product); err != nil {
		fmt.Println(err)
}

	database.DB.Model(&product).Updates(product)

	c.IndentedJSON(http.StatusOK , product)


}

func DeleteProduct(c *gin.Context) {
	queryId := c.Param("id")

	id ,_ := strconv.Atoi(queryId)

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)


}