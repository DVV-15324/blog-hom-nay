package user

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiDeleteUser() func(c *gin.Context) {
	return func(c *gin.Context) {
<<<<<<< HEAD

		id := common.GetRequestContext(c.Request.Context())

		uid := common.DecodeFromBase58(id.GetSub())

=======
		var id = c.Param("user-id")
		uid := common.DecodeFromBase58(id)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		er := api.bz.DeleteUserById(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
