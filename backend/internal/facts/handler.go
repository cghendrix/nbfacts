package facts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	GetFacts(c *gin.Context)
	GetFact(c *gin.Context)
	AddFact(c *gin.Context)
	UpdateFact(c *gin.Context)
	DeleteFact(c *gin.Context)
}

type handler struct {
	service Service
}

func NewHandler(server *gin.Engine, service Service) Handler {
	handler := &handler{service: service}
	handler.RegisterRoutes(server)
	return handler
}

func (h handler) RegisterRoutes(server *gin.Engine) {
	routes := server.Group("/facts")
	routes.POST("/", h.AddFact)
	routes.GET("/", h.GetFacts)
	routes.GET("/:id", h.GetFact)
	routes.PUT("/:id", h.UpdateFact)
	routes.DELETE("/:id", h.DeleteFact)
}

// GetFacts godoc
// @Summary Get the latest facts
// @Schemes
// @Description Gets the most recent nickleback facts sorted by date added
// @Tags facts
// @Accept json
// @Produce json
// @Success 200 {array} models.Fact
// @Router /facts/ [get]
func (h handler) GetFacts(c *gin.Context) {
	facts, err := h.service.GetFacts(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, &facts)
}

// GetFact godoc
// @Summary Gets a fact by id
// @Schemes
// @Description Gets a single fact by its id
// @Param id   path string true "Fact Id"
// @Tags facts
// @Accept json
// @Produce json
// @Success 200 {object} models.Fact
// @Router /fact/{id} [get]
func (h handler) GetFact(c *gin.Context) {
	id := c.Param("id")
	fact, err := h.service.GetFact(c, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, &fact)
}

// AddFact godoc
// @Summary Adds a fact
// @Schemes
// @Description Adds a fact
// @Param request body facts.CreateFactRequest true "Fact to create"
// @Tags facts
// @Accept json
// @Produce json
// @Success 201
// @Router /fact/ [post]
func (h handler) AddFact(c *gin.Context) {
	var requestBody CreateFactRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err := h.service.AddFact(c, requestBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusCreated)
}

// UpdateFact godoc
// @Summary Updates a fact
// @Schemes
// @Description Updates a fact
// @Param request body facts.UpdateFactRequest true "Fact to create"
// @Tags facts
// @Accept json
// @Produce json
// @Success 204
// @Router /fact/ [put]
func (h handler) UpdateFact(c *gin.Context) {
	var requestBody UpdateFactRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err := h.service.UpdateFact(c, requestBody)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusNoContent)
}

// DeleteFact godoc
// @Summary Deletes a fact
// @Schemes
// @Description Deletes a fact
// @Param id   path string true "Fact Id"
// @Tags facts
// @Accept json
// @Produce json
// @Success 204
// @Router /fact/{id} [delete]
func (h handler) DeleteFact(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteFact(c, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.Status(http.StatusNoContent)
}
