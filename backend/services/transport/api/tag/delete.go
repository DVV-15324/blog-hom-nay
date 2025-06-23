package tag

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiTag) ApiDeleteTag() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("tag-id")
		uid := common.DecodeFromBase58(id)
		er := api.bz.BusinessDeleteTag(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
