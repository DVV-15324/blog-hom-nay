package comment

<<<<<<< HEAD
/*
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

<<<<<<<< HEAD:backend/services/transport/api/img/get_img.go
func (api *ApiImg) ApiGetImg() func(c *gin.Context) {
========
<<<<<<< HEAD
func (api *ApiComment) ApiGetComment() func(c *gin.Context) {
=======
func (api *ApiComment) ApiGetTags() func(c *gin.Context) {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
>>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033:backend/services/transport/api/comment/get.go
	return func(c *gin.Context) {
		id := common.GetRequestContext(c.Request.Context())
		uid := common.DecodeFromBase58(id.GetSub())
		comment, er := api.bz.BusinessGetImg(c, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(comment); i++ {
			comment[i].Mask()
		}
		common.NewSuccessH(c, comment)

	}
}
<<<<<<< HEAD
*/
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
