package route

import (
	"github.com/gorilla/mux"
)

func Router(r *mux.Router) *mux.Router {
	OnBoardRouter(r)
	TransactionRouter(r)
	return r
}
