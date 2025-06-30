package user

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiDeleteUser() func(c *gin.Context) {
	return func(c *gin.Context) {

		id := common.GetRequestContext(c.Request.Context())

		uid := common.DecodeFromBase58(id.GetSub())

		er := api.bz.DeleteUserById(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
