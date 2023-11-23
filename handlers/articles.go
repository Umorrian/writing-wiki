package handlers

import (
	tmpl "arnesteen.de/writing-wiki/templates"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	awc := app.Db.GetArticleByName(r.Context(), name)
	content := tmpl.TArticle(&awc.Article, &awc.ArticleText)
	tmpl.TLayout("Here: "+awc.Article.Title, content).Render(r.Context(), w)
}

func (app *Application) GetArticleList(w http.ResponseWriter, r *http.Request) {
	articles := app.Db.GetArticleList(r.Context())
	content := tmpl.TArticlesList(articles)
	tmpl.TLayout("Title to come", content).Render(r.Context(), w)
}
