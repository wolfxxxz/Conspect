package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Wolfxxxz/Dictionary/internal/app/apiserver"
)

var (
	configPath string //= "configs/api.toml"
	//configPath string = "configs/.env"
)

func init() {

	flag.StringVar(&configPath, "pathtoml", "configs/api.toml", "path to config file in .toml format")
	// - path to config file in .env format - В каком файде хранится инфо (дескриптор)
	//flag.StringVar(&configPath, "pathenv", "configs/.env", "path to config file in .env format")
	// Параметры path - Должны отличатся
}

func main() {
	//Запускаем функцию init и flag.StringVar()
	flag.Parse()
	log.Println("It works")
	//Server instance initialization
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Using default value", err)
	}
	//Теперь нужно попробовать прочитать из .toml/.env , так как там может быть новая информация
	server := apiserver.New(config)

	//И сохранить в config
	/*if configPath == "configs/.env" {
		//Если на FLAG пришол запрос на puthenv
		err := godotenv.Load("configs/.env")
		if err != nil {
			log.Fatal("Could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
		//fmt.Println(config)
	} else if configPath == "configs/api.toml" {
		//Если на FLAG пришол запрос на puthtoml
		//Достаём конфиг из томл var configPath string = "configs/api.toml"
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("Can not find configs file. Using default value", err)
		}
	}*/

	//api server Start
	//log.Fatal(server.Start())

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
