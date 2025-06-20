package categories

import (
	"bloghomnay-project/common"
	entityCategories "bloghomnay-project/services/entity/categories"
	"github.com/gin-gonic/gin"
)

func (api *ApiCategories) ApiCreateCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityCategories.CreateCategory
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		er := api.bz.BusinessCreateCategories(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tao thanh cong")

	}
}
