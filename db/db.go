package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateDBConnection(*StorageParams) (*PostgresDB, error)
}

type StorageParams struct {
	driver   string
	user     string
	port     string
	password string
	dbname   string
	sslmode  string
}

func NewStorageParams(driver, user, port, password, dbname, sslmode string) *StorageParams {
	return &StorageParams{
		driver:   driver,
		user:     user,
		port:     port,
		password: password,
		dbname:   dbname,
		sslmode:  sslmode,
	}
}

type PostgresDB struct {
	db            *sql.DB
	storageParams StorageParams
}

func CreateDBConnection(sp *StorageParams) (*PostgresDB, error) {
	connStr := NewConnString(sp)
	db, err := sql.Open(sp.driver, connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New("failed to ping DB")
	}
	return &PostgresDB{
		db:            db,
		storageParams: *sp,
	}, nil
}

func NewConnString(sp *StorageParams) string {
	return fmt.Sprintf("user=%s port=%s dbname=%s password=%s sslmode=%s", sp.user, sp.port, sp.dbname, sp.password, sp.sslmode)
}
