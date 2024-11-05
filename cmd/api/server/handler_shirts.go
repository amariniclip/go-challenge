package server

import (
	"errors"
	"fmt"
	"go-challenge/internal/domain"
	"go-challenge/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShirtsHandler struct {
	ShirtService *services.ShirtService
}

func (h *ShirtsHandler) HandleNewShirt(ctx *gin.Context) {
	var body *domain.Shirt
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("error trying to bind shirt body: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	s, err := h.ShirtService.Create(ctx.Request.Context(), body)
	if err != nil {
		var vErrs domain.ShirtValidationError
		if errors.As(err, &vErrs) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message":     fmt.Sprintf("invalid shirt: %s", vErrs.Error()),
				"status_code": http.StatusBadRequest,
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to create shirt body: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, s)
}

func (h *ShirtsHandler) HandleUpdateShirt(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("invalid id: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	var body *services.UpdateShirtParams
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("error trying to bind shirt update params: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	s, err := h.ShirtService.Update(ctx.Request.Context(), id, body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to update shirt: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

func (h *ShirtsHandler) HandleGetShirt(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("invalid id: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	s, err := h.ShirtService.Read(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to get shirt: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, s)
}

func (h *ShirtsHandler) HandleDeleteShirt(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     fmt.Sprintf("invalid id: %s", err.Error()),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	if err := h.ShirtService.Delete(ctx.Request.Context(), id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     fmt.Sprintf("error trying to delete shirt: %s", err.Error()),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "deleted successfully",
		"status_code": http.StatusOK,
	})
}
