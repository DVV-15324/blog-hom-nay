package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostById() func(c *gin.Context) {
	return func(c *gin.Context) {
		var postId = c.Param("post-id")
		uidPostId := common.DecodeFromBase58(postId)
<<<<<<< HEAD

		id := common.GetRequestContext(c.Request.Context())
		uidUserId := 0
		if id != nil {
			uidUserId = int(common.DecodeFromBase58(id.GetSub()).LocalID)

=======
		userId := common.GetRequestContext(c.Request.Context())
		uidUserId := 0
		if userId != nil {
			uidUserId = int(common.DecodeFromBase58(userId.GetSub()).LocalID)
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
		}

		post, er := api.bz.BusinessGetPostByID(c, uidUserId, int(uidPostId.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		post.Mask()
		common.NewSuccessH(c, post)

	}
}
