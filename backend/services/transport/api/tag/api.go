package tag

import (
	"bloghomnay-project/common"
	entityTag "bloghomnay-project/services/entity/tag"
	"context"
)

type BusinessTag interface {
	BusinessCreateTag(ctx context.Context, data *entityTag.CreateTagForm) *common.AppError
	BusinessDeleteTag(ctx context.Context, id int) *common.AppError
	BusinessGetTagByPostId(ctx context.Context, postID int) ([]entityTag.Tag, *common.AppError)
	BusinessGetALLTag(ctx context.Context) ([]entityTag.Tag, *common.AppError)
	BusinessUpdateTag(ctx context.Context, data *entityTag.UpdateTagForm, id int) *common.AppError
}
type ApiTag struct {
	bz BusinessTag
}

func NewApiTag(bz BusinessTag) *ApiTag {
	return &ApiTag{
		bz: bz,
	}
}
