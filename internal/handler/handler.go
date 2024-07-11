package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"net/http"
	"proxy-server/internal/proxy"
	"proxy-server/internal/store"
	"proxy-server/pkg/entity"
	"proxy-server/pkg/utils"
)

// HandleRequest godoc
// @Summary Proxy HTTP request
// @Description Proxies an HTTP request to a specified URL.
// @Accept  json
// @Produce  json
// @Param   request body models.Request true "Request payload"
// @Success 200 {object} models.Response
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Failed to proxy request"
// @Router /proxy [post]
func HandleRequest(c *gin.Context) {
	var req entity.ProxyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		slog.Error("Invalid request", slog.Any("error", err))
		return
	}

	if err := utils.ValidateRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Invalid request data", slog.Any("error", err))
		return
	}

	reqID := uuid.New().String()
	store.SaveRequest(reqID, req)

	res, err := proxy.SendRequest(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to proxy request"})
		slog.Error("Failed to proxy request", slog.Any("error", err))
		return
	}

	res.ID = reqID
	store.SaveResponse(reqID, res)

	c.JSON(http.StatusOK, res)
	slog.Info("Request proxied successfully", slog.String("request_id", reqID))
}
