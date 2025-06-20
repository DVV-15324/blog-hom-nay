package postlikes

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPostLike) ApiDeletePostLike() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id_post = c.Param("post-id")
		uid_post := common.DecodeFromBase58(id_post)
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		er := api.bz.BusinessDeletePostLike(c, int(uid.LocalID), int(uid_post.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xóa thanh cong")

	}
}
