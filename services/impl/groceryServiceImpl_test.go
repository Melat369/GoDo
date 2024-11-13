package impl

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

// Set up PostgreSQL connection for tests (replace with your DB credentials)
func getTestDBConnection(t *testing.T) *pgx.Conn {

	connStr := "postgres://postgres:postgres@localhost:5432/grocery-store-db-test?sslmode=disable"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Unable to connect to the database: %v", err)
	}
	return conn
}

// Clear test data in the grocery table before running each test
func clearTestData(t *testing.T, db *pgx.Conn) {
	_, err := db.Exec(context.Background(), "DELETE FROM groceries")
	if err != nil {
		t.Fatalf("Failed to clear test data: %v", err)
	}
	// Reset the sequence of the ID column to 1 (or the appropriate starting value)
	_, err = db.Exec(context.Background(), `SELECT setval('groceries_id_seq', 1, false)`)
	t.Log("The sequence ID is set to 1")
}

func TestAddGrocery(t *testing.T) {
	// Get DB connection for the test
	conn := getTestDBConnection(t)
	defer conn.Close(context.Background())

	// Create grocery service
	service := NewGroceryService(conn)

	// Clear previous data
	clearTestData(t, conn)

	// Add a grocery item
	task, err := service.AddGrocery("Broccoli")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check if the ID is 1 (assuming this is the first entry)
	if task.ID != 1 {
		t.Errorf("expected task ID 1, got %d", task.ID)
	}

	// Check the title of the grocery item
	if task.Title != "Broccoli" {
		t.Errorf("expected title 'Broccoli', got %s", task.Title)
	}

	// Ensure the grocery item is not marked as done
	if task.IsDone {
		t.Errorf("expected task to be not done, got done")
	}
}

func TestListGroceriesTasks(t *testing.T) {
	// Get DB connection for the test
	conn := getTestDBConnection(t)
	defer conn.Close(context.Background())

	// Create grocery service
	service := NewGroceryService(conn)

	// Clear previous data
	clearTestData(t, conn)

	// Add groceries to the database
	service.AddGrocery("Yogurt")
	service.AddGrocery("Cream Cheese")
	service.AddGrocery("Cookies")
	service.AddGrocery("Coffee")
	service.AddGrocery("Frozen Fish")

	// List all groceries
	groceries, err := service.ListGrocery()
	t.Log("All groceries for test", groceries)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Ensure we added 5 items, not 2
	t.Log("Total groceries for test", len(groceries))
	if len(groceries) != 5 {
		t.Errorf("expected 5 groceries, got %d", len(groceries))
	}

	// Check that the grocery titles match what was added
	if groceries[0].Title != "Yogurt" || groceries[1].Title != "Cream Cheese" ||
		groceries[2].Title != "Cookies" || groceries[3].Title != "Coffee" ||
		groceries[4].Title != "Frozen Fish" {
		t.Errorf("expected grocery titles to match added values")
	}
}
