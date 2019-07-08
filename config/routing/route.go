package routing

import (
	httpHandlers "github.com/nandaryanizar/golang-webservice-example/handlers/http"

	"github.com/gorilla/mux"
)

// NewRouter factory
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	initRoute(r)

	return r
}

func initRoute(r *mux.Router) {
	r.HandleFunc("/token", httpHandlers.Token).Methods("POST")
	r.HandleFunc("/api/v1/user/{id:[0-9]+}", httpHandlers.FindUser).Methods("GET")
}
