package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
"math"
)

func Paginate(db *gorm.DB , entity Entity , page int) map[string]interface{} {

		// return this many records
		limit:= 15
		
		// if page = 2 , we will return records 6 and onwards (skip first 5)
		offset := (page-1) * limit

		data:= entity.Take(db , limit , offset)
		total := entity.Count(db)

	
		result := map[string]interface{}{"data":data,
		"meta":gin.H{"total":total,"page":page,
		"last_page":math.Ceil(float64(int(total)/limit)),},}


		return result 

		// c.IndentedJSON(http.StatusOK , gin.H{"data":products,
		// "meta":gin.H{"total":total,"page":page,
		// "last_page":math.Ceil(float64(int(total)/limit)),},})
}