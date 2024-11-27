package pg

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbAdapter *PgAdapter

type PgAdapter struct {
	PgClient *gorm.DB
}

func NewPgAdapter(connectionString string) (*PgAdapter, error) {
	if dbAdapter == nil {
		pgClient, err := NewPgClient(connectionString)
		if err != nil {
			return nil, fmt.Errorf("failed to create postgres client: %w", err)
		}
		dbAdapter = &PgAdapter{
			PgClient: pgClient,
		}
	}
	return dbAdapter, nil
}

func NewPgClient(connectionString string) (*gorm.DB, error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (pg *PgAdapter) Migrate(dst ...interface{}) error {
	return pg.PgClient.AutoMigrate(dst...)
}
