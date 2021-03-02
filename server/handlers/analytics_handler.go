package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/analytics"
	"bookstore/log"
	"bookstore/utils"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	log       *log.Factory
	analytics *analytics.Factory
}

func NewAnalyticsHandler(log *log.Factory, analytics *analytics.Factory) *AnalyticsHandler {
	return &AnalyticsHandler{
		analytics: analytics,
		log:       log,
	}
}

func (b *AnalyticsHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/analytics", func(c *gin.Context) { b.getAnalyticsInfo(c.Writer, c.Request) })
}

func (b *AnalyticsHandler) getAnalyticsInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	analytics, err := b.analytics.GetAnalytics()
	if err != nil {
		b.log.Error("error fetching analytics", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(analytics)
	if err != nil {
		b.log.Error("error marshalling analytics", err)
		utils.WriteError(http.StatusInternalServerError, w, err)
		return
	}
}
