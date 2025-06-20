package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
	//"fmt"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiUpdateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityUser.UpdateUserForm

		//var id = c.Param("user-id")

		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		//fmt.Println(data)

		er := api.bz.BzUpdateUser(c, &data, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Cap nhat thanh cong")

	}
}
