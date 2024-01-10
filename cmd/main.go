package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	todoback "github.com/moxicom/todo-back"
	config "github.com/moxicom/todo-back/configs"
	"github.com/moxicom/todo-back/pkg/handlers"
	"github.com/moxicom/todo-back/pkg/repository"
	"github.com/moxicom/todo-back/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Todo application back
// @version 0.1
// @description This is a simple backend for todo application using golang

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx); err != nil {
		logrus.Fatal(err)
	}
}

func runServer(ctx context.Context) error {
	// Config init
	if err := config.Init(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}
	// dotenv init
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	// Db opening
	cfg, err := config.InitDbConfig()
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	db, err := repository.NewDbInit(cfg)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	//Db Migration
	err = repository.NewMigration(db)
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}

	// Dependency injection
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handlers.NewHandler(service)
	server := todoback.NewServer()

	go func() {
		if err = server.Run(viper.GetString("server_port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("Listen and serve: %s", err.Error())
		}
	}()

	<-ctx.Done()

	logrus.Println("Shutting down gracefully")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
