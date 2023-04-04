package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vovah1a/my-money/docs"
	"github.com/vovah1a/my-money/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		board := api.Group("/board")
		{
			board.POST("/", h.createBoard)
			board.GET("/", h.getAllBoards)
			board.GET("/:id", h.getBoardById)
			board.PUT("/:id", h.updateBoard)
			board.DELETE("/:id", h.deleteBoard)

			category := board.Group("/:id/category")
			{
				category.POST("/", h.createCategory)
				category.GET("/", h.getAllCategories)
			}
		}

		category := api.Group("/category")
		{
			category.GET("/:id", h.getCategoryById)
			category.PUT("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}
	}

	return router
}
