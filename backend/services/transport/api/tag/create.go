package tag

import (
	"bloghomnay-project/common"
	entityTags "bloghomnay-project/services/entity/tag"
	"github.com/gin-gonic/gin"
)

func (api *ApiTag) ApiCreateTag() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityTags.CreateTagForm
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		er := api.bz.BusinessCreateTag(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tao thanh cong")

	}
}
