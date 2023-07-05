package api

import (
	"ServerAndDB2/storage"

	_ "github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаемся отконфигурировать маршрутизатор (а конкретнее поле роутер)
func (a *API) configreRouterField() {
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
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
