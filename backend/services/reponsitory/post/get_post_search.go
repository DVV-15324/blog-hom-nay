package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (t *PostServiceSQL) SearchsPost(ctx context.Context, tagsName []string, title string) ([]entity.Posts, error) {
	if len(tagsName) == 0 {
		return []entity.Posts{}, nil
	}

	args := []interface{}{}
	placeholders := []string{}
	for _, name := range tagsName {
		paramName := fmt.Sprintf("p%s", name)
		placeholders = append(placeholders, "@"+paramName)
		args = append(args, sql.Named(paramName, name))
	}

	// Add title as named parameter safely
	args = append(args, sql.Named("title", "%"+title+"%"))

	query := fmt.Sprintf(`
		SELECT p.id, p.user_id, p.category_id, p.description, p.title, p.content, COUNT(DISTINCT pl.id) AS [like], COUNT(DISTINCT c.id) AS [count_comment], p.created_at,
		p.updated_at
		FROM posts AS p
		JOIN post_tags AS pt ON pt.post_id = p.id
		JOIN tags AS t ON pt.tag_id = t.id
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.content LIKE @title AND t.name IN (%s)
		GROUP BY p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.created_at, p.updated_at
		HAVING COUNT(DISTINCT pt.tag_id) = %d;
	`, strings.Join(placeholders, ", "), len(tagsName))

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment, &data.CreatedAt, &data.UpdatedAt); err != nil {
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	return listPosts, nil
}
