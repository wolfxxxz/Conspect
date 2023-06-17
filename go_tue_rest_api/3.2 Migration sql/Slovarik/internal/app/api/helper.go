package api

import (
	"net/http"

	"Slovarik/storage"

	_ "github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

// Пытаемся отконфигурировать наш API инстанс а точнее поле Logger
// Установить уровень logger
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	//Если нет ошибки то уровень логирования == log_level,
	// который может достать func logrus.SetLevel
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router)
func (a API) configreRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api! Slovarik"))
	})
}

// Пытаемся сконфигурировать хранилище
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Питаемся установить соединение если невозможно возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
