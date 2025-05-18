package categories

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiCategories) ApiDeleteCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		var id = c.Param("categories-id")
		uid := common.DecodeFromBase58(id)
		er := api.bz.BusinessDeleteCategories(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Xoa thanh cong")

	}
}
