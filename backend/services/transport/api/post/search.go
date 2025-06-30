package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiSearch() func(c *gin.Context) {
	return func(c *gin.Context) {
		var search = c.Query("search")
		post, er := api.bz.BussinessSearchPost(c, search)
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
