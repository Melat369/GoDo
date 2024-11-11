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
		fmt.Fprintln(writer, "\nChoose an action: add, complete, delete, list, or exit")
		fmt.Fprint(writer, "-> ")
		writer.Flush()

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "add":
			h.AddGrocery(scanner, writer)
		case "complete":
			h.completeTask(scanner, writer)
		case "delete":
			h.DeleteGrocery(scanner, writer)
		case "list":
			h.ListGrocery(writer)
		case "exit":
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
		}
		fmt.Fprintf(writer, "ID: %d, Title: %s, Status: %s\n", grocery.ID, grocery.Title, status)
	}
	writer.(*bufio.Writer).Flush()
}
