package comment

<<<<<<< HEAD
/*
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
import (
	"bloghomnay-project/common"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
func (api *ApiComment) ApiGetComment() func(c *gin.Context) {
=======
func (api *ApiComment) ApiGetTags() func(c *gin.Context) {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	return func(c *gin.Context) {
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		comment, er := api.bz.BusinessGetComment(c, int(uid.LocalID))
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
