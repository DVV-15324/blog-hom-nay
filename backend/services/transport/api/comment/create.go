package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
<<<<<<< HEAD

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	"github.com/gin-gonic/gin"
)

func (api *ApiComment) ApiCreateComment() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityComment.CreateComment
<<<<<<< HEAD

=======
		//lấy dữ liệu
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		err := c.ShouldBindJSON(&data)
		if err != nil {
			return
		}
<<<<<<< HEAD

		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		data.UserID = int(uid.LocalID)
		data.PostID = int(common.DecodeFromBase58(data.FakePostID).LocalID)
		comment, er := api.bz.BusinessCreateComment(c, &data)
=======
		er := api.bz.BusinessCreateComment(c, &data)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
<<<<<<< HEAD
		comment.Mask()
		common.NewSuccessH(c, comment)
=======
		common.NewSuccessH(c, "Tao thanh cong")
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8

	}
}
