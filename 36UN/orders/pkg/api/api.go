package api

import (
	"encoding/json"
	"net/http"
	"skillfactory_codespace/36UN/orders/pkg/db"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	r  *mux.Router
	db *db.DB
}

func New(db *db.DB) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/orders", api.ordersHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/orders", api.newOrderHandler).Methods(http.MethodPost)
	api.r.HandleFunc("/orders/{id}", api.updateOrderHandler).Methods(http.MethodPatch)
	api.r.HandleFunc("/orders/{id}", api.deleteOrderHandler).Methods(http.MethodDelete)
}

func (api *API) ordersHandler(w http.ResponseWriter, r *http.Request) {
	orders := api.db.Orders()
	json.NewEncoder(w).Encode(orders)
}

func (api *API) newOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o db.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := api.db.NewOrder(o)
	w.Write([]byte(strconv.Itoa(id)))
}

func (api *API) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	//считывание параметра {id} из пути запроса
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//декодирование в переменную тела запроса
	var o db.Order
	err = json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	o.ID = id
	api.db.UpdateOrder(o)
	w.WriteHeader(http.StatusOK)
}

func (api *API) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.db.DeleteOrder(id)
	w.WriteHeader(http.StatusOK)
}
