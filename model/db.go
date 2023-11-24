package model

import (
	"arnesteen.de/writing-wiki/config"
	_ "arnesteen.de/writing-wiki/model/migrations"
	"database/sql"
	"embed"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"os"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type sqlWork func(sqldb *sql.DB, gqdb *goqu.Database)

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

	sqldb, err := sql.Open("sqlite3", "file:"+db.Cfg.VolumePath+"/db.sqlite")
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(sqldb, "migrations"); err != nil {
		panic(err)
	}

	sqlDialect := goqu.Dialect("sqlite3")
	gqdb := sqlDialect.DB(sqldb)

	for {
		work := <-db.WorkChannel
		work(sqldb, gqdb)
	}
}
