package post

import (
	"bloghomnay-project/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByTag() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Query("tag-id")
		listTag := strings.Split(id, "&")
		listTagDe := []int{}
		for i := 0; i < len(listTag); i++ {
			uid := common.DecodeFromBase58(id)
			listTagDe = append(listTagDe, int(uid.LocalID))
		}

		post, er := api.bz.BusinessGetPostByTag(c, listTagDe)
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
