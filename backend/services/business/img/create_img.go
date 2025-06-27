package img

import (
	"bloghomnay-project/common"
	entityImg "bloghomnay-project/services/entity/img"
	"context"
	"net/http"
)

func (p *BusinessImg) BusinessCreateImg(ctx context.Context, data *entityImg.CreateImg) *common.AppError {
	err := data.Validate()
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	er := p.bzImg.CreateImg(ctx, data.Img, data.UserId)
	if er != nil {
		app := common.NewAppError(500, http.StatusText(500), er)
		return app
	}
	return nil
}
