
package main

import (
	"Go-Blog-Project-master/controllers"
	"Go-Blog-Project-master/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)



func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {


	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/templates", "./templates")


	router.GET("/index", indexHandler)
	router.GET("/signup", signupHandler)
	router.GET("/login", loginHandler)
	router.GET("/support", supportHandler)
	router.GET("/private", privateHandler)

	article := router.Group("/")
	{
		post := new(controllers.PostController)
	article.GET("/write", writeHandler)
	article.POST("/write", post.NewPost)
	}



	auth := router.Group("/")
	{

		hello := new(controllers.HelloWorldController)


		auth.GET("/hello", hello.Default)


		user := new(controllers.UserController)

		auth.POST("/signup", user.Signup)

		auth.POST("/login", user.Login)
		// Password reset
		auth.PUT("/password-reset", user.PasswordReset)
		// Send reset link
		auth.PUT("/reset-link", user.ResetLink)
		// Send verify link
		auth.PUT("/verify-link", user.VerifyLink)
		// Verify account
		auth.PUT("/verify-account", user.VerifyAccount)
		// Refresh token
		auth.GET("/refresh", user.RefreshToken)
	}


	router.NoRoute(func(c *gin.Context) {

		c.JSON(404, gin.H{"message": "Not found"})
	})


	router.Run(":8080")
}



func indexHandler(c *gin.Context)  {

	c.HTML(http.StatusOK, "index.html", models.GetAllPosts())
}


func signupHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"title" : "main website",
	})
}
func loginHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title" : "main website", //ignore
	})
}

func privateHandler(с *gin.Context)  {
	с.HTML(http.StatusOK, "private.html", models.GetMyPage())
}

func supportHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "support.html", gin.H{
		"title" : "main website", //ignore
	})
}

func writeHandler(c *gin.Context)  {
	c.HTML(http.StatusOK, "write.html", gin.H{
		"title" : "main website", //ignore
	})
}
