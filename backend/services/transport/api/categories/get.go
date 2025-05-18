package categories

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiCategories) ApiGetCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		cate, er := api.bz.BusinessGetALLCategories(c)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(cate); i++ {
			cate[i].Mask()
		}
		common.NewSuccessH(c, cate)

	}
}
