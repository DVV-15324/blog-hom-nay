package post

import (
	"bloghomnay-project/common"
	"context"
	"net/http"
)

func (p *BusinessPost) BusinessDeletePost(ctx context.Context, id int) *common.AppError {
	er := p.bzPost.DeletePostById(ctx, id)
	if er != nil {
		app := common.NewAppError(500, http.StatusText(500), er)
		return app
	}

	return nil
}
