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
<<<<<<< HEAD

=======
	err := p.bzPostTag.DeletePostTagByPost(ctx, id)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
	return nil
}
