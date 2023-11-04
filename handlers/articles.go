package handlers

import (
	tmpl "arnesteen.de/writing-wiki/templates"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	article := app.Db.GetArticleByName(r.Context(), name)
	content := tmpl.TArticle(article)
	tmpl.TLayout("Here: "+article.Name, content).Render(r.Context(), w)
}

func (app *Application) GetArticleList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	articles := app.Db.GetArticleList(r.Context())
	content := tmpl.TArticlesList(articles)
	tmpl.TLayout("Title to come", content).Render(r.Context(), w)
}
