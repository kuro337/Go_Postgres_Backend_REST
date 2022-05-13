<link rel="stylesheet" href="style.css" />

- [Go with React & Postgres](#go-with-react--postgres)
	- [Go with Fiber](#go-with-fiber)
	- [Database](#database)
		- [Modularizing Packages](#modularizing-packages)
		- [JWT Integration for Go Backend](#jwt-integration-for-go-backend)

# Go with React & Postgres

## Go with Fiber

<ul>

<li>Starting a server using Fiber and setting a GET route   </li>

</br>

```
main.go :

package main

import (
"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/" , func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":8000")
}

Steps to run :
go mod init main.go
go get github... // fiberpackage
go run main.go

Starts server at localhost:8000

// Note : In Go , error is a type

```

</br>

<li><k>Optional: </k> Starting web server using Gin  </li>

</br>

```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", hello)

	router.Run("localhost:8080")
}

func hello( c *gin.Context) {
	c.IndentedJSON(http.StatusOK , "Hello Susma")
}

```

</br>

</br>

```
GIN Framework Notes:

Reading Request Query String Params
c.Param("id")

Ex:
If we had a route set up such as

app.GET("/api/users/:id" , GetUser)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	msg:= "user id is " + id

	c.IndentedJSON(http.StatusOK , msg )


}

Checking Method of Request

c.Request.Method

if c.Request.Method() == "GET" {
	// do something
}
```

</br>

<li><k>Go Pointer basics </k>  </li>

</br>

```
name := "Chin"

// &name will print memory address
fmt.Println(name , &name)

// Chin 0xc0009200


var name *string = new(string)
// setting value
*name = "Chin"

fmt.Println(name , *name)

// 0xc0009200 Chin


```

</br>

## Database

</br>

<li> Note: Using <k>Postgres + Golang ORM</k> instead of  <k>MySQL + GORM</k> as the <k>DB+ORM</k> as used in the course </li>

<li>Download PSQL for creating DB and pgAdmin for managing  </li>

</br>

```
Default username for Postgres after installing is postgres

Connecting to DB :

psql -U postgres

// Create DB

CREATE DATABASE firstdb WITH ENCODING 'UTF8' LC_COLLATE='English_United States' LC_CTYPE='English_United States';

// Create User

CREATE USER AND GRANT PERMISSIONS

create user testuser with encrypted password 'password';
grant all privileges on database firstDB to testuser;

// Connect to DB

\c firstdb

// Showing all tables in DB

\dt

```

</br>

</br>

```
Code to Connect to DB and start server at localhost

package main

import (
"fmt"
"net/http"

"github.com/gin-gonic/gin"

"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=testuser password=password dbname=firstdb port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("no db connection")
	}

	fmt.Println(db)








	router := gin.Default()
	router.GET("/", hello)

	router.Run("localhost:8080")
}

func hello( c *gin.Context) {
	c.IndentedJSON(http.StatusOK , "Hello Susma")
}
```

</br>

### Modularizing Packages

<ol>
<li>Run <code>go mod init main</code> to create a go.mod file for dependency and module tracking  </li>
<li>Create folders for code and import them from our <span>main.go</span> file by the package name </li>

</ol>

</br>

```
Note: we ran command go mod init main hence we are importing packages using "main/packagename" for local imports

File Structure is :
Go_Project /
			main.go
			go.mod
			go.sum
			controllers/authController.go
			database/connect.go
			routes/routes.go

--------------
--------------

main.go :

package main

import (
	"main/database"
	"main/routes"
"github.com/gin-gonic/gin"
)

func main() {

database.Connect()

	router := gin.Default()
	routes.Setup(router)
	router.Run("localhost:8080")
}

--------------
connect.go:

package database

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"fmt"
)

func Connect() {
	dsn := "host=localhost user=testuser password=password dbname=firstdb port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("no db connection")
	}

	fmt.Println(db)
}
--------------

authController.go:

package controllers

import 	(
	"net/http"
	"github.com/gin-gonic/gin"
)


func Hello( c *gin.Context) {
	c.IndentedJSON(http.StatusOK , "Hello Susma")
}

--------------
routes.go:

package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
)

func Setup(app *gin.Engine) {
app.GET("/" , controllers.Hello)
}

```

</br>

<li><k>Creating a DB Schema and adding entries to DB from Go</k>  </li>

</br>

```
Create a User Struct as this will define a table in our DB

package models

type User struct {
	Id uint
	FirstName string
	LastName string
	Email string `gorm:"unique`
	Password string
}

Email string `gorm:"unique` represents that the Email field value should be UNIQUE

Hashing a string in Go : use bcrypt package

import "golang.org/x/crypto/bcrypt"

secretstring:= "secretpassword"

password , _ := bcrypt.GenerateFromPassword([]byte(secretstring) , 14)

// Syntax : bcrypt.GenerateFromPassword(byte array , hash cost )

fmt.Println(password) /// 1313ahbdcmab1313njbn1m3b1m2b3

```

</br>

<li> Code to Register User and Log In User from Go Web Server </li>

</br>

```
File Structure is :
Go_Project /
			main.go
			go.mod
			go.sum
			controllers/authController.go
			database/connect.go
			routes/routes.go
			models/user.go

--------------
--------------

main.go :

package main

import (
	"main/database"
	"main/routes"
"github.com/gin-gonic/gin"
)

func main() {

database.Connect()

	router := gin.Default()
	routes.Setup(router)
	router.Run("localhost:8080")
}

--------------
connect.go:

package database

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"fmt"
	"main/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=testuser password=password dbname=firstdb port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil {
		panic("no db connection")
	}

	// pointer to DB
	DB = db

	fmt.Println(db)
	db.AutoMigrate(&models.User{})
}
--------------

authController.go:

package controllers

import 	(
	"net/http"
	"github.com/gin-gonic/gin"
	"main/models"
	"main/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

// using bcrypt to encrypt password passed by user and converting to byte array
password , _ := bcrypt.GenerateFromPassword([]byte(data["password"]) , 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		Password: password,
	}

	// This will create and store the record in our DB
	database.DB.Create(&user)

	// c.IndentedJSON(http.StatusOK , "Hello Susma")
	 c.IndentedJSON(http.StatusOK , user)
}

func Login(c *gin.Context) {

	var data map[string]string

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
if err :=  bcrypt.CompareHashAndPassword(user.Password , []byte(data["password"])) ; err!=nil {
	c.IndentedJSON(http.StatusBadRequest , "Incorrect Password")
	return
}

c.IndentedJSON(http.StatusOK , user)


}

--------------
routes.go:

package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
)

// Registering User

func Setup(app *gin.Engine) {
app.POST("/api/register" , controllers.Register)
app.POST("/api/login" , controllers.Login)

}
--------------

user.go :
package models

// defines the columns for a Table in our DB

type User struct {
	Id uint
	FirstName string
	LastName string
	Email string `gorm:"unique`
	Password []byte
}
```

</br>

### JWT Integration for Go Backend

<li>This is a JWT</li>

<li>
<span style="color:#ffa8a8">eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9</span>.<span style="color:#d0bfff">JzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.</span>
<span style="color:#c3fae8">SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c</span>
</li>

<li> Represents :  </li>

</br>

```
Header
{
  "alg": "HS256",
  "typ": "JWT"
}

Payload
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}

Signature

HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),

your-256-bit-secret

```

</br>

<li>Integrating JWT   </li>

</br>

```
Get Package - 	"github.com/dgrijalva/jwt-go"


```

</br>

<li> <k>Storing the JWT in a Cookie</k> </li>

<li> <span>func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
</span> </li>

</br>

```

```

</br>

<li> Creating an Image Upload API  </li>

</br>

```
Will be POST request to localhost:8080/api/upload

In Postman within Body instead of raw we need to use form-data

Click on Key and Change Text to File

Set Key to Image and for Value we will get File Upload Dropdown

Key - Image , Upload file to Value and Hit Send

imageController.go :

package controllers

import (
		"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func Upload(c *gin.Context) {

	// Loads the form used by Client
	form , err := c.MultipartForm()

	if err!= nil {
		fmt.Println("multi" , err)
		return
	}

// Loads the files , we used key of image

files := form.File["image"]
filename := ""

// Loop through files and save the image to /uploads/imagename on our backend

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

routes.go :

app.POST("/api/upload" , controllers.Upload)




```

</br>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>

<li>  </li>
</ul>
