package handlers

import (
	tmpl "arnesteen.de/writing-wiki/templates"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	article, _ := app.Db.GetArticleByName(name)
	content := tmpl.TArticle(article, article.CurrentText)
	tmpl.TLayout("Here: "+article.Title, content).Render(r.Context(), w)
}

func (app *Application) GetArticleList(w http.ResponseWriter, r *http.Request) {
	articles := app.Db.GetArticleList()
	content := tmpl.TArticlesList(articles)
	tmpl.TLayout("Title to come", content).Render(r.Context(), w)
}
