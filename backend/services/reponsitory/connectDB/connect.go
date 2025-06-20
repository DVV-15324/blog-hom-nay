package connectdb

import (
	"database/sql"
)

func Connectdb() (*sql.DB, error) {
	db, err := sql.Open("sqlserver", "sqlserver://sa:123@localhost:1503?database=bloghomnay&connection+timeout=30")
	//db, err := sql.Open("sqlserver", "sqlserver://sa:YourStrong!Passw0rd@localhost:1503?database=bloghomnay&connection+timeout=30")
	if err != nil {
		return nil, err
	}
	return db, nil
}
