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
}

func GetExcute() *cobra.Command {
	return root
}
