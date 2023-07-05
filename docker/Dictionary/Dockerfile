# Используем базовый образ Ubuntu
FROM ubuntu

# Устанавливаем необходимые зависимости
# Качаем необходимый софт
# && apt-get install -y build-essential
RUN apt-get update && apt-get install -y gcc && apt-get install make 

# Копируем файлы проекта внутрь контейнера
# Создаём папку Dictionary в контейнере и копируем туда все файлы
COPY . /Dictionary

# Устанавливаем рабочую директорию
WORKDIR /Dictionary

# Компилируем и запускаем программу
# RUN gcc -o Dictionary Dictionary.c
# Если запустить контейнер с открытием bash то последняя команда не будет выполнена
CMD ["./Dictionary"]
