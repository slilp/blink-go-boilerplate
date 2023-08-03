package middleware

import (
	"blink-go-gin-boilerplate/models"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

func AuthorizedUser(roles  []models.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		
		if headerToken == ""{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var user models.UserEntity
		if err := validateToken(token,&user); err != nil || !slices.Contains(roles, user.Role){
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

	}
}

func RefreshUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		
		if headerToken == ""{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var user models.UserEntity
		if err := validateRefreshToken(token,&user); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)

	}
}


func validateToken(token string,user *models.UserEntity) error {
	claims := jwt.MapClaims{}
	 _  , err := jwt.ParseWithClaims(token,claims, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET_KEY")), nil
    })

	if err != nil{
		return err
	}
	
	for key, val := range claims {
		if key == "sub"{
			 json.Unmarshal([]byte(val.(string)), user)
		}
	}

	return nil

}

func validateRefreshToken(token string,user *models.UserEntity) error {
	claims := jwt.MapClaims{}
	 _  , err := jwt.ParseWithClaims(token,claims, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("REFRESH_JWT_SECRET_KEY")), nil
    })

	if err != nil{
		return err
	}
	
	for key, val := range claims {
		if key == "sub"{
			 json.Unmarshal([]byte(val.(string)), user)
		}
	}

	return nil

}

func ExtractUserToken(ctx *gin.Context) models.UserEntity {
	return ctx.MustGet("user").(models.UserEntity)
}