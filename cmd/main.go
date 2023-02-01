package main

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/handler"
	"github.com/Kirnata/User_balance/internal/repository"
	"github.com/Kirnata/User_balance/internal/service"
	"github.com/Kirnata/User_balance/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

//var _ http.Handler = gin.New()

func main() {
	zapLog, err := logger.InitializeLogger()
	if err != nil {
		log.Fatalf("Error in creation zapLogger: %s", err.Error())
	}
	if err := initConfig(); err != nil {
		zapLog.Fatalf("error initializion configs: %s", err.Error())

	}

	if err := godotenv.Load(); err != nil {
		zapLog.Fatalf("error loading env variables: %s", err.Error())
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
		zapLog.Fatalf("failed to initialized db: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(User_balance.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		zapLog.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
