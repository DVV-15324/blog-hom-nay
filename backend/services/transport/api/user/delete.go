package user

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiDeleteUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("user-id")
		uid := common.DecodeFromBase58(id)
		er := api.bz.DeleteUserById(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
