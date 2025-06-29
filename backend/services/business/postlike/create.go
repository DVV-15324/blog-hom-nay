package postlike

import (
	"bloghomnay-project/common"
	entityPostLike "bloghomnay-project/services/entity/postlikes"
	"context"
	"net/http"
)

func (p *BusinessPostLike) BusinessCreatePostLike(ctx context.Context, data *entityPostLike.CreatePostLikes) *common.AppError {

	er := p.bzPostLike.CreatePostLikes(ctx, data.UserId, data.PostId)
	if er != nil {
		app := common.NewAppError(500, http.StatusText(500), er)
		return app
	}
	return nil
}
