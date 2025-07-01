package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiDeletePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		er := api.bz.BusinessDeletePost(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
