package post

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("categories-id")
		uid := common.DecodeFromBase58(id)
<<<<<<< HEAD

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		post, er := api.bz.BusinessGetPostByCategories(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(post); i++ {
			post[i].Mask()
		}
<<<<<<< HEAD

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		common.NewSuccessH(c, post)

	}
}
