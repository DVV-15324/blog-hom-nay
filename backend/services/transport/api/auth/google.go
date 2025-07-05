package auth

import (
	entityAuth "bloghomnay-project/services/entity/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *ApiAuth) ApiGoogleLogin() func(c *gin.Context) {
	return func(c *gin.Context) {
		var form entityAuth.GoogleLoginForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		token, err := a.bz.LoginWithGoogle(c.Request.Context(), &form)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Message})
			return
		}

		c.JSON(http.StatusOK, token)
	}
}
