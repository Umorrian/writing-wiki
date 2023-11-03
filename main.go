package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"

	"arnesteen.de/writing-wiki/sqlite_gen"
)

type SqlWork func(ctx context.Context, queries *sqlite_gen.Queries)

var globalChannel = make(chan SqlWork, 32)

func sqlWarden() {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "file:/home/arne/Projects/writing-wiki/test1.sqlite")
	if err != nil {
		log.Printf("Error opening sqlite db %v", err)
		return
	}
	queries := sqlite_gen.New(db)

	for {
		work := <-globalChannel
		work(ctx, queries)
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func greet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	component := hello(ps.ByName("name"))
	component.Render(r.Context(), w)
}

func fromdb(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dbId, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		log.Printf("Error converting id")
		return
	}

	result := make(chan sqlite_gen.Author)
	globalChannel <- func(ctx context.Context, queries *sqlite_gen.Queries) {
		fetchedAuthor, err := queries.GetAuthor(ctx, dbId)
		if err != nil {
			log.Printf("Error fetching data")
			return
		}
		result <- fetchedAuthor
	}
	author := <-result

	fmt.Fprint(w, author.Bio.String)
}

func middleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		// call registered handler
		n(w, r, ps)
	}
}

func main() {
	go sqlWarden()

	router := httprouter.New()
	router.GET("/", middleware(index))
	router.GET("/hello/:name", middleware(greet))
	router.GET("/db/:id", middleware(fromdb))

	log.Fatal(http.ListenAndServe(":8080", router))
}
