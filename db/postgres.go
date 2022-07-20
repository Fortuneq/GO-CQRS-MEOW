package db

import "database/sql"

type PostgresRepository struct{
	db * sql.DB
}
