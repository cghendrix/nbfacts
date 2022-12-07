package facts

import (
	"cghendrix/nbfacts/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetFacts(c *gin.Context) ([]models.Fact, error)
	GetFactById(c *gin.Context, id string) (models.Fact, error)
	AddFact(c *gin.Context, request CreateFactRequest) error
	UpdateFact(c *gin.Context, request UpdateFactRequest) error
	DeleteFact(c *gin.Context, id string) error
}

type repository struct {
	DB *sqlx.DB
}

func NewRepository(DB *sqlx.DB) Repository {
	return &repository{DB: DB}
}

func (r repository) GetFacts(c *gin.Context) ([]models.Fact, error) {
	var facts []models.Fact
	query := `
		SELECT 
		    * 
		FROM fact 
		ORDER BY date_added 
		DESC LIMIT 10;
	`
	err := r.DB.SelectContext(c, &facts, query)
	if err != nil {
		return nil, err
	}
	return facts, nil
}

func (r repository) GetFactById(c *gin.Context, id string) (models.Fact, error) {
	var fact models.Fact
	query := `
		SELECT 
		    * 
		FROM fact 
		WHERE fact_id = ?;
	`
	err := r.DB.GetContext(c, &fact, query, id)
	if err != nil {
		return models.Fact{}, err
	}
	return fact, nil
}

func (r repository) AddFact(c *gin.Context, request CreateFactRequest) error {
	query := `
		INSERT INTO 
		    fact 
		    (body, info)
		VALUES
		    (?, ?)
	`
	_, err := r.DB.ExecContext(c, query, request.Body, request.Info)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) UpdateFact(c *gin.Context, request UpdateFactRequest) error {
	query := `
		UPDATE 
		    fact 
		SET body = ?, info = ?, date_updated = ?
		WHERE fact_id = ?;
	`
	_, err := r.DB.ExecContext(c, query, request.Body, request.Info, request.Updated, c.Param("id"))
	if err != nil {
		return err
	}
	return nil
}

func (r repository) DeleteFact(c *gin.Context, id string) error {
	query := `
		DELETE 
		FROM 
		    fact 
		WHERE fact_id = ?;
	`
	_, err := r.DB.ExecContext(c, query, id)
	if err != nil {
		return err
	}
	return nil
}
