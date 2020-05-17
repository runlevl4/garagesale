// Package main provides the executable for our sales-admin.
package main

import (
	"flag"
	"log"

	"github.com/runlevl4/garagesale/internal/platform/database"
	"github.com/runlevl4/garagesale/internal/schema"
)

func main() {

	// =========================================================================
	// Setup database
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	flag.Parse()
	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Fatal(err)
		}
		log.Println("main : migrate complete")
		return
	case "seed":
		if err := schema.Seed(db); err != nil {
			log.Fatal(err)
		}
		log.Println("main : seed complete")
		return
	}

}
