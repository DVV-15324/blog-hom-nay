package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostById() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		post, er := api.bz.BusinessGetPostByID(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		post.Mask()
		common.NewSuccessH(c, post)

	}
}
