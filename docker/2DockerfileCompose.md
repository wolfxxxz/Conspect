# Проверки версий 
docker --version
docker-compouse version
node --version
# First project node.js
## docker-compose.yml
version: '3'

services:
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    command: npm run start
    ports:
      - 3000:3000
## ./api/Dockerfile
FROM node:13

#Создать папку внутри контейнера
WORKDIR /usr/src/app

#Чтоб не копировать прям всё
#Скопировать папки в созданную дирикторию
COPY package*.json ./

#Запустить внутри контейнера прогу
RUN npm install

#Эта комманда копирует прям всё
COPY . .
```
# ----эти действия берёт на себя docker-compose--------------
# Открыть порт в контейнере
#EXPOSE 3000

# Запустить приложение
#CMD [ "node", "run start" ]
#-----------------------------------------------------------
```
## cmd: docker-compose build
Собрать
## cmd: docker-compose up
Запустить
## docker-compose up --build
Собрать и запустить
## environment переменные окружения
### как получить порт автоматически из контейнера
version: '3'

services:
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    command: npm run start
    restart: unless-stopped
    ports:
      - 3000:3000
    environment:
      - PORT=3000
### Node.js
const port = process.env.PORT;
### golang (Не проверено)
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Значение по умолчанию
	}
	
	fmt.Printf("Server is listening on port %s\n", port)
	// Ваш код обработки запросов на указанном порту
}
## restart: unless-stopped
version: '3'

services:
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    command: npm run start
    restart: unless-stopped
    ports:
      - 3000:3000
    environment:
      - PORT=3000
# Second project golang:1.20-alpine
## struct
docker-compose.yml
./api
    main.go
    Dockerfile
## docker-compose.yml
version: '3'

services:
  api:
    build:
      context: ./api
      dockerfile: Dockerfile
    command: ./main
    restart: unless-stopped
    ports:
      - 3000:3000
    environment:
      - PORT=3000
      - HOST=http://realworld-docker.com
## main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Значение по умолчанию
	}
	host := os.Getenv("HOST")
	fmt.Println(host)

	fmt.Printf("Server is listening on port %s\n", port)
	// Ваш код обработки запросов на указанном порту
}
## Dockerfile
FROM golang:1.20-alpine

#Создать папку внутри контейнера
WORKDIR /app

#Скопировать файлы в созданную дирикторию
COPY . .
#COPY go.mod ./
#COPY main.go ./

#Запустить внутри контейнера прогу
#RUN go mod tidy
RUN go mod init app
#Построить и переименовать в main
RUN go build -o main
#Чтоб не копировать прям всё
#COPY . .

#----эти действия берёт на себя docker-compose--------------
#Открыть порт в контейнере
#EXPOSE 3000

#Запустить приложение
#CMD [ "./main" ]
#-----------------------------------------------------------
## docker-compose up --build
# docker logs <project> - выдаёт всю инфо по логам одного проэкта

