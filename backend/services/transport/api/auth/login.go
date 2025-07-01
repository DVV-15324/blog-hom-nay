package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiAuth) ApiLoginAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityAuth.LoginForm
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			app := common.NewAppError(400, http.StatusText(400), err)
			common.NewErrorH(c, app)
			return
		}
		claims, er := api.bz.LoginAuth(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, claims)

	}
}
