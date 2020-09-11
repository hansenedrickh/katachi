package repository

import "database/sql"

type Repository interface {

}

type repository struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) Repository {
	return &repository{
		DB: DB,
	}
}

