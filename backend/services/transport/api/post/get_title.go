package post

/*
import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByTitles() func(c *gin.Context) {
	return func(c *gin.Context) {
		var title = c.Query("titles")
		post, er := api.bz.BusinessGetPostByTitles(c, title)
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
*/
