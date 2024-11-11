package impl

import (
	"errors"
	"fmt"

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

//	func (g *GroceryServiceImpl) DeleteGrocery(id int) error {
//		for i, grocery := range g.groceries {
//			if grocery.ID == id {
//				g.groceries = append(g.groceries[:i], g.groceries[i+1:]...)
//				return nil
//			}
//		}
//		return errors.New("grocery not found")
//	}
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
			fmt.Printf("\033[31m\033[9m%s\033[0m has been deleted.\n", grocery.Title)

			g.groceries = append(g.groceries[:i], g.groceries[i+1:]...)
			return nil
		}
	}
	return errors.New("grocery not found")
}

func (u *GroceryServiceImpl) ListGrocery() []entities.Grocery {
	return u.groceries
}

// func (u *GroceryServiceImpl) ListGrocery() {
// 	for _, grocery := range u.groceries {
// 		if grocery.Deleted {
// 			// Print deleted groceries in red with strikethrough
// 			fmt.Printf("\033[31m\033[9m%s\033[0m\n", grocery.Title) // Red, Strikethrough
// 		} else {
// 			// Print undeleted groceries in green
// 			fmt.Printf("\033[32m%s\033[0m\n", grocery.Title) // Green
// 		}
// 	}
// }
