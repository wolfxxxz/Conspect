# docker for testUsers
## 1 install docker on ubuntu and configure
sudo apt-get update
1:
sudo apt install docker.io
2: автоматический запуск при старте системы
sudo systemctl start docker
sudo systemctl enable docker
3:
docker --version
Docker version 20.10.24, build 297e128
4: docker-compose
sudo curl -L https://github.com/docker/compose/releases/download/2.18.1/docker-compose-`uname -s`-`uname -m` -o/usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
или
sudo curl -L "https://github.com/docker/compose/releases/download/2.18.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
--и тут лажа
## 3 Commands simple
docker -v - version docker
docker --help
docker container --help
docker image --help
docker volume --help
## 4 container
```
docker container --help
docker run --help

sudo docker run hello-world - <sudo> повысить права юзера <docker run> - запустить контейнер <hello-world> - название контейнера
Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/
For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```
## 5 docker run and exit вход и выход из контейнера
### docker pull
sudo docker pull ubuntu - скачать образ ubuntu последней версии
[sudo] пароль для mvmir: 
Using default tag: latest
latest: Pulling from library/ubuntu
dbf6a9befcde: Pull complete 
Digest: sha256:dfd64a3b4296d8c9b62aa3309984f8620b98d87e47492599ee20739e8eb54fbf
Status: Downloaded newer image for ubuntu:latest
docker.io/library/ubuntu:latest
### sudo docker run -i -t ubuntu bash
sudo docker run -i -t ubuntu bash - <docker run> - запустить контейнер, <-i> -вывести всё из контейнера <-t> - принимать все комманты в контейнер <ubuntu> - контейнер, <bash> - программа из контейнера (bash запустится как дефолтная комманда)
root@9519bd31f6cc:/#  - теперь мы внутри контейнера
### exit - выйти из ...
## 6 ls, ls -a, rm
### docker container ls
sudo docker container ls - show working containers
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
### sudo docker container ls -a == docker ps -a
show all downloaded containers
CONTAINER ID   IMAGE         COMMAND    CREATED         STATUS                     PORTS     NAMES
9519bd31f6cc   ubuntu        "bash"     7 minutes ago   Exited (0) 5 minutes ago             admiring_brahmagupta
5e567b7e6064   hello-world   "/hello"   3 hours ago     Exited (0) 3 hours ago               reverent_sinoussi
83e9dccbe66f   hello-world   "/hello"   3 hours ago     Exited (0) 3 hours ago               beautiful_chebyshev
### sudo docker container rm id or name
delete container 
удалить можно только остановленные контейнеры
восстановить контейнер невозможно - удалять нужно осторожно
sudo docker container rm 83e9dccbe66f
sudo docker container rm admiring_brahmagupta
## 7 фоновый запуск и остановка, удаление всех остановленных контейнеров
### -d - фоновый режим
sudo docker run -it -d ubuntu bash
### attach - зайти в фоновый контейнер
sudo docker container attach eeebb78bda74
### <ctrl + p> -> <ctrl + q> - выйти из сеанса в фоновый режим
root@eeebb78bda74:/# read escape sequence
### docker container stop <id> or <name>
sudo docker container stop eeebb78bda74
sudo docker container ls
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
### docker container prune - удалить все контейнеры (остановленные)
sudo docker container prune
WARNING! This will remove all stopped containers.
Are you sure you want to continue? [y/N] y
Deleted Containers:
eeebb78bda74b8cb17a70321ecedea6bc1fcea6b99b308ada79b06f890373d12
Total reclaimed space: 0B
sudo docker container prune -f  <-f> - можно не спрашивать об удалении
## 8 --rm (после использования уничтожить)
sudo docker run -it --name ubuntu_1 --rm ubuntu bash
<--rm> - удалить после остановки
## 9 prune -f (удалить все контейнеры без подтверждения)
1: sudo docker run -it --name ubuntu ubuntu bash
2: echo thirty one > text.txt
3: ls
4: exit
5: sudo docker ps -a
6: sudo docker container start ubuntu
7: sudo docker container attach ubuntu
8: ls
9: cat text.txt
10: exit
11: sudo docker ps
12: sudo docker container prune -f
    Deleted Containers:
    63f55b7f6553089ad3ec524bc80a7dfd4d0fa06f66d29fde5631936dc1a00ce4
    Total reclaimed space: 67B
