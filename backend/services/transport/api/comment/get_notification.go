package comment

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiGetNotification() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id_post = c.Param("post-id")
		uid := common.GetRequestContext(c.Request.Context()).GetSub()
		uid_post := common.DecodeFromBase58(id_post)
		uid_user := common.DecodeFromBase58(uid)

		postLike, er := api.bz.BusinessGetNotification(c, int(uid_user.LocalID), int(uid_post.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(postLike); i++ {
			postLike[i].Mask()
		}

		common.NewSuccessH(c, postLike)

	}
}
