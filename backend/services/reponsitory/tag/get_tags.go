package tag

import (
	entity "bloghomnay-project/services/entity/tag"
	"context"
)

func (t *TagServiceSQL) GetALLTags(ctx context.Context) ([]entity.Tag, error) {
	query := "SELECT id, name FROM tags"
	rows, err := t.db.QueryContext(ctx, query)
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
