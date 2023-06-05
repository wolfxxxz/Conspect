# Основы
## 1 pwd, cd
### pwd print working directory
pwd    - где я нахожусь // /home/mvmir
### cd change directory
cd /   - перейти в корневая папка (root)
cd -   - вернутся в предидущее место // /home/mvmir
cd ~   - перейти в домашний каталог юзера
cd == cd ~
cd dir/bash - перейти в каталог по адрессу dir/bash
cd ~/dir/bash - перейти в домашний каталог а потом dir/bash
cd ..   - родительская директория (на папку в верх)
cd ../.. - на две папки вверх
clear   - очистить всё
## 2 basename, dirname (отображают части пути)
man basename
basename mvmir/go/src/github.com/Wolfxxxz/SlovarNV
SlovarNV
dirname mvmir/go/src/github.com/Wolfxxxz/SlovarNV
mvmir/go/src/github.com/Wolfxxxz
## 3 ls, ls -R, ls <folder>
ls   - показать содержимое папки (в линию)
ls one  - где 'one' - название папки что мы хотим увидеть
ls -a    - показать скрытые папки и файлы
ls -A    - показать скрытые несистемные файлы
ls -a -l bash    - показать все скрытые файлы в столбик с папки one
ls -a -l bash == ls -al bash == ls -la bash
ls -lh  - файлы смотрятся приятнее и опрятнее 
ls -R  - show al the papki and categories
ls -Ralh - выдаст все секреты венского леса
ls -i  - inot number
ls -l   - показать данные по файлам (показать полностью дата создания, владелец ...)
ls -l bash   - показать все данные по файлу
man ls - справка по ls
man <name> - справка что угодно
## 4 input output >, >> , error Перенаправленный вывод, if
> - reWrite
>> - append
ls 1>output.txt  - если ls что то выдаст типа true то записать это в файл outout (при этом 1 это типа true)
ls >>file.txt
***1 можно упускать она по умолчанию стоит**
ls /sdsd/sdsd 2>error.txt - если что то выдаст (2 если error) то записать это в error.txt
### if:)
ls > output.txt 2> error.txt
ls /sdsd/sdd > output.txt 2> error.txt
### и оутпут и эрор в один файл
ls &>> output.txt
## 5 mkdir/rmdir/rm/touch/ Create and delete
### mkdir (makedir) - create directory
mkdir folder   - create folder <folder>
mkdir -p files/newFolder  - создать папку files и в ней создать папку newFolder
### rmdir remove an EMPTY(пустая) directory
rmdir folder   - удалить ***пустой** delete folder <folder>
### rm (remove) - удаляет папки и файлы
man rm - many options in manual
rm  - delete
rm file.avi   - удалить файл <file.avi>
rm directory  - удалить папку <directory> - ***не работает**
rm -r directory  -  **работает** удалить папку <directory> **ВНИМАТЕЛЬНО** удаляет мимо корзины
rm -r ./*    - удалить всё из текущей папки
rm -rf ***УДАЛЯЕТ ВСЁ НАХРЕН*** вместе с системными файлами и папками <f> - force
rm -ri <folder>  - спрашивает по каждому файлу в папке, хочешь удалить или нет - сам <folder> - тоже будет удалён если <y>
rm -rv CopyTest - <-v> показать лог
### touch
touch file.txt   - обновляет дату изминения файла **как побочка** создаёт файл если его раньше не было create <file1.txt>
## 6 mv and cp  - move, rename, Copy file
### mv - move or rename
mv one/file1.txt tue  - переместить file1.txt с папки one в папку tue
mv one/file2.txt one/file.txt - переименовать file2.txt в file.txt - команда меняет только ссылку на файл (работает только в границах одной файловой системы)
mv -v TestCopy CopyTest - переименовать и показать лог
renamed 'TestCopy' -> 'CopyTest'
mv OpenFolder/newFolder/file.avi .  - перенесёт ссылку в текущую папку (где сейчас находится пользователь)
mv newFolder/file.avi ~/dir  - переместить в папку home/dir
mv -i -v file.txt output.txt CopyTest - переместить в папку CopyTest <-i> - если есть такой файл то спросить
renamed 'file.txt' -> 'CopyTest/file.txt'
renamed 'output.txt' -> 'CopyTest/output.txt'
mv -i -v output.txt CopyTest 
mv: переписать 'CopyTest/output.txt'? n
### cp  - Copy file
cp  - копирует файл (именно копирует в новое место а не создаёт ссылку и не перемещает)
cp one/* tue  - копирует все файлы с one (кроме папок) в каталог tue
cp -vR test4 copyTest  - копирует всё содержимое -R и показывает лог -v
'test4' -> 'copyTest'
'test4/test5' -> 'copyTest/test5'
'test4/test.txt' -> 'copyTest/test.txt'
cp -vR bind/. SlovarNV - копировать всё содержимое папки bind в папку SlovarNV
cp -R SlovarNV/* newSlovarNV - копировать всё содержимое SlovarNV в папку newSlovarNV
## 7 df (inot), ln (LINK)
### df - display free disk space
man df
ls -i  -(inot number) уникальный номер файла
ls -l   - показать данные по файлам (показать полностью дата создания, владелец ...)
df -hi  - количество Инодов (адрессов ссылок) и занятое ними пространство на диске
### ln - make links between files
#### hard link - не может ссылатся на папку, зато можно перемещать и связь сохранится ln
ln files/kelvin.file hard_link_kelvin
cat hard_link_kelvin
Kelvin
mv hard_link_kelvin test4/
cat test4/hard_link_kelvin
Kelvin
#### soft link - может ссылатся на что угодно ln -s, при перемещении в другое место связь обрывается
ln -s output.txt soft_link_output.txt - создать софт ссылку на output.txt
ls -l - покажет что soft_link_output.txt -> ссылка
echo how are you >> soft_link_output.txt
cat output.txt - запись по ссылке попадёт в файл
Hello
how are you
ln -s ~/go/src/github.com/Wolfxxxz/ Link_Wolfxxxz
## 8 Просмотр содержимого cat & less & wc & nl & head & tail & ll
### cat - показать файл
cat <file.txt> - посмотреть файл
cat <file.txt> <file2.txt> - смотреть прям два файла одновременно
### less
less <file.txt> - прям можно дажо что то искать и листать и ...огого <q> - для выхода
/Петя  - поиск символов в файле
man less - мануал less 
man cat - есть и даже можно глянуть (открывает в cat)
### head, tail, nl, wc
head <file> - показать первые 10 строк
tail <file> - показать последние 10 строк
nl <file>  - выводит нумерованные строки
wc <file>  - выводит статистику по файлу, количество символов , строк...
## 9 nano, vi, emacs Текстовые редакторы
### nano
nano first_file - или откроет или создаст новый(если такого нет) - выходы и т,д - через ctrl +
### vi
vi first_file - открыть
- i - insert - что бы писать
- esc :w  - что бы сохранить
- esc :q  - выйти
- esc :q! - выйти без сохранения
#### чуть больше инфо по vi
Некоторые основные команды в командном режиме:

Сохранение и выход: :w - сохранить изменения, :q - выйти из редактора, :wq - сохранить и выйти.
Выделение текста: В режиме команд воспользуйтесь клавишами h, j, k, l для перемещения курсора по тексту.
Удаление текста: x - удалить символ под курсором, dd - удалить текущую строку.
Копирование и вставка: yy - скопировать текущую строку, p - вставить скопированный текст после курсора.
Поиск и замена текста:

В режиме команд введите / и затем введите текст, который вы хотите найти, и нажмите Enter. Чтобы перейти к следующему вхождению, нажмите n.
Для замены текста в режиме команд используйте команду :%s/старый_текст/новый_текст/g, где % означает применение замены ко всему файлу, g - замена всех вхождений.
### emacs
## 10 file text manipulation /cut, paste, tr, sort, uniq, tee, grep
### cut - извлечь колонку
cut -d "," -f2 cutFile - где <-d ","> - делить колонки по запятой, <-f2> - вывести вторую колонку
ls -l | tail -n +2 > dir/ls_output.txt - где <tail -n +2> - взять все последние строки начиная со второй <> dir/ls_output.txt> и записать их в файл
cut -d ' ' -f3 ls_output.txt - где <-d ' '> - делить колонки по пробелу (-w походу работает на макОС)
paste - добавить колонку
tr - преобразовать один символ в другой
### sort
cut -d "," -f2 cutFile | sort
cut -d "," -f2 cutFile | sort -r  - <-r> - sort revers
cut -d "," -f2 cutFile | sort -n  - <-n> - приводит строки к числу и потом сортирует
### uniq - уникальные
cut -d "," -f2 cutFile | sort -n | uniq -c - <-c> - показать количество дубликатов
### tee - передать в output + записать
ls -l | tail -n +2 | cut -d ' ' -f 7 | sort -n -r | uniq -c | tee dir/uniq_numb.txt - и пишет в файл и выводит в консоль или можно что то дальше делать с потоком
     2 3935
      1 50
      2 17
      3 16
      3 15
     10 
### grep - поиск в текстовых файлах
#### Команды
grep <text> <file.txt> - поиск в файле
grep <text> <directoria/file.txt> поиск в файле с другой директории
grep -i <text> <file.txt>  - поиск без учета регИстрА
grep -v <text>  <file>  - показать всё кроме <text> 
grep -v -i петя one/file.txt
grep -c <text> <file.txt>  - count - посчитать все записи <text>
grep -R <text> <directoria/> - искать в папке
grep -R <text> <directoria>  - тоже самое
grep -R <text>  - и в текущей папке тоже
grep -R <text> .  - с точкой покажет ещё и файлы
#### egrep - работает лучше с регулярными выражениями
#### https://docs.oracle.com/javase/8/docs/api/java/util/regex/Pattern.html
#### Регулярные выражения в grep
Ниже приведены некоторые основные элементы регулярных выражений в grep:

***Символы:**
. (точка): Соответствует любому одиночному символу, кроме символа новой строки.
* (звездочка): Соответствует нулю или более повторений предыдущего символа или выражения. Например, a* соответствует a, aa, aaa и так далее.
***Символьные классы:**

[abc]: Соответствует любому из символов a, b или c.
[a-z]: Соответствует любому символу от a до z.
[0-9]: Соответствует любой цифре от 0 до 9.
[^abc]: Соответствует любому символу, кроме a, b или c.
***Альтернатива:**
| (вертикальная черта): Соответствует одному из двух выражений. Например, apple|orange соответствует apple или orange.
***Квантификаторы:**
? (вопросительный знак): Соответствует нулю или одному повторению предыдущего символа или выражения.
+ (плюс): Соответствует одному или более повторений предыдущего символа или выражения.
{n}: Соответствует ровно n повторениям предыдущего символа или выражения.
{n,}: Соответствует не менее n повторений предыдущего символа или выражения.
{n,m}: Соответствует от n до m повторений предыдущего символа или выражения.
***Символы начала и конца строки:**
```
^ (канаточка): Соответствует началу строки.
$ (доллар): Соответствует концу строки.
```
***Экранирование символов:**
Если вы хотите использовать специальные символы (например, ., *, [, ] и т. д.) как часть шаблона, вам может понадобиться экранировать их с помощью обратного слеша \. Например, \. будет соответствовать точке, а не любому символу.
## 11 Find <-type d> - папки, <-type f> - файлы
find . -name file1.txt  - <.> - начать с текущей директории, <-name> - искать имя, <file1.txt> - само имя файла
find . -name *txt - где <*txt> - найти все файлы название которых заканчивается на txt
find . -type d  - найти все папки (директории) <.> - начиная с текущей директории
find . -type d -name "on*" - найти <-type d> - все папки кот начинаются с <"one*">
find . -type f -name "*e*" - найти все <-type f> - файлы 
find . -name "*f*"  -  найти все файлы И папки в которых есть f
find . -name "*avi" -delete - найти и удалить
find . -name "*.txt" -delete -print - показать что удалилось
## 12.1 echo Ввод текста и запись в файл
echo "hello world" - напечатать в ответе hello world
echo "hello world" > file.txt  - write "hello world" in file.txt - команда перезаписует файл полностью
echo "this is a text" >> file2.txt  - ДОБАВИТЬ запись
**если файл не существует создаёт файл**
echo "This is file1" > 1.file - создать 1.file с записью
echo "This is file2" > 2.file - ...
cat 1.file 2.file >> 3.file - создать 3.file со строчками с файла 1 + 2
## 12.2 PAIP, &&, || - логические комманды
### paip |
cat 3.file | grep i - прочитать файл и найти в нём i
ls | grep file - вывести только те файлы где grep найдёт file
### ; == && == логическое И
cd dir; ls
### || == or == или
cd Movies || cat file.txt
## 13 unix command Процессы
ps - process status - process id, tty, time, cmd
ps x - процессы
ps u  - только процессы пользователя
ps au - ещё процессы
ps aux - вообще все процессы
top - диспетчер приложений <q> - exit
kill 2597 - убить процесс <2597> - pid (process id)
kill -l - справка по кил
9 - рубит как попало, 15 - завершает штатно процесс
## 14 Сети (Host Location) host, ping, curl, traceroute
### 14.1 host
host google.com
google.com has address 216.58.208.206
google.com has IPv6 address 2a00:1450:401b:800::200e
google.com mail is handled by 10 smtp.google.com.
host localhost
localhost has address 127.0.0.1
### 14.2 ping (проверяет доступность конект)
ping 127.0.0.1
PING 127.0.0.1 (127.0.0.1) 56(84) bytes of data.
64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.040 ms
ping google.com - выдаёт пинг по запросам
ping -c 3 google.com  - отправить три запроса 
ping -c 3 -i 2 google.com - где -i - интервал отправки запросов в секунду
ping -c 3 -a google.com  - <-a> - пикает при каждой отправке
### traceroute - показывает путь и время затраченное на подключение
traceroute google.com
traceroute to google.com (216.58.208.206), 30 hops max, 60 byte packets
 1  _gateway (192.168.3.1)  4.739 ms  4.595 ms  4.505 ms
 2  10.191.99.1 (10.191.99.1)  7.502 ms  7.443 ms  7.385 ms
 3  29.51.209.91.it-tv.org (91.209.51.29)  7.968 ms  7.913 ms  7.857 ms
 4  74.51.209.91.it-tv.org (91.209.51.74)  7.878 ms  8.882 ms  8.830 ms
 5  108.170.248.155 (108.170.248.155)  8.888 ms  8.784 ms 108.170.248.138 (108.170.248.138)  8.669 ms
 6  142.251.242.35 (142.251.242.35)  21.068 ms 142.251.242.39 (142.251.242.39)  21.608 ms  22.290 ms
 7  142.250.46.55 (142.250.46.55)  22.208 ms  18.589 ms 142.250.37.193 (142.250.37.193)  18.489 ms
 8  142.250.224.91 (142.250.224.91)  19.225 ms  20.162 ms 142.250.37.209 (142.250.37.209)  20.108 ms
 9  142.250.224.91 (142.250.224.91)  20.053 ms 142.250.224.89 (142.250.224.89)  20.735 ms par10s21-in-f14.1e100.net (216.58.208.206)  20.024 ms
### 14. curl
curl google.com - получается ответ html прям в bash
curl -L ya.ru  - <-L> - перейти на ссылку переадресации
curl https://regres.in/api/users - главная страница
curl https://regres.in/api/users?page=2 - вторая страница
## 15 ssh - regestration на удалённом сервере
sshcurl https://regres.in/api/users?page=2
exit
scp - copy something на удалённом сервере
## 16 sed - потоковый текстовый редактор
echo 123123 | sed 's/1/5/' - вывод 523123
sed s/hi/Hello/ file2.txt - заменить hi на Hello с файла file2.txt <s> - значит заменить и вывести в консоль - в файле останется старое значение
sed s/hi/Hello/ file2.txt >> hello.file - тоже самое что выше только записать ещё всё в hello.file
sed "2 s/is/Hello/" file2.txt - <2> означает вторая строка сверху
sed "2 s/is/Hello/g" file2.txt - <g> - все подстроки <is> заменить
sed "2 s/is/Hello/g; 1 s/hi/brown/" file2.txt - несколько действий одновременно <;>
sed "s/is/Hello/w hello2.file" file2.txt - записать всё что получится в hello2.file
### sed -f  - перенаправить команды из файла
```
echo s/test/HiAgain/ >> sed_command.file - создать файл sed_command.file с первой строчкой s/test/HiAgain/
echo s/9/000/ >> sed_command.file - записать вторую строчку s/9/000/
echo test1 >> my.file
echo test2 >> my.file
echo test9 >> my.file
----------------------------
cat my.file
test1
test2
test3
test9
----------------------------
sed -f sed_command.file my.file
----------------------------
HiAgain1
HiAgain2
HiAgain3
HiAgain000
```
## 17 awk - работа с табличными значениями
echo "first second" | awk '{print $2}' - типа принт второго значения строки 
second
echo "first second" | awk '{print $0}' - где 0 - это вся строчка
first second
### file.table
cat file.txt
test1	test2
test3	test4
test5	test6
test7	test8
test9	
cat file.txt | awk '{print $1}' - schow first stolbik
test1
test3
test5
test7
test9
cat file.txt | awk '{print $2}' - schow second stolbik
test2
test4
test6
test8
### меняем разделитель который по умолчанию пробел
echo "this,is,a,table" | awk -F, '{print $1}' <-F,> - разделителем теперь будет <,>
this
echo "this,is,a,table" | awk -F, '{print $4}'
table
echo "this;is;a;table" | awk -F\; '{print $4}' - <;> без \ не будет работать тк ; нужно экранировать ... вдруг что то не работает можно экранировать \
table
### типа первый элемент = оне и принт
echo "first second" | awk '{$1="one";print $1}'
one
## 18.1 which find File location commands
### which  - locate executable files in env PATH
which cat echo mkdir rm
/bin/cat
/bin/echo
/bin/mkdir
/bin/rm
### type    - locate executable files in env PATH
***в отличии от which - type понимает где alias**
type cat echo mkdir rm
для cat вычислен хэш (/bin/cat)
echo — это встроенная команда bash
для mkdir вычислен хэш (/bin/mkdir)
для rm вычислен хэш (/bin/rm)
### whereis - locate executable files in env PATH
whereis cat echo mkdir rm
cat: /bin/cat /usr/share/man/man1/cat.1.gz
echo: /bin/echo /usr/share/man/man1/echo.1.gz
mkdir: /bin/mkdir /usr/share/man/man2/mkdir.2.gz /usr/share/man/man1/mkdir.1.gz
rm: /bin/rm /usr/share/man/man1/rm.1.gz
### locate - ищет по индексу
### find - ищет файл в иерархии
#### find simple
Find <-type d> - папки, <-type f> - файлы
find . -name file1.txt  - <.> - начать с текущей директории, <-name> - искать имя, <file1.txt> - само имя файла
find . -name *txt - где <*txt> - найти все файлы название которых заканчивается на txt
find . -type d  - найти все папки (директории) <.> - начиная с текущей директории
find . -type d -name "on*" - найти <-type d> - все папки кот начинаются с <"one*">
find . -type f -name "*e*" - найти все <-type f> - файлы 
find . -name "*f*"  -  найти все файлы И папки в которых есть f
find . -name "*avi" -delete - найти и удалить
find . -name "*.txt" -delete -print - показать что удалилось
#### find +
man find - очень много опций
find . -name "info" 2> /dev/null
file ./go/src/github.com/Wolfxxxz/SlovarNV/.git/objects/info - как узнать папка это или файл - <file ...>
./go/src/github.com/Wolfxxxz/SlovarNV/.git/objects/info: directory
find . -user mvmir -type f -name "info" -empty 2> /dev/null
<-user mvmir> - конкретного пользователя
<-type f>   - искать файлы
<-name "info">  - с названием
<-empty> - пустые
<2> /dev/null>  - ошибки не показывать
## 18.2 xargs
echo a b c d e f | xargs  - просто отобразит последовательность
a b c d e f
echo 1.file 2.file | xargs - отобразит файлы
1.file 2.file
echo 1.file 2.file | xargs cat - отобразит содержимое файлов
This is file1
This is file2
ls | sort -n | xargs  - отобразить отсортированные файлы 
bash bed file2.txt file.txt got hello2.file hello.file sed_command.file 1.file 2.file 3.file
ls | sort -n | xargs cat  - отобразить содержимое отсортированных файлов 
cat: bash: Это каталог
hi again
this is a text
test1	test2
test3	test4
ls *.file | sort -n | xargs cat  - вывести всё кроме *.file
thHello is a text
thHello is a text
s/test/HiAgain/
s/9/000/
This is file1
find . -name '*.new_file'  - найти где есть такое в имени файла
./71.new_file
./72.new_file
find . -name '*.new_file' | xargs cat
new file content 71
new file content 72

grep file -R   - найти file все строки
bash/one/one/new.file:Hello ! My name is nano, i'm a redactor file .txt 
bash/one/one/new.file:I can open, write, and save .txt files
1.file:This is file1
71.new_file:new file content 71
2.file:This is file2
3.file:This is file1
3.file:This is file2
72.new_file:new file content 72

grep file -R | awk -F: '{print $1}'  - найти все строки и выдать только первый элемент до разделителя (только название файла)
bash/one/one/new.file
bash/one/one/new.file
1.file
71.new_file
2.file
3.file
3.file
72.new_file

grep file -R | awk -F: '{print $1}' | sort -n  - сортировать по названию
bash/one/one/new.file
bash/one/one/new.file
1.file
2.file
3.file
3.file
71.new_file
72.new_file

grep file -R | awk -F: '{print $1}' | sort -n | xargs cat   - выдать всё содержимое этих файлов
Hello ! My name is nano, i'm a redactor file .txt 
I can open, write, and save .txt files
ctrl - my favorit command
Hello ! My name is nano, i'm a redactor file .txt 
I can open, write, and save .txt files
ctrl - my favorit command
This is file1
This is file2
This is file1
This is file2
This is file1
This is file2
new file content 71
new file content 72

grep This -R | awk -F: '{print $1}' | sort -n | xargs cat   - или только те где есть This
This is file1
This is file2
This is file1
This is file2
This is file1
This is file2
## 19 bashrc, переменные, PATH
### 19.1 Переменные в bash
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/dir$ a="Hello World"
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/dir$ echo $a
Hello World
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/dir$ b="Bye World"
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/dir$ echo $a $b
Hello World Bye World
-- интерполяция
value=25
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ echo $value
25
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ test="Hello World $value"
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ echo $test
Hello World 25
- переменные регистрозависимые 
- переменные живут только в текущей сессии
bash - запускает дочернюю сессию
exit - завершает дочернюю сессию
### Глобальные переменные printenv & export
printenv  - показать все глобальные переменные
export b='Hello again' - создание експортированной переменной (видимая с дочерней сессии)
### 19.2 bashrc (bash run commands) - скрытый каталог в домашней дирректории 
cat ~/.bashrc
Если хочется юзать переменные с любой сесии bash то переменную нужно поместить в bashrc
**cat ~/.bashrc >> .bashrc_bak -- бекап**
echo 'echo "Hello world for every bash session"' >> ~/.bashrc  - добавил запись которая при старте каждой сессии выводит надпись
bash
Hello world for every bash session

echo "A='Hello world'" >> ~/.bashrc  - добавить переменную в bashrc
echo $A    - вызов переменной (только с домашней папки ~./)
Hello world
### 19.3 bash_profile
### 19.4 переменная как команда
whoami   - user name
ME=$(whoami) - если хочется что б прям команда была а не просто стринга
ME=$(whoami)" :)" - такая себе канкатенация ответа + смайл
echo $ME
mvmir :)
Ещё один atempt
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ ME_IN_DIR=$(whoami)" now in "$(pwd)
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ echo $ME_IN_DIR
mvmir now in /home/mvmir
### 19.4 PATH
printenv PATH - показать все глобальные переменные
***добавить директорию в PATH для текущей сессии**
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ export PATH=$PATH:/home/mvmir
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~$ printenv PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/home/mvmir - теперь комманды будут искать и в /home/mvmir
***где лежат все команды**
ls /usr/bin  - много
ls /bin - основные
## 20.1 echo $PS1  строка состояний (mvmir@mvmir-Lenovo-ideapad-320-15IKB:)
\[\e]0;\u@\h: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$
\u@\h: \w\a\ - интересует только эта часть
/u  - user name
/h  - computer name
/w  - current directory
mvmir@mvmir-Lenovo-ideapad-320-15IKB:$
 PS1="[\t][\w] $"
[09:26:13][~] $ - вот такая лажа
[09:26:13][~] $PS1="\t j\j \u@\h:\w\n\$ "

09:28:54 j0 mvmir@mvmir-Lenovo-ideapad-320-15IKB:~
$ 
PS1="# " - теперь будет просто решотка
PS1="$(whoami) in $(pwd)"  - user name and path
- return original
PS1="\[\e]0;\u@\h: \w\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$"
## 20.2 $PS2 == >  разделитель многострочных комманд
ls \<enter>
...
ls \
> |\  - <\enter> - будет перебрасывать на новый рядок с возможностью добавления новых частей комманд
----- две одинаковые комманды
1. ls | xargs -n 1 cat
This is file 1
This is file 2
2. ls \
> |\
> xargs \
> -n 1\
>  cat
This is file 1
This is file 2
------ чему равен PS2?
echo $PS2
>
----- меняем значение PS2
$PS2="this is PS2>"
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/dir/bash/one/two
$ls \
this is PS2>
## 21 Alias
alias зцв='pwd'
работает только в одном сеансе
## 22.1 Source - запускает скрипты
cat bash.run - тут лежит скрипт
echo "Start";
echo "This is a calendar:"
cal -y
echo "And this is a current path:"
pwd
echo "Have a nice day!"
- source bash.run
## 22.2 Source - обновляет файл bashrc (alias)
1. Создание алиаса для текущей сессии
alias goWolfxxxz='cd ~/go/src/github.com/Wolfxxxz'
goWolfxxxz
mvmir@mvmir-Lenovo-ideapad-320-15IKB:~/go/src/github.com/Wolfxxxz$ 
2. Запись алиаса в bashrc
echo "alias goWolfxxxz='cd ~/go/src/github.com/Wolfxxxz'" >> ~/.bashrc
source ~/.bashrc  -- обновить bashrc
3. 
echo "alias goSlovar='cd ~/go/src/github.com/Wolfxxxz/SlovarNV ; ./SlovarNV'" >> ~/.bashrc
source ~/.bashrc  -- обновить bashrc
# 4 shell and bash Удалённое подключение TCP/IP (ssh, scp, sftp,...)
## 23 Установка удалённого контейнера docker
### /ubuntu-ssh/
#### Dockerfile
```
FROM ubuntu:latest

##########################################################
# Install openssh-server
RUN apt update && apt install openssh-server sudo -y

##########################################################
# Create a group “dmdev”
RUN groupadd dmdev

##########################################################
# Create a user “ivan”
RUN useradd -d /home/ivan -m -s /bin/bash -g dmdev ivan
# Set default password 123 for user ivan
RUN echo "123\n123" | passwd ivan

##########################################################
# Create a user “denis”
RUN useradd -d /home/denis -m -s /bin/bash -g dmdev denis
# Create .ssh directory in denis home directory
RUN mkdir -p /home/denis/.ssh
# Copy the ssh public key in the authorized_keys file.
# The id_rsa.pub below is a public key file you get from ssh-keygen.
# They are under ~/.ssh directory by default.
COPY id_rsa.pub /home/denis/.ssh/authorized_keys
# change ownership of the key file.
RUN chown denis /home/denis/.ssh/authorized_keys && chgrp dmdev /home/denis/.ssh/authorized_keys && chmod 640 /home/denis/.ssh/authorized_keys

##########################################################
# Start SSH service
RUN service ssh start
# Expose docker port 22
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
```
#### id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDhcwGAxO3u6Exmg63cLjmYAvWD4PSDOjq+vpgCvKc1gE2ZO469GITPiVbquwh4H4lBkylCuMUzDRFj1N/ZYbsccTGMQRUpIZ0ERdlX3V7sVjWIZbKO0AMZKPkRLWy/P2BXNlpl0EIndW1Fyg6ware7cuATsKzhApD0j8uhQhl+kbBLSZnuJgTrwWeCmX2+KHPiWm+NA4T/dZJRiWXXeIFMk7OmOnMMHPirbwwzlwuiEY7PV4/Sj3P99RpyBo2XjjfMZ7xQdm0tC5wk5XOK28R4JhhevpbJbSgKS9/fCOrDMHLSBu0YNW1Q73pD/RHjVO2StsKKTjNseFuTA54/bDBCM4tKh6jepcyJlo82whD6/ufCoek5690Eao2A+PTujaEwMX5Xn7R5yV80rYJjrWU3L/DUld8huXyG0twJMU7rz3PwGnXbRCHYOpKMfLmBQxuqBYeLhccDuAvyUDn4MvDI9UIQgbml2j+S47df0rNoRlFqJO9wgM3EkrYrnZAfbck= dmatveyenka@Dzianiss-MacBook-Pro-2.local
### docker build -t ubuntu-ssh . 
### docker run -d -p 2022:22 ubuntu-ssh
### ssh -p 2022 ivan@localhost
<ssh -p 2022> порт 2022
<ivan@localhost> name 

### password 123
### exit
### запрос без захода на сервер
ssh -p 2022 ivan@localhost "ls -a"
ivan@localhost's password: 
.
..
.bash_history
.bash_logout
.bashrc
.cache
.profile
## 24 Public key auth
Есть пара ключей один из которых хранится на своеё машине а второй в строчке файла на сервере
### privat key - хранится на своей машине
Где хранится ключ? 
ls ~/.ssh
### public key - лежит на удалённой машине
Где хранится ключ на сервере ?
~/.ssh/authorized_keys
### Где взять ключи?
ssh-keygen
Генерирует пару ключей в папку по умолчанию ~/.ssh
### У меня они уже есть
cat ~/.ssh/id_ed25519.pub
### Заменяю в папке /ubuntu-ssh/ содержимое публичного ключа
ls ~/.ssh
id_rsa.pub
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOOBJCn775RhckiyQs6MRWqYMg5SSf4BqjhpmZGFbEo8 wolfxxxz@gmail.com
### docker build -t ubuntu-ssh .
### docker run -d -p 2022:22 ubuntu-ssh
# 5 bash scripting
## 27 bash script (.sh) или без разширения (best practice)
### first script
touch start-app
nano start-app
#!/bin/bash   - комментарий
echo "Application is starting..." - комманда
source start-app  - запустить скрипт
 Application is starting...
 ```
 chmod +x start-app  - ченж мод для всех в start-app
 chmod - это команда для изменения прав доступа.
 +x - это опция, которая добавляет право на выполнение (исполнение) для файла.
```
Теперь можно запускать скрипт как приложение
./start-app
 Application is starting...
export PATH=$PATH:$PWD
 добавляем в path текущую дерикторию
 теперь из любого места на пк можно вызвать этот скрипт
start-app
 Application is starting...
если хочется запускать всегда то нужно добавить  PATH=$PATH:$PWD в .bashrc или в path
### second script
#### start_app
```
#!/bin/bash

echo "Application is starting..."
# Переменная default_app_name
default_app_name="Spring"
# принято все переменные оборачивать в {}
# Интерполяция переменных (канкатенация)
echo "Application name is ${default_app_name}!"
# Неявные переменные
# Строки "I am string"
# Числа
value=25
# Массивы
array=(1 5 22 190 1240 2050)
# Вывести весь массив @
echo "${array[@]}"
# Вывести третий елемент массива
echo "${array[2]}"
echo "${array[@]}"
# Узнать размер массива #
echo "${#array[@]}"
# Изменить значение в массиве 
array[2]=222
echo "${array[@]}"
# Добавить елементы в массив
array+=(10999 34)
echo "${array[@]}"
# Scan
echo "Enter application name"
read app_name
echo "New application name is ${app_name}!"
```
#### bash
start-app
 Application is starting...
 Application name is Spring!
 1 5 22 190 1240 2050
 22
 1 5 22 190 1240 2050
 6
 1 5 222 190 1240 2050
 1 5 222 190 1240 2050 10999 34
 Enter application name
 second app
 New application name is second app!
## 28 Основы bash script (shell expansions)
### Brace expansions          -> {x...y} {a,b,c}
#### {x..y}
ll
 итого 8
 drwxrwxr-x  2 mvmir mvmir 4096 мая 23 16:47 ./
 drwxr-xr-x 37 mvmir mvmir 4096 мая 23 16:47 ../
 -rw-rw-r--  1 mvmir mvmir    0 мая 23 16:47 file1
 -rw-rw-r--  1 mvmir mvmir    0 мая 23 16:47 file2
 -rw-rw-r--  1 mvmir mvmir    0 мая 23 16:47 file3
 -rw-rw-r--  1 mvmir mvmir    0 мая 23 16:47 file4
ll file{1..4}
-rw-rw-r-- 1 mvmir mvmir 0 мая 23 16:47 file1
-rw-rw-r-- 1 mvmir mvmir 0 мая 23 16:47 file2
-rw-rw-r-- 1 mvmir mvmir 0 мая 23 16:47 file3
-rw-rw-r-- 1 mvmir mvmir 0 мая 23 16:47 file4
#### range{x..y}
echo {a..d}
a b c d
#### {a,b,c}
rm file{1,2,4,5}
ll
 итого 8
 drwxrwxr-x  2 mvmir mvmir 4096 мая 23 17:02 ./
 drwxr-xr-x 37 mvmir mvmir 4096 мая 23 16:47 ../
 -rw-rw-r--  1 mvmir mvmir    0 мая 23 16:47 file3
### Tilde expansions          ->  ... $HOME
var=~/foo
echo ${var}
 /home/mvmir/foo
### Parameter & variable expansions ->  $var ${var} ${var:=default}
#### переменные всегда в скобках
echo $varb lac k/l ist
 lac k/l ist
echo ${var}/black/list
 /home/mvmir/foo/black/list
#### дефолт значения
value=
echo "${value}"
Если value пустое то вывести 45
echo "${value:=45}"
45
value=24
echo "${value:=45}"
24
### Command substitution      -> $(command) `command`
echo "The first file in directory ${PWD} is $(ls | head -n1)"
The first file in directory /home/mvmir/lek28 is bestFile
### Arithmetic expansion      -> $((expression))
типа калькулятор
echo "$(( 5 * 3 / 7 ))"
2
echo "$(( (5 **2 + 3) / 7 ))"
4
### Filename expansion        -> * ? [..]
ls log* - <*> - любые символы
loga  logb  logroisdos
ls log? - только один символ
loga  logb
ls log[b,r] - только один символ из []
logb
### Word splitting            -> <space><tab><newline>
разбивает строку на пробелы - всячески избегать
## 29 bool
### logic $?
0 - true
!=0 - false
$? - show bool
cat start-app  - существующий файл
echo $?
 0
cat start-app2  - не существующий
 cat: start-app2: Нет такого файла или каталога
echo $?
 1
cat start-app | grep 1240
 array=(1 5 22 190 1240 2050)
echo $?
 0
### use logic <test>
#### man test > <
test 10 -lt 5 - <-lt> == 10 < 5
echo $?
 1
#### -f (file)
test -f Dockerfile - <-f> - есть такой файл в папке?
echo $?
 1
[[ -f new.file ]]  - тоже самое - другой синтаксис
echo $?
 0
#### -d (folder)
test -d newFolder  - <-d> - для папок
echo $?
 0
[[ -d newFolder ]]  - тоже самое - другой синтаксис
echo $?
 0
#### (( char ))
(( 10 > 5 * 3 ))
echo $?
 1
## 30 if else and case
### elif
value=0
if (( value > 0 )); then
	echo "${value} is positive"
`# elif == else if`
elif (( value < 0 )); then
	echo "${value} is negative"
else
	echo "${value} is zero"
fi
### switch case
echo "Hello bash switch case"
echo "Enter any program:"
read program
case "${program}" in
clean)
	echo "Clean is invoked"
	;;
build)
	echo "Build is invoked"
	;;
*)
	echo "${program} is not supported by this application"
	exit 2
	;;
esac
## 31 циклы
### while
counter=10
while (( counter > 0 )); do
	echo "Counter value: ${counter}"
	(( counter-- ))
done
### until (перевёрнутый while) <=
counter=10
until (( counter <= 0 )); do
	echo "Counter value: ${counter}"
	(( counter-- ))
done
### for
```
array=(5 90 83 23 67)
for (( i = 0; i < "${#array[@]}"; i++ )); do
   echo "Index value: ${i}. Array value: ${array[i]}"
done
```
### for each
```
array=(5 90 83 23 67)
for value in {1..5}; do
   echo "For each loop: ${value}"
done
--------------
array=(5 90 83 23 67)
for value in "${array[@]}"; do
   echo "For each loop: ${value}"
done
### for + if
files=(loga logb logc new.file)
for file in "${files[@]}"; do
	if [[ -f "${file}" ]]; then
		ls -l "${file}"
	else
	    echo "File doesn't exist: ${file}"
	fi
done
-----------------
files=(loga logb logc new.file)
for file in "${files[@]}"; do
	if [[ -f "${file}" ]]; then
		ls -l "${file}"
	else
		echo "File doesn't exist: ${file}"
		echo "Do you want to create ${file}? y/n"
		read createFileAnswer
		if [[ "${createFileAnswer}" = "y" ]]; then
			touch "${file}"
		fi
	fi
done
```
### for Word splitting
for key in ~/keys/id_rsa* ; do
    echo "Key file: ${key}"
    cat "${key}"
done
## 32 func
### Примичания
1. Функцию можно вызвать после обьявления (порядок имеет значение)
2. Входящие переменные нельзя обьявить в функции
3. При обьявлении функции можно не писать function hello(){}, hello() {}
4. Все переменные которые находятся в скрипте (вне зависимости от того где они находятся(внутри функции или цикла)) **ДОСТУПНЫ ДЛЯ ВСЕГО СКРИПТА** == **переменные лучше обьявлять local**
5. return может вернуть только bool статус код $?
### Устаревший способ
function make() {
  ...
}
### Рабочий способ
**file:func**

echo "Script all arguments $*"

clean() {
	echo "going to clean directory..."
	echo "Function first argument $1" # первый входящий аргумент
	echo "Function all arguments $@" # аргументы разделены
	echo "Function all arguments $*" # все аргументы как строка
	index=0
	for arg in "$@" ; do 
		echo "Index: ${index}. Argument: ${arg}"
		(( index++ ))
	done
}

clean program1 25 98 arg4

**bash: source func**
Script all arguments 
going to clean directory...
Function first argument program1
Function all arguments program1 25 98 arg4
Function all arguments program1 25 98 arg4
Index: 0. Argument: program1
Index: 1. Argument: 25
Index: 2. Argument: 98
Index: 3. Argument: arg4
### Непонятная передача аргументов прям в bash
**bash: source func f1 f2 gt**
Script all arguments f1 f2 gt
going to clean directory...
Function first argument program1
Function all arguments program1 25 98 arg4
Function all arguments program1 25 98 arg4
Index: 0. Argument: program1
Index: 1. Argument: 25
Index: 2. Argument: 98
Index: 3. Argument: arg4
### Возвращает статус
echo "Script all arguments $*"

clean() {
	echo "going to clean directory..."
	echo "Function first argument $1" # первый входящий аргумент
	echo "Function all arguments $@" # аргументы разделены
	echo "Function all arguments $*" # все аргументы как строка
	index=0
	for arg in "$@" ; do 
		echo "Index: ${index}. Argument: ${arg}"
		(( index++ ))
	done
	return 5
}

clean program1 25 98 arg4
echo "Clean function status: $?"
### Странный возврат (просто вывод строки из функции)
echo "Script all arguments $*"

clean() {
	echo "going to clean directory..."
	echo "Function first argument $1" # первый входящий аргумент
	echo "Function all arguments $@" # аргументы разделены
	echo "Function all arguments $*" # все аргументы как строка
	index=0
	for arg in "$@" ; do 
		echo "Index: ${index}. Argument: ${arg}"
		(( index++ ))
	done
	echo "return value from function" # типа любое значение из функции передать через echo
}

result=$(clean program1 25 98 arg4) #result = тому что передало echo
echo "Clean function result: ${result}" 
### local переменные
echo "Script all arguments $*"

clean() {
	echo "going to clean directory..."
	echo "Function first argument $1" # первый входящий аргумент
	echo "Function all arguments $@" # аргументы разделены
	echo "Function all arguments $*" # все аргументы как строка
	local index=0
	for arg in "$@" ; do 
		echo "Index: ${index}. Argument: ${arg}"
		(( index++ ))
	done
	echo "return value from function"
}

result=$(clean program1 25 98 arg4)
echo "Clean function result: ${result}"
### comments
m ##############
## 33 Practice
https://github.com/dmdev2020/bash-starter/tree/lesson-33
### application file
```
#!/bin/bash

if [[ -f "./env.sh" ]]; then
  echo "Use env variables from file ${PWD}/env.sh"
  source ./env.sh
fi

DB_CONTAINER_NAME="spring-postgres"
workDir="${WORKING_DIRECTORY:=~/Workspace}"

help() {
  echo "
  Usage:
    ./application init - init working directory and database
    ./application clean - clean working directory and stop database
    ./application build - run JUnit tests to check app health (-skipTests arg to skip tests) and build jar
    ./application up - launch application
  "
}

up() {
  cd "${workDir}/spring-starter/build/libs" || exit
  java -jar spring-starter-*.jar
}

build() {
  cd "${workDir}/spring-starter" || exit

  ./gradlew clean

  if [[ "$1" = "-skipTests" ]] || ./gradlew test; then
    echo "Application is building..."
    ./gradlew bootJar
  else
    echo "Tests failed. See test report or send -skipTests arg to skip tests"
  fi
}

clean() {
#  remove working directory
  echo "Removing working directory ${workDir}..."
  rm -rf "${workDir}"

#  stop docker container (PostgreSQL)
  if docker ps | grep "${DB_CONTAINER_NAME}"; then
    echo "Stopping container name ${DB_CONTAINER_NAME}..."
    docker stop "${DB_CONTAINER_NAME}"
  fi
}

init() {
#  init working directory
  mkdir -p "${workDir}"
  cd "${workDir}" || exit

# git clone
  if [[ ! -d "spring-starter" ]]; then
    git clone git@github.com:dmdev2020/spring-starter.git
  fi
  cd "spring-starter" || exit
  git checkout lesson-125

#  PostgreSQL
  docker pull postgres
#  $(docker ps -a | grep "${DB_CONTAINER_NAME}")
  if docker ps -a | grep "${DB_CONTAINER_NAME}"; then
    docker start "${DB_CONTAINER_NAME}"
  else
    docker run --name "${DB_CONTAINER_NAME}" \
        -e POSTGRES_PASSWORD=pass \
        -e POSTGRES_USER=postgres \
        -e POSTGRES_DB=postgres \
        -p 5433:5432 \
        -d postgres
  fi
}

case $1 in
help)
  help
  ;;
init)
  init
  ;;
clean)
  clean
  ;;
build)
  build $2
  ;;
up)
  up
  ;;
*)
  echo "$1 command is not valid"
  exit 1
  ;;
esac
```
### env
WORKING_DIRECTORY=~/Workspace
### Ссылки
https://www.gnu.org/software/bash/manual/bash.html
https://google.github.io/styleguide/shellguide.html
## scp - копировать файлы с удалённого сервера
man scp
scp — secure copy (remote file copy program)
## history
```
tail ~/.bash_history
history - история
!! - повтор последней комманды
!-3  - третья команда от последней
<ctr + r> - поиск по истории комманд
```
echo $HISTSIZE  - количество запросов в памяти
history -c    - history clear
## gpt
chmod - Команда используется для изменения прав доступа к файлам и директориям
chmod u+x script.sh  - Установить право выполнения для владельца файла
chmod g-w файл.txt   - Удалить право записи для группы:
chmod u+rwx,g+r,o+r файл.txt  - Установить права чтения, записи и выполнения для владельца, и права чтения для группы и остальных пользователей:
chmod -R u+r,go-w /путь/к/директории  - Рекурсивно изменить права доступа для всех файлов и поддиректорий в указанной директории:
## /dev/null - чёрная дыра (поток просто исчезает)
find . -name "info" 2> /dev/null
## My notes 
ctr + alt + t  - bash
exit - close bash
code SlovarNV - vsCode open SlovarNV
## update ubuntu
sudo apt-get update  - <sudo> - подымает права до superUser
sudo apt-get dist-upgrade - делает тоже самое, что upgrade плюс выполяет «умное» разрешения конфликтов версий пакетов
sudo do-release-upgrade  - обновляет версию системы, например с Ubuntu 13.04 до Ubuntu 13.10.
sudo do-release-upgrade -d  - включая бетта версии **DANGER** - есть подводные камни при установке
https://tyapk.ru/blog/post/how-to-upgrade-ubuntu-in-terminal
## Телеграмм
nohup telegram-desktop >/dev/null 2>&1 &
>/dev/null 2>&1 - перенаправть все логи и тп в жопу
& - отрывает телегу от терминала совсем
## Crome
google-chrome
## backup bashrc
cp ~/.bashrc ~/bashrc_save
## link 
ln -s ~/go/src/github.com/Wolfxxxz/ Link_Wolfxxxz
## все установленные проги
1: Для просмотра всех установленных программ:
apt list --installed
2: Для поиска конкретной программы:
apt search <название_программы>
3: Для просмотра списка доступных программ с постраничным выводом:
apt list | less
## 2: Настроить нужный софт (в нашем случае vim)
apt-get update
apt-get install apt-file
apt-file update
apt-get install vim
echo hello vim > test.txt
vim test.txt
# в 
command 'win' from deb wily
  command 'link' from deb coreutils
  command 'din' from deb din
  command 'tin' from deb tin
  command 'ln' from deb coreutils
  command 'lie' from deb lie
  command 'lein' from deb leiningen
  command 'lid' from deb id-utils
#в