# Instal
## Скачать и установить git
Желательно для своей ОС
## Информация + комманды
cmd
* gitInfo
>git
* Конкретная информация 
>git help
>git help status

## Регистрация
git config --global user.name "wolfxxxz"
git config --global user.email "wolfxxxz@gmail.com"
* Меняем цветовую схему
>git config --global color.ui true

## Создание проэкта "git init"
cmd
Создать папку Slovar
>mkdir Slovar
Зайти в эту папку
>cd Slovar
Инициализировать папку для git
>git init

## git status
>git status
Проверка статуса - информация о текущей ситуации

- Для инициализированной директории
On branch master //на ветке мастер
No commits yet //коментов нет
nothing to commit (create/copy files and use "git add" to track) // нет изминений для сохранения

- Не инициализированной
fatal: not a git repository (or any of the parent directories): .git

# Основы git init, Status, git add, git commit, git log, git reset, git diff, git clean, git checkout

## Статусы хранения
**Просто созданный "untracked"**
Не отслеживается gitom
**Изминения не сохранены "modified"**
Файл создан но не добавлен - еще не было git add .
**Подготовленный "staged"**
В связи с тем что git хранит снимки (копии) только изменений 
- Изменённые файлы нужно пометить командой
>git add file.go file2.exe
**Зафиксированный "committed"**
- После чего сохранить 
>git commit -m"Message"

