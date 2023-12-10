package db

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		{name: "first test",
			want: &DB{
				id:    1,
				store: map[int]Order{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Orders(t *testing.T) {
	tests := []struct {
		name string
		db   *DB
		want []Order
	}{
		{
			name: "first test",
			db:   New(),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.Orders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Orders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewOrder(t *testing.T) {
	type args struct {
		o Order
	}
	tests := []struct {
		name string
		db   *DB
		args args
		want int
	}{
		{
			name: "first test",
			db:   New(),
			args: args{},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.NewOrder(tt.args.o); got != tt.want {
				t.Errorf("DB.NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateOrder(t *testing.T) {
	type args struct {
		o Order
	}
	tests := []struct {
		name string
		db   *DB
		args args
	}{
		{
			name: "first test",
			db:   New(),
			args: args{o: Order{ID: 1, IsOpen: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.db.UpdateOrder(tt.args.o)
		})
		if !tt.db.store[1].IsOpen {
			t.Errorf("%v", tt.db.store[1].IsOpen)
		}
	}
}

func TestDB_DeleteOrder(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		db   *DB
		args args
	}{
		{
			name: "first test",
			db: &DB{
				id: 1,
				store: map[int]Order{
					1: {},
					2: {},
				},
			},
			args: args{id: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.db.DeleteOrder(tt.args.id)
		})
		if _, ok := tt.db.store[tt.args.id]; ok {
			t.Errorf("%v", tt.db)
		}
	}
}
