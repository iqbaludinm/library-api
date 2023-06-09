package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/library-api/config"
	"github.com/iqbaludinm/library-api/handler"
	"github.com/iqbaludinm/library-api/repositories"
	"github.com/iqbaludinm/library-api/routes"
	"github.com/iqbaludinm/library-api/services"
)

var router = gin.New()

func StartApp() {
	repo := repositories.NewRepo(config.PSQL.DB)
	service := services.NewService(repo)
	server := handler.NewHttpServer(service)

	routes.RegisterAPI(router, server)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))

}