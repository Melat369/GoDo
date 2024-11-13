package impl

import (
	"context"
	"fmt"

	"github.com/Melat369/GoDo/entities"
	"github.com/Melat369/GoDo/services"
	"github.com/jackc/pgx/v4"
)

type GroceryServiceImpl struct {
	db *pgx.Conn
}

func NewGroceryService(db *pgx.Conn) services.GroceryService {
	return &GroceryServiceImpl{
		db: db,
	}
}

// AddGrocery adds a grocery item to the database
func (s *GroceryServiceImpl) AddGrocery(title string) (entities.Grocery, error) {
	// Insert the grocery into the database
	var id int
	err := s.db.QueryRow(context.Background(), "INSERT INTO groceries(title) VALUES($1) RETURNING id", title).Scan(&id)
	if err != nil {
		return entities.Grocery{}, fmt.Errorf("error inserting grocery: %v", err)
	}

	grocery := entities.Grocery{
		ID:     id,
		Title:  title,
		IsDone: false,
	}

	return grocery, nil
}

// CompleteGrocery marks a grocery item as completed
func (s *GroceryServiceImpl) CompleteGrocery(id int) error {
	// Update the grocery item to mark it as completed
	_, err := s.db.Exec(context.Background(), "UPDATE groceries SET is_done = TRUE WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error completing grocery: %v", err)
	}
	return nil
}

// DeleteGrocery marks a grocery item as deleted
func (s *GroceryServiceImpl) DeleteGrocery(id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE groceries WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting grocery: %v", err)
	}
	return nil
}

// ListGrocery retrieves all grocery items from the database
func (s *GroceryServiceImpl) ListGrocery() ([]entities.Grocery, error) {
	rows, err := s.db.Query(context.Background(), "SELECT * FROM groceries")
	if err != nil {
		return nil, fmt.Errorf("error listing groceries: %v", err)
	}
	defer rows.Close()

	var groceries []entities.Grocery
	for rows.Next() {
		var g entities.Grocery
		if err := rows.Scan(&g.ID, &g.Title, &g.IsDone, &g.Deleted); err != nil {
			return nil, fmt.Errorf("error scanning grocery: %v", err)
		}
		groceries = append(groceries, g)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return groceries, nil
}
