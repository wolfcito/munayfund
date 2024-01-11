package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupHealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

func (ra *RestAPI) setupEndpoints(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/login", ra.Components.UserHandler.Login)
	apiGroup.POST("/signup", ra.Components.UserHandler.SigUp)

	userAPIGroup := apiGroup.Group("/user")
	userAPIGroup.Use(basicTokenMiddleware())

	userAPIGroup.DELETE("/:id", ra.Components.UserHandler.Delete)
	userAPIGroup.PUT("/:id", ra.Components.UserHandler.Update)

	projectsAPIGroup := apiGroup.Group("/projects")
	projectsAPIGroup.GET("/", ra.Components.ProjectHandler.GetProjects)
	projectsAPIGroup.GET("/search", ra.Components.ProjectHandler.GetProjectsBySimilarName)
	projectsAPIGroup.GET("/:id", ra.Components.ProjectHandler.GetProjectDetails)

	projectsAuthAPIGroup := projectsAPIGroup.Group("/")
	projectsAuthAPIGroup.Use(basicTokenMiddleware())

	projectsAuthAPIGroup.POST("/create", ra.Components.ProjectHandler.CreateProject)
	projectsAuthAPIGroup.PUT("/:id/media", ra.Components.ProjectHandler.UpdateProjectMultimedia)
	projectsAuthAPIGroup.PUT("/:id", ra.Components.ProjectHandler.UpdateProject)
	projectsAuthAPIGroup.DELETE("/:id", ra.Components.ProjectHandler.DeleteProject)

	projectsHistoryAPIGroup := projectsAPIGroup.Group("/history")

	projectsHistoryAPIGroup.PUT("/:id", ra.Components.HistoryHandler.UpdateHistory)
	projectsHistoryAPIGroup.GET("/:id", ra.Components.HistoryHandler.GetHistory)

	apiGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
