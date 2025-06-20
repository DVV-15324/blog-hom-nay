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
<<<<<<< HEAD
	if posts.Description != nil {
		placeholders = append(placeholders, "description=@description")
		args = append(args, sql.Named("description", posts.Description))
	}
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	query := fmt.Sprintf("UPDATE posts SET %s Where id=@id", strings.Join(placeholders, ","))

	args = append(args, sql.Named("id", id))

	_, err := p.db.ExecContext(cxt, query, args...)
	if err != nil {
		return err
	}
	return nil
}
