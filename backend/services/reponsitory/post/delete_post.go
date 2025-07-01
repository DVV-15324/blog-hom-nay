package post

import (
	"context"
	"database/sql"
)

func (p *PostServiceSQL) DeletePostById(cxt context.Context, id int) error {
	query := "DELETE FROM posts WHERE id = @id"
	_, err := p.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {

		return err
	}
	return nil
}
