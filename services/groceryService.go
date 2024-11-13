package services

import "github.com/Melat369/GoDo/entities"

type GroceryService interface {
	AddGrocery(title string) (entities.Grocery, error)
	CompleteGrocery(id int) error
	DeleteGrocery(id int) error
	ListGrocery() ([]entities.Grocery, error)
}
