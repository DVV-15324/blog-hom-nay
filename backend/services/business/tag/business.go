package tag

import (
	entityTag "bloghomnay-project/services/entity/tag"
	"context"
)

type ResponsitoryTag interface {
	CreateTag(ctx context.Context, data *entityTag.CreateTagForm) error
	DeleteTagById(cxt context.Context, id int) error
	GetTagByPostId(ctx context.Context, postId int) ([]entityTag.Tag, error)
	GetALLTags(ctx context.Context) ([]entityTag.Tag, error)
	UpdateTag(cxt context.Context, tag *entityTag.UpdateTagForm, id int) error
}
type ResponsitoryPostTag interface {
	DeletedPostTagByTag(ctx context.Context, tagIdId int) error
}
type BusinessTag struct {
	bzTag     ResponsitoryTag
	bzPosttag ResponsitoryPostTag
}

func NewBusinessTag(bzTag ResponsitoryTag, bzPosttag ResponsitoryPostTag) *BusinessTag {
	return &BusinessTag{
		bzTag:     bzTag,
		bzPosttag: bzPosttag,
	}
}
