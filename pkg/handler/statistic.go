package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	urlshortener "github.com/ramil66/url-shortener"
)

type getAllMetricResponse struct {
	Data []urlshortener.Statistic `json:"data"`
}

func (h *Handler) GetMetric(c *gin.Context) {
	alias, err := getAlias(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	metric, err := h.services.Statistic.GetMetric(alias)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllMetricResponse{
		Data: metric,
	})
}
