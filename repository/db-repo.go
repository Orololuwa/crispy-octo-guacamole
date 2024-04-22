package repository

import (
	"context"
	"database/sql"
)

type dbRepo struct {
	DB *sql.DB
}

func NewDBRepo(conn *sql.DB) DBRepo {
	return &dbRepo{
		DB: conn,
	}
}

func (m *dbRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
    tx, err := m.DB.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
	
    defer func() error{
        if err != nil {
            tx.Rollback()
            return err
        }

        if err := tx.Commit(); err != nil {
            return err
        }

        return nil
    }()

    if err := operation(ctx, tx); err != nil {
        return err
    }

    return nil
}