package handlers

import (
	"arnesteen.de/writing-wiki/config"
	"arnesteen.de/writing-wiki/queries"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Application struct {
	Cfg *config.Config
	Db  *queries.DB
}

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}
