package comment

import (
	"bloghomnay-project/common"
	entityImg "bloghomnay-project/services/entity/img"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *ApiImg) ApiCreateImg() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityImg.CreateImg

		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		fmt.Println(data)
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		data.UserId = int(uid.LocalID)
		er := api.bz.BusinessCreateImg(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tao thanh cong")

	}
}
