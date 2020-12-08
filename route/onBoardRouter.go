package route

import (
	"encoding/json"
	"net/http"

	"ewallet/model"

	"github.com/gorilla/mux"
)

func OnBoardRouter(r *mux.Router) *mux.Router {
	r = r.PathPrefix("/onBoard").Subrouter()
	r.HandleFunc("/check_device", func(res http.ResponseWriter, req *http.Request) {
		customerModel := model.CustomerModel{}
		cModel := customerModel.Select()
		json.NewEncoder(res).Encode(cModel)
	}).Methods(http.MethodPost)

	return r
}