13: 
## 10 bind Как обмениватся файлами пк с контейнером
По сути с помощью bint контейнер получает доступ к папке на диске пк
*** sudo docker run -it --mount type=bind,src=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind,target=/bind/ ubuntu bash**
1: <sudo docker run -it>
2: <--mount>
3: <type=bind> - может быть bind volume tmpfs
4: <src=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind> - путь к директории на хост машине (смотрим pwd)
5: <target=/bind/> - путь внутри контейнера
6: <ubuntu bash>  - ещё может быть <readonly> - который позволяет только читать
## 11 volume - папка которая не удаляется после удаления контейнера
Папка которая доступна для разных контейнеров но при этом не доступна с пк (такая себе папка вне компа), ещё она не удаляется при остановке или удалении контейнера
docker volume - мануал по volume
1: sudo docker volume create my_volume
2: sudo docker volume ls
3: sudo docker run -it --rm --mount type=volume,src=my_volume,target=/volume/ ubuntu bash
4: echo hi volume! I know you are very interesting fi4a > volume/test.txt
5: exit
6: sudo docker run -it --mount type=volume,src=my_volume,target=/volume/ ubuntu bash
7: cat volume/test.txt
  hi volume! I know you are very interesting fi4a
sudo docker container prune -f
***sudo docker volume rm my_volume** - удалить volume файл с пк
## 12 Примеры bind , volume
### 12.1 запуск программы из папки на компе с помощью контейнера python через bind + readonly
sudo docker pull python - скачать контейнер с пайтон
*** sudo docker run --rm --mount type=bind,src=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind,target=/bind/,readonly python python /bind/counter.py**
```
<sudo docker run --rm --mount type=bind,src=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind,target=/bind/> - смотри 10 главу
<readonly python > - запустить контейнер пайтон в режим readonly
<python /bind/counter.py> - открыть в пайтоне файл counter.py (.п игрик) в нём лежит скрипт: for i in range(1, 6): print(i)
```
### 12.2 volume Создаём скрипт "container ubuntu" запускаем "container python"
1: Создаём скрипт питона в папке volume(которая создаётся при запуске автоматически) для этого запускаем контейнер ubuntu и открываем bash
***sudo docker run -it --rm --mount type=volume,src=python_program,target=/volume/ ubuntu bash**
2: cd volume
3: echo "for i in range(1, 6): print(i)" > counter.py
3: или так > echo "for i in range(1, 6): print(i, end=' ')" > counter.py
4: exit - (--rm удаляет после выхода контейнер)
5: Запускаем программу с помощью контейнера с python
***sudo docker run --rm --mount type=volume,src=python_program,target=/src/,readonly python python /src/counter.py**
1 2 3 4 5 
## 13 образы разных версий программ
1: Тестим код пайтона
sudo docker run --rm --mount trc=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind,target=/bind/,readonly python python /bind/test_version.py
[sudo] пароль для mvmir: - ебаная шляпа
Hello, Ivan
2: скачать контейнер питон версии 3.5
sudo docker pull python:3.5
3: пробуем запустить файл на python:3.5
sudo docker run --rm --mount type=bind,src=/home/mvmir/LearnQA_Docker/LearnQA_Docker/dir_for_bind,target=/bind/,readonly python:3.5 python /bind/test_version.py
3.1: вывод
File "/bind/test_version.py", line 2
    print(f'Hello, {name}')
SyntaxError: invalid syntax
## 14 образ image pull and rm (delete)
1: Скачать образ
sudo docker pull ubuntu:16.04
2: Проверяем скачанные на пк образы контейнеров
sudo docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
python       latest    815c8c75dfc0   2 weeks ago     920MB
ubuntu       latest    3b418d7b466a   3 weeks ago     77.8MB
ubuntu       16.04     b6f507652425   20 months ago   135MB
python       3.5       3687eb5ea744   2 years ago     871MB
3: Проверим запустив контейнер
sudo docker run -it --rm ubuntu:16.04 bash
root@0603a137b85c:/# exit
exit
4: Delete image
sudo docker image rm b6f507652425
5: полезные ссылки
hub.docker.com
6: По умолчанию качается latest version
sudo docker pull ubuntu == ubuntu:latest
## 15 Create image (создаём образ)
1: Создать контейнер от основного образа
sudo docker run -it ubuntu bash
2: Настроить нужный софт (в нашем случае vim)
apt-get update
apt-get install apt-file
apt-file update
apt-get install vim
echo hello vim > test.txt
vim test.txt
3: Выходим из контейнера
exit - (при запуске без --rm)
sudo docker ps -a
CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS                          PORTS     NAMES
59f77556abfc   ubuntu    "bash"    11 minutes ago   Exited (0) About a minute ago             confident_dirac
4: Создаём образ на основе контейнера
sudo docker commit 59f77556abfc ubuntu_with_vim
5: Проверяем
sudo docker image ls
REPOSITORY        TAG       IMAGE ID       CREATED          SIZE
ubuntu_with_vim   latest    378e484d2560   48 seconds ago   350MB
python            latest    815c8c75dfc0   2 weeks ago      920MB
ubuntu            latest    3b418d7b466a   3 weeks ago      77.8MB
python            3.5       3687eb5ea744   2 years ago      871MB
6: Запускаем наш new image
sudo docker run -it --rm ubuntu_with_vim
vim test.txt
## 16 



## Пропали заметки
2: Настроить нужный софт (в нашем случае vim)
apt-get update
apt-get install apt-file
apt-file update
apt-get install vim
echo hello vim > test.txt
vim test.txt