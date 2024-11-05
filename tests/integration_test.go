package tests

import (
	"encoding/json"
	"fmt"
	"go-challenge/cmd/api/server"
	"go-challenge/internal/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestIntegration(t *testing.T) {
	s := server.New()

	resShirtExample := doRequest(s, httptest.NewRequest(http.MethodPost, "/shirts", strings.NewReader(`{
		"brand": "Nike",
		"size": "XL",
		"description": "sarasa",
		"available_units": 10,
		"price_per_unit": 2.0
	}`)))

	require.Equal(t, http.StatusCreated, resShirtExample.Code)
	var shirtExample domain.Shirt
	require.Nil(t, json.Unmarshal(resShirtExample.Body.Bytes(), &shirtExample))

	t.Run("crud shirt", func(t *testing.T) {
		resShirt := doRequest(s, httptest.NewRequest(http.MethodPost, "/shirts", strings.NewReader(`{
			"brand": "Nike",
			"size": "XL",
			"description": "sarasa",
			"available_units": 10,
			"price_per_unit": 2.0
		}`)))

		require.Equal(t, http.StatusCreated, resShirt.Code)

		var shirt domain.Shirt
		require.Nil(t, json.Unmarshal(resShirt.Body.Bytes(), &shirt))

		resShirt = doRequest(s, httptest.NewRequest(http.MethodGet, fmt.Sprintf("/shirts/%s", shirt.ID), nil))
		require.Equal(t, http.StatusOK, resShirt.Code)
	})

	t.Run("sale test", func(t *testing.T) {
		resSale := doRequest(s, httptest.NewRequest(http.MethodPost, "/sales", strings.NewReader(fmt.Sprintf(`{
			"customer": {
				"name": "Ayrton",
				"email": "amarini@gmail.com",
				"phone": "+5492664016024"
			},
			"cart": [
				{
					"id": "%s"
				}
			]
		}`, shirtExample.ID,
		))))

		require.Equal(t, http.StatusCreated, resSale.Code)

		var sale domain.Sale
		require.Nil(t, json.Unmarshal(resSale.Body.Bytes(), &sale))

		resSale = doRequest(s, httptest.NewRequest(http.MethodGet, fmt.Sprintf("/sales/%s", sale.ID), nil))
		require.Equal(t, http.StatusOK, resSale.Code)

		resSale = doRequest(s, httptest.NewRequest(http.MethodPost, fmt.Sprintf("/sales/%s/refund", sale.ID), nil))
		require.Equal(t, http.StatusOK, resSale.Code)
	})
}

func doRequest(e *gin.Engine, r *http.Request) *httptest.ResponseRecorder {
	writer := httptest.NewRecorder()
	e.ServeHTTP(writer, r)

	return writer
}
