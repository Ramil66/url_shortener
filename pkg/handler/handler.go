package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ramil66/url-shortener/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(a *gin.RouterGroup) *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}
	redirect := router.Group("")
	{
		redirect.GET("/:alias", h.RedirectUrl)
	}
	api := router.Group("/api")
	{
		api.POST("/shorten", h.ShorteningUrl)
		user := router.Group("/api/user", h.userIdentity)
		{
			user.GET("/metric/:alias", h.GetMetric)
			user.POST("/custom-url", h.CustomUrl)
			user.POST("/shorten", h.ShorteningUrlUsers)
			user.GET("/urls", h.GetAllUrls)
			user.DELETE("/urls/:alias", h.DeleteUrl)
		}
	}
	return router
}
