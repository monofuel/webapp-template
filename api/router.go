package api

import (
	"github.com/gorilla/mux"
	"github.com/monofuel/webapp-template/api/v1"
)

func AddRoutes(r *mux.Router) {
	v1.AddRoutes(r.PathPrefix("/v1").Subrouter())
}
