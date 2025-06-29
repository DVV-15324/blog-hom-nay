package comment

import (
	"bloghomnay-project/common"
	entityImg "bloghomnay-project/services/entity/img"
	"context"
)

type BusinessImg interface {
	BusinessCreateImg(ctx context.Context, data *entityImg.CreateImg) *common.AppError
	BusinessGetImg(ctx context.Context, userID int) ([]entityImg.Img, *common.AppError)
}
type ApiImg struct {
	bz BusinessImg
}

func NewApiImg(bz BusinessImg) *ApiImg {
	return &ApiImg{
		bz: bz,
	}
}
