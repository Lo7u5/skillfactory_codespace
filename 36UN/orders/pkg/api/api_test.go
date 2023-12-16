package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"skillfactory_codespace/36UN/orders/pkg/db"
	"testing"
)

func TestAPI_ordersHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	b, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	var data []db.Order
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	const wantLen = 1
	if len(data) != wantLen {
		t.Fatalf("получено %d записей, ожидалось %d", len(data), wantLen)
	}
}

func TestAPI_newOrderHandler(t *testing.T) {
	dbase := db.New()
	api := New(dbase)
	order := db.Order{}
	orderJson, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("не удалось сериализовать объект заказа: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(orderJson))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	b, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	var id int
	err = json.Unmarshal(b, &id)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	if id != 1{
		t.Fatalf("получен %d id, ожидалось %d", id, 2)
	}
}

func TestAPI_updateOrderHandler(t *testing.T) {
	dbase := db.New()
	api := New(dbase)
	dbase.NewOrder(db.Order{})
	var order db.Order = db.Order{IsOpen: true}
	orderJson, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("не удалось сериализовать объект заказа: %v", err)
	}
	req := httptest.NewRequest(http.MethodPatch, "/orders/1", bytes.NewBuffer(orderJson))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_deleteOrderHandler(t *testing.T) {
	dbase := db.New()
	api := New(dbase)
	dbase.NewOrder(db.Order{})
	req := httptest.NewRequest(http.MethodDelete, "/orders/1", nil)
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}