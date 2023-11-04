package main

import (
	"arnesteen.de/writing-wiki/config"
	"arnesteen.de/writing-wiki/handlers"
	"arnesteen.de/writing-wiki/queries"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	app := handlers.Application{
		Cfg: cfg,
		Db:  queries.NewDB(cfg),
	}

	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/articles/:name", app.Middleware(app.GetArticle))
	router.GET("/articles", app.Middleware(app.GetArticleList))

	log.Fatal(http.ListenAndServe(cfg.GetHostWithPort(), router))
}
