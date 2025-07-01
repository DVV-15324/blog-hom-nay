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
			fmt.Println(data.Tag[i].FakeId)
			fmt.Println(data.Tag[i].Id)
		}
		// Gọi business logic cập nhật post
		api.bz.BusinessUpdatePostByID(c, &data, int(uid.LocalID))

		common.NewSuccessH(c, "Cập nhật thành công")
	}
}
