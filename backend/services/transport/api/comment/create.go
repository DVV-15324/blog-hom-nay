package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"

	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiCreateComment() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityComment.CreateComment

		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}

		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		data.UserID = int(uid.LocalID)
		data.PostID = int(common.DecodeFromBase58(data.FakePostID).LocalID)
		comment, er := api.bz.BusinessCreateComment(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		comment.Mask()
		common.NewSuccessH(c, comment)

	}
}
