package server

import (
	"fmt"
	"go-challenge/internal/domain"
	"go-challenge/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SalesHandler struct {
	SalesService *services.SalesService
}

func (sh *SalesHandler) HandleNewSale(ctx *gin.Context) {
	var sale *domain.Sale
	if err := ctx.ShouldBindJSON(&sale); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("error trying to bind sale body: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	sale, err := sh.SalesService.Create(ctx, sale)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to create sale: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, sale)
}

func (sh *SalesHandler) HandleRefundSale(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("invalid id: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	s, err := sh.SalesService.Refund(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to refund sale: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

func (sh *SalesHandler) HandleGetSale(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("invalid id: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	s, err := sh.SalesService.Read(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to get sale: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}
