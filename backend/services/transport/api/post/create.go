package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (api *ApiPosts) ApiCreatePost() func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entityPosts.CreatePost

		// 👇 BIND TRƯỚC
		err := c.ShouldBindJSON(&data)
		if err != nil {
			common.NewErrorH(c, common.NewAppError(400, "Bad request", err))
			return
		}

		// 👇 Sau khi có dữ liệu rồi mới decode
		data.CategoryId = int(common.DecodeFromBase58(data.FakeCategoryId).LocalID)
		data.UserId = int(common.DecodeFromBase58(data.FakeUserId).LocalID)
		for i := 0; i < len(data.Tag); i++ {
			data.Tag[i].Id = int(common.DecodeFromBase58(data.Tag[i].FakeId).LocalID)
		}

		// Kiểm tra dữ liệu đã bind và decode
		fmt.Println("Create post data:", data)

		// Gọi business logic để tạo post
		er := api.bz.BusinessCreatePost(c, &data)
		if er != nil {
			common.NewErrorH(c, er)
			return
		}
		common.NewSuccessH(c, "Tạo thành công")
	}
}
