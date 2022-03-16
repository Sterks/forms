package http

import (
	"errors"
	"forms/internal/config"
	"forms/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	router.GET("/", h.setClientFromRequest)

	return router
}

type setClientFromRequest struct {
	Firstname  string `json:"firstname" binding:"required,min=2,max=64"`
	Lastname   string `json:"lastname" binding:"required,min=2,max=64"`
	Patronymic string `json:"patronymic" binding:"required,min=2,max=64"`
	Position   string `json:"position" binding:"required,min=2,max=64"`
	Company    string `json:"company" binding:"required,min=2,max=64"`
	Phone      string `json:"phone" binding:"required,min=2,max=14"`
	Email      string `json:"email" binding:"required,email,max=64"`
}

func (h *Handler) setClientFromRequest(c *gin.Context) {
	var sc setClientFromRequest
	if err := c.BindJSON(&sc); err != nil {
		c.AbortWithStatusJSON(500, errors.New("Введенные данные некорректны"))
	}

	_, err := h.services.ClientService.Create(c.Request.Context(), service.ClientInput{
		Firstname:  sc.Firstname,
		Lastname:   sc.Lastname,
		Patronomic: sc.Patronymic,
		Position:   sc.Patronymic,
		Company:    sc.Company,
		Phone:      sc.Phone,
		Email:      sc.Email,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, errors.New("Введенные данные некорректны"))
	}

	c.Status(http.StatusCreated)
}
