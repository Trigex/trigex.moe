package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trigex/trigex.moe/internal/app"
)

const (
	Port = 8080
)

func main() {
	e, err := app.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize app: %v\n", err)
		os.Exit(1)
	}

	portString := ":" + strconv.Itoa(Port)

	e.Logger.Info("Starting server on http://localhost" + portString + " ...")
	e.Logger.Info("Access static assets at http://localhost" + portString + "/static/")

	if err := e.Start(portString); err != nil {
		e.Logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
