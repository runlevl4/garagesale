// Package product contains all details about dealing with our products.
package product

import "github.com/jmoiron/sqlx"

func List(db *sqlx.DB) ([]Product, error) {
	// Empty list of products to populate from database
	list := []Product{}

	const q = "select * from products"
	if err := db.Select(&list, q); err != nil {
		return nil, err
	}
	return list, nil
}
