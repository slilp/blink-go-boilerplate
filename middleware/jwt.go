package middleware

import (
	"net/http"
	"strings"

	user "github.com/slilp/blink-go-boilerplate/app/user/api"
	"github.com/slilp/blink-go-boilerplate/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func AuthorizedUser(roles  []utils.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		
		if headerToken == ""{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var user user.UserEntity		
		if err := utils.ValidateToken(token,&user); err != nil || !slices.Contains(roles, user.Role){
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
		var user user.UserEntity		
		if err := utils.ValidateRefreshToken(token,&user); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)

	}
}



func ExtractUserToken(ctx *gin.Context) user.UserEntity {
	return ctx.MustGet("user").(user.UserEntity)
}