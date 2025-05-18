package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiUpdatePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityPosts.UpdatePost
		var id = c.Param("post-id")
		uid := common.DecodeFromBase58(id)
		//lấy dữ liệu
		err := c.ShouldBindJSON(&data)
		if err != nil {
			fmt.Println("JSON bind error:", err)

			return
		}
		fmt.Println(data)
		if data.CategoryId != nil {
			*data.CategoryId = int(common.DecodeFromBase58(*data.FakeCategoryId).LocalID)
		}
		for i := 0; i < len(data.Tag); i++ {
			data.Tag[i].Id = int(common.DecodeFromBase58(data.Tag[i].FakeId).LocalID)
		}
		er := api.bz.BusinessUpdatePostByID(c, &data, int(uid.LocalID))
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Cap nhat thanh cong")

	}
}
