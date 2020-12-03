package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"Walter0697/GinBackend/middleware"
	"Walter0697/GinBackend/base" 
	"Walter0697/GinBackend/controllers"
	"Walter0697/GinBackend/service"
)


func main() {

	fmt.Println("server starting")
	r := gin.New()

	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	fmt.Println("Connecting to Database...")
	base.ConnectDataBase()

	//only initializing the admin account so that it has one authorized account to create another account
	newuser := service.CreateAccount("system", "123456", 1)
	fmt.Println(newuser)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthorizeJWT())
	{
		//account controller
		authorized.POST("/account", controllers.CreateAccount)
		authorized.PATCH("/password", controllers.ChangePassword)
		authorized.GET("/profile", controllers.Profile)

		//book controller
		authorized.POST("/books", controllers.CreateBook)
		authorized.GET("/books/:id", controllers.FindBook)
		authorized.PATCH("/books/:id", controllers.UpdateBook)
		authorized.DELETE("/books/:id", controllers.DeleteBook)
	}

	r.GET("/books", controllers.FindBooks)
	r.GET("/allusers", controllers.FindAllUsers)
	
	r.POST("/login", controllers.Login)

	fmt.Println("Running the server:")
	r.Run(":5000")
}