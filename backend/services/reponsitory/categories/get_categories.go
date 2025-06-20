package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
)

func (c *CategoriesServiceSQL) GetALLCategories(ctx context.Context) ([]entity.Categories, error) {
	query := "SELECT id, name, img, description FROM categories"
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listCategories []entity.Categories
	for rows.Next() {
		var data entity.Categories
		if err := rows.Scan(&data.Id, &data.Name, &data.Img, &data.Description); err != nil {
			return nil, err
		}
		listCategories = append(listCategories, data)
	}

	// Check lỗi sau khi duyệt xong (ví dụ: lỗi mạng trong lúc đọc)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listCategories, nil
}
