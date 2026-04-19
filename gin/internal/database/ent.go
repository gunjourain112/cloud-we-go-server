package database

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/ent"
)

func NewEntClient(db *sql.DB) *ent.Client {
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}
