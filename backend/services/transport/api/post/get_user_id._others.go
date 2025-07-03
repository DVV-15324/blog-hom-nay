package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByUserOtherId() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("id")
		uid := common.DecodeFromBase58(id)
		post, er := api.bz.BusinessGetPostByUserId(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(post); i++ {
			post[i].Mask()
		}
		common.NewSuccessH(c, post)

	}
}
