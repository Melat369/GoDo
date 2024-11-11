package main

import (
	"os"

	"github.com/Melat369/GoDo/interfaces/cli"
	"github.com/Melat369/GoDo/services/impl"
)

func main() {
	taskUseCase := impl.NewGroceryService()
	cliHandler := cli.NewCLIHandler(taskUseCase)
	cliHandler.Start(os.Stdin, os.Stdout)
}
