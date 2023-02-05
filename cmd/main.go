package main

import (
	"context"
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/handler"
	"github.com/Kirnata/User_balance/internal/repository"
	"github.com/Kirnata/User_balance/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	zapLogger, errZap := zap.NewProduction()
	if errZap != nil {
		log.Println("Error in creation zapLogger")
	}
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Println(err)
		}
	}(zapLogger)
	logger := zapLogger.Sugar()

	if err := initConfig(); err != nil {
		logger.Fatalf("error initializion configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Fatalf("failed to initialized db: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(User_balance.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
			logger.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()
	logger.Info("User_balance started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Info("User_balance finished")
	if err := srv.ShutDown(context.Background()); err != nil {
		logger.Errorf("error occurred on server shutting down: %s", err)
	}
	if err := db.Close(); err != nil {
		logger.Errorf("error occurred on db connection close: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
