package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/middleware"
)

// Registering User

func Setup(app *gin.Engine) {
app.POST("/api/register" , controllers.Register)
app.POST("/api/login" , controllers.Login)

// Next routes become available only if authenticated
app.Use(middleware.IsAuthenticated)

app.PUT("api/users/info" , controllers.UpdateInfo)
app.PUT("api/users/password" , controllers.UpdatePassword)

app.GET("/api/user" , controllers.User)
app.POST("/api/logout" , controllers.Logout)


// User CRUD
app.GET("/api/users" , controllers.AllUsers)
app.POST("/api/users" , controllers.CreateUser)
app.GET("/api/users/:id" , controllers.GetUser)
app.PUT("/api/users/:id" , controllers.UpdateUser)
app.DELETE("/api/users/:id" , controllers.DeleteUser)

// Role CRUD
app.GET("/api/roles" , controllers.AllRoles)
app.POST("/api/roles" , controllers.CreateRole)
app.GET("/api/roles/:id" , controllers.GetRole)
app.PUT("/api/roles/:id" , controllers.UpdateRole)
app.DELETE("/api/roles/:id" , controllers.DeleteRole)

// Permissions
app.GET("/api/permissions" , controllers.AllPermissions)

// Products CRUD 
app.GET("/api/products" , controllers.AllProducts)
app.POST("/api/products" , controllers.CreateProduct)
app.GET("/api/products/:id" , controllers.GetProduct)
app.PUT("/api/products/:id" , controllers.UpdateProduct)
app.DELETE("/api/products/:id" , controllers.DeleteProduct)

// Upload Images 
app.POST("/api/upload" , controllers.Upload)
app.Static("/api/uploads/" ,"./uploads" )

// Orders
app.GET("/api/orders" , controllers.AllOrders)

// Export CSV 
app.POST("/api/export" , controllers.Export)

// Viewing Sales Chart 
app.GET("/api/chart" , controllers.Chart)

}



// func Hello(app *gin.Engine) {
// 	app.GET("/" , func () {
// 		c.IndentedJSON(http.StatusOK , "Hello Susma")
// 	})
// 	}