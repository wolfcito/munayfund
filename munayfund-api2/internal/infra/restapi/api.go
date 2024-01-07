package restapi

import (
	"context"
	"fmt"
	"log"
	"munayfund-api2/internal/core/service"
	mongodb "munayfund-api2/internal/infra/db"
	"munayfund-api2/internal/infra/handler"
	"munayfund-api2/internal/infra/ipfs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	APIName                string = "munayfund"
	DBName                 string = "MunayFundDB"
	UserCollectionName     string = "Users"
	ProjectCollectionName  string = "Projects"
	UserCollectionIndex    string = "_id"
	ProjectCollectionIndex string = "_id"
)

type RestAPI struct {
	Server        *gin.Engine
	ServerReady   chan bool
	ServerStopped chan bool
	Components    *Components
}

type Components struct {
	mongoConnections []*mongodb.MongoImpl
	UserHandler      *handler.UserHandler
	ProjectHandler   *handler.ProjectHandler
	HistoryHandler   *handler.HistoryHandler
}

func NewAPI() *RestAPI {
	ra := &RestAPI{}

	return ra.setupComponents().setupServer()
}

func (ra *RestAPI) StartAPI() {
	go ra.startRestServer()

	if ra.ServerReady != nil {
		ra.ServerReady <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ra.stopAPIServer()
}

func (ra *RestAPI) startRestServer() {
	fmt.Printf("Starting MunayFund API on port: %s\n", os.Getenv("PORT"))

	if err := ra.Server.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting the server, shutting down... %s, %v", "error", err)
	}
}

func (ra *RestAPI) setupComponents() *RestAPI {
	userRepo, err := mongodb.NewMongoDBCollection(context.Background(), DBName, UserCollectionName, UserCollectionIndex)
	if err != nil {
		log.Fatal(err)
	}

	projectRepo, err := mongodb.NewMongoDBCollection(context.Background(), DBName, ProjectCollectionName, ProjectCollectionIndex)
	if err != nil {
		log.Fatal(err)
	}

	userSrv := service.NewUserService(userRepo)
	projectSrv := service.NewProjectService(projectRepo, ipfs.NewIPFSService())

	projectHistorySrv := service.NewHistoryService(projectRepo, ipfs.NewIPFSService())

	ra.Components = &Components{
		UserHandler:      handler.NewUserHandler(userSrv),
		ProjectHandler:   handler.NewProjectHandler(projectSrv),
		HistoryHandler:   handler.NewHistoryHandler(projectHistorySrv),
		mongoConnections: []*mongodb.MongoImpl{userRepo, projectRepo},
	}

	return ra
}

func (ra *RestAPI) setupServer() *RestAPI {
	ra.Server = NewServer(APIName)

	baseGroup := ra.Server.Group(fmt.Sprintf("/api/%s", APIName))

	setupHealthCheckRoute(baseGroup)

	apiGroup := baseGroup.Group("/v1")

	ra.setupEndpoints(apiGroup)

	return ra
}

func (ra *RestAPI) stopAPIServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, mongoConn := range ra.Components.mongoConnections {
		defer mongoConn.Disconnect(ctx)
	}

	if ra.ServerStopped != nil {
		ra.ServerStopped <- true
	}
}
