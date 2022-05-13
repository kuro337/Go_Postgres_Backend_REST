package controllers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/database"
	"main/models"

)

func AllPermissions(c *gin.Context) {
	var permissions []models.Permission 

	database.DB.Find(&permissions)

	c.IndentedJSON(http.StatusOK , permissions)

}
