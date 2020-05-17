// Package main provides the executable for our sales-admin.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/runlevl4/garagesale/internal/platform/conf"
	"github.com/runlevl4/garagesale/internal/platform/database"
	"github.com/runlevl4/garagesale/internal/schema"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error | shutting down | %s", err)
		os.Exit(1)
	}
}

func run() error {

	var cfg struct {
		DB struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:localhost"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:false"`
		}
		Args conf.Args
	}

	// =========================================================================
	// Get configuration
	if err := conf.Parse(os.Args[1:], "sales", &cfg); err != nil {
		// Allows user to display help
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("SALES", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// =========================================================================
	// Setup database
	db, err := database.Open(database.Config{
		Host:       cfg.DB.Host,
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		return errors.Wrap(err, "error opening database")
	}
	defer db.Close()

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			return errors.Wrap(err, "migrating database")
		}
		log.Println("main : migrate complete")
	case "seed":
		if err := schema.Seed(db); err != nil {
			return errors.Wrap(err, "seeding database")
		}
		log.Println("main : seed complete")
	}
	return nil
}
