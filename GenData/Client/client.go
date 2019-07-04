package Client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewClient(db string) (client *sql.DB,err error) {
	client, err= sql.Open("mysql", "root:123456@/"+db)
	return
}
