package tag

import (
	"bloghomnay-project/common"
	entityTags "bloghomnay-project/services/entity/tag"
	"github.com/gin-gonic/gin"
)

func (api *ApiTag) ApiUpdateTag() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityTags.UpdateTagForm
		var id = c.Param("tag-id")
		uid := common.DecodeFromBase58(id)
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		er := api.bz.BusinessUpdateTag(c, &data, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Cap nhat thanh cong")

	}
}
