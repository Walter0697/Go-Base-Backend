package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"

	"Walter0697/GinBackend/middleware"
	"Walter0697/GinBackend/service"
	"Walter0697/GinBackend/models"
)

func GetUserByContext(c * gin.Context) (*models.User, error) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := middleware.JWTAuthService().ValidateToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)
	user, err := service.FindUserByName(claims["name"].(string))
	return user, err
}