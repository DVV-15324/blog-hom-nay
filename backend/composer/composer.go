package composer

import (
	"bloghomnay-project/common"
	bzAuth "bloghomnay-project/services/business/auth"
	bzCategories "bloghomnay-project/services/business/categories"
	bzComment "bloghomnay-project/services/business/comment"
	bzPost "bloghomnay-project/services/business/post"
	bzTag "bloghomnay-project/services/business/tag"
	bzUser "bloghomnay-project/services/business/user"
	responsitoryAuth "bloghomnay-project/services/reponsitory/auth"
	responsitoryCategories "bloghomnay-project/services/reponsitory/categories"
	responsitoryComment "bloghomnay-project/services/reponsitory/comment"
	responsitoryPost "bloghomnay-project/services/reponsitory/post"
	responsitoryPostTag "bloghomnay-project/services/reponsitory/post_tag"
	responsitoryTag "bloghomnay-project/services/reponsitory/tag"
	responsitoryUser "bloghomnay-project/services/reponsitory/user"
	apiAuth "bloghomnay-project/services/transport/api/auth"
	apiCategories "bloghomnay-project/services/transport/api/categories"
	apiComment "bloghomnay-project/services/transport/api/comment"
	apiPost "bloghomnay-project/services/transport/api/post"
	apiTag "bloghomnay-project/services/transport/api/tag"
	apiUser "bloghomnay-project/services/transport/api/user"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/gin-gonic/gin"
)

type BzIntroSpectToken interface {
	BzIntrospectToken(ctx context.Context, accessToken string) (*common.JwtClaims, error)
}
type ApiAuth interface {
	ApiLoginAuth() func(c *gin.Context)
	ApiRegisterAuth() func(c *gin.Context)
}
type ApiUser interface {
	ApiDeleteUser() func(c *gin.Context)
	ApiGetUserById() func(c *gin.Context)
	ApiGetUserBySDT() func(c *gin.Context)
	ApiUpdateUser() func(c *gin.Context)
}
type ApiCategories interface {
	ApiCreateCategories() func(c *gin.Context)
	ApiDeleteCategories() func(c *gin.Context)
	ApiGetCategories() func(c *gin.Context)
	ApiUpdateCategories() func(c *gin.Context)
}
type ApiPosts interface {
	ApiCreatePost() func(c *gin.Context)
	ApiDeletePost() func(c *gin.Context)
	ApiGetAll() func(c *gin.Context)
	ApiGetPostByCategories() func(c *gin.Context)
	ApiGetPostById() func(c *gin.Context)
	ApiGetPostByTag() func(c *gin.Context)
	ApiGetPostByTitles() func(c *gin.Context)
	ApiGetPostByUserId() func(c *gin.Context)
	ApiUpdatePost() func(c *gin.Context)
}
type ApiTag interface {
	ApiCreateTag() func(c *gin.Context)
	ApiDeleteTag() func(c *gin.Context)
	ApiGetALLTag() func(c *gin.Context)
	ApiGetTagByPostId() func(c *gin.Context)
	ApiUpdateTag() func(c *gin.Context)
}
type ApiComment interface {
	ApiCreateComment() func(c *gin.Context)
	ApiGetTags() func(c *gin.Context)
}

type ApiServer struct {
	BzIntrospect  BzIntroSpectToken
	ApiAuth       ApiAuth
	ApiUser       ApiUser
	ApiCategories ApiCategories
	ApiPost       ApiPosts
	ApiComment    ApiComment
	ApiTag        ApiTag
}

func ComposerService() *ApiServer {
	db, er := Connectdb()
	if er != nil {
		fmt.Println("loi ket loi db")
	}
	//repon
	rAuth := responsitoryAuth.NewAuthServiceSQL(db)
	rUser := responsitoryUser.NewUserServiceSQL(db)
	rCategories := responsitoryCategories.NewCategoriesServiceSQL(db)
	rPost := responsitoryPost.NewPostServiceSQL(db)
	rTag := responsitoryTag.NewTagServiceSQL(db)
	rPostTag := responsitoryPostTag.NewPostTagServiceSQL(db)
	rComment := responsitoryComment.NewCommentServiceSQL(db)
	// business
	jwt := common.NewJwtServer("vu-dep-trai-nhat-the-gioi", 604800)
	hash := new(common.Hash)
	bzUser := bzUser.NewBusinessUser(rUser)
	bzAuth := bzAuth.NewBusinessAuth(jwt, bzUser, hash, rAuth)
	bzTag := bzTag.NewBusinessTag(rTag, rPostTag)
	bzCategories := bzCategories.NewBusinessCategories(rCategories)
	bzPost := bzPost.NewBusinessPost(rPost, bzUser, rPostTag, bzTag)
	bzComment := bzComment.NewBusinessComment(rComment, bzUser)
	// api
	apiUser := apiUser.NewApiUser(bzUser)
	apiAuth := apiAuth.NewApiAuth(bzAuth)
	apiTag := apiTag.NewApiTag(bzTag)
	apiCategories := apiCategories.NewApiCategories(bzCategories)
	apiPost := apiPost.NewApiPosts(bzPost)
	apiComment := apiComment.NewApiComment(bzComment)
	return &ApiServer{
		BzIntrospect:  bzAuth,
		ApiAuth:       apiAuth,
		ApiUser:       apiUser,
		ApiCategories: apiCategories,
		ApiPost:       apiPost,
		ApiComment:    apiComment,
		ApiTag:        apiTag,
	}
}

func Connectdb() (*sql.DB, error) {
	db, err := sql.Open("sqlserver", "sqlserver://sa:123@192.168.219.1:1503?database=bloghomnay&encrypt=disable&connection+timeout=30")
	if err != nil {
		return nil, err
	}
	return db, nil
}
