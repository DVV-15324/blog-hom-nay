package middleware

import (
	"bloghomnay-project/common"
	"context"
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//import "github.com/gin-gonic/gin"

type BzAuth interface {
	BzIntrospectToken(ctx context.Context, accessToken string) (*common.JwtClaims, error)
}

func RequiredAuth(bz BzAuth) func(c *gin.Context) {
	return func(c *gin.Context) {
		//lấy token ở header
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			app := common.NewAppError(403, http.StatusText(http.StatusForbidden), err)
			common.NewErrorH(c, app)
			c.Abort()
			return
		}
		// xác thực token
		claims, er := bz.BzIntrospectToken(c, token)
		if er != nil {
			app := common.NewAppError(403, http.StatusText(http.StatusForbidden), er)
			common.NewErrorH(c, app)
			c.Abort()
			return
		}
		//Lưu vào context
		c.Request = c.Request.WithContext(common.SaveRequestContext(c, common.NewRequestResponse(claims.Subject, claims.ID)))
		c.Next()
	}

}

func extractTokenFromHeader(accessToken string) (string, error) {
	args := strings.Split(accessToken, " ")
	//Thiếu bearer, thiếu token, bị nhiều " "
	if args[0] != "Bearer" || len(args) != 2 {
		return "", errors.New("loi token")
	}
	return args[1], nil
}
