package categories

import (
	"bloghomnay-project/common"
	entityCategories "bloghomnay-project/services/entity/categories"
	"github.com/gin-gonic/gin"
)

func (api *ApiCategories) ApiUpdateCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityCategories.UpdateCategory
		var id = c.Param("categories-id")
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		uid := common.DecodeFromBase58(id)
		er := api.bz.BusinessUpdateCategories(c, &data, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Cap nhat thanh cong")

	}
}
