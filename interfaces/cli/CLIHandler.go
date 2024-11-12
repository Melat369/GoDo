package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Melat369/GoDo/services"
)

type CLIHandler struct {
	service services.GroceryService
}

func NewCLIHandler(g services.GroceryService) *CLIHandler {
	return &CLIHandler{service: g}
}

// Start takes io.Reader and io.Writer to facilitate testing.
func (h *CLIHandler) Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)

	for {
		fmt.Fprintln(writer, "\033[34m\nChoose an action: 'a' to add, 'c' to complete, 'd' to delete, 'l' to list, or 'e' to exit.\033[0m")
		fmt.Fprint(writer, "-> ")
		writer.Flush()

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "a":
			h.AddGrocery(scanner, writer)
		case "c":
			h.completeTask(scanner, writer)
		case "d":
			h.DeleteGrocery(scanner, writer)
		case "l":
			h.ListGrocery(writer)
		case "e":
			fmt.Fprintln(writer, "Exiting...")
			writer.Flush()
			return
		default:
			fmt.Fprintln(writer, "Unknown command. Please choose add, complete, delete, list, or exit.")
			writer.Flush()
		}
	}
}

func (h *CLIHandler) AddGrocery(scanner *bufio.Scanner, writer io.Writer) {
	fmt.Fprint(writer, "Enter grocery title: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	title := strings.TrimSpace(scanner.Text())
	grocery, err := h.service.AddGrocery(title)
	if err != nil {
		fmt.Fprintln(writer, "Error adding grocery:", err)
	} else {
		fmt.Fprintf(writer, "Grocery added with ID %d\n", grocery.ID)
	}
	writer.(*bufio.Writer).Flush()
}
func (h *CLIHandler) completeTask(scanner *bufio.Scanner, writer io.Writer) {
	fmt.Fprint(writer, "Enter task ID to complete: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(writer, "Invalid ID format")
		return
	}
	if err := h.service.CompleteGrocery(id); err != nil {
		fmt.Fprintln(writer, "Error completing grocery:", err)
	} else {
		fmt.Fprintln(writer, "Grocery marked as complete")
	}
	writer.(*bufio.Writer).Flush()
}
func (h *CLIHandler) DeleteGrocery(scanner *bufio.Scanner, writer io.Writer) {
	fmt.Fprint(writer, "Enter Grocery ID to delete: ")
	writer.(*bufio.Writer).Flush()

	if !scanner.Scan() {
		fmt.Fprintln(writer, "Error reading input.")
		return
	}

	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(writer, "Invalid ID format")
		return
	}
	if err := h.service.DeleteGrocery(id); err != nil {
		fmt.Fprintln(writer, "Error deleting grocery:", err)
	} else {
		fmt.Fprintln(writer, "Grocery deleted")
	}
	writer.(*bufio.Writer).Flush()
}
func (h *CLIHandler) ListGrocery(writer io.Writer) {
	groceries := h.service.ListGrocery()
	if len(groceries) == 0 {
		fmt.Fprintln(writer, "No groceries found.")
		writer.(*bufio.Writer).Flush()
		return
	}

	fmt.Fprintln(writer, "Groceries:")
	for _, grocery := range groceries {
		status := "Pending"
		if grocery.IsDone {
			status = "Completed"
		} else if grocery.Deleted {
			fmt.Fprintf(writer, "\033[31m\033[9mID: %d, Title: %s, Status: %s\033[0m\n", grocery.ID, grocery.Title, status)
		} else {
			fmt.Fprintf(writer, "\033[32mID: %d, Title: %s, Status: %s\033[0m\n", grocery.ID, grocery.Title, status)
		}
	}
	writer.(*bufio.Writer).Flush()
}
