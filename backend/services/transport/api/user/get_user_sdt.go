package user

import (
	"bloghomnay-project/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiGetUserBySDT() func(c *gin.Context) {
	return func(c *gin.Context) {
		var phone = c.Query("phone")

		user, er := api.bz.BzGetUsersBySDT(c, phone)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		fmt.Println(user)
		user.Mask()
		common.NewSuccessH(c, user)

	}
}
