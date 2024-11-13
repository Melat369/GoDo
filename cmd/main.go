package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Melat369/GoDo/config"
	"github.com/Melat369/GoDo/interfaces/cli"
	"github.com/Melat369/GoDo/services/impl"
	"github.com/jackc/pgx/v4"
)

func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Using DB connection string: %s\n", cfg.DBConnectionString)
	fmt.Println("Using API Key:", cfg.ApiKey)

	conn, err := pgx.Connect(context.Background(), cfg.DBConnectionString)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	taskUseCase := impl.NewGroceryService(conn)
	cliHandler := cli.NewCLIHandler(taskUseCase)
	cliHandler.Start(os.Stdin, os.Stdout)
}
