package controllers

import (
	"Go-Blog-Project-master/forms"
	"Go-Blog-Project-master/helpers"
	"Go-Blog-Project-master/models"
	"Go-Blog-Project-master/services"
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"

)

var userModel = new(models.UserModel)

type UserController struct{}

func (u *UserController) Signup(c *gin.Context) {

	var err error
	Name := c.PostForm("name")
	Email := c.PostForm("email")
	Password := c.PostForm("password")
	About := c.PostForm("about")
	Image := c.PostForm("image")

	data := forms.SignupUserCommand{
		Name:     Name,
		Email:    Email,
		Password: Password,
		About: About,
		Image: Image,
	}

	if err != nil{

		c.JSON(406, gin.H{"message": "Provide relevant fields"})

		c.Abort()

		return
	}


	result, _ := userModel.GetUserByEmail(data.Email)

	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err = userModel.Signup(data)

	resetToken, _ := services.GenerateNonAuthToken(data.Email)

	link := "http://localhost:8080/verify-account?verify_token=" + resetToken
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	email := services.SendMail("Verify Account", body, data.Email, html, data.Name)

	// If email fails while sending
	if !email {
		c.JSON(500, gin.H{"message": "An issue occured sending you an email"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}


	c.Redirect(http.StatusMovedPermanently, "/index")
}

func (u *UserController) Login(c *gin.Context) {

	var err error

	Email := c.PostForm("email")
	Password := c.PostForm("password")

	data := forms.LoginUserCommand{
		Email:    Email,
		Password: Password,
	}

	if err != nil{
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if !result.IsVerified {
		c.JSON(403, gin.H{"message": "Account is not verified"})
		c.Abort()
		return
	}

	hashedPassword := []byte(result.Password)

	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)



	jwtToken, refreshToken, err2 := services.GenerateToken(data.Email)


	if err2 != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}
	fmt.Println(jwtToken, refreshToken)//шоб не ругалось
	c.Redirect(http.StatusMovedPermanently, "/index")

}





func (u *UserController) PasswordReset(c *gin.Context) {
	var data forms.PasswordResetCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	if data.Password != data.Confirm {
		c.JSON(400, gin.H{"message": "Passwords do not match"})
		c.Abort()
		return
	}

	resetToken, _ := c.GetQuery("reset_token")

	userID, _ := services.DecodeNonAuthToken(resetToken)

	result, err := userModel.GetUserByEmail(userID)

	if err != nil {

		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	newHashedPassword := helpers.GeneratePasswordHash([]byte(data.Password))


	_err := userModel.UpdateUserPass(userID, newHashedPassword)

	if _err != nil {

		c.JSON(500, gin.H{"message": "Something happened while updating your password try again"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "Password has been updated, log in"})
	c.Abort()
	return
}


func (u *UserController) ResetLink(c *gin.Context) {
	var data forms.ResendCommand

	if (c.BindJSON(&data)) != nil {
		c.JSON(400, gin.H{"message": "Provided all fields"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	resetToken, _ := services.GenerateNonAuthToken(result.Email)

	link := "http://localhost:5000/api/v1/password-reset?reset_token=" + resetToken
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	email := services.SendMail("Reset Password", body, result.Email, html, result.Name)

	if email == true {
		c.JSON(200, gin.H{"messsage": "Check mail"})
		c.Abort()
		return
	} else {
		c.JSON(500, gin.H{"message": "An issue occured sending you an email"})
		c.Abort()
		return
	}
}

func (u *UserController) VerifyLink(c *gin.Context) {
	var data forms.ResendCommand

	if (c.BindJSON(&data)) != nil {
		c.JSON(400, gin.H{"message": "Provided all fields"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	resetToken, _ := services.GenerateNonAuthToken(result.Email)

	link := "http://localhost:5000/api/v1/verify-account?verify_token=" + resetToken
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	email := services.SendMail("Verify Account", body, result.Email, html, result.Name)

	if email == true {
		c.JSON(200, gin.H{"messsage": "Check mail"})
		c.Abort()
		return
	} else {
		c.JSON(500, gin.H{"message": "An issue occured sending you an email"})
		c.Abort()
		return
	}
}

func (u *UserController) VerifyAccount(c *gin.Context) {
	verifyToken, _ := c.GetQuery("verify_token")

	userID, _ := services.DecodeNonAuthToken(verifyToken)

	result, err := userModel.GetUserByEmail(userID)

	if err != nil {

		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User accoun was not found"})
		c.Abort()
		return
	}


	_err := userModel.VerifyAccount(userID)

	if _err != nil {
		c.JSON(500, gin.H{"message": "Something happened while verifying you account, try again"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "Account verified, log in"})
}

func (u *UserController) RefreshToken(c *gin.Context) {
	refreshToken := c.Request.Header["Refreshtoken"]

	if refreshToken == nil {
		c.JSON(403, gin.H{"message": "No refresh token provided"})
		c.Abort()
		return
	}

	email, err := services.DecodeRefreshToken(refreshToken[0])

	if err != nil {
		c.JSON(500, gin.H{"message": "Problem refreshing your session"})
		c.Abort()
		return
	}


	accessToken, _refreshToken, _err := services.GenerateToken(email)

	if _err != nil {
		c.JSON(500, gin.H{"message": "Problem creating new session"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Log in success", "token": accessToken, "refresh_token": _refreshToken})
}