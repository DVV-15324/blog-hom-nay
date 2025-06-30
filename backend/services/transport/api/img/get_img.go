package comment

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiImg) ApiGetImg() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		comment, er := api.bz.BusinessGetImg(c, int(uid.LocalID))
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
