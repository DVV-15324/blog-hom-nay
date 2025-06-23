package img

import (
	"bloghomnay-project/common"
	entityImg "bloghomnay-project/services/entity/img"
	"context"
	"net/http"
)

func (c *BusinessImg) BusinessGetImg(ctx context.Context, userID int) ([]entityImg.Img, *common.AppError) {
	img, err := c.bzImg.GetImgByUserId(ctx, userID)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return img, nil
}
