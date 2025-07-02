package tag

import (
	entity "bloghomnay-project/services/entity/tag"
	"context"
	"database/sql"
)

func (t *TagServiceSQL) GetTagByPostId(ctx context.Context, postId int) ([]entity.Tag, error) {
	query := "SELECT t.id, t.name FROM tags as t join post_tags as pt on pt.tag_id=t.id Where post_id=@post_id"
	rows, err := t.db.QueryContext(ctx, query, sql.Named("post_id", postId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listTag []entity.Tag
	for rows.Next() {
		var data entity.Tag
		if err := rows.Scan(&data.Id, &data.Name); err != nil {
			return nil, err
		}
		listTag = append(listTag, data)
	}

	// Check lỗi sau khi duyệt xong (ví dụ: lỗi mạng trong lúc đọc)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listTag, nil
}
