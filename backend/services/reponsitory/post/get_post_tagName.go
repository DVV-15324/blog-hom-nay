package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (t *PostServiceSQL) GetsPostByTagsName(ctx context.Context, tagName []string) ([]entity.Posts, error) {
	if len(tagName) == 0 {
		return []entity.Posts{}, nil
	}

	args := []interface{}{}
	placeholders := []string{}
	for _, name := range tagName {
		paramName := fmt.Sprintf("p%s", name)
		placeholders = append(placeholders, "@"+paramName)
		args = append(args, sql.Named(paramName, name))
	}

	// Dùng GROUP BY + HAVING COUNT(DISTINCT tag_id) = số lượng tag cần tìm
	query := fmt.Sprintf(`
		SELECT p.id, p.user_id, p.category_id, p.description, p.title, p.content, COUNT(DISTINCT pl.id) AS [like], COUNT(DISTINCT c.id) AS [count_comment]
		FROM posts AS p
		JOIN post_tags AS pt ON pt.post_id = p.id
		JOIN tags AS t ON pt.tag_id = t.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		WHERE t.name IN (%s)
		GROUP BY p.id, p.user_id, p.category_id, p.description, p.title, p.content
		HAVING COUNT(DISTINCT pt.tag_id) = %d
	`, strings.Join(placeholders, ", "), len(tagName))

	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment); err != nil {
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	return listPosts, nil
}
