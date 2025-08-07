package services

import (
	"net/http"

	"github.com/KhanhLinh2810/5G-core/udm/internal/models"
	"github.com/KhanhLinh2810/5G-core/udm/internal/types"
)

func GetSDMDetail(req types.GetSDMDetailType, resultChan chan any) {
	session, err := models.GetSessionBySupi(req.Supi)
	if err != nil {
		resultChan <- map[string]any{
			"status": http.StatusNotFound,
			"data": map[string]any{
				"error":   "No session found for SUPI",
				"supi":    req.Supi,
				"details": err.Error(),
			},
		}
		return
	}

	resultChan <- map[string]any{
		"status": http.StatusOK,
		"data": map[string]any{
			"supi": session.Supi,
			"dnn":  session.Dnn,
			"sNssai": map[string]any{
				"sst": session.Sst,
				"sd":  session.Sd,
			},
		},
	}
}
