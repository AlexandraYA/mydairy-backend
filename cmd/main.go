package main

import (
	"context"
	"fmt"
	"os"

	dairy "github.com/alexandraya/mydairy-backend"
	"github.com/alexandraya/mydairy-backend/pkg/handler"
	"github.com/alexandraya/mydairy-backend/pkg/repository"
	"github.com/alexandraya/mydairy-backend/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Print("This is my dairy")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, errDB := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if errDB != nil {
		logrus.Fatalf("failed to initialize DB: %s", errDB.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(dairy.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("DairyApp Started")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
