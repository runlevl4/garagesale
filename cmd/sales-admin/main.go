// Package main provides the executable for our sales-admin.
package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/runlevl4/garagesale/schema"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	// =========================================================================
	// Setup database
	db, err := openDB()
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

func openDB() (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("postgres", "postgres"),
		Host:     "localhost",
		Path:     "postgres",
		RawQuery: q.Encode(),
	}

	return sqlx.Open("postgres", u.String())
}
