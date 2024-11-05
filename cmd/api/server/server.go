package server

import (
	"go-challenge/internal/repository"
	"go-challenge/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// storage initialization
	salesRepo := repository.NewSalesLocal()
	shirtsRepo := repository.NewShirtsLocal()

	r := gin.Default()

	// health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ------------ shirts CRUD ---------------

	h := &ShirtsHandler{
		ShirtService: services.NewShirtService(shirtsRepo),
	}

	r.POST("/shirts", h.HandleNewShirt)
	r.PATCH("/shirts/:id", h.HandleUpdateShirt)
	r.GET("/shirts/:id", h.HandleGetShirt)
	r.DELETE("/shirts/:id", h.HandleDeleteShirt)

	// ------------ shirts CRUD ---------------

	// ------------ sales manager ---------------

	sh := &SalesHandler{
		SalesService: services.NewSalesService(salesRepo, shirtsRepo),
	}

	r.POST("/sales", sh.HandleNewSale)
	r.POST("/sales/:id/refund", sh.HandleRefundSale)
	r.GET("/sales/:id", sh.HandleGetSale)

	// ------------ sales manager ---------------

	return r
}
