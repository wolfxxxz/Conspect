package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"slovarik/internal/app/api"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

var variable string //переменная принимает значение
var defaultVal = ":8080"
var description = "path to config file in toml or .env format"
var path = "configs/"

func init() {
	flag.StringVar(&variable, "path", defaultVal, description)
}

func main() {
	//flag.Pars читает данные с настройками flag.Stringvar(...)
	flag.Parse()
	log.Println("It works")
	//fmt.Println(configPath)

	config := api.NewConfig()

	//И сохранить в config
	if variable == "env" {
		err := godotenv.Load(path + ".env")
		if err != nil {
			log.Fatal("Could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")

	} else if variable == "toml" {
		_, err := toml.DecodeFile(path+"api.toml", config)
		if err != nil {
			log.Println("Can not find configs file. Using default value", err)
		}
	} else {
		fmt.Println("you wrote wrong file - config ", variable)
	}

	fmt.Println(config)

	server := api.New(config)

	//api server Start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}

//var l []library.Library

/*func init() {

	//-----Зачитываем содержимое файла библиотека-----
	l = library.Takejson("library.json")
}
		rttr = library.TakeTXT("library.txt")
		library.Savejson(rttr, "library.json")
		library.WriteArr(rttr, "library.txt")
		library.SortLibraryTheme(rttr, "library.txt")
*/

/* SemiTrash Server
log.Println("Starting REST API server and port: ", port)
//Инициализация роутера
router := mux.NewRouter()
utils.BuildWordResource(router, wordResourcePrefix)
utils.BuildManyWordsResource(router, manyWordsResourcePrefix)
log.Println("Router initalizing successfully")
//Безконечный цикл
//log.Println("Router configured successfully! Let's go!")
log.Fatal(http.ListenAndServe(":"+port, router))
*/

/*
	log.Println("Trying to start Slovarik")
	router := mux.NewRouter()
	//mux.NewRouter == http://localhost - вероятно используется по умолчанию...
	log.Println("Router configured successfully! Let's go!")
	//Если на вход пришел запрос /themes
	router.HandleFunc("/themes", httpSend.ThemesWordsPrint).Methods("GET")

	//Если на вход пришел запрос  /words/{id}
	router.HandleFunc("/words/{id}", httpSend.GetWordById).Methods("GET")

	//Безконечный цикл

	log.Fatal(http.ListenAndServe(":"+port, router))

	//if err := nil http.ListenAndServe(":"+port, nil); err != nil {
	//log.Fatal(err)
	//}
*/

//Тесты и т.д.
/*
	fmt.Println("Тест слов по количеству")
	testink.TestKnowlig(rttr)

	fmt.Println("Тест знаний по теме")
	testink.ThemesOfWords(rttr)

	fmt.Println("Добавить новые слова СПИСКОМ - введите 1")
	library.UpdateLibrary("newWords.txt", rttr)

	fmt.Println("Ввести новые слова в ручном режиме     2")
	library.NewWordRukamy(rttr)

	fmt.Println("Добавить тему или изменить слова       3")
	library.AddTheme(rttr)

	fmt.Println("Сортировать библиотеку порядок изменится  4")
	library.SortLibrary(rttr, "library.txt")
*/
