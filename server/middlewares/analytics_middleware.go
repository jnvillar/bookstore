package middlewares

import (
	"bookstore/analytics"
	"bookstore/log"

	"github.com/gin-gonic/gin"
)

type AnalyticsMiddleware struct {
	analytics *analytics.Factory
	log       *log.Factory
}

func NewAnalyticsMiddleware(analytics *analytics.Factory, log *log.Factory) *AnalyticsMiddleware {
	return &AnalyticsMiddleware{
		analytics: analytics,
		log:       log,
	}
}

func (m *AnalyticsMiddleware) RegisterMiddleware(router *gin.Engine) {
	router.Use(m.AnalyticsMiddleWare)
}

func (m *AnalyticsMiddleware) AnalyticsMiddleWare(c *gin.Context) {
	m.log.Info("registering visit")
	if err := m.analytics.AddVisit(c.Request.RemoteAddr); err != nil {
		m.log.Error("error logging visit", err)
	}
	c.Next()
}
