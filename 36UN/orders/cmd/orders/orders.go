package main

import (
	"log"
	"net/http"
	"skillfactory_codespace/36UN/orders/pkg/api"
	"skillfactory_codespace/36UN/orders/pkg/db"
)

func main() {
	db1 := db.New()
	var o db.Order = db.Order{
		ID:              23,
		IsOpen:          true,
		DeliveryTime:    2365,
		DeliveryAddress: "address",
		Products: []db.Product{
			{ID: 1,
				Name:  "product",
				Price: 23.4},
		},
	}
	db1.NewOrder(o)
	api := api.New(db1)
	err := http.ListenAndServe(":80", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
