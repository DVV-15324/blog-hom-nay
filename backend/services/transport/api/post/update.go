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
<<<<<<< HEAD

		// Lấy dữ liệu JSON từ body request
		err := c.ShouldBindJSON(&data)
		if err != nil {
			fmt.Println("JSON bind error:", err)
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		if data.FakeCategoryId != nil {
			val := int(common.DecodeFromBase58(*data.FakeCategoryId).LocalID)
			data.CategoryId = &val
		}

		for i := 0; i < len(data.Tag); i++ {
			data.Tag[i].Id = int(common.DecodeFromBase58(data.Tag[i].FakeId).LocalID)
		}
		// Gọi business logic cập nhật post
		api.bz.BusinessUpdatePostByID(c, &data, int(uid.LocalID))

		common.NewSuccessH(c, "Cập nhật thành công")
=======
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

>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	}
}
