package queries

import (
	"arnesteen.de/writing-wiki/config"
	sqlg "arnesteen.de/writing-wiki/queries/gen"
	_ "arnesteen.de/writing-wiki/queries/migrations"
	"database/sql"
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"os"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type sqlWork func(queries *sqlg.Queries)

type DB struct {
	WorkChannel chan sqlWork
	Cfg         *config.Config
}

func NewDB(cfg *config.Config) *DB {
	db := &DB{
		WorkChannel: make(chan sqlWork, 128),
		Cfg:         cfg,
	}

	go db.sqlWarden()

	return db
}

func (db *DB) sqlWarden() {
	f, err := os.OpenFile(db.Cfg.VolumePath+"/db.sqlite", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	if f.Close() != nil {
		panic(err)
	}

	conn, err := sql.Open("sqlite3", "file:"+db.Cfg.VolumePath+"/db.sqlite")
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(conn, "migrations"); err != nil {
		panic(err)
	}

	queries := sqlg.New(conn)

	for {
		work := <-db.WorkChannel
		work(queries)
	}
}
