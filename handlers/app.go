package handlers

import (
	"arnesteen.de/writing-wiki/config"
	"arnesteen.de/writing-wiki/model"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type Application struct {
	Cfg *config.Config
	Db  *model.DB
}

func Index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
