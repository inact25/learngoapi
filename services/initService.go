package services

import "database/sql"

type SchoolService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *SchoolService {
	return &SchoolService{db}
}
