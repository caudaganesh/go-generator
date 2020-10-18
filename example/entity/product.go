package entity

import "time"

type ProductStatus int

// Product partner entity
type Product struct {
	ID            int       `db:"id"         json:"id"`
	Name          string    `db:"name"       json:"name"`
	Active        bool      `db:"active"     json:"active"`
	Discount      float32   `db:"discount"   json:"discount"`
	Price         int       `db:"price"      json:"price"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	ProductStatus ProductStatus
}
