package model

import (
	"database/sql"
	q "github.com/doug-martin/goqu/v9"
	"log"
)

func (db *DB) GetArticleList() []Article {
	c := make(chan []Article)
	db.WorkChannel <- func(sqldb *sql.DB, gqdb *q.Database) {
		var articles []Article
		err := gqdb.Select("id", "title", "created_at", "updated_at").
			From("articles").
			ScanStructs(&articles)
		if err != nil {
			log.Printf("Error fetching data")
			c <- []Article{}
		}
		c <- articles
	}
	return <-c
}

func (db *DB) GetArticleByName(name string) (*Article, error) {
	c := make(chan *Article)
	db.WorkChannel <- func(sqldb *sql.DB, gqdb *q.Database) {
		var article Article
		query := gqdb.Select(
			"a.id", "a.title", "a.created_at", "a.updated_at",
			q.I("at.id").As(q.C("current_text.id")),
			q.I("at.article_id").As(q.C("current_text.article_id")),
			q.I("at.content").As(q.C("current_text.content")),
			q.I("at.version").As(q.C("current_text.version")),
			q.I("at.created_at").As(q.C("current_text.created_at")),
			q.I("at.updated_at").As(q.C("current_text.updated_at")),
		).
			From(q.T("articles").As("a")).
			Join(q.T("article_texts").As("at"),
				q.On(q.I("a.id").Eq(q.I("at.article_id")))).
			Where(q.Ex{"a.title": name}).
			Order(q.I("at.version").Desc()).
			Limit(1)

		exists, err := query.ScanStruct(&article)
		switch {
		case err != nil:
			log.Printf("Error fetching data")
			c <- nil
		case !exists:
			c <- nil
		default:
			c <- &article
		}
	}
	result := <-c
	if c == nil {
		return nil, nil
	} else {
		return result, nil
	}
}
