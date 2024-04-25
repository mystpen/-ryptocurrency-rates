package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Service interface {
}

func (h *Handler) Routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", nil)

	return router
}
