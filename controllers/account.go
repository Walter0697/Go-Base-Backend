package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"Walter0697/GinBackend/middleware"
	"Walter0697/GinBackend/apibody"
	"Walter0697/GinBackend/helper"
	"Walter0697/GinBackend/constant"
	"Walter0697/GinBackend/service"
	"Walter0697/GinBackend/utility"
)

// Login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService middleware.LoginService
	jWtService   middleware.JWTService
}

type ProfileInfo struct {
	Username    string	`json:"username"`
	Userrole    uint	`json:"userrole"`
}

func LoginHandler(loginService middleware.LoginService,
	jWtService middleware.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credential apibody.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := c.loginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return c.jWtService.GenerateToken(credential.Username)

	}
	return ""
}

// POST /login
// Login to get jwt token
func Login(c *gin.Context) {
	var credential apibody.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}
	var loginService middleware.LoginService = middleware.FetchingLoginService(credential.Username)
	var jwtService middleware.JWTService = middleware.JWTAuthService()
	var loginController LoginController = LoginHandler(loginService, jwtService)

	if loginService == nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

	token := loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// GET /profile
// For getting the user information, by jwt
func Profile(c * gin.Context) {
	user, err := helper.GetUserByContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	returnProfile := &ProfileInfo{
		Username: user.Username,
		Userrole: user.Userrole,
	}
	c.JSON(http.StatusOK, gin.H{"data": returnProfile})
}

// POST /account
// To create a new account
func CreateAccount(c * gin.Context) {
	var input apibody.CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if account username already exists
	_, err2 := service.FindUserByName(input.Username)
	if (err2 == nil) {
		// If there is no error, then the database found something
		c.JSON(http.StatusBadRequest, gin.H{"error": "This username has already been taken"})
		return
	}
	
	// Check if the user who request this account is admin
	currentuser, _ := helper.GetUserByContext(c)
	adminrole := constant.GetAdminRole()
	if (!utility.IsAuthorized(currentuser.Userrole, adminrole)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "this account has no authorization to create account"})
		return
	}

	// Create an account in database
	user := service.CreateAccount(input.Username, input.Password, input.Userrole)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /password
// To change the password for yourself
func ChangePassword(c * gin.Context) {
	var input apibody.UpdateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the user is the account owner
	currentuser, _ := helper.GetUserByContext(c)
	if currentuser.Username != input.Username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: the user who requested this password changed is not the account owner"})
		return
	}

	//update the password in database
	isChanged, err2 := service.ChangePassword(input)
	if (!isChanged) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

// GET /allusers
// To list all the users
func FindAllUsers(c *gin.Context) {
	users := service.FindAllUsers()

	c.JSON(http.StatusOK, gin.H{"data": users})
}