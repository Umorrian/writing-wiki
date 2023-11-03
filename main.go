package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"log"
	"net/http"
	"os"

	_ "arnesteen.de/writing-wiki/migrations"
	sqlg "arnesteen.de/writing-wiki/sqlite_gen"
	tmpl "arnesteen.de/writing-wiki/templates"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type SqlWork func(ctx context.Context, queries *sqlg.Queries)

var globalChannel = make(chan SqlWork, 32)

func sqlWarden() {
	f, err := os.OpenFile("/home/steen/Privat/writing-wiki/bin/db.sqlite", os.O_RDONLY|os.O_CREATE, 0666)
	f.Close()

	ctx := context.Background()
	db, err := sql.Open("sqlite3", "file:/home/steen/Privat/writing-wiki/bin/db.sqlite")
	if err != nil {
		log.Printf("Error opening sqlite db %v", err)
		return
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	queries := sqlg.New(db)

	for {
		work := <-globalChannel
		work(ctx, queries)
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func get_article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")

	result := make(chan sqlg.Article)
	globalChannel <- func(ctx context.Context, queries *sqlg.Queries) {
		fetchedArticle, err := queries.GetArticleByName(ctx, name)
		if err != nil {
			log.Printf("Error fetching data")
			return
		}
		result <- fetchedArticle
	}
	article := <-result

	content := tmpl.TArticle(article)
	tmpl.TLayout("Here: "+article.Name, content).Render(r.Context(), w)
}

func get_article_list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	result := make(chan []sqlg.Article)
	globalChannel <- func(ctx context.Context, queries *sqlg.Queries) {
		articles, err := queries.GetAllArticles(ctx)
		if err != nil {
			log.Printf("Error fetching data")
			return
		}
		result <- articles
	}
	articles := <-result

	content := tmpl.TArticlesList(articles)
	tmpl.TLayout("Coolest Title", content).Render(r.Context(), w)
}

func middleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// call registered handler
		n(w, r, ps)
	}
}

func main() {
	go sqlWarden()

	router := httprouter.New()
	router.GET("/", middleware(index))
	router.GET("/articles/:name", middleware(get_article))
	router.GET("/articles", middleware(get_article_list))

	log.Fatal(http.ListenAndServe(":8080", router))
}
