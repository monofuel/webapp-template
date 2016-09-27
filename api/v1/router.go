package v1

import (
	"github.com/gorilla/mux"
	"github.com/monofuel/webapp-template/api/v1/test"
)

func AddRoutes(r *mux.Router) {
	test.AddRoutes(r.PathPrefix("/test").Subrouter())
}
