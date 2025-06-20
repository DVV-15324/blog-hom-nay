package command

import (
	"bloghomnay-project/composer"
	"bloghomnay-project/middleware"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "app",
	Short: "Bắt đầu khởi động phần mềm",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		port_api := os.Getenv("PORT_REST_API")
		fmt.Println(port_api)
		g := gin.Default()
<<<<<<< HEAD
		g.Use(middleware.Cors())
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
		StartService(g)
		g.Run(":3000")
	},
}

func StartService(r *gin.Engine) {
	comp := composer.ComposerService()
	v1 := r.Group("v1")
	v1.Use(middleware.Cors())
	//auth
	auth := v1.Group("auth")
<<<<<<< HEAD
	auth.POST("/login", comp.ApiAuth.ApiLoginAuth())
	auth.POST("/register", comp.ApiAuth.ApiRegisterAuth())

	v2 := r.Group("v2")
	v2.Use(middleware.Cors())

	userV2 := v2.Group("user").Use(middleware.RequiredAuth(comp.BzIntrospect))
	userV2.POST("/get_user_sdt", comp.ApiUser.ApiGetUserBySDT())
	userV2.POST("/get_user_id", comp.ApiUser.ApiGetUserById())
	userV2.POST("/delete_user_id", comp.ApiUser.ApiDeleteUser())
	userV2.PATCH("/update_user_id", comp.ApiUser.ApiUpdateUser())

	categoriesV1 := v1.Group("categories")
	categoriesV1.POST("/get_categories", comp.ApiCategories.ApiGetCategories())

	categoriesV2 := v2.Group("categories").Use(middleware.RequiredAuth(comp.BzIntrospect))
	categoriesV2.POST("/create_categories", comp.ApiCategories.ApiCreateCategories())
	categoriesV2.POST("/delete_categories/:categories-id", comp.ApiCategories.ApiDeleteCategories())
	categoriesV2.PATCH("/update_categories/:categories-id", comp.ApiCategories.ApiUpdateCategories())

	postV1 := v1.Group("post")
	postV1.POST("/get_post_all", comp.ApiPost.ApiGetAll())
	postV1.POST("/get_post_categories_id/:categories-id", comp.ApiPost.ApiGetPostByCategories())
	postV1.POST("/get_post_id/:post-id", comp.ApiPost.ApiGetPostById())
	postV1.POST("/get_post_tag", comp.ApiPost.ApiSearch())

	postV2 := v2.Group("post").Use(middleware.RequiredAuth(comp.BzIntrospect))
	postV2.POST("/get_post_user_id", comp.ApiPost.ApiGetPostByUserId())
	postV2.POST("/create_post", comp.ApiPost.ApiCreatePost())
	postV2.POST("/delete_post/:post-id", comp.ApiPost.ApiDeletePost())
	postV2.POST("/update_post_id/:post-id", comp.ApiPost.ApiUpdatePost())

	tagV1 := v1.Group("tag")
	tagV1.POST("/get_all_tag", comp.ApiTag.ApiGetALLTag())

	tagV2 := v2.Group("tag").Use(middleware.RequiredAuth(comp.BzIntrospect))
	tagV2.POST("/create_tag", comp.ApiTag.ApiCreateTag())
	tagV2.POST("/delete_tag/:tag-id", comp.ApiTag.ApiDeleteTag())
	tagV2.POST("/get_tag_post_id/:tag-id", comp.ApiTag.ApiGetTagByPostId())

	commentV2 := v2.Group("comment").Use(middleware.RequiredAuth(comp.BzIntrospect))
	commentV2.POST("/create_comment", comp.ApiComment.ApiCreateComment())
	commentV2.POST("/get_notication/:post-id", comp.ApiComment.ApiGetNotification())

	postlikesV2 := v2.Group("postlikes").Use(middleware.RequiredAuth(comp.BzIntrospect))
	postlikesV2.POST("/create_postlike/:post-id", comp.ApiPostLike.ApiCreatePostLike())
	postlikesV2.POST("/delete_postlike/:post-id", comp.ApiPostLike.ApiDeletePostLike())

	imgV2 := v2.Group("img").Use(middleware.RequiredAuth(comp.BzIntrospect))
	imgV2.POST("/create_img", comp.ApiImg.ApiCreateImg())
	imgV2.POST("/get_img", comp.ApiImg.ApiGetImg())

=======
	{
		auth.POST("/login", comp.ApiAuth.ApiLoginAuth())
		auth.POST("/register", comp.ApiAuth.ApiRegisterAuth())
	}
	v2 := r.Group("v2")
	v2.Use(middleware.Cors())
	v2.Use(middleware.RequiredAuth(comp.BzIntrospect))
	user := v2.Group("user")
	{
		user.POST("/get_user_sdt", comp.ApiUser.ApiGetUserBySDT())
		user.POST("/get_user_id/:user-id", comp.ApiUser.ApiGetUserById())
		user.POST("/delete_user_id/:user-id", comp.ApiUser.ApiDeleteUser())
		user.PATCH("/update_user_id/:user-id", comp.ApiUser.ApiUpdateUser())
	}
	categories := v2.Group("categories")
	{
		categories.POST("/create_categories", comp.ApiCategories.ApiCreateCategories())
		categories.POST("/get_categories", comp.ApiCategories.ApiGetCategories())
		categories.POST("/delete_categories/:categories-id", comp.ApiCategories.ApiDeleteCategories())
		categories.PATCH("/update_categories/:categories-id", comp.ApiCategories.ApiUpdateCategories())
	}
	post := v2.Group("post")
	{
		post.POST("/create_post", comp.ApiPost.ApiCreatePost())
		post.POST("/delete_post/:post-id", comp.ApiPost.ApiDeletePost())
		post.POST("/get_post_all", comp.ApiPost.ApiGetAll())
		post.POST("/get_post_categories_id/:categories-id", comp.ApiPost.ApiGetPostByCategories())
		post.POST("/get_post_id/:post-id", comp.ApiPost.ApiGetPostById())
		post.POST("/get_post_tag", comp.ApiPost.ApiGetPostByTag())
		post.POST("/get_post_titles", comp.ApiPost.ApiGetPostByTitles())
		post.POST("/get_post_user_id/:user-id", comp.ApiPost.ApiGetPostByUserId())
		post.PATCH("/update_post_id/:post-id", comp.ApiPost.ApiUpdatePost())
	}
	tag := v2.Group("tag")
	{
		tag.POST("/create_tag", comp.ApiTag.ApiCreateTag())
		tag.POST("/delete_tag/:tag-id", comp.ApiTag.ApiDeleteTag())
		tag.POST("/get_all_tag", comp.ApiTag.ApiGetALLTag())
		tag.POST("/get_tag_post_id/:tag-id", comp.ApiTag.ApiGetTagByPostId())
		tag.PATCH("/update_tag_id/:tag-id", comp.ApiTag.ApiUpdateTag())
	}
	comment := v2.Group("comment")
	{
		comment.POST("/create_comment", comp.ApiComment.ApiCreateComment())
		comment.POST("/get_comment/:post-id", comp.ApiComment.ApiGetTags())
	}
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

func GetExcute() *cobra.Command {
	return root
}
