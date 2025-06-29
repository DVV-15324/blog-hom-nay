package postlikes

import (
	"bloghomnay-project/common"
	entityPostLike "bloghomnay-project/services/entity/postlikes"

	"github.com/gin-gonic/gin"
)

func (api *ApiPostLike) ApiCreatePostLike() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityPostLike.CreatePostLikes
		var id_post = c.Param("post-id")
		uid_post := common.DecodeFromBase58(id_post)

		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		data.UserId = int(uid.LocalID)
		data.PostId = int(uid_post.LocalID)
		er := api.bz.BusinessCreatePostLike(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tao thanh cong")

	}
}
