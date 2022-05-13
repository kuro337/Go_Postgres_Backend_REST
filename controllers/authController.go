package controllers

import 	(
	"net/http"
	"github.com/gin-gonic/gin"
	"main/util"
	"main/models"
	"main/database"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	
)


func Register( c *gin.Context)  {
	// data will hold request sent to API endpoint
	var data map[string]string

	// parsing data from request
	
	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err)
}

if data["password"] != data["password_confirm"] {
	c.IndentedJSON(http.StatusBadRequest , "Passwords do not match")
	return 
}



	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		RoleId : 1,
	}

	// Defined SetPassword on User Struct in models folder

	user.SetPassword(data["password"])

	// This will create and store the record in our DB 
	database.DB.Create(&user)

	// c.IndentedJSON(http.StatusOK , "Hello Susma")
	 c.IndentedJSON(http.StatusOK , user)
}

func Login(c *gin.Context) {

	var data map[string]string 

	// Maps user request to data

	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest , "Bad Request")
		return
}

    var user models.User 


		// Getting the email based on the Email 
		database.DB.Where("email = ?" , data["email"]).First(&user)

		if user.Id == 0 {
			c.IndentedJSON(http.StatusBadRequest , "User Not Found")
		  return 
		}

		// compares password stored in DB and password passed to backend
if err := user.ComparePassword(data["password"]); err!=nil {
	c.IndentedJSON(http.StatusBadRequest , "Incorrect Password")
	return 
}


token, err:=  util.GenerateJwt(strconv.Itoa(int(user.Id)))


if err!=nil {
	c.IndentedJSON(http.StatusBadRequest , "Internal Server Error JWT")
	return 
}

// SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)


  c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

 // note second last arg must be false for Dev in order for cookie to work 
 // https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies#Secure_and_HttpOnly_cookies


c.IndentedJSON(http.StatusOK , gin.H{
	"message" : "success",
})

}

type Claims struct {
	jwt.StandardClaims
}

func User (c *gin.Context) {

	// pulls Cookie that has been set by Register Func

	cookie , err := c.Cookie("jwt")
	
	// Parses JWT cookie into a Claims Struct so we can access encoded data
	// check util module 
	id , err := util.ParseJwt(cookie)

	if err!=nil {
		c.IndentedJSON(http.StatusForbidden  , gin.H{
			"message" : "user not found",
		})
	}


	// // Casting token.Claims as our Claims struct
	// claims := token.Claims.(*Claims) 

	// We find the user from our DB that matches the ID we parsed from cookie
	var user models.User 

	database.DB.Where("id=?", id ).First(&user)

c.IndentedJSON(http.StatusOK , user)
}

func Logout(c *gin.Context) {
	// This function invalidates the cookie so any request we make to the GET user API will not succeed
	// create cookie with expiration in the past 
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", "", 0, "/", "localhost", false, true)

	c.IndentedJSON(http.StatusOK , gin.H{
		"message" : "logged out",
	})
}

func UpdateInfo(c *gin.Context) {

	// getting data from req
	var data map[string]string

	// parsing data from request
	
	if err := c.BindJSON(&data); err != nil {
		fmt.Println(err)
}

	// same code from func User() to pull the User
	// -->
	// pulls Cookie that has been set by Register Func

	cookie , err := c.Cookie("jwt")
	
	// Parses JWT cookie into a Claims Struct so we can access encoded data
	// check util module 
	id , err := util.ParseJwt(cookie)

	// setting body from request to user struct 

	userId , _ := strconv.Atoi(id)

	user := models.User{
		Id : uint(userId),
		FirstName : data["first_name"],
		LastName : data["last_name"],
		Email : data["email"],
	}

	if err!=nil {
		c.IndentedJSON(http.StatusForbidden  , gin.H{
			"message" : "user not found",
		})
	}


	// // Casting token.Claims as our Claims struct
	// claims := token.Claims.(*Claims) 


// Updates user in DB 
	database.DB.Model(&user).Updates(user)

	// <--

	c.IndentedJSON(http.StatusOK  , user)

}

func UpdatePassword(c *gin.Context) {
		// getting data from req
		var data map[string]string

		// parsing data from request
		
		if err := c.BindJSON(&data); err != nil {
			fmt.Println(err)
	}

	if data["password"] != data["password_confirm"] {
		c.IndentedJSON(http.StatusBadRequest , "Passwords do not match")
		return 
	}

		// pulls Cookie that has been set by Register Func

		cookie , err := c.Cookie("jwt")
	
		// Parses JWT cookie into a Claims Struct so we can access encoded data
		// check util module 
		id , err := util.ParseJwt(cookie)
	
		if err!=nil {
			c.IndentedJSON(http.StatusForbidden  , gin.H{
				"message" : "user not found",
			})
		}

		
		
	
	
	// Updates Password in DB 
	userId , _ := strconv.Atoi(id)

	user := models.User{
		Id : uint(userId),
	}

		user.SetPassword(data["password"])
	
		database.DB.Model(&user).Updates(user)
	
		// <--
	
		c.IndentedJSON(http.StatusOK  , user)

}