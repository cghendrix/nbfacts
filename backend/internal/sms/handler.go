package sms

import (
	"cghendrix/nbfacts/internal/facts"
	"context"
	"database/sql"
)

type Handler interface {
	AddSMSWithFact(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) error
}

type handler struct {
	smsService   Service
	factsService facts.Service
}

func NewHandler(smsService Service, factsService facts.Service) Handler {
	return &handler{smsService: smsService, factsService: factsService}
}

func (h handler) AddSMSWithFact(ctx context.Context, tx *sql.Tx, request CreateSMSRequest) error {
	lastInsertedId, err := h.smsService.AddSMS(ctx, tx, request)
	if err != nil {
		return err
	}
	err = h.factsService.AddFactFromSMS(ctx, tx, facts.CreateFactFromSMSRequest{
		Body:  request.Body,
		Info:  "",
		SmsId: lastInsertedId,
	})
	if err != nil {
		return err
	}
	return nil
}
