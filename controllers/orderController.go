package controllers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/database"
	"main/models"
"strconv"
"os"
"encoding/csv"
"fmt"
)

func AllOrders(c *gin.Context) {

	// if we make request to localhost:8080/api/products?page=1 returns 1st 5 records
	page , _ := strconv.Atoi(c.Query("page"))

	// if user makes req to localhost:8080/api/products without passing page query param , we set page to 1
	if page < 1 {
		page = 1
	}

  // we created a function to Paginate results at models/paginate.go 

	c.IndentedJSON(http.StatusOK , gin.H(models.Paginate( database.DB ,&models.Order{}, page )))

}

func Export(c *gin.Context)  {
	filePath := "./csv/orders.csv"

	if err:= CreateFile(filePath);err!=nil {
		fmt.Println(err)
		return 
	}
	 
	c.File(filePath)
	return  

}

func CreateFile(filePath string) error {
	file , err := os.Create(filePath)

	if err != nil {
		return err
	}

	// the defer expression executes once rest of the code executes 
	defer file.Close() 

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID","Name","Email","Product Title" , "Price" , "Quantity",
	})

	for _,order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"",
			"",
			"",
		}
		if err:= writer.Write(data);err!=nil {
			fmt.Println(err)
			return err
		}

		for _,orderItem:= range order.OrderItems {
			data := []string{
			"",
				"",
				"",
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),

			}
			if err:= writer.Write(data);err!=nil {
				fmt.Println(err)
				return err
			}
		}
	}
 return nil
}

// creating a struct to assign the result of below query to 

type Sales struct {
	Date string `json:"date"`
	Sum string `json:"sum"`
}

func Chart(c *gin.Context) {
	var sales []Sales 

	database.DB.Raw(`
	SELECT o.created_at as date, SUM(oi.price * oi.quantity) as sum
	FROM orders o
	JOIN order_items oi on o.id = oi.order_id
	GROUP BY date
	`).Scan(&sales)

	c.IndentedJSON(http.StatusOK , sales)

}

