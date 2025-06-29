package post

import (
	"context"
	"database/sql"
)

func (p *PostServiceSQL) DeletePostById(cxt context.Context, id int) error {
<<<<<<< HEAD
	query := "DELETE FROM posts WHERE id = @id"
	_, err := p.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {

=======
	query := "DELETE posts Where id = @id"
	_, err := p.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
		return err
	}
	return nil
}
