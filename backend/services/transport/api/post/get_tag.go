package post

<<<<<<< HEAD
/*
import (
	"bloghomnay-project/common"
	"fmt"
=======
import (
	"bloghomnay-project/common"
	"strings"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiGetPostByTag() func(c *gin.Context) {
	return func(c *gin.Context) {
<<<<<<< HEAD
		var id = c.QueryArray("tag-id")
		listTag := id
		listTagDe := []int{}
		for i := 0; i < len(listTag); i++ {
			uid := common.DecodeFromBase58(listTag[i])
			fmt.Println(listTag[i])
			fmt.Println(len(listTag[i]))
			fmt.Println(uid)
=======
		var id = c.Query("tag-id")
		listTag := strings.Split(id, "&")
		listTagDe := []int{}
		for i := 0; i < len(listTag); i++ {
			uid := common.DecodeFromBase58(id)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
			listTagDe = append(listTagDe, int(uid.LocalID))
		}

		post, er := api.bz.BusinessGetPostByTag(c, listTagDe)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		for i := 0; i < len(post); i++ {
			post[i].Mask()
		}
		common.NewSuccessH(c, post)

	}
}
<<<<<<< HEAD
*/
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
