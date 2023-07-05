package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Wolfxxxz/JWTAuth/internal/app/middleware"
	"github.com/Wolfxxxz/JWTAuth/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *APIServer) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	articles, err := api.store.Article().SelectAll()
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing articles in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("Get All Articles GET /articles")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (api *APIServer) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /articles")
	var article models.Article
	//взять с постмана json и декодировать в article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	// отправить в pgsql
	a, err := api.store.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *APIServer) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *APIServer) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by Id DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.store.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Article with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

// Update Article by ID---------------------------------------------------------------------------------------------
func (api *APIServer) PutArticleById(writer http.ResponseWriter, req *http.Request) {
	//now its get
	initHeaders(writer)
	api.logger.Info("PUT Article by ID /api/v1/articles/{id}")
	// Если ID не инт
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Проверка на наличие в дб
	_, ok, err := api.store.Article().FindArticleById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Достать из постман json
	var article models.Article
	error := json.NewDecoder(req.Body).Decode(&article)
	if error != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article.ID = id
	//fmt.Println(article.ID)
	_, errar := api.store.Article().Update(&article)
	if errar != nil {
		api.logger.Info("Troubles while updating new article:", errar)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *APIServer) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register POST /api/v1/user/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Пытаемся найти пользователя с таким логином в бд
	_, ok, err := api.store.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Смотрим, если такой пользователь уже есть - то никакой регистрации мы не делаем!
	if ok {
		api.logger.Info("User with that ID already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exists in database",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Теперь пытаемся добавить в бд
	userAdded, err := api.store.User().Create(&user)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *APIServer) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth /api/v1/user/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	// Если не json or some troubles
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Пытаемся найти пользователя в бд
	userInDB, ok, err := api.store.User().FindByLogin(userFromJson.Login)
	// Проблема доступа к бд
	if err != nil {
		api.logger.Info("Can not make user search in database", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//Если нет такого пользователя в бд
	if !ok {
		api.logger.Info("User with that Login does not exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does'nt exists in database. Try register first",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Проверяем что пароль совпадает
	if userInDB.Password != userFromJson.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Your password is invalid",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Теперь выбиваем токен как знак успешной аутентификации
	token := jwt.New(jwt.SigningMethodHS256)             // Тот же метод подписания что и в JwtMiddleware
	claims := token.Claims.(jwt.MapClaims)               // Дополнительные действия (в формате map) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //время после которого токен будет считатся протухшим (здесь 2 часа)
	claims["admin"] = true
	claims["name"] = userInDB.Login
	// получаем токен
	tokenString, err := token.SignedString(middleware.SecretKey)
	// Если токен выбить не удалось по каким либо причинам
	if err != nil {
		api.logger.Info("Can not claim jwt-token")
		msg := Message{
			StatusCode: 500,
			// Не указывать что проблема с получением токена иначе могут быть траблы
			Message: "We have some troubles. Try again",
			IsError: true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Если токен успешно выбит отдаём его клиенту
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
