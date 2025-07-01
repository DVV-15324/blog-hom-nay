package auth

import (
	"bloghomnay-project/common"
	entityAuth "bloghomnay-project/services/entity/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *ApiAuth) ApiRegisterAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityAuth.RegisterForm
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			app := common.NewAppError(400, http.StatusText(400), err)
			common.NewErrorH(c, app)
			return
		}
		// dang ki
		er := api.bz.BzRegisterAuth(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Dang ki thanh cong")
	}
}
