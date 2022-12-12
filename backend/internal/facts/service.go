package facts

import (
	"cghendrix/nbfacts/models"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetFacts(c *gin.Context) ([]models.Fact, error)
	GetFact(c *gin.Context, id string) (models.Fact, error)
	AddFact(c *gin.Context, request CreateFactRequest) error
	AddFactFromSMS(ctx context.Context, tx *sql.Tx, request CreateFactFromSMSRequest) error
	UpdateFact(c *gin.Context, request UpdateFactRequest) error
	DeleteFact(c *gin.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetFacts(c *gin.Context) ([]models.Fact, error) {
	facts, err := s.repo.GetFacts(c)
	if err != nil {
		return nil, err
	}
	return facts, nil
}

func (s service) GetFact(c *gin.Context, id string) (models.Fact, error) {
	facts, err := s.repo.GetFactById(c, id)
	if err != nil {
		return facts, err
	}
	return facts, nil
}

func (s service) AddFact(c *gin.Context, request CreateFactRequest) error {
	err := s.repo.AddFact(c, request)
	if err != nil {
		return err
	}
	return nil
}

func (s service) AddFactFromSMS(ctx context.Context, tx *sql.Tx, request CreateFactFromSMSRequest) error {
	err := s.repo.AddFactFromSMS(ctx, tx, request)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateFact(c *gin.Context, request UpdateFactRequest) error {
	err := s.repo.UpdateFact(c, request)
	if err != nil {
		return err
	}
	return nil
}

func (s service) DeleteFact(c *gin.Context, id string) error {
	err := s.repo.DeleteFact(c, id)
	if err != nil {
		return err
	}
	return nil
}
