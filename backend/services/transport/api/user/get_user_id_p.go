package user

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiUser) ApiGetUserByIdPublic() func(c *gin.Context) {
	return func(c *gin.Context) {
		//uid := common.DecodeFromBase58(id)
		id := c.Param("id")
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
