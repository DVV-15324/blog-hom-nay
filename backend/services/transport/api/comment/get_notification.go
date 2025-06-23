package comment

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiGetNotification() func(c *gin.Context) {
	return func(c *gin.Context) {
		uid := common.GetRequestContext(c.Request.Context()).GetSub()
		uid_user := common.DecodeFromBase58(uid)

		n, er := api.bz.BusinessGetNotification(c, int(uid_user.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(n); i++ {
			n[i].Mask()
		}

		common.NewSuccessH(c, n)

	}
}
