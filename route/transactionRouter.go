package route

import (
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func TransactionRouter(r *mux.Router) *mux.Router {
	r = r.PathPrefix("/transaction").Subrouter()

	r.HandleFunc("/transfer", func(res http.ResponseWriter, req *http.Request) {
		json := simplejson.New()
		json.Set("transfer", "1")
		payload, _ := json.MarshalJSON()

		res.Write(payload)
	}).Methods(http.MethodPost)

	r.HandleFunc("/payment", func(res http.ResponseWriter, req *http.Request) {
		json := simplejson.New()
		json.Set("payment", "2")
		payload, _ := json.MarshalJSON()

		res.Write(payload)
	}).Methods(http.MethodPost)

	return r
}
