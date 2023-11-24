package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up, Down)
}

func Up(ctx context.Context, tx *sql.Tx) error {
	qs := []string{
		"INSERT INTO articles (title, created_at) VALUES ('testname', DATETIME());",
		"INSERT INTO article_texts (article_id, content, version, created_at) VALUES (1, 'This is some content', 1, DATETIME());",
		"INSERT INTO article_texts (article_id, content, version, created_at) VALUES (1, 'This is some content, and now even more', 2, DATETIME());",
		"INSERT INTO articles (title, created_at) VALUES ('secondtitle', DATETIME());",
		"INSERT INTO article_texts (article_id, content, version, created_at) VALUES (2, 'ABCDE', 1, DATETIME());",
		"INSERT INTO article_texts (article_id, content, version, created_at) VALUES (2, 'AB', 2, DATETIME());",
	}
	for _, q := range qs {
		_, err := tx.ExecContext(ctx, q)
		if err != nil {
			return err
		}
	}
	return nil
}

func Down(_ context.Context, _ *sql.Tx) error {
	return nil
}
