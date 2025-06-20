package tag

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiTag) ApiGetTagByPostId() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Query("post-id")
		uid := common.DecodeFromBase58(id)
		tag, er := api.bz.BusinessGetTagByPostId(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(tag); i++ {
			tag[i].Mask()
		}
		common.NewSuccessH(c, tag)

	}
}
