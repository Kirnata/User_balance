package main

import (
	"github.com/Kirnata/User_balance"
	"github.com/Kirnata/User_balance/internal/handler"
	"github.com/spf13/viper"
	"log"
)

//var _ http.Handler = gin.New()

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializion configs: %s", err.Error())
	}

	handlers := handler.NewHandler()
	srv := new(User_balance.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
