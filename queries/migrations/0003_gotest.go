package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up0003, Down0003)
}

func Up0003(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO main.articles (name, content) VALUES ('autoadd', 'LOLOLOL');")
	if err != nil {
		return err
	}
	return nil
}

func Down0003(_ context.Context, _ *sql.Tx) error {
	return nil
}
