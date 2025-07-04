package post

import (
	"bloghomnay-project/common"
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (p *PostServiceSQL) UpdatePosts(cxt context.Context, posts *entity.UpdatePost, id int) error {
	if posts == nil {
		return fmt.Errorf("post is invalid")
	}
	var (
		placeholders = []string{}
		args         = []interface{}{}
	)
	if posts.CategoryId != nil {
		placeholders = append(placeholders, "category_id=@category_id")
		args = append(args, sql.Named("category_id", posts.CategoryId))
	}
	if posts.Content != nil {
		placeholders = append(placeholders, "content=@content")
		args = append(args, sql.Named("content", posts.Content))
	}
	if posts.Title != nil {
		placeholders = append(placeholders, "title=@title")
		args = append(args, sql.Named("title", posts.Title))
	}
	if posts.Description != nil {
		placeholders = append(placeholders, "description=@description")
		args = append(args, sql.Named("description", posts.Description))
	}

	postIdFake := common.NewUID(uint32(id), 3)
	slug := fmt.Sprintf("%s-%s", common.GenerateSlug(*posts.Title), postIdFake.ToBase58())
	placeholders = append(placeholders, "slug=@slug")
	args = append(args, sql.Named("slug", slug))

	query := fmt.Sprintf("UPDATE posts SET %s, updated_at= GETDATE() Where id=@id", strings.Join(placeholders, ","))

	args = append(args, sql.Named("id", id))

	_, err := p.db.ExecContext(cxt, query, args...)
	if err != nil {
		return err
	}
	return nil
}
