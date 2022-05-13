package controllers 

import (
		"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func Upload(c *gin.Context) {
	form , err := c.MultipartForm()

	if err!= nil {
		fmt.Println("multi" , err)
		return
	}

files := form.File["image"]
filename := ""

for _,file := range files {
	filename = file.Filename

	if err:=c.SaveUploadedFile(file , "./uploads/" + filename); err!=nil {
		fmt.Println(err , "location")
		return 
	}
}

c.IndentedJSON(http.StatusOK , gin.H{
	"url" : "http://localhost:8080/api/uploads/" + filename,
})

}