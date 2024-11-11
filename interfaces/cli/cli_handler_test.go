package cli

import (
	"errors"

	"github.com/Melat369/GoDo/entities"
)

type MockGroceryService struct {
	groceries []entities.Grocery
	nextID    int
}

func (m *MockGroceryService) AddGrocery(title string) (entities.Grocery, error) {
	task := entities.Grocery{ID: m.nextID, Title: title, IsDone: false}
	m.groceries = append(m.groceries, task)
	m.nextID++
	return task, nil
}
func (m *MockGroceryService) DeleteGrocery(id int) error {
	for i, grocery := range m.groceries {
		if grocery.ID == id {
			m.groceries = append(m.groceries[:i], m.groceries[i+1:]...)
			return nil
		}
	}
	return errors.New("grocery not found")
}
func (m *MockGroceryService) CompleteGrocery(id int) error {
	for i, grocery := range m.groceries {
		if grocery.ID == id {
			m.groceries[i].IsDone = true
			return nil
		}
	}
	return errors.New("grocery not found")
}
func (m *MockGroceryService) ListGrocery() []entities.Grocery {
	return m.groceries
}
