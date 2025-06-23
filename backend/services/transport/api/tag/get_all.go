package tag

import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

func (api *ApiTag) ApiGetALLTag() func(c *gin.Context) {
	return func(c *gin.Context) {
		tag, er := api.bz.BusinessGetALLTag(c)
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
