package categories

import (
	entity "bloghomnay-project/services/entity/categories"
	"context"
)

func (c *CategoriesServiceSQL) GetALLCategories(ctx context.Context) ([]entity.Categories, error) {
<<<<<<< HEAD
	query := "SELECT id, name, img, description FROM categories"
=======
	query := "SELECT id, name, description FROM categories"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listCategories []entity.Categories
	for rows.Next() {
		var data entity.Categories
<<<<<<< HEAD
		if err := rows.Scan(&data.Id, &data.Name, &data.Img, &data.Description); err != nil {
=======
		if err := rows.Scan(&data.Id, &data.Name, &data.Description); err != nil {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
