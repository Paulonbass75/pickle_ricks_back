package models

import (
	"database/sql"
	"time"
)

//DBModel is a struct that holds the database connection
type DBModel struct {
	DB *sql.DB
}

//Models is wrapper for all models
type Models struct {
	DB DBModel
}
//NewModels returns a new instance of the models with db connection pool
func NewModels(db *sql.DB) *Models {
	return &Models{
		DB: DBModel{DB: db},
	}
}


//Meseeks is a struct that holds all the merch
type Meseeks struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}