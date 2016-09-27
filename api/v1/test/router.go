package test

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/monofuel/webapp-template/util"
)

//AddRoutes adds the test endpoints
func AddRoutes(r *mux.Router) {
	r.Methods("GET").Path("/random").Handler(util.AppHandler(randomHandler))
}

func randomHandler(w http.ResponseWriter, r *http.Request) *util.AppError {

	js, err := json.Marshal(map[string]int{"result": rand.Int()})
	if err != nil {
		return util.NewAppError(err, "failed to marshal json", 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
