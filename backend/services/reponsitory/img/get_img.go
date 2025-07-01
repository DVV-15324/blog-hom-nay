package img

import (
	entity "bloghomnay-project/services/entity/img"
	"context"
	"database/sql"
)

// Get nhiều comment của 1 post
func (c *ImgServiceSQL) GetImgByUserId(ctx context.Context, userID int) ([]entity.Img, error) {
	query := "SELECT id, user_id, img FROM img where user_id = @user_id"
	rows, err := c.db.QueryContext(ctx, query, sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listComment []entity.Img
	for rows.Next() {
		var data entity.Img
		if err := rows.Scan(&data.Id, &data.UserId, &data.Img); err != nil {
			return nil, err
		}
		listComment = append(listComment, data)
	}

	// Check lỗi sau khi duyệt xong (ví dụ: lỗi mạng trong lúc đọc)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listComment, nil
}
