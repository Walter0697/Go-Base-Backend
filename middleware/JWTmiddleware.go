package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		length := len(BEARER_SCHEMA)
		authLength := len(authHeader)
		if authLength == 0 || length >= authLength {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			token, _:= JWTAuthService().ValidateToken(tokenString)
			if !token.Valid {
				c.AbortWithStatus(http.StatusUnauthorized)
				//claims := token.Claims.(jwt.MapClaims)
				//fmt.Println(claims)
			} 
		}
		
	}
}