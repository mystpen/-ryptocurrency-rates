package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mystpen/cryptocurrency-rates/internal/model"
	"github.com/mystpen/cryptocurrency-rates/internal/repository/api"
)

type Service interface {
	GetInfoByName(name string) (*model.CryptoCoin, error)
}

func (h *Handler) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/currency", h.getCurrencyHandler)

	return router
}

func (h *Handler) getCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	name := strings.ToLower(qs.Get("name"))
	if name == ""{
		h.errorResponse(w,r, http.StatusBadRequest, "bad request")
		return
	} 

	info, err := h.service.GetInfoByName(name)
	if err != nil{
		if errors.Is(err, api.ErrNoData){
			h.errorResponse(w,r, http.StatusNotFound, fmt.Sprintf("no data for %v", name))
			return
		} else {
			h.errorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}
	}

	err = h.writeJSON(w, http.StatusOK, envelope{name: info}, nil)
	if err != nil{
		h.errorResponse(w, r, http.StatusInternalServerError, err.Error())
	}
}