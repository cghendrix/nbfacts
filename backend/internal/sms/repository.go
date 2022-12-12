package sms

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AddSMS(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) (string, error)
}

type repository struct {
	f  *sqlx.DB
	DB *sqlx.DB
}

func NewRepository(DB *sqlx.DB) Repository {
	return &repository{DB: DB}
}

func (r repository) AddSMS(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) (string, error) {

	query := `
		INSERT INTO 
		    sms 
		    (body, sender, message_sid)
		VALUES
		    (?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query, request.Body, request.From, request.MessageSID)
	if err != nil {
		return "", err
	}
	lastInsertedId, err := r.GetLastInsertedId(ctx, tx)
	if err != nil {
		return "", err
	}
	return lastInsertedId, nil
}

func (r repository) GetLastInsertedId(ctx context.Context, tx *sql.Tx) (string, error) {
	var lastInsertedId string
	query := `
		SELECT @last_sms_uuid;
	`
	row := tx.QueryRowContext(ctx, query)
	err := row.Scan(&lastInsertedId)
	if err != nil {
		return "", err
	}
	return lastInsertedId, nil
}