## git add command
Добавить файлы списком                    - git add file1.exe file2.go file3.java 
Добавить все файлы из текущей папки       - git add .
Добавить все файлы с расширением          - git add *.java
Добавить все файлы из папки с расширением - git add someDir/*.java
Добавить все файлы из папки               - git add someDir/
Добавить все файлы в проэкте с расширением - git add "*.java"

## git log
git log - показывает полную историю комитов с хеш меткой
git log --oneline - короткая история комитов

## Практика "git init", "git status", "git add", "git commit", 
### git init
* cmd
Создать папку Slovar
>mkdir Slovar
Зайти в эту папку
>cd Slovar
Инициализировать папку для git
>git init
### git status
* git bash
```
$ git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   main.go
```
### git add
```
$ git add main.go
```
### git status
On branch master
No commits yet
Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   main.go
### git commit -m "Add file main"
```
$ git commit -m "Add file main"
[master (root-commit) 6d0cb72] Add file main
 1 file changed, 7 insertions(+)
 create mode 100644 main.go
```
### git add b.go Makefile
```
$ git add b.go Makefile
```
### git commit -m "hey"
```
$ git commit -m "hey"
[master 244b0f0] hey
 2 files changed, 11 insertions(+)
 create mode 100644 Makefile
 create mode 100644 b.go
```
### git add main.exe main.go
```
$ git add main.exe main.go
```
### git commit -m "add mainu"
```
$ git commit -m "add mainu"
[master 6d8c76e] add mainu
 2 files changed, 1 insertion(+)
 create mode 100644 main.exe
```
### git status
```
$ git status
On branch master
nothing to commit, working tree clean
```
### git log
```
git log
commit 6d8c76e16a329cd8c2c596d78efc19cc197e3835 (HEAD -> master)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 15:15:02 2023 +0200

    add mainu

commit 244b0f0135d36bf13b8b15b5c7116ae7a91393c8
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 15:13:41 2023 +0200

    hey

commit 6d0cb726a707c156859ba2a9c0aa80d6bbcdbf38
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 14:36:57 2023 +0200

    Add file main
```

## git diff: показывает различия между версиями файлов
### git diff работает до выделения git add в статусе modified
git diff
### работает после git add в статусе staged
git diff --staged - разница между текущим и последним сохранённым 
### git diff COMMIT_ID - все изменения между текущим и указанным 
git diff COMMIT_ID 
COMMIT_ID - Длинный хеш 
### git diff Практика
#### git diff
- Вносим изминения в файл и сохраняем его

- $ git diff
diff --git a/Makefile b/Makefile
index f695377..282634a 100644
--- a/Makefile                     // Изменённый файл
+++ b/Makefile
@@ -5,4 +5,4 @@ build:
        go build main.go b.go

 run_file:
-       ./main                     // Изменения в файле
+       ./main1
#### git diff --staged
//Добавляем в modified
$ git add Makefile
// Смотрим изминения
$ git diff --staged
diff --git a/Makefile b/Makefile
index f695377..282634a 100644
--- a/Makefile
+++ b/Makefile
@@ -5,4 +5,4 @@ build:
        go build main.go b.go

 run_file:
-       ./main
+       ./main1
#### git diff COMMIT_ID (Все изменения с даты хеша)
##### Нужен хеш для этого git log => 9af5e5cca64989fc89f99d6bfe49dd779ca60cb2
```
$ git log
commit 9af5e5cca64989fc89f99d6bfe49dd779ca60cb2 (HEAD -> master)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 16:55:00 2023 +0200

commit 9af5e5cca64989fc89f99d6bfe49dd779ca60cb2 (HEAD -> master)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 16:55:00 2023 +0200

    add1

commit 67268207393e2cc8a944b320b3e9343fdbb3313f
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 16:33:02 2023 +0200

    delete lajja
```
##### git diff 6d0cb726a707c156859ba2a9c0aa80d6bbcdbf38
```
$ git diff 6d0cb726a707c156859ba2a9c0aa80d6bbcdbf38
diff --git a/Makefile b/Makefile
new file mode 100644
index 0000000..282634a
--- /dev/null
+++ b/Makefile
@@ -0,0 +1,8 @@
+run:
+       go run main.go b.go
+
+build:
+       go build main.go b.go
+
+run_file:
+       ./main1
diff --git a/b.go b/b.go
new file mode 100644
index 0000000..9673f70
--- /dev/null
+++ b/b.go
@@ -0,0 +1,3 @@
+package main
+
+var get = "Hello world"
diff --git a/main.exe b/main.exe
new file mode 100644
index 0000000..e503eb6
Binary files /dev/null and b/main.exe differ
diff --git a/main.go b/main.go
index 1f80fce..badd008 100644
--- a/main.go
+++ b/main.go
```

## git reset
### Теория
Имеет три режима в зависимости от радикальности отката
- soft
- mixed
- hard 
git reset [soft | mixed | hard] [ точка отката ]
                режим             может быть хеш или head 
#### git reset - по умолчанию используется mixed HEAD^
Отменить последний git add . (удалить последнее добавление из unstage)
git reset = git reset --mixed HEAD^
#### git reset --mixed переводит в неотслежеваемую зону (unstaged)
Отменить git add .
git reset --mixed
git reset --mixed HEAD^^ - ...
#### git reset HEAD^^ - по умолчанию используется mixed
git reset HEAD^^ = git reset --mixed HEAD^^
#### git reset --hard удаляет без возвратно изминения и на git и на компютере
git reset --hard
git reset --hard HEAD^^ - удалить два последних коммита
#### git reset --soft переводит из commit в отслежуемую зону (added)
git reset --soft
git reset --soft HEAD^^ - ...
### Практика
#### git reset = git reset --mixed HEAD^
- Изменим файл и добавим его в unstage
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
- Проверим статус
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile

- git reset
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset
Unstaged changes after reset:
M       Makefile

- to discard changes in working directory
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")
#### git reset --hard
- Есть два изменения 
* Первое добавлено to unstage
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile
* Второе только на компьютере
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

* git reset --hard -  чистит оба
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset --hard
HEAD is now at 42069aa 1

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
#### git reset --soft HEAD^^ возвращает из commit в add зону
- git add .
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
- git commit -m
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "abra"
[master 178606b] abra
 1 file changed, 1 insertion(+)
- git reset --soft HEAD^^
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset --soft HEAD^^

* возвращает из commit в add зону *
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile
#### git reset --hard HEAD^^
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
6c4304f (HEAD -> master, origin/master) False push
c585ce0 SSH push
8d724e5 firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git reset --hard HEAD^^**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset --hard HEAD^^
HEAD is now at c585ce0 SSH push
**git reset --hard HEAD^^**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset --hard HEAD^^
HEAD is now at 8d724e5 firs push
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
8d724e5 (HEAD -> master) firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend

## git clean - удаляет untracket файлы т.е. файлы тоько созданные
### Теория
**git clean -f** - удаляет untracket файлы т.е. файлы тоько созданные 
Показать untracket файлы
>git clean -n
Удалить
>git clean -f
### Практика
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
        file.txt

nothing added to commit but untracked files present (use "git add" to track)
**git clean -n**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git clean -n
Would remove file.txt
**git clean -f**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git clean -f
Removing file.txt

## git checkout: переключается между ветками, коммитами и версиями отдельных файлов
### Теория 
#### посмотреть ка выглядел проэкт в каkом то снимке в прошлом
- git checkout хеш
- git checkout HEAD^^
- git checkout HEAD~2
При работе с предыдущей версией программы (предыдущий commit) - данные не сохраняются а предыдущая версия остаётся без изминений
- git checkout master - возвращает HEAD на последний commit
#### вернуть файл к предыдущему commit версии
- git checkout -- . **вернуть все файлы в это состояние**
- git checkout -- file **вернуть только file в это состояние**
#### Для отмены не добавленных изминений
- git checkout -- file
- git checkout -- . 
#### Для отмены добавленых изминений
- git reset 
- git checkout -- .
### Практика
#### вернутся к логу git checkout HEAD^^
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
5231e55 (HEAD -> master) 5
42069aa 1
6726820 delete lajja
85a660f "added"
6d8c76e add mainu
244b0f0 hey
6d0cb72 Add file main

**git checkout HEAD^^**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout HEAD^^
Note: switching to 'HEAD^'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c <new-branch-name>
Or undo this operation with:

  git switch -
Turn off this advice by setting config variable advice.detachedHead to false
HEAD is now at 42069aa 1

**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
42069aa (HEAD) 1
6726820 delete lajja
85a660f "added"
6d8c76e add mainu
244b0f0 hey
6d0cb72 Add file main

**git checkout master**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout master
Previous HEAD position was 42069aa 1
Switched to branch 'master'

**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
5231e55 (HEAD -> master) 5
42069aa 1
6726820 delete lajja
85a660f "added"
6d8c76e add mainu
244b0f0 hey
6d0cb72 Add file main
#### Перейти к сохранению и вернуть файл к этому значению "git checkout -- ."
##### вернуть один файл из лога
**git log**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log
commit 5231e55bd6a065819ed02108d7198dfce5eaa0ed (HEAD -> master)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Mon Mar 13 09:24:21 2023 +0200

    5

commit 42069aa42acb1c925a3d9317e3591d6caf3bcdd4
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 10 09:38:58 2023 +0200

    1

commit 67268207393e2cc8a944b320b3e9343fdbb3313f
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 16:33:02 2023 +0200

    delete lajja

commit 85a660f5e600a59da8feab586e0f8f097d996fd2
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 16:23:59 2023 +0200

    "added"

commit 6d8c76e16a329cd8c2c596d78efc19cc197e3835
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 15:15:02 2023 +0200

    add mainu

commit 244b0f0135d36bf13b8b15b5c7116ae7a91393c8
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 9 15:13:41 2023 +0200

    hey

commit 6d0cb726a707c156859ba2a9c0aa80d6bbcdbf38
Author: wolfxxxz <wolfxxxz@gmail.com>
**git checkout 67268207393e2cc8a944b320b3e9343fdbb3313f -- Makefile**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout 67268207393e2cc8a944b320b3e9343fdbb3313f -- Makefile

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile

**git commit -m "return Makefile to its original stage"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "return Makefile to its original stage"
[master 27e4940] return Makefile to its original stage
 1 file changed, 3 deletions(-)
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
##### вернуть все файлы из лога
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
**git checkout 244b0f0135d36bf13b8b15b5c7116ae7a91393c8 -- .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout 244b0f0135d36bf13b8b15b5c7116ae7a91393c8 -- .
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile
        modified:   main.go

**git commit -m "return All files to its original stage"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "return All files to its original stage"
[master eadb340] return All files to its original stage
 2 files changed, 1 insertion(+), 2 deletions(-)
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
#### Вернуть последние не добавленные изминения
##### git checkout -- .
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   main.go

no changes added to commit (use "git add" and/or "git commit -a")

**git checkout -- .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout -- .

**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>
##### git checkout -- Makefile
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout -- Makefile

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
#### git reset + git checkout -- . = удалить add снимки
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   Makefile

**git reset**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git reset
Unstaged changes after reset:
M       Makefile
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")
**git checkout -- .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout -- .
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean

## git commit
### git commit -m "message"
git add .
git commit -m "message"
### git commit -a -m "message"
Добавить все изменённые файлы в commit пропустив при этом git add .
!!!новые файлы не добавляет, только изменённые!!!
-a = git add .
### git commit --amend -m "message"
- Сохранить изминения в последний commite и изменить message
git commit --amend -m "message"
- Сообщение менять не обязательно тогда 
git commit --amend
:wq - выйти из редактора сообщений
### Практика
#### git commit -a -m "message" *commit в одну комманду*
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")
**git commit -a -m "git commit -a -m"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -a -m "git commit -a -m"
[master 48235a4] git commit -a -m
 1 file changed, 3 insertions(+), 1 deletion(-)
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
#### git commit --amend -m "message" *commit с сейвом в старый commit*
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   Makefile

no changes added to commit (use "git add" and/or "git commit -a")
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
48235a4 (HEAD -> master) git commit -a -m
eadb340 return All files to its original stage
27e4940 return Makefile to its original stage
5231e55 5
42069aa 1
6726820 delete lajja
85a660f "added"
6d8c76e add mainu
244b0f0 hey
6d0cb72 Add file main
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
**git commit --amend -m "git commit --amend"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit --amend -m "git commit --amend"
[master 349c5e9] git commit --amend
 Date: Mon Mar 13 11:42:10 2023 +0200
 1 file changed, 4 insertions(+), 1 deletion(-)
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
349c5e9 (HEAD -> master) git commit --amend
eadb340 return All files to its original stage
27e4940 return Makefile to its original stage
5231e55 5
42069aa 1
6726820 delete lajja
85a660f "added"
6d8c76e add mainu
244b0f0 hey
6d0cb72 Add file main



# Совместная Работа с удалённым репозиторием git remote add, git push, git pull, SSH, git clone

## git remote add *привязать проэкт к ссылке*
### Теория
**git remote -v** - просмотр списка репозиториев
**git remote add <название адресс>** - добавить новый репозиторий (привязать проэкт к ссылке)
>git remote add origin https://github.com/Wolfxxxz/Some-Name.git
origin - название репозитория (любое)
**git remote remove** - удалить репозиторий
### Практика
#### git remote add origin https://github.com/wolfxxxz/FirstRepository.git** - привязать проэкт к ссылке
- Создать репозиторий на github
- Скопировать ссылку
**git remote add origin https://github.com/wolfxxxz/FirstRepository.git** - привязать проэкт к ссылке
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote add origin https://github.com/wolfxxxz/FirstRepository.git
**git remote -v** - просмотр возможностей
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote -v
origin  https://github.com/wolfxxxz/FirstRepository.git (fetch)
origin  https://github.com/wolfxxxz/FirstRepository.git (push)
**git push origin master** - первый push
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git push origin master
info: please complete authentication in your browser...
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 8 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (9/9), 709 bytes | 354.00 KiB/s, done.
Total 9 (delta 2), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (2/2), done.
To https://github.com/wolfxxxz/FirstRepository.git
 * [new branch]      master -> master

## git push отправка commit на ветку
### Теория
>git push ИМЯ ВЕТКА
>git push origin master
### Практика
**git commit -a -m "firs push"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -a -m "firs push"
[master 8d724e5] firs push
 1 file changed, 5 insertions(+), 2 deletions(-)
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch master
nothing to commit, working tree clean
**git push origin master**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git push origin master
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 332 bytes | 332.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/wolfxxxz/FirstRepository.git
   0de16ec..8d724e5  master -> master

## git pull синхронизировать файлы с удалённого репозитория
### Теория
git pull - синхронизировать актуальные версии файлов со всеми commit 
- создать папку
- cmd
- привязать папку к ссылке
**git remote add origin https://github.com/wolfxxxz/FirstRepository.git**
D:\Onedrive\Рабочий стол\second>git remote add origin https://github.com/wolfxxxz/FirstRepository.git
- синхронизировать файлы с ветки master
**git pull origin master**
D:\Onedrive\Рабочий стол\second>git pull origin master
remote: Enumerating objects: 12, done.
remote: Counting objects: 100% (12/12), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 12 (delta 2), reused 12 (delta 2), pack-reused 0
Unpacking objects: 100% (12/12), 989 bytes | 0 bytes/s, done.
From https://github.com/wolfxxxz/FirstRepository
 * branch            master     -> FETCH_HEAD
 * [new branch]      master     -> origin/master

**все файлы с локального сервера синхронизированы**
**git log --oneline**
D:\Onedrive\Рабочий стол\second>git log --oneline
8d724e5 (HEAD -> master, origin/master) firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**логи читаются**

## SSH настройка пароля
### Теория 
SSH - Secure SHell - безопасная оболочка
### Настройка SSH
- https://docs.github.com/en/authentication/connecting-to-github-with-ssh/checking-for-existing-ssh-keys
> git bash
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/second (master)
$ ls -al ~/.ssh
bash: $'\302\226\302\226\302\226\302\226\302\226\302\226\302\202ls': command not found

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/second (master)
$ ssh-keygen -t ed25519 -C "wolfxxxz@gmail.com"
Generating public/private ed25519 key pair.
Enter file in which to save the key (/c/Users/Mvmir/.ssh/id_ed25519):
Created directory '/c/Users/Mvmir/.ssh'.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /c/Users/Mvmir/.ssh/id_ed25519
Your public key has been saved in /c/Users/Mvmir/.ssh/id_ed25519.pub
The key fingerprint is:
SHA256:M5+BUBsZIs0WZox2QGwGbWhKwwxMFDRHZ/tR3zKCf4o wolfxxxz@gmail.com
The key's randomart image is:
+--[ED25519 256]--+
|X*+O+O=.++       |
| *= @+Bo+o. .    |
|.o.= +.o.. + .   |
|.     ..o.. o    |
|       .S...     |
|        .+oo     |
|       E .o      |
|                 |
|                 |
+----[SHA256]-----+

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/second (master)
$ eval "$(ssh-agent -s)"
Agent pid 684

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/second (master)
$ ssh-add ~/.ssh/id_ed25519
Identity added: /c/Users/Mvmir/.ssh/id_ed25519 (wolfxxxz@gmail.com)

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/second (master)
$ clip < ~/.ssh/id_ed25519.pub
**Изменить адресс в cmd с http на ssh**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote -v
origin  https://github.com/wolfxxxz/FirstRepository.git (fetch)
origin  https://github.com/wolfxxxz/FirstRepository.git (push)

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote remove origin

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote -v
**git remote add origin git@github.com:wolfxxxz/FirstRepository.git**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote add origin git@github.com:wolfxxxz/FirstRepository.git

C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git remote -v
origin  git@github.com:wolfxxxz/FirstRepository.git (fetch)
origin  git@github.com:wolfxxxz/FirstRepository.git (push)
### Проверка
**>git commit -a -m "False push"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -a -m "False push"
[master 6c4304f] False push
 1 file changed, 1 insertion(+)
**git push origin master**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git push origin master
The authenticity of host 'github.com (140.82.121.3)' can't be established.
ED25519 key fingerprint is SHA256:+DiY3wvvV6TuJJhbpZisF/zLDA0zPMSvHdkr4UvCOqU.
This key is not known by any other names.
**yes**
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added 'github.com' (ED25519) to the list of known hosts.
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 305 bytes | 305.00 KiB/s, done.
Total 3 (delta 1), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To github.com:wolfxxxz/FirstRepository.git
   c585ce0..6c4304f  master -> master
**D:\Onedrive\Рабочий стол\second**
**git pull origin master**
D:\Onedrive\Рабочий стол\second>git pull origin master
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (1/1), done.
remote: Total 3 (delta 1), reused 3 (delta 1), pack-reused 0
Unpacking objects: 100% (3/3), 285 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/FirstRepository
 * branch            master     -> FETCH_HEAD
 * [new branch]      master     -> origin/master
Updating c585ce0..6c4304f
Fast-forward
 main.go | 1 +
 1 file changed, 1 insertion(+)

**git log --oneline**
D:\Onedrive\Рабочий стол\second>git log --oneline
6c4304f (HEAD -> master, origin/master) **False push**
c585ce0 SSH push
8d724e5 firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend

## git clone
### Теория 
git clone <ссылка> = git remote add <ссылка> + git pull
### Практика
**git clone git@github.com:wolfxxxz/FirstRepository.git**
D:\Onedrive\Рабочий стол\remoute repository clone>git clone git@github.com:wolfxxxz/FirstRepository.git
Cloning into 'FirstRepository'...
remote: Enumerating objects: 18, done.
remote: Counting objects: 100% (18/18), done.
remote: Compressing objects: 100% (8/8), done.
remote: Total 18 (delta 4), reused 18 (delta 4), pack-reused 0
Receiving objects: 100% (18/18), done.
Resolving deltas: 100% (4/4), done.
**cd firstrepository**
D:\Onedrive\Рабочий стол\remoute repository clone>cd firstrepository
**git log --oneline**
D:\Onedrive\Рабочий стол\remoute repository clone\FirstRepository>git log --oneline
6c4304f (HEAD -> master, origin/master, origin/HEAD) False push
c585ce0 SSH push
8d724e5 firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git remote -v**
D:\Onedrive\Рабочий стол\remoute repository clone\FirstRepository>git remote -v
origin  git@github.com:wolfxxxz/FirstRepository.git (fetch)
origin  git@github.com:wolfxxxz/FirstRepository.git (push)

# Ветвление (Branching)
## git branch (create, schow, delete) Fast - Forward merge, ort strategy
### Теория
Ветвление - создание паралельных версий (бета версий) не изменяя при этом основную рабочуюю версию
git branch <name> - создание новой ветки
git branch - показывает название ветки HEAD
git branch -d  - удаление ветки
git branch -D  - удаление ветки
git push --delete <origin newLine1> - удаление удалённой ветки
git checkout <name> - переключение между ветками
git merge - сливает одну ветку с другой
git branch -r - показывает доступные удалённые ветки
Fast - Forward merge - при этом ветка мастер не изменяется в процессе редактирования паралельной ветки
       если в ветке мастер произойдут изминения то появятся ошибки
Recursive merge - если изминения происходят в обеих ветках (у меня получился ort)
### Практика
#### Fast - Forward merge с одним commit
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
8d724e5 (HEAD -> master) firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git branch newLine** - *Создание*
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch newLine
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
* master
  newLine
**git checkout newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout newLine
Switched to branch 'newLine'
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
  master
* newLine
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch newLine
Untracked files:
  (use "git add <file>..." to include in what will be committed)
        newDoc.txt

nothing added to commit but untracked files present (use "git add" to track)
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
**git commit -m "Hello newLine"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "Hello newLine"
[newLine 7988218] Hello newLine
 1 file changed, 1 insertion(+)
 create mode 100644 newDoc.txt
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
**git status**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git status
On branch newLine
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        new file:   newDoc2.txt

**git commit -m "i need more branches"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "i need more branches"
[newLine 8f4ff28] i need more branches
 1 file changed, 2 insertions(+)
 create mode 100644 newDoc2.txt
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
8f4ff28 (HEAD -> newLine) i need more branches
7988218 Hello newLine
8d724e5 (master) firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git checkout master**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout master
Switched to branch 'master'
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
8d724e5 (HEAD -> master) firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git merge newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git merge newLine
Updating 8d724e5..8f4ff28
**Fast-forward**
 newDoc.txt  | 1 +
 newDoc2.txt | 2 ++
 2 files changed, 3 insertions(+)
 create mode 100644 newDoc.txt
 create mode 100644 newDoc2.txt
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
8f4ff28 (HEAD -> master, newLine) i need more branches
7988218 Hello newLine
8d724e5 firs push
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git branch -d newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch -d newLine
Deleted branch newLine (was 8f4ff28).
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
* master
#### ort strategy
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
0de16ec (HEAD -> master) amendation
e757718 ammend
d8d9e5b git commit --amend
**git brash newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git brash newLine
git: 'brash' is not a git command. See 'git --help'.

The most similar command is
        branch
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
0de16ec (HEAD -> master) amendation
e757718 ammend
d8d9e5b git commit --amend
**git branch newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch newLine
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
* master
  newLine
**git checkout newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout newLine
Switched to branch 'newLine'
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
  master
* newLine
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
**git commit -m "model"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "model"
[newLine 398d38f] model
 1 file changed, 5 insertions(+)
 create mode 100644 model.go
**git checkout master**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git checkout master
Switched to branch 'master'
**git add .**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git add .
**git commit -m "some delete"**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git commit -m "some delete"
[master e357730] some delete
 1 file changed, 2 deletions(-)
**git merge newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git merge newLine
**Merge made by the 'ort' strategy.**
 model.go | 5 +++++
 1 file changed, 5 insertions(+)
 create mode 100644 model.go
**git log --oneline**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git log --oneline
2af6a79 (HEAD -> master) Merge branch 'newLine'
e357730 some delete
398d38f (newLine) model
0de16ec amendation
e757718 ammend
d8d9e5b git commit --amend
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
* master
  newLine
**git branch -d newLine**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch -d newLine
Deleted branch newLine (was 398d38f).
**git branch**
C:\Users\Mvmir\go\src\github.com\Wolfxxxz\Slovar>git branch
* master

## git pull and git push удалённо
### Теория git fetch, git merge original/master
git fetch - скачивает удалённые ветки с репозитория не делая слияния
### Практика
#### Fast - Forward original master
##### Создать репозиторий на сайте github и запушить два commit
D:\Onedrive\Рабочий стол\newLine>git init
Initialized empty Git repository in D:/Onedrive/Рабочий стол/newLine/.git/

D:\Onedrive\Рабочий стол\newLine>git remote add origin git@github.com:wolfxxxz/newLine.git

D:\Onedrive\Рабочий стол\newLine>git remote -v
origin  git@github.com:wolfxxxz/newLine.git (fetch)
origin  git@github.com:wolfxxxz/newLine.git (push)

D:\Onedrive\Рабочий стол\newLine>git add .

D:\Onedrive\Рабочий стол\newLine>git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   file1.txt


D:\Onedrive\Рабочий стол\newLine>git commit -m "first commit"
[master (root-commit) f6cf9b8] first commit
 1 file changed, 1 insertion(+)
 create mode 100644 file1.txt

D:\Onedrive\Рабочий стол\newLine>git push origin master
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Writing objects: 100% (3/3), 214 bytes | 214.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
 * [new branch]      master -> master

D:\Onedrive\Рабочий стол\newLine>git add .

D:\Onedrive\Рабочий стол\newLine>git commit -m "file2 commit"
[master 11148d3] file2 commit
 1 file changed, 2 insertions(+)
 create mode 100644 file 2.txt

D:\Onedrive\Рабочий стол\newLine>git push origin master
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 279 bytes | 279.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   f6cf9b8..11148d3  master -> master

D:\Onedrive\Рабочий стол\newLine>git log --oneline
11148d3 (HEAD -> master, origin/master) file2 commit
f6cf9b8 first commit
##### User1 and User2 git clone url
**git clone git@github.com:wolfxxxz/newLine.git**
D:\Onedrive\Рабочий стол\us2>git clone git@github.com:wolfxxxz/newLine.git
Cloning into 'newLine'...
remote: Enumerating objects: 6, done.
remote: Counting objects: 100% (6/6), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 6 (delta 0), reused 6 (delta 0), pack-reused 0
Receiving objects: 100% (6/6), done.
**cd newLine**
D:\Onedrive\Рабочий стол\us2>cd newLine
##### User2 change file2
D:\Onedrive\Рабочий стол\us2\newLine>**git status**
On branch master
Your branch is up to date with 'origin/master'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   file 2.txt

no changes added to commit (use "git add" and/or "git commit -a")

D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "user2 add string3"**
[master 56e3138] user2 add string3
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>git status
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 303 bytes | 303.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   11148d3..56e3138  master -> master
##### User1 обновляет данные git pool
D:\Onedrive\Рабочий стол\us1\newLine>**git pull origin master**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 283 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            master     -> FETCH_HEAD
   11148d3..56e3138  master     -> origin/master
Updating 11148d3..56e3138
Fast-forward
 file 2.txt | 3 ++-
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
56e3138 (HEAD -> master, origin/master, origin/HEAD) user2 add string3
11148d3 file2 commit
f6cf9b8 first commit
#### Merge made by the 'ort' strategy.
##### user1 changed file1 and push
D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "user1 chenge file1"**
[master fb269ee] user1 chenge file1
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin master**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 289 bytes | 289.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   56e3138..fb269ee  master -> master
##### user2 changed file2 and push ***'ort' strategy***
D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "user2 change file2"**
[master f479183] user2 change file2
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
To github.com:wolfxxxz/newLine.git
 ! [rejected]        master -> master (fetch first)
error: failed to push some refs to 'github.com:wolfxxxz/newLine.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
**Эта ошибка говорит о том что были изминения и сразу их нужно скачать**
D:\Onedrive\Рабочий стол\us2\newLine>**git pull origin master**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 269 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            master     -> FETCH_HEAD
   56e3138..fb269ee  master     -> origin/master
**Merge made by the 'ort' strategy.**
 file1.txt | 3 ++-
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
Enumerating objects: 9, done.
Counting objects: 100% (8/8), done.
Delta compression using up to 8 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 591 bytes | 591.00 KiB/s, done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   fb269ee..f090e8d  master -> master
##### user1 git pull
D:\Onedrive\Рабочий стол\us1\newLine>**git pull origin master**
remote: Enumerating objects: 9, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 5 (delta 0), reused 5 (delta 0), pack-reused 0
Unpacking objects: 100% (5/5), 571 bytes | 0 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            master     -> FETCH_HEAD
   fb269ee..f090e8d  master     -> origin/master
Updating fb269ee..f090e8d
**Fast-forward**
 file 2.txt | 3 ++-
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
f090e8d (HEAD -> master, origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine
f479183 user2 change file2
fb269ee user1 chenge file1
56e3138 user2 add string3
11148d3 file2 commit
f6cf9b8 first commit

## конфликты слияния
### Теория
Конфликты не случаются если
-  Fast - Forward - комиты и пуши происходят по очереди со стороны пользователей
-  Рекурсия ('ort' strategy) но файлы при этом изменены разные без разрыва связей между ними

### Практика
#### Конфликт при слиянии веток
##### Провоцируем конфликт
D:\Onedrive\Рабочий стол\us2\newLine>**git branch line2**
D:\Onedrive\Рабочий стол\us2\newLine>**git branch
  line2**
* master
D:\Onedrive\Рабочий стол\us2\newLine>**git checkout line2**
Switched to branch 'line2'
D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* line2
  master
D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "line2 change"**
[line2 e84cc24] line2 change
 1 file changed, 2 insertions(+), 1 deletion(-)
D:\Onedrive\Рабочий стол\us2\newLine>**git checkout master**
Switched to branch 'master'
Your branch is up to date with 'origin/master'.
D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "newLine changed Master"**
[master 5d2d634] newLine changed Master
 1 file changed, 2 insertions(+), 1 deletion(-)
D:\Onedrive\Рабочий стол\us2\newLine>**git log --oneline**
5d2d634 (HEAD -> master) newLine changed Master
f090e8d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine
f479183 user2 change file2
fb269ee user1 chenge file1
56e3138 user2 add string3
11148d3 file2 commit
f6cf9b8 first commit
D:\Onedrive\Рабочий стол\us2\newLine>**git checkout line2**
Switched to branch 'line2'
D:\Onedrive\Рабочий стол\us2\newLine>**git log --oneline**
e84cc24 (HEAD -> line2) line2 change
f090e8d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine
f479183 user2 change file2
fb269ee user1 chenge file1
56e3138 user2 add string3
11148d3 file2 commit
f6cf9b8 first commit
##### Переход на мастер и начало слияния
D:\Onedrive\Рабочий стол\us2\newLine>**git checkout master**
Switched to branch 'master'
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)
***Слияние git merge line2*** 
D:\Onedrive\Рабочий стол\us2\newLine>**git merge line2**
Auto-merging file1.txt
CONFLICT (content): Merge conflict in file1.txt
**Automatic merge failed; fix conflicts and then commit the result.**
D:\Onedrive\Рабочий стол\us2\newLine>**git status**
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)
You have unmerged paths.
  (fix conflicts and run "git commit")
  (use "git merge --abort" to abort the merge)
Unmerged paths:
  (use "git add <file>..." to mark resolution)
        both modified:   file1.txt
no changes added to commit (use "git add" and/or "git commit -a")
##### Идём в file.txt что б исправить ошибку
// Hello!!!
// im user1

```<<<<<<< HEAD
line1 changed
=======
line2 changed
>>>>>>> line2```
```
**нужно изменить файл в котором произошёл конфликт**
Hello!!!
im user1
line1 changed
**удалив всё лишнее**
##### Закончить слияние вручную
D:\Onedrive\Рабочий стол\us2\newLine>**git add .**

D:\Onedrive\Рабочий стол\us2\newLine>**git status**
On branch master
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

All conflicts fixed but you are still merging.
  (use "git commit" to conclude merge)


D:\Onedrive\Рабочий стол\us2\newLine>**git commit**
[master 34c9644] Merge branch 'line2'
                                         **:wq**
D:\Onedrive\Рабочий стол\us2\newLine>**git log --oneline**
34c9644 (HEAD -> master) Merge branch 'line2'
5d2d634 newLine changed Master
e84cc24 (line2) line2 change
f090e8d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine
f479183 user2 change file2
fb269ee user1 chenge file1
56e3138 user2 add string3
11148d3 file2 commit
f6cf9b8 first commit
D:\Onedrive\Рабочий стол\us2\newLine>**git branch -d line2**
Deleted branch line2 (was e84cc24).

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* master

#### Конфликт на удалённом репозитории
##### User2 git push
D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* master

D:\Onedrive\Рабочий стол\us2\newLine>**git log --oneline**
34c9644 (HEAD -> master, origin/master, origin/HEAD) Merge branch 'line2'
5d2d634 newLine changed Master
e84cc24 line2 change
f090e8d Merge branch 'master' of github.com:wolfxxxz/newLine
f479183 user2 change file2
fb269ee user1 chenge file1
56e3138 user2 add string3
11148d3 file2 commit
f6cf9b8 first commitgit push origin newLine1*

D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "user2 line3"**
[master 6098931] user2 line3
 1 file changed, 1 insertion(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 295 bytes | 295.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   34c9644..6098931  master -> master
##### User1 git push
D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "user1 line3"**
[master d53536a] user1 line3
 1 file changed, 1 insertion(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin master**
To github.com:wolfxxxz/newLine.git
 ! [rejected]        master -> master (fetch first)
**error: failed to push some refs to 'github.com:wolfxxxz/newLine.git'**
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
**hint: (e.g., 'git pull ...') before pushing again.**
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

D:\Onedrive\Рабочий стол\us1\newLine>**git pull origin master**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 275 bytes | 0 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            master     -> FETCH_HEAD
   34c9644..6098931  master     -> origin/master
Auto-merging file1.txt
CONFLICT (content): Merge conflict in file1.txt
**Automatic merge failed; fix conflicts and then commit the result.**
**Исправляем ошибки как в примере с ветками**
D:\Onedrive\Рабочий стол\us1\newLine>**git status**
On branch master
Your branch and 'origin/master' have diverged,
and have 1 and 1 different commits each, respectively.
  (use "git pull" to merge the remote branch into yours)

You have unmerged paths.
  (fix conflicts and run "git commit")
  (use "git merge --abort" to abort the merge)

Unmerged paths:git push origin newLine1*
  (use "git add <file>..." to mark resolution)
        both modified:   file1.txt

no changes added to commit (use "git add" and/or "git commit -a")

D:\Onedrive\Рабочий стол\us1\newLine>**git add .**

D:\Onedrive\Рабочий стол\us1\newLine>**git commit**
[master eaef84d] Merge branch 'master' of github.com:wolfxxxz/newLine
                                      **:wq**
D:\Onedrive\Рабочий стол\us1\newLine>git status
On branch master
Your branch is ahead of 'origin/master' by 2 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin master**
Enumerating objects: 8, done.
Counting objects: 100% (8/8), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 502 bytes | 502.00 KiB/s, done.
Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   6098931..eaef84d  master -> master
##### User2 git pull
D:\Onedrive\Рабочий стол\us2\newLine>**git pull origin master**
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 4 (delta 0), reused 4 (delta 0), pack-reused 0
Unpacking objects: 100% (4/4), 482 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            master     -> FETCH_HEAD
   6098931..eaef84d  master     -> origin/master
Updating 6098931..eaef84d
**Fast-forward**

## Удалённые ветки (ветки на сервере)
### Теория
git branch -r - показывает доступные удалённые ветки
git remote show <origin> - показывает инфо про все ветки и их актуальность
git push --delete <origin> <newLine1> - удаление удалённой ветки
### Практика
#### Добавление ветки на сервер и подключение к ней второго пользователя
##### User1
D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master

D:\Onedrive\Рабочий стол\us1\newLine>**git branch newLine1**

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master
  newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git checkout newLine1**
Switched to branch 'newLine1'

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
  master
* newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin newLine1**
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0
remote:
remote: Create a pull request for 'newLine1' on GitHub by visiting:
remote:      https://github.com/wolfxxxz/newLine/pull/new/newLine1
remote:
To github.com:wolfxxxz/newLine.git
 * [new branch]      newLine1 -> newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "line4 user1"**
[newLine1 6801ff6] line4 user1
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin newLine1**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 314 bytes | 314.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   eaef84d..6801ff6  newLine1 -> newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
  origin/newLine1
##### User2
D:\Onedrive\Рабочий стол\us2\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
**git pull origin newLine1 - скачает только ветку newLine1**
**git pull - загрузит всю информацию по "дереву"**

D:\Onedrive\Рабочий стол\us2\newLine>**git pull origin newLine1**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 294 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
 * branch            newLine1   -> FETCH_HEAD
 * [new branch]      newLine1   -> origin/newLine1
Updating eaef84d..6801ff6
Fast-forward
 file1.txt | 3 ++-
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git pull**
Already up to date.

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* master
  newLine1

D:\Onedrive\Рабочий стол\us2\newLine>**git checkout newLine1**
Switched to branch 'newLine1'

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
  master
* newLine1

D:\Onedrive\Рабочий стол\us2\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
  origin/newLine1

#### git remote show <origin>
D:\Onedrive\Рабочий стол\us2\newLine>**git remote show origin**
* remote origin
  Fetch URL: git@github.com:wolfxxxz/newLine.git
  Push  URL: git@github.com:wolfxxxz/newLine.git
  HEAD branch: master
  Remote branches:
    master   tracked
    newLine1 tracked
  Local branch configured for 'git pull':
    master merges with remote master
  Local refs configured for 'git push':
    master   pushes to master   (fast-forwardable)
    newLine1 pushes to newLine1 (up to date)

#### Удаление удалённой ветки
##### User1 - удаление ветки с сервера
D:\Onedrive\Рабочий стол\us1\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
  origin/newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git push --delete origin newLine1**
To github.com:wolfxxxz/newLine.git
 - [deleted]         newLine1

D:\Onedrive\Рабочий стол\us1\newLine>git checkout master
Switched to branch 'master'
Your branch is up to date with 'origin/master'.

D:\Onedrive\Рабочий стол\us1\newLine>**git branch -d newLine1**
**error: The branch 'newLine1' is not fully merged.**
If you are sure you want to delete it, run 'git branch -D newLine1'.

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master
  newLine1

D:\Onedrive\Рабочий стол\us1\newLine>**git branch -D newLine1**
Deleted branch newLine1 (was 6801ff6).

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master

D:\Onedrive\Рабочий стол\us1\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
##### User2 - удаление локальной копии
D:\Onedrive\Рабочий стол\us2\newLine>**git pull**
There is no tracking information for the current branch.
Please specify which branch you want to merge with.
See git-pull(1) for details.

    git pull <remote> <branch>

If you wish to set tracking information for this branch you can do so with:

    git branch --set-upstream-to=origin/<branch> newLine1


D:\Onedrive\Рабочий стол\us2\newLine>**git branch -r**
  origin/HEAD -> origin/master
  origin/master
  origin/newLine1

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
  master
* newLine1

D:\Onedrive\Рабочий стол\us2\newLine>**git checkout master**
Switched to branch 'master'
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* master
  newLine1

D:\Onedrive\Рабочий стол\us2\newLine>**git branch -D newLine1**
Deleted branch newLine1 (was 6801ff6).

D:\Onedrive\Рабочий стол\us2\newLine>**git branch**
* master

## git rebase - сдвигает момент создания ветки на последний комит ветки
### Теория git rebase <nameBrash>
git rebese - сдвигает момент создания ветки на последний комит ветки мастер
             за счёт этого получается слияние типа fast forward
            *команда запускается из ветки, комиты в которой нужно сдвинуть*
git fetch - скачать обновление с сервера но не применять
if conflict == true {
   git rebase --continue 
   git rebase --skip
   git rebase --abort
   }
### Практика
#### локально без конфликтов 
D:\Onedrive\Рабочий стол\us1\newLine>**git branch newLine**

D:\Onedrive\Рабочий стол\us1\newLine>**git checkout newLine**
Switched to branch 'newLine'

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
  master
* newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "create new func newLine file2"**
[newLine 8eac49e] create new func file2
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git checkout master**
Switched to branch 'master'
Your branch is up to date with 'origin/master'.

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master
  newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "repair some bag master file1"**
[master e5d0ab0] repair some bag master file1
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git checkout newLine**
Switched to branch 'newLine'

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
8eac49e (HEAD -> newLine) create new func file2
eaef84d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git rebase master**
Successfully rebased and updated refs/heads/newLine.

**Ощущение такое что ветку newLine создали после создания комита в ветке мастер**
D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
74c4c39 (HEAD -> newLine) create new func newLine file2
**e5d0ab0 (master) repair some bag master file1**
eaef84d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git checkout master**
Switched to branch 'master'
Your branch is ahead of 'origin/master' by 1 commit.
  (use "git push" to publish your local commits)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
e5d0ab0 (HEAD -> master) repair some bag master file1
eaef84d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git merge newLine**
Updating e5d0ab0..74c4c39
**Fast-forward**
 file 2.txt | 3 ++-
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
74c4c39 (HEAD -> master, newLine) create new func file2
e5d0ab0 repair some bag master file1
eaef84d (origin/master, origin/HEAD) Merge branch 'master' of github.com:wolfxxxz/newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git branch**
* master
  newLine

D:\Onedrive\Рабочий стол\us1\newLine>**git branch -d newLine**
Deleted branch newLine (was 74c4c39).
#### Через сервер 2 users без конфликта
##### User1
D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "user1 file1 origin master"**
[master 4b89a81] user1 file1 origin master
 1 file changed, 2 insertions(+)

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin master**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 334 bytes | 334.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   f94d5a3..4b89a81  master -> master
##### User2 git fetch, git rebase
D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "user2 file2 origin master"**
[master 9e6e09c] user2 file2 origin master
 1 file changed, 3 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
To github.com:wolfxxxz/newLine.git
 ! [rejected]        master -> master (fetch first)
error: failed to push some refs to 'github.com:wolfxxxz/newLine.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

D:\Onedrive\Рабочий стол\us2\newLine>**git fetch**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 314 bytes | 0 bytes/s, done.
From github.com:wolfxxxz/newLine
   f94d5a3..4b89a81  master     -> origin/master

D:\Onedrive\Рабочий стол\us2\newLine>**git rebase origin/master**
Successfully rebased and updated refs/heads/master.

D:\Onedrive\Рабочий стол\us2\newLine>**git log --oneline**
**fb5454a (HEAD -> master) user2 file2 origin master** - **commit просто занял следующее место**
4b89a81 (origin/master, origin/HEAD) user1 file1 origin master
f94d5a3 line5 user2
6abf5d6 repair
c99987d line5 master
eaef84d Merge branch 'master' of github.com:wolfxxxz/newLine


D:\Onedrive\Рабочий стол\us2\newLine>git push origin master
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 351 bytes | 351.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   4b89a81..fb5454a  master -> master
##### User1
D:\Onedrive\Рабочий стол\us1\newLine>**git pull**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 331 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
   4b89a81..fb5454a  master     -> origin/master
Updating 4b89a81..fb5454a
Fast-forward
 file 2.txt | 4 +++-
 1 file changed, 3 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
fb5454a (HEAD -> master, origin/master, origin/HEAD) user2 file2 origin master
4b89a81 user1 file1 origin master
f94d5a3 line5 user2
6abf5d6 repair
c99987d line5 master
eaef84d Merge branch 'master' of github.com:wolfxxxz/newLine
#### через сервер с конфликтом 2 users
##### User1 file1
D:\Onedrive\Рабочий стол\us1\newLine>**git pull**
Already up to date.

D:\Onedrive\Рабочий стол\us1\newLine>**git commit -a -m "line5 user1 file1 master"**
[master c99987d] line5 user1 file1 master
 1 file changed, 3 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git push origin master**
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 8 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 319 bytes | 319.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:wolfxxxz/newLine.git
   eaef84d..c99987d  master -> master
##### User2 file1
D:\Onedrive\Рабочий стол\us2\newLine>**git commit -a -m "line5 user2 file1 master"**
[master af44015] line5 user2 file2 master
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
To github.com:wolfxxxz/newLine.git
 ! [rejected]        **master -> master (fetch first)**
**error:** failed to push some refs to 'github.com:wolfxxxz/newLine.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

D:\Onedrive\Рабочий стол\us2\newLine>**git fetch**
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 3 (delta 0), reused 3 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), 299 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
   eaef84d..c99987d  master     -> origin/master

D:\Onedrive\Рабочий стол\us2\newLine>**git rebase origin/master**
Auto-merging file1.txt
CONFLICT (content): Merge conflict in file1.txt
**error:** could not apply 6801ff6... line4 user1
hint: Resolve all conflicts manually, mark them as resolved with
hint: "git add/rm <conflicted_files>", then run "git rebase --continue".
hint: You can instead skip this commit: run "git rebase --skip".
hint: To abort and get back to the state before "git rebase", run "git rebase --abort".
Could not apply 6801ff6... line4 user1

D:\Onedrive\Рабочий стол\us2\newLine>**git rebase --continue**
file1.txt: needs merge
You must edit all merge conflicts and then
mark them as resolved using **git add**

D:\Onedrive\Рабочий стол\us2\newLine>**git add .**

D:\Onedrive\Рабочий стол\us2\newLine>**git commit -m "repair"**
[detached HEAD 6abf5d6] repair
 1 file changed, 2 insertions(+), 1 deletion(-)

D:\Onedrive\Рабочий стол\us2\newLine>**git rebase --continue**
Successfully rebased and updated refs/heads/master.

D:\Onedrive\Рабочий стол\us2\newLine>**git push origin master**
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 8 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (6/6), 587 bytes | 17.00 KiB/s, done.
Total 6 (delta 1), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To github.com:wolfxxxz/newLine.git
   c99987d..f94d5a3  master -> master
##### User1 git pull
D:\Onedrive\Рабочий стол\us1\newLine>**git pull**
remote: Enumerating objects: 9, done.
remote: Counting objects: 100% (9/9), done.
remote: Compressing objects: 100% (5/5), done.
remote: Total 6 (delta 1), reused 6 (delta 1), pack-reused 0
Unpacking objects: 100% (6/6), 567 bytes | 1024 bytes/s, done.
From github.com:wolfxxxz/newLine
   c99987d..f94d5a3  master     -> origin/master
Updating c99987d..f94d5a3
Fast-forward
 file 2.txt | 3 ++-
 file1.txt  | 3 ++-
 2 files changed, 4 insertions(+), 2 deletions(-)

D:\Onedrive\Рабочий стол\us1\newLine>**git log --oneline**
f94d5a3 (HEAD -> master, origin/master, origin/HEAD) line5 user2 file1 master
6abf5d6 repair user2
c99987d line5 user1 file1 master
eaef84d Merge branch 'master' of github.com:wolfxxxz/newLine

## Интерактивный rebase
### Теория
Что может git rebase -i? git rebase -i HEAD~3
- менять местами commit
- переименовывать коммиты
- обьеденить комиты `squash`
- добавить изменения в любой комит `edit`
### Практика
#### VIM redactor и изминение очерёдности комитов
git bash
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
> **git log --oneline**
10da7ef (HEAD -> master) file4
cfe6c4e file3
a2c7680 file2
1cc5e4f file1

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
> **git rebase -i HEAD~3**
##### VIM redactor
```
pick 27e7901 file2
pick ec2f7de file3
pick 6438893 file4

# Rebase 1cc5e4f..6438893 onto 1cc5e4f (3 commands)
#
# Commands:
# p, pick <commit> = use commit
# r, reword <commit> = use commit, but edit the commit message
# e, edit <commit> = use commit, but stop for amending
# s, squash <commit> = use commit, but meld into previous commit
# f, fixup [-C | -c] <commit> = like "squash" but keep only the previous
#                    commit's log message, unless -C is used, in which case
#                    keep only this commit's message; -c is same as -C but
#                    opens the editor
# x, exec <command> = run command (the rest of the line) using shell
# b, break = stop here (continue rebase later with 'git rebase --continue')
# d, drop <commit> = remove commit
# l, label <label> = label current HEAD with a name
# t, reset <label> = reset HEAD to a label
# m, merge [-C <commit> | -c <commit>] <label> [# <oneline>]
#         create a merge commit using the original merge commit's
#         message (or the oneline, if no original merge commit was
#         specified); use -c <commit> to reword the commit message
# u, update-ref <ref> = track a placeholder for the <ref> to be updated
#                       to this position in the new commits. The <ref> is
#                       updated at the end of the rebase
#
```
**Ввести команду `i` => --INSERT-- Включить режим редактирования**
git bash
copy  `ctr` + `shift` + `c`
paste `shift` + `0`

**вверх ногами ВНИМАТЕЛЬНО**
``` vim redactor
pick 6438893 file4
pick ec2f7de file3
pick 27e7901 file2

# Rebase 1cc5e4f..6438893 onto 1cc5e4f (3 commands)
```
exit INSERT `esc`
save and exit `:wq`
`enter`

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
> git rebase -i HEAD~3
**Successfully rebased and updated refs/heads/master.**
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)

> **git log --oneline**
a2c7680 (HEAD -> master) file2
cfe6c4e file3
10da7ef file4
1cc5e4f file1
#### Переименовать комит vim redactor
>git rebase -i HEAD~2
```
pick e5c66a7 file3
pick 966a8d1 file2

# Rebase 1f87207..966a8d1 onto 1f87207 (2 commands)
#
# Commands:
```
`i` => `pick` = `reword`
`esc` => `:wq`
```
file3

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
#
# Date:      Thu Mar 16 14:18:47 2023 +0200
#
# interactive rebase in progress; onto 1f87207
# Last command done (1 command done):
#    reword e5c66a7 file3
# Next command to do (1 remaining command):
#    pick 966a8d1 file2
# You are currently editing a commit while rebasing branch 'master' on '1f87207'.
#
# Changes to be committed:
#       new file:   file3.txt
```
`i` => `file3` rename `file888`
`esc` => `:wq` - save and exit

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$ git rebase -i HEAD~2**
[detached HEAD 16def64] file88
 Date: Thu Mar 16 14:18:47 2023 +0200
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 file3.txt
**Successfully rebased and updated refs/heads/master.**

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$ git log --oneline**
72957c5 (HEAD -> master) file2
**16def64 file88**
1f87207 file4
1cc5e4f file1
#### Обьеденить два коммита в один
>git rebase -i HEAD~3
```
pick 1f87207 file4
pick 16def64 file88
pick 72957c5 file2

# Rebase 1cc5e4f..72957c5 onto 1cc5e4f (3 commands)
```
`i` + `squash 72957c5 file2`
```
pick 1f87207 file4
pick 16def64 file88
sguash 72957c5 file2

# Rebase 1cc5e4f..72957c5 onto 1cc5e4f (3 commands)
```
`esc` => `:wq`
После обьединения получился новый комит его нужно обозвать new commit name
```
# This is a combination of 2 commits.
# This is the 1st commit message:
file88
# This is the commit message #2:
file2
# Please enter the commit message for your changes. Lines starting
```
```
# This is a combination of 2 commits.
# This is the 1st commit message:

# This is the commit message #2:
new files 2 and 3
# Please enter the commit message for your changes. Lines starting
```
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$ git rebase -i HEAD~3**
[detached HEAD 124276e] new files 2 and 3
 Date: Thu Mar 16 14:18:47 2023 +0200
 2 files changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 file2.txt
 create mode 100644 file3.txt
`Successfully rebased and updated refs/heads/master.`

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$ git log --oneline**
124276e (HEAD -> master) `new files 2 and 3`
1f87207 file4
1cc5e4f file1
#### Изменить любой комит
**git rebase -i HEAD~2**
```
pick 1f87207 file4
pick 124276e new files 2 and 3

# Rebase 1cc5e4f..124276e onto 1cc5e4f (2 commands)
```
`i` + `edit`
```
edit 1f87207 file4
pick 124276e new files 2 and 3

# Rebase 1cc5e4f..124276e onto 1cc5e4f (2 commands)
```
`esc` => `:wq`

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$ git rebase -i HEAD~2**
Stopped at 1f87207...  file4
You can amend the commit now, with

  git commit --amend

Once you are satisfied with your changes, run

  git rebase --continue

**Вносим изминения в файл 4**
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master|REBASE 1/2)
**$ git add .**

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master|REBASE 1/2)
**$ git commit --amend**
[detached HEAD 9b9f283] file4 and like
 Date: Thu Mar 16 14:19:15 2023 +0200
 1 file changed, 2 insertions(+)
 create mode 100644 file4.txt
**Предлагает изменить название комита**
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master|REBASE 1/2)
**$ git rebase --continue**
Successfully rebased and updated refs/heads/master.

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/number (master)
**$git log --oneline**
13ca811 (HEAD -> master) new files 2 and 3
9b9f283 file4 and like
1cc5e4f file1

## git cherry pick 
### Теория
git cherry-pick - ипользуется для переноса одного комита из второстипенной ветки в основную
git cherry-pick <хеш комита>
git cherry-pick --no-commit - слияние из двух комитов и перенос его в основную ветку
git cherry-pick -x - хеш комита который был добавлен
git cherry-pick -signoff
git cherry-pick --edit - Добавление комита из соседней ветки с изминением названия
get cherry-pick -x <хеш> - добавляет приписку с хешем от куда был коммит
git cherry-pick --signoff - добавляет подпись (автора) к chery-picky
### Практика
#### Перенос одного комита из соседней ветки git cherry-pick <хеш>
##### branch development
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (development)
`$ git log`
commit 39802cc1fa5b44635c75d8b1cc8c601b6cd74be6 (HEAD -> development)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:34:34 2023 +0200

    E

commit 7a5e96765d22d073a478c1223c071f21a6e0e09d
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:34:11 2023 +0200

    D

commit 95bc2cd31fdb47101f362767e6ca2d84cd877568
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:33:52 2023 +0200

    C

commit `ab6170ffe621cd4c65e6819874ab65f10cc09f7d`
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:33:29 2023 +0200

    B
##### branch production
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (development)
`$ git checkout production`
Switched to branch 'production'

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git log`
commit a560033a848a1dd330becaf40cb636eda79b0706 (HEAD -> production)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 16 19:00:34 2023 +0200

    just one commit

commit b1b1538819e71fed4b3eba2bb22aa9b34d6919c8
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 16 18:44:47 2023 +0200

    without files don't create branches
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git cherry-pick ab6170ffe621cd4c65e6819874ab65f10cc09f7d`
[production c72baa9] ` B `
 Date: Fri Mar 17 09:33:29 2023 +0200
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 B.txt

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git log`
commit c72baa91212d799ff052fe404fd6904294a8352a (HEAD -> production)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:33:29 2023 +0200

   ` B `

commit a560033a848a1dd330becaf40cb636eda79b0706
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 16 19:00:34 2023 +0200

    just one commit

commit b1b1538819e71fed4b3eba2bb22aa9b34d6919c8
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Thu Mar 16 18:44:47 2023 +0200

    without files don't create branches
#### git cherry-pick --edit (Добавление комита из соседней ветки с изминением названия)
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (development)
`$ git log --oneline`
39802cc (HEAD -> development) E
7a5e967 D `7a5e96765d22d073a478c1223c071f21a6e0e09d`
95bc2cd C
ab6170f B 
33fc169 A
b1b1538 without files don't create branches

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (development)
`$ git checkout production`
Switched to branch 'production'

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git log --oneline`
c72baa9 (HEAD -> production) B
a560033 just one commit
b1b1538 without files don't create branches

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production|CHERRY-PICKING)
`$ git cherry-pick --edit 7a5e96765d22d073a478c1223c071f21a6e0e09d`
VIM `i` => редактировать название `esc :wq`
[production 4c172f5] commit D with changed commit message
 Date: Fri Mar 17 09:34:11 2023 +0200
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 D.txt

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git log --oneline`
4c172f5 (HEAD -> production) **commit D with changed commit message**
c72baa9 B
a560033 just one commit
b1b1538 without files don't create branches
#### git cherry-pick --no-commit (Добавляет файлы из комита соседней ветки в незакомиченную зону (unstage zone))
Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (`development`)
`$ git log`
commit `39802cc1fa5b44635c75d8b1cc8c601b6cd74be6` (HEAD -> development)
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:34:34 2023 +0200

    E

commit 7a5e96765d22d073a478c1223c071f21a6e0e09d
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:34:11 2023 +0200

    D

commit `95bc2cd31fdb47101f362767e6ca2d84cd877568`
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:33:52 2023 +0200

    C

commit ab6170ffe621cd4c65e6819874ab65f10cc09f7d
Author: wolfxxxz <wolfxxxz@gmail.com>
Date:   Fri Mar 17 09:33:29 2023 +0200

    B

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (development)
`$ git checkout production`
Switched to branch 'production'

Mvmir@MVMIR MINGW64 /d/Onedrive/Рабочий стол/git/firm (production)
`$ git log --oneline`
4c172f5 (HEAD -> production) commit D with changed commit message
c72baa9 B
a560033 just one commit
b1b1538 without files don't create branches
**go to cmd because git bash - глючная херь**
D:\Onedrive\Рабочий стол\git\firm>**git log --oneline**
4c172f5 (HEAD -> production) commit D with changed commit message
c72baa9 B
a560033 just one commit
b1b1538 without files don't create branches

D:\Onedrive\Рабочий стол\git\firm>**git cherry-pick --no-commit 39802cc1fa5b44635c75d8b1cc8c601b6cd74be6 95bc2cd31fdb47101f362767e6ca2d84cd877568**

D:\Onedrive\Рабочий стол\git\firm>**git log --oneline**
4c172f5 (HEAD -> production) commit D with changed commit message
c72baa9 B
a560033 just one commit
b1b1538 without files don't create branches

D:\Onedrive\Рабочий стол\git\firm>**git status**
On branch production
Changes to be committed:
  (use "git restore --staged <file>..." **to unstage**)
        new file:   C.txt
        new file:   E.txt


D:\Onedrive\Рабочий стол\git\firm>**git commit -m "some add from cherry-pick --no-commit"**
[production f936869] some add from cherry-pick --no-commit
 2 files changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 C.txt
 create mode 100644 E.txt

D:\Onedrive\Рабочий стол\git\firm>**git log --oneline**
f936869 (HEAD -> production) **some add from cherry-pick --no-commit**
4c172f5 commit D with changed commit message
c72baa9 B
a560033 just one commit
b1b1538 without files don't create branches
#

















#  Kомманды
## Конфигурация и установка
git config: устанавливает настройки Git
git init: инициализирует репозиторий Git

## Основные команды
### git add: добавляет измененные файлы в индекс
Добавить файлы списком                    - git add file1.exe file2.go file3.java 
Добавить все файлы из текущей папки       - git add .
Добавить все файлы с расширением          - git add *.java
Добавить все файлы из папки с расширением - git add someDir/*.java
Добавить все файлы из папки               - git add someDir/
Добавить все файлы в проэкте с расширением - git add "*.java"

### git commit: фиксирует изменения и создает новый коммит
#### git commit -m"Message"
> git commit -m"Message"
создаёт резервную копию + сообщение о ней, Сообщение ОБЯЗАТЕЛЬНО будет добавлено

### git clone: клонирует репозиторий
git clone <ссылка> = git remote add <ссылка> + git pull
git clone git@github.com:wolfxxxz/FirstRepository.git

### git pull: забирает изменения из удаленного репозитория и сливает их с текущей веткой
git remote add origin https://github.com/wolfxxxz/FirstRepository.git
git pull - синхронизировать актуальные версии файлов со всеми commit 

### git push: отправляет изменения из локального репозитория в удаленный репозиторий
git push ИМЯ ВЕТКА
git push origin master
>git commit -a -m "firs push"
>git push origin master

### git reset - откат
Имеет три режима в зависимости от радикальности отката
- soft
- mixed
- hard 
git reset [soft | mixed | hard] [ точка отката ]
                режим             может быть хеш или head 
#### git reset - по умолчанию используется HEAD^
Отменить последний git add . (удалить последнее добавление из unstage)
> git reset = > git reset --mixed HEAD^
#### git reset HEAD^^ - по умолчанию используется mixed
> git reset HEAD^^ = > git reset --mixed HEAD^^
#### git reset --hard удаляет без возвратно изминения и на git и на компютере
> git reset --hard - удалить полность последний commit
> git reset --hard HEAD^^ - удалить два последних commit
#### git reset --mixed переводит в неотслежеваемую зону (unstaged)
Отменить git add .
> git reset --mixed
> git reset --mixed HEAD^^ - удалить последних два add
#### git reset --soft переводит в отслежуемую зону
* возвращает из commit в add зону *
git reset --soft
git reset --soft HEAD^^ - вернёт из commit в add
#### ВАЖНО - не исползовать при совместных проэктах

### git clean 
**git clean -f** - удаляет untracket файлы т.е. файлы тоько созданные 
Показать untracket файлы
>git clean -n
Удалить
>git clean -f

## Ветки
### git branch: показать, создать, удалить ветку
git branch <name> - создание новой ветки
git branch - показывает название ветки HEAD
git branch -d  - удаление ветки

### git checkout: откат и переключения между ветками 
#### Откат
##### посмотреть ка выглядел проэкт в каkом то снимке в прошлом
- git checkout хеш
- git checkout HEAD^^
- git checkout HEAD~2
При работе с предыдущей версией программы (предыдущий commit) - данные не сохраняются а предыдущая версия остаётся без изминений
- git checkout master - возвращает HEAD на последний commit
##### вернуть файл к предыдущему commit версии
- git checkout -- . **вернуть все файлы в это состояние**
- git checkout -- file **вернуть только file в это состояние**
##### Для отмены не добавленных изминений
- git checkout -- file
- git checkout -- . 
##### Для отмены добавленых изминений
- git reset 
- git checkout -- .
#### Переключение между ветками
- git checkout <name> - переключение между ветками

### git merge: сливает две или более веток

## Информация и проверка
### git status: показывает текущее состояние репозитория
**Просто созданный "untracked"**
Не отслеживается gitom
**Изминения не сохранены "modified"**
Файл создан но не добавлен - еще не было git add
**Подготовленный "staged"**
В связи с тем что git хранит снимки (копии) только изменений 
- Изменённые файлы нужно пометить командой
>git add file.go file2.exe
**Зафиксированный "committed"**
- После чего сохранить 
>git commit -m"Message"

### git log: показывает историю коммитов
- git log
- git log --oneline

### git diff: показывает различия между версиями файлов
#### git diff работает до выделения git add в статусе modified
git diff

#### работает после git add в статусе staged
git diff --staged - разница между текущим и последним сохранённым 

#### git diff COMMIT_ID - все изменения между текущим и указанным 
git diff COMMIT_ID 
COMMIT_ID - Длинный хеш 

### git blame: показывает, кто и когда внес изменения в каждую строку файла

#
