package user

import (
	"bloghomnay-project/common"
	entityUser "bloghomnay-project/services/entity/user"
<<<<<<< HEAD
	//"fmt"
=======
	"fmt"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiUpdateUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityUser.UpdateUserForm

<<<<<<< HEAD
		//var id = c.Param("user-id")
=======
		var id = c.Param("user-id")
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8

		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
<<<<<<< HEAD
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		//fmt.Println(data)

=======
		fmt.Println(data)
		uid := common.DecodeFromBase58(id)
		//lấy dữ liệu
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		er := api.bz.BzUpdateUser(c, &data, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Cap nhat thanh cong")

	}
}
