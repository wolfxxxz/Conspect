package api

import (
	"net/http"

	"ServerAndDB/storage"

	_ "github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a API) configreRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})
}

// Пытаемся сконфигурировать хранилище
func (a *API) configreStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пітаемся установить соединение если невозможно возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
