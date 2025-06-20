package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostById() func(c *gin.Context) {
	return func(c *gin.Context) {
<<<<<<< HEAD
		var postId = c.Param("post-id")
		uidPostId := common.DecodeFromBase58(postId)
		userId := common.GetRequestContext(c.Request.Context())
		uidUserId := 0
		if userId != nil {
			uidUserId = int(common.DecodeFromBase58(userId.GetSub()).LocalID)
		}

		post, er := api.bz.BusinessGetPostByID(c, uidUserId, int(uidPostId.LocalID))
=======
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		post, er := api.bz.BusinessGetPostByID(c, int(uid.LocalID))
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
<<<<<<< HEAD

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		post.Mask()
		common.NewSuccessH(c, post)

	}
}
