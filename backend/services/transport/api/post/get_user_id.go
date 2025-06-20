package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByUserId() func(c *gin.Context) {
	return func(c *gin.Context) {
<<<<<<< HEAD

		id := common.GetRequestContext(c.Request.Context())

		uid := common.DecodeFromBase58(id.GetSub())
=======
		var id = c.Param("user-id")
		uid := common.DecodeFromBase58(id)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
