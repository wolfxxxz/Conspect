package api

import (
	"net/http"

	"ServerAndDB/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instanse description
type API struct {
	//UNEXPORTED FIELD
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавления поля для работы с хранилищем
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure Loggers, router, database connection and etc...
func (api *API) Start() error {
	//Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//Подтверждение того что логер построен
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configreRouterField()

	//Конфигурируем хранилище
	if err := api.configreStorageField(); err != nil {
		return err
	}

	//На этапе валидного завершения стартуем http server
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
