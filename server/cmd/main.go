package main

import (
	"context"
	"os"
	"os/signal"
	"stcTask/docs"
	"stcTask/server"
	"stcTask/server/pkg/handler"
	"stcTask/server/pkg/logging"
	"stcTask/server/pkg/repository"
	"stcTask/server/pkg/service"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.GetLooger()

	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewRedisDB(viper.GetViper())
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewReposiroty(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(server.Server)
	runHttp(srv, handler, logger)
}

func runHttp(srv *server.Server, handler *handler.Handler, logger logging.Logger) {
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Print("stcTask Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Print("stcTask Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("server/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "server API"
	docs.SwaggerInfo.Description = "stcTask swagger api examples"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
