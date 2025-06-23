package img

import (
	entity "bloghomnay-project/services/entity/img"
	"context"
)

type ResponsitoryImg interface {
	CreateImg(ctx context.Context, linkImg string, userId int) error
	GetImgByUserId(ctx context.Context, userID int) ([]entity.Img, error)
}

type BusinessImg struct {
	bzImg ResponsitoryImg
}

func NewBusinessImg(bzImg ResponsitoryImg) *BusinessImg {
	return &BusinessImg{

		bzImg: bzImg,
	}
}
