package user

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiGetUserById() func(c *gin.Context) {
	return func(c *gin.Context) {
		//uid := common.DecodeFromBase58(id)
		id := common.GetRequestContext(c.Request.Context()).GetSub()
		uid := common.DecodeFromBase58(id)
		user, er := api.bz.BzGetUsersById(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}

		user.Mask()

		common.NewSuccessH(c, user)

	}
}
