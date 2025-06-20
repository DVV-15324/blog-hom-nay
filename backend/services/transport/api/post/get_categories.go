package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("categories-id")
		uid := common.DecodeFromBase58(id)

		post, er := api.bz.BusinessGetPostByCategories(c, int(uid.LocalID))
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
