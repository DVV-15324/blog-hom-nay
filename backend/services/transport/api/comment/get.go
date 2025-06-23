package comment

/*
import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiGetComment() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		comment, er := api.bz.BusinessGetComment(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(comment); i++ {
			comment[i].Mask()
		}
		common.NewSuccessH(c, comment)

	}
}
*/
