package impl

import (
	"errors"

	"github.com/Melat369/GoDo/entities"
	"github.com/Melat369/GoDo/services"
)

type GroceryServiceImpl struct {
	groceries []entities.Grocery
	nextID    int
}

func NewGroceryService() services.GroceryService {
	return &GroceryServiceImpl{
		groceries: []entities.Grocery{},
		nextID:    1,
	}
}

func (u *GroceryServiceImpl) AddGrocery(title string) (entities.Grocery, error) {
	grocery := entities.Grocery{ID: u.nextID, Title: title, IsDone: false}
	u.groceries = append(u.groceries, grocery)
	u.nextID++
	return grocery, nil
}

func (u *GroceryServiceImpl) CompleteGrocery(id int) error {
	for i, grocery := range u.groceries {
		if grocery.ID == id {
			u.groceries[i].IsDone = true
			return nil
		}
	}
	return errors.New("grocery not found")
}

func (g *GroceryServiceImpl) DeleteGrocery(id int) error {
	for i, grocery := range g.groceries {
		if grocery.ID == id {
			g.groceries[i].Deleted = true
			return nil
		}
	}
	return errors.New("grocery not found")
}

func (u *GroceryServiceImpl) ListGrocery() []entities.Grocery {
	return u.groceries
}
