package main

import (
	"fmt"

	"billing-service/helpers"
	"billing-service/httphandler"
	"billing-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func main() {
	logger.Debug("Service starting")

	// регистрация общих метрик
	helpers.RegisterCommonMetrics()

	// настройка всех конфигов
	config, err := models.LoadConfig()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Config error: %v", err))
	}

	// настройка коннекта к БД
	_, err = helpers.InitDataBase(config.Database)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Cant initialize DB: %v", err))
	}

	// выполнение миграций
	helpers.CreateTables()

	// инициализация REST-api
	httphandler.InitHTTPServer()

	logger.Info("Service started")
}
