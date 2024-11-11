package impl

import (
	"testing"
)

func TestAddGrocery(t *testing.T) {
	service := NewGroceryService()

	task, err := service.AddGrocery("Broccoli")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected task ID 1, got %d", task.ID)
	}

	if task.Title != "Broccoli" {
		t.Errorf("expected title 'Broccoli', got %s", task.Title)
	}

	if task.IsDone {
		t.Errorf("expected task to be not done, got done")
	}
}

func TestListGroceriesTasks(t *testing.T) {
	useCase := NewGroceryService()
	useCase.AddGrocery("Yogurt")
	useCase.AddGrocery("Cream Cheese")
	useCase.AddGrocery("Cookies")
	useCase.AddGrocery("Coffee")
	useCase.AddGrocery("Frozen Fish")

	groceries := useCase.ListGrocery()
	if len(groceries) != 2 {
		t.Errorf("expected 5 Groceries, got %d", len(groceries))
	}

	if groceries[0].Title != "Grocery 1" || groceries[1].Title != "Grocery 2" {
		t.Error("expected groceries to match added titles")
	}
}
