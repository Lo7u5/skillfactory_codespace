package db

import "sync"

type Order struct {
	ID              int
	IsOpen          bool
	DeliveryTime    int64
	DeliveryAddress string
	Products        []Product
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type DB struct {
	m     sync.Mutex
	id    int
	store map[int]Order
}

func New() *DB {
	db := DB{
		id:    1,
		store: map[int]Order{},
	}
	return &db
}

func (db *DB) Orders() []Order {
	db.m.Lock()
	defer db.m.Unlock()
	var data []Order
	for _, v := range db.store {
		data = append(data, v)
	}
	return data
}

func (db *DB) NewOrder(o Order) int {
	db.m.Lock()
	defer db.m.Unlock()
	o.ID = db.id
	db.store[o.ID] = o
	db.id++
	return o.ID
}

func (db *DB) UpdateOrder(o Order) {
	db.m.Lock()
	defer db.m.Unlock()
	db.store[o.ID] = o
}

func (db *DB) DeleteOrder(id int) {
	db.m.Lock()
	defer db.m.Unlock()
	delete(db.store, id)
}
