package post

import (
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
	query := fmt.Sprintf("UPDATE posts SET %s Where id=@id", strings.Join(placeholders, ","))

	args = append(args, sql.Named("id", id))

	_, err := p.db.ExecContext(cxt, query, args...)
	if err != nil {
		return err
	}
	return nil
}
