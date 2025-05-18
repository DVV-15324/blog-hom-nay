package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiCreateComment() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityComment.CreateComment
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
		er := api.bz.BusinessCreateComment(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tao thanh cong")

	}
}
