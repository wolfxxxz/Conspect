package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/vlasove/go2/12.GinGormAPI/models"
	"github.com/vlasove/go2/12.GinGormAPI/routers"
	"github.com/vlasove/go2/12.GinGormAPI/storage"
)

var err error

func main() {
	//go get -u girhub.com/jinzhu/gorm/<dialects>
	storage.DB, err = gorm.Open("postgres", "host=... user=... password=... dbname= ...")
	if err != nil {
		log.Println("error while accessing database:", err)
	}
	defer storage.DB.Close()                  // не забудем закрыть соединение
	storage.DB.AutoMigrate(&models.Article{}) // в этот момент орм сама сгенерит все запросы, миграции и их применит

	r := routers.SetupRouter()

	// r - gin маршрутизатор
	r.Run()
}
