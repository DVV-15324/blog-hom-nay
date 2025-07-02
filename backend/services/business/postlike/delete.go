package postlike

import (
	"bloghomnay-project/common"

	"context"
	"net/http"
)

func (p *BusinessPostLike) BusinessDeletePostLike(ctx context.Context, userId int, postId int) *common.AppError {

	er := p.bzPostLike.DeletePostLikeByPost(ctx, userId, postId)
	if er != nil {
		app := common.NewAppError(500, http.StatusText(500), er)
		return app
	}
	return nil
}
