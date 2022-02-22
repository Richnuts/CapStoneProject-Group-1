package checkinandout

import (
	"database/sql"
)

type CheckRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CheckRepository {
	return &CheckRepository{db: db}
}
