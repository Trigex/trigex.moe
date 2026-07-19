package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/trigex/trigex.moe/internal/app"
	"github.com/trigex/trigex.moe/internal/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	e, err := app.New(cfg.Admin.Username, cfg.Admin.Password, cfg.Database.Path, cfg.Paths.DataDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize app: %v\n", err)
		os.Exit(1)
	}

	portString := ":" + strconv.Itoa(cfg.Server.Port)

	e.Logger.Info("Starting server on http://localhost" + portString + " ...")
	e.Logger.Info("Access static assets at http://localhost" + portString + "/static/")

	if err := e.Start(portString); err != nil {
		e.Logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
