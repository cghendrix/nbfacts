package sms

import (
	"database/sql"
	"golang.org/x/net/context"
)

type Service interface {
	AddSMS(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) AddSMS(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) (string, error) {
	lastInsertedId, err := s.repo.AddSMS(ctx, tx, request)
	if err != nil {
		return "", err
	}
	return lastInsertedId, nil
}
