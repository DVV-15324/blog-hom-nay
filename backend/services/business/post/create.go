package post

import (
	"bloghomnay-project/common"
	entityPost "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (p *BusinessPost) BusinessCreatePost(ctx context.Context, data *entityPost.CreatePost) *common.AppError {
	err := data.Validate()
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	uid, er := p.bzPost.CreatePost(ctx, data)
	if er != nil {
		app := common.NewAppError(500, http.StatusText(500), er)
		return app
	}
	for i := 0; i < len(data.Tag); i++ {
		er := p.bzPostTag.CreatePostTag(ctx, uid, data.Tag[i].Id)
		if er != nil {
			app := common.NewAppError(500, http.StatusText(500), er)
			return app
		}
	}
	return nil
}
