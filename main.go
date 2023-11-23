package main

import (
	"arnesteen.de/writing-wiki/config"
	"arnesteen.de/writing-wiki/handlers"
	"arnesteen.de/writing-wiki/queries"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	app := handlers.Application{
		Cfg: cfg,
		Db:  queries.NewDB(cfg),
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	handlers.FileServer(router, "/static", http.Dir(cfg.StaticPath))

	router.Group(func(router chi.Router) {
		router.Use(app.SetContentType)
		router.Get("/", handlers.Index)
		router.Get("/articles/{name}", app.GetArticle)
		router.Get("/articles", app.GetArticleList)
	})

	log.Fatal(http.ListenAndServe(cfg.GetHostWithPort(), router))
}
