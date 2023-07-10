# Установка SQL and PGAdmin 4 
## Настройка PgAdmin4 ubuntu
### Подключение локально
sudo netstat -tlnp | grep postgres
Из вывода команды sudo netstat -tlnp | grep postgres видно, что PostgreSQL прослушивает три разных порта на локальном хосте (127.0.0.1):

Порт 5432, связанный с процессом 1004/postgres.
Порт 5433, связанный с процессом 7038/postgres.
Порт 5434, связанный с процессом 1005/postgres.
В этом случае, вам необходимо выбрать один из этих портов для подключения к PostgreSQL. Например, если вы хотите использовать порт 5432, следующие параметры подключения могут быть использованы:

Host name/address (Имя хоста/адрес): 127.0.0.1
Port (Порт): 5432 (или любой другой выбранный порт)
Maintenance database (Обслуживаемая база данных): Имя базы данных
Username (Имя пользователя): Имя пользователя
Password (Пароль): Пароль пользователя
Пожалуйста, убедитесь, что вы используете правильные параметры подключения, включая выбранный порт, для подключения к PostgreSQL через pgAdmin или другой инструмент управления базами данных.
### Подключение к контейнеру
521647f03130   my-postgres-image   "docker-entrypoint.s…"   23 seconds ago   Up 21 seconds   0.0.0.0:5435->5432/tcp, :::5435->5432/tcp   my-postgres-container
можешь конкретно пример показывать на этих данных
ChatGPT
Исходя из информации, которую вы предоставили, у вас есть контейнер с идентификатором 521647f03130 и именем my-postgres-container, основанный на образе my-postgres-image. Контейнер пробрасывает порт 5432 в контейнере на порт 5435 хоста.

Чтобы подключиться к базе данных PostgreSQL внутри этого контейнера с помощью pgAdmin, выполните следующие шаги:

Запустите pgAdmin на вашем компьютере.

Щелкните правой кнопкой мыши на "Servers" (Серверы) в панели навигации слева и выберите "Create" (Создать) -> "Server" (Сервер).

В окне "Create - Server" (Создать - Сервер) вкладка "General" (Общее), укажите имя сервера в поле "Name" (Имя).

Перейдите на вкладку "Connection" (Подключение) и введите следующую информацию:

Host name/address (Имя хоста/адрес): localhost
Port (Порт): 5435
Maintenance database (Обслуживаемая база данных): Имя базы данных, к которой вы хотите подключиться.
Username (Имя пользователя): Имя пользователя для доступа к базе данных.
Password (Пароль): Пароль пользователя для доступа к базе данных.
Нажмите кнопку "Save" (Сохранить) для создания сервера.
## Скрипт с бд по ссылке
https://onedrive.live.com/?authkey=%21AGZmnLWD0KNfvho&id=D0486487037850AB%21955870&cid=D0486487037850AB&parId=root&parQt=sharedby&o=OneUp

## Качаем с сайта
https://www.postgresql.org/download/

# Типы данных ОСНОВНЫЕ и приобразование
## Типы данных
### Integral Numbers (Целые числа)
Name          Bytes        Description
smallint      2            small range integer             от -32.768 до 32.767
integer       4            typical choice for integer      от -2.147.483.648 до 2.147.483.648
bigint        8            large range integer             от -9.223.372.036.854.775.808 до 9.223.372.036.854.775.807
smallserial   2           Autoincrementing serial of table  от 0 до 32.767
serial        4           Autoincrementing serial of table  от 0 до 2.147.483.648
bigserial     8           Autoincrementing serial of table  от 0 до 9.223.372.036.854.775.807
### Real Numbers (с плавающей точкой)
Name                                Bytes        Description
decimal/numeric                     variable     user specified - ИСПОЛЬЗОВАТЬ для ДЕНЕГ
real/ float4                        4            user specified - точность 6 знаков после запятой
double precision/ float 8 / float   8            user specified - точность 15 знаков после запятой
### Characters (Символы) char(5) - 5 символов
Name          Bytes        Description
char          variable     лажа с длинной
varchar       variable     лажа с длинной
text          variable    Любая длинна
### Logical 
boolean true/false
### Temporal (date, time, timestamp, interval, timestamptz)
Name            Bytes        Description     Range
date            4            only date
time            8            only time
timestamp       8            date and time
interval        16           between timestamp (разница между timestamp)
timestamptz     8            timestamp + timezona (хранит данные по часовым поясам)
### Другие типы данных + NULL
Arrays
json
xml
Геометрические
Custom
NULL - на него нельзя множить:)
## Преобразование
SELECT float4 '0.5' AS float8;  --преобразование типа данных
SELECT 50! AS big_factorial;  ----преобразование типа данных (факториал ставится после числа, int и все остальные после)
SELECT CAST(50 AS bigint)! AS big_factorial;
SELECT 'abc' || 1; -- cancatенирование - склеивание (abc1 text)ж
SELECT ' 10  ' = 10; -- true boolean
## ТИПЫ ДАННЫХ and func
```
 ОПЕРАТОРЫ    -- SELECT * FROM pg_operator  -- '+', '-' и тд. и тп.
 ТИПЫ ДАННЫХ ---SELECT * FROM pg_proc --- там лежат все функции + наверное процессы - странная таблица

 ЯВНОЕ ПРЕОБРАЗОВАНИЕ 
   -- SELECT (expression AS target_type)  -- SELECT float4 '0.5' AS float8;  SELECT 50! AS big_factorial
   --CAST (expression AS target_type) -- совместимо стандартом  (SELECT 50! AS big_factorial) или (CAST(50 AS bigint)! AS big_factorial)
   --expression::target_type -- несовместимо со стандартом    
float8 - 8 знаков после запятой
integer(int) - целые числа
```
### Пример 1
```
CREATE OR REPLACE FUNCTION type_testing(money_val float8) RETURNS void AS $$
BEGIN
   RAISE NOTICE 'ran %', money_val;       ---обработка ошибок (RAICE NOTICE) замечания = (RAISE[level]'message(%)', arg_name;)
END
$$ LANGUAGE plpgsql;

SELECT type_testing(0.5); -- смотреть messages неявное приобразование -- (SELECT float8 '0.5')
SELECT type_testing(0.5::float4); -- конвертация в float4 --(SELECT float8 '0.5'::float4)
```
### Пример 2
```
CREATE OR REPLACE FUNCTION type_testing2(money_val int) RETURNS void AS $$
BEGIN
   RAISE NOTICE 'ran %', money_val; 
END
$$ LANGUAGE plpgsql;

SELECT type_testing2(0.5::int); -- конвертация в int(без плавающей точки, только целые) принудительная конвертация приводит к потери точности

SELECT float8 '2'; -- тоже что -- SELECT type_testing(0.5);

SELECT float4 '0.5'::float8 -- преобразование то же что для функции (SELECT type_testing(0.5::float4))

SELECT type_testing2(CAST(0.5 AS int));

SELECT type_testing2(0.4::int) -- 0 - int 0.4 = 0 ; 0.5 = 1

SELECT type_testing2(1.5);
SELECT type_testing2(1.5::int);
SELECT type_testing2(1.5::numeric::int);

SELECT type_testing2(3::int);

SELECT float4 '0.5' AS float8;  --преобразование типа данных

SELECT 50! AS big_factorial;  ----преобразование типа данных (факториал ставится после числа, int и все остальные после)

SELECT CAST(50 AS bigint)! AS big_factorial;

SELECT 'abc' || 1; -- cancatенирование - склеивание (abc1 text)ж

SELECT ' 10  ' = 10; -- true boolean
```

# Create and DELETE db 
## Create and Drop 
### Create
CREATE DATABASE "Dictionary"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

COMMENT ON DATABASE "Dictionary"
    IS 'Create db pgAdmin 4';

Или командой

CREATE DATABASE Dictionary

### DROP db
- Stop all connection db Dictionary1
SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'Dictionary1'
AND pid <> pg_backend_pid()

- DROP DATABASE Dictionary1

## запустить db northwind
- CREATE DATABASE northwind
- run sql script northwind
Скопировать в query editor скрипт базы 


# Table Create, Delete, Insert into, ALTER TABLE  
## Create table
CREATE TABLE book
(
	book_id int PRIMARY KEY,
	title text NOT NULL,
	isbn text NOT NULL
);

CREATE TABLE author
(
	author_id int PRIMARY KEY,
	full_name text NOT NULL,
	rating real
);

- PRIMARY KEY
CREATE TABLE book_author
(
	book_id int REFERENCES book(book_id),
	author_id int REFERENCES author(author_id),
	
	CONSTRAINT book_author_pkey PRIMARY KEY (book_id, author_id) -- composite key
);
## DROP TABLE IF EXISTS book;
DROP TABLE IF EXISTS book;
## Внести данные INSERT into book
INSERT into book
values
(1, 'Book for Dumnies', '123456'),
(2, 'Book for Smart Guyis', '7890123'),
(3, 'Book for Happy Peaple', '4567890'),
(4, 'Book for Unhappy Peaple', '1234567');

INSERT INTO author
VALUES
(1, 'Bob', 4.5),
(2, 'Alice', 4.8),
(3, 'John', 4.7);

INSERT INTO book_author
VALUES
(1,1),
(2,1),
(3,1),
(3,2),
(4,1),
(4,2),
(4,3);

INSERT into book
VALUES
(1, 'The Diary of a Young Girl', '0199535566'),
(2, 'Pride and Prejudice', '9780307594006'),
(3, 'To kill a Mockingbird', '0446310786'),
(4, 'Gutsy Women', '1501178415'),
(5, 'War and Peace', '1788886526');

INSERT INTO publisher
VALUES
(1, 'Everyman``s Library', 'NY'),
(2, 'Oxford U P', 'NY'),
(3, 'Grand Central Pub', 'Washington'),
(4, 'Simon & Schluster', 'Chikago')
## Добавить поля и зависимости ALTER TABLE
ALTER TABLE book
ADD COLUMN fk_publisher_id;

ALTER TABLE book
ADD CONSTRAINT fk_book_publisher FOREIGN KEY(fk_publisher_id) REFERENCES publisher(publisher_id)

# SELECT (выборки)
## SELECT  (Показать __ с таблицы __)
SELECT *
FROM customers;

SELECT contact_name, city
FROM customers;
## DISTINCT (уникальные)
SELECT DISTINCT country, city
FROM employees;

SELECT DISTINCT city
FROM customers;

SELECT DISTINCT city, country
FROM customers;
## количество строк COUNT(*)
SELECT COUNT(*) всего
FROM orders;

SELECT COUNT(customer_id)
FROM customers;
## количество уникальных COUNT DISTINCT
SELECT COUNT(DISTINCT country)
FROM customers;
## math + - sum avg(среднее) min max
SELECT product_id, product_name, unit_price * units_in_stock
FROM products

SELECT order_id, shipped_date - order_date
FROM orders;

SUM
SELECT SUM(units_in_stock)
FROM products
WHERE discontinued <>1

AVG среднее
SELECT AVG(unit_price)
FROM products
WHERE discontinued <> 1  

SELECT MIN(order_date)
FROM orders
WHERE ship_city = 'London'

SELECT MAX(order_date)
FROM orders
WHERE ship_city = 'London'

# WHERE (Фильтр)
## Простые выборки <, >, =, >=, <>  
SELECT company_name, contact_name, phone
FROM customers
WHERE country = 'USA';

SELECT *
FROM products
WHERE unit_price > 20;


SELECT *
FROM customers
WHERE city <>'London';


SELECT *
FROM orders
WHERE order_date >= '1998-03-01';
## AND (логическое И)
если оба 'true' то 

SELECT *
FROM table
WHERE condition1 AND condition2

SELECT *
FROM products
WHERE unit_price > 20 AND units_in_stock > 40;

- скобки вовремя
SELECT *
FROM orders
WHERE shipped_date > '1998-04-30' AND (freight < 75 OR freight > 150);
##  OR (логичиское или)
если один из 'true'

SELECT *
FROM table
WHERE condition1 OR condition2

SELECT *
FROM customers
WHERE city = 'Berlin' OR city = 'London' OR city = 'San Francisco';
## BETWEEN (логическое между)
            Между или от 20.. до 40..
SELECT *
FROM orders
WHERE freight BETWEEN 20 AND 40;
               включая 20 и 40 >= ... =<
SELECT *
FROM orders
WHERE freight >= 20 AND freight <= 40;
## IN, NOT IN (все из списка (где категория = или != id))

- IN(...) - все из списка где ...
SELECT *
FROM products
WHERE category_id IN(1, 2, 3, 5, 7)

- NOT IN(,,,) - кроме этих
SELECT *
FROM customers
WHERE country NOT IN('Mexico', 'Germany', 'USA', 'Canada')
## LIKE поиск в строке
             LIKE 'U%'     строка начинается с U...
             LIKE '%a'     строка заканчивается на ...a
             LIKE '%John%  внутри строки  ...John...
             LIKE '_oh_    фиксировано - .oh. 
             LIKE '_oh%'   фиксировано с одной стороны - .oh... 
1 SELECT *
FROM orders
WHERE ship_country LIKE'U%'
## IS NULL & NOT NULL
SELECT ship_city, ship_region, ship_country
FROM orders
WHERE ship_region IS NULL

SELECT ship_city, ship_region, ship_country
FROM orders
WHERE ship_region IS NOT NULL

# HAVING фильтр  после WHERE (второй фильтр)

SELECT category_id, SUM(unit_price * units_in_stock)
FROM products
WHERE discontinued <> 1
GROUP BY category_id
HAVING SUM(unit_price * units_in_stock) > 5000
ORDER BY SUM(unit_price * units_in_stock) DESC

# ORDER BY and GROUP BY (упорядочить и группировать)
## ORDER BY - упорядочить (ASC - стоит по умолчанию)
    ASC по возростанию
    DESC по убыванию
SELECT country
FROM customers
ORDER BY country ASC

SELECT DISTINCT country
FROM customers
ORDER BY country DESC

SELECT DISTINCT country, city
FROM customers
ORDER BY country DESC, city ASC
##  GROUP BY
SELECT ship_country, COUNT(*)                         COUNT(*) 
FROM orders
WHERE freight > 50
GROUP BY ship_country                               GROUP BY ship_country
ORDER BY COUNT(*) DESC                              упорядочить

SELECT ship_country, COUNT(*)
FROM orders                                          по кол-ву текста
GROUP BY ship_country

SELECT category_id, SUM(units_in_stock)            по одной сумма
FROM products
GROUP BY category_id                               по второй групировка

SELECT supplier_id, SUM(unit_price)
FROM products
GROUP BY supplier_id
ORDER BY SUM(unit_price) DESC
LIMIT 10
## GROUP BY ROLLUP GRAND TOTAL
SELECT supplier_id, SUM(units_in_stock)
FROM products
GROUP BY ROLLUP (supplier_id);  --Появилась сумма всех строк GRAND TOTAL
## GROUP BY CUBE 
добавилась отдельно групировка по category_id
SELECT supplier_id, category_id, reorder_level, SUM(units_in_stock)
FROM products
GROUP BY CUBE (supplier_id, category_id, reorder_level)
ORDER BY supplier_id, category_id NULLS FIRST;
## GROUPING SET, ROLLUP, CUBE Подитоги и общие итоги
### Термины
 -- Как сформировать подитоги и общие итоги
 -- GROUPING SET, ROLLUP, CUBE

- Набор столбцов в GROUP BY - и есть GROUPING SET
- GROUP BY GROUPING SETS((col_a), (col_a, col_b)) -
  вернёт группировку по (col_a) и по (col_a, col_b)
- ROLLUP генерирует aгрегативный набор для иерархии значений
  в столбцах указанных в скобках(в порядке следования) - 
  просто сокращённый вариант GROUPING SET:
  ROLLUP(col_a, col_b) - полный эквивалент предидущей группировки
- CUBE генерирует агрегативный набор для всех комбинаций значений
  в столбцах указанных в скобках (порядок следования не важен)

SELECT *
FROM products;
### группировка по одному столбцу
SELECT supplier_id, SUM(units_in_stock)
FROM products
GROUP BY supplier_id
ORDER BY supplier_id;
### групировка по двум столбцам
SELECT supplier_id, category_id, SUM(units_in_stock)
FROM products
GROUP BY supplier_id, category_id
ORDER BY supplier_id;
### группировка одновременно и по (supplier_id), (supplier_id, category_id)
- то есть добавилась общая сумма по supplier_id
- NULLS FIRST - сумма идёт первой
SELECT supplier_id, category_id, SUM(units_in_stock)
FROM products
GROUP BY GROUPING SETS ((supplier_id), (supplier_id, category_id))
ORDER BY supplier_id, category_id NULLS FIRST
### ROLLUP
Появилась сумма всех строк GRAND TOTAL
SELECT supplier_id, SUM(units_in_stock)
FROM products
GROUP BY ROLLUP (supplier_id);

SELECT supplier_id, category_id, SUM(units_in_stock)
FROM products
GROUP BY ROLLUP (supplier_id, category_id)
ORDER BY supplier_id, category_id NULLS FIRST;
--+reorder_level --
SELECT supplier_id, category_id, reorder_level, SUM(units_in_stock)
FROM products
GROUP BY ROLLUP (supplier_id, category_id, reorder_level)
ORDER BY supplier_id, category_id NULLS FIRST;
### CUBE
--добавилась отдельно групировка по category_id
SELECT supplier_id, category_id, SUM(units_in_stock)
FROM products
GROUP BY CUBE (supplier_id, category_id);
--+reorder_level--
SELECT supplier_id, category_id, reorder_level, SUM(units_in_stock)
FROM products
GROUP BY CUBE (supplier_id, category_id, reorder_level)
ORDER BY supplier_id, category_id NULLS FIRST;




# LIMIT  количество результатов
SELECT product_name, unit_price
FROM products
LIMIT 10

SELECT product_name, unit_price
FROM products
WHERE discontinued <> 1
ORDER BY unit_price DESC                                упорядочить по DESC
LIMIT 10                                                10 результатов

# UNION and UNION ALL ; INTERSEKT and EXCEPT  (обьединение и исключения) 
## UNION and UNION ALL - обьединение 
SELECT country
FROM customers
UNION
SELECT country
FROM employees

SELECT country
FROM customers
UNION ALL
SELECT country
FROM employees
## INTERSEKT 
 пересечение только те что есть в обоих таблицах
SELECT country
FROM customers
INTERSECT
SELECT country
FROM suppliers
## EXCEPT and EXPECT ALL - включить повторы
 исключение из кастом исключили саплаер 
 
SELECT country
FROM customers
EXCEPT
SELECT country
FROM suppliers

SELECT country
FROM customers
EXCEPT ALL
SELECT country
FROM suppliers

# СОЕДИНЕНИЯ JOIN
- INNERN JOIN - внутреннее соединение DISTINCT
- LEFT JOIN, RIGHT JOIN - правое левое соединение (внешние)
- FULL JOIN полное соединение
- CROSS JOIN декартово соединение
- SELF JOIN соединение на саму себя
## INNERN JOIN = JOIN (он же простой JOIN)
SELECT product_name, suppliers.company_name, units_in_stock
FROM products
INNER JOIN suppliers ON products.supplier_id = suppliers.supplier_id           
ORDER BY   DESC

SELECT category_name, SUM(unit_price*units_in_stock)
FROM products
INNER JOIN categories ON products.category_id = categories.category_id
WHERE discontinued <> 1
GROUP BY category_name
HAVING SUM(unit_price * units_in_stock) > 5000
ORDER BY SUM(unit_price * units_in_stock) DESC
## LEFT JOIN, RIGHT JOIN - правое левое соединение (внешние)
SELECT company_name, product_name
FROM suppliers
LEFT JOIN products ON suppliers.supplier_id = products.supplier_id

SELECT company_name, order_id 
FROM customers
LEFT JOIN orders ON orders.customer_id = customers.customer_id
WHERE order_id IS NULL

SELECT company_name, order_id
FROM orders
RIGHT JOIN customers ON orders.customer_id = customers.customer_id
WHERE order_id IS NULL
## FULL JOIN
SELECT company_name, order_id
FROM orders
FULL JOIN customers ON orders.customer_id = customers.customer_id
WHERE order_id IS NULL
## CROSS JOIN
SELECT company_name, order_id
FROM orders
CROSS JOIN customers
## SELF JOIN иерархия соединение на саму себя (РЕКУРСИЯ)
CREATE TABLE employee (
	employee_id INT PRIMARY KEY,
	first_name VARCHAR (255) NOT NULL,
	last_name VARCHAR (255) NOT NULL,
	manager_id INT,
	FOREIGN KEY (manager_id) REFERENCES employee (employee_id));
INSERT INTO employee (
	employee_id,
	first_name,
	last_name,
	manager_id)
VALUES
	(1, 'Windy', 'Hays', NULL),
	(2, 'Ava', 'CHristensen', 1),
	(3, 'Hassan', 'Conner', 1),
	(4, 'Anna', 'Reeves', 2),
	(5, 'Sau', 'Norman', 2),
	(6, 'Kelsie', 'Hays', 3),
	(7, 'Tory', 'Goff', 3),
	(8, 'Salley', 'Lester', 3);

SELECT e.first_name || ' ' || e.last_name AS employee,
       m.first_name || ' ' || m.last_name AS manager
FROM employee e
LEFT JOIN employee m ON m.employee_id = e.manager_id
ORDER BY manager
## USING (сократить написание JOIN)
______________________________________________
SELECT COUNT(company_name)
FROM orders
JOIN customers ON orders.customer_id = customers.customer_id
- ИЛИ использовать JOIN USING
SELECT COUNT(company_name)
FROM orders
JOIN customers USING(customer_id)
- При это название полей должно совпадать
_______________________________________________

"USING(order_id)" = "ON orders.order_id = order_details.order_id"

SELECT contact_name, company_name, phone, first_name, last_name, title,
       order_date, product_name, ship_country, products.unit_price, quantity, discount
FROM orders
JOIN order_details USING(order_id) --- ON orders.order_id = order_details.order_id
JOIN products USING(product_id) --- ON order_details.product_id = products.product_id
JOIN customers USING(customer_id) --- ON orders.customer_id = customers.customer_id
JOIN employees USING(employee_id) --- ON orders.employee_id = employees.employee_id
WHERE ship_country = 'USA'

# AS псевдоним (называем столбик) 
 WHERE & HAVING нельзя задать через псевдоним

SELECT COUNT(*) AS emploise_count
FROM employees

SELECT COUNT(DISTINCT country) AS country_distinct
FROM employees

SELECT category_id, SUM(unit_price * units_in_stock) AS total_price
FROM products
WHERE discontinued <> 1
GROUP BY category_id
HAVING SUM(unit_price * units_in_stock) > 5000
ORDER BY total_price DESC

 CONCAT(first_name ' ', last_name)
 Обьединить имя и фамилию в одну ячейку

# ПОДЗАПРОС
## 1 Пример подзапрос в WHERE
### найти все компании поставщиков (suppliers) из тех стран в которых есть заказчики (customers)
--шаг 1 подзапрос
SELECT country
FROM customers

--шаг 2 вставляеm в WHERE
SELECT company_name
FROM suppliers
WHERE country IN (SELECT DISTINCT country
                  FROM customers); -- из списка стран customers
--тоже самое с помощью JOIN
SELECT DISTINCT suppliers.company_name
FROM suppliers
JOIN customers USING(country); --останутся только те что есть в обеих country (INNER JOIN)
### Вывести позиции по которым товара на остатке больше чем в среднем 
WHERE подзапрос
SELECT AVG(units_in_stock)
FROM products

SELECT product_name, units_in_stock
FROM products
WHERE units_in_stock > (SELECT AVG(units_in_stock)
                        FROM products);
ORDER BY units_in_stock
## 2 ПОДЗАПРОС в LIMIT
--  вывести сумму единиц товара разбитых на группы и 
--  лимитировать результирующий набор числом которое необходимо
--  вычислить найдя наименьший product_id + 4
SELECT category_name, SUM(units_in_stock)
FROM products
JOIN categories USING(category_id)
GROUP BY category_name
ORDER BY SUM(units_in_stock)
LIMIT  (SELECT MIN(product_id) + 4 FROM products )
## WHERE EXISTS (подзапрос когда каждая последующая строка сравнивается с подзапросом (одновременно делает DISTINCT)
SELECT company_name, contact_name
FROM customers
WHERE EXISTS (SELECT customer_id FROM orders
			  WHERE customer_id = customers.customer_id -- отсеяли которых нет в customers
			 AND freight BETWEEN 50 AND 100); --второе условие
----переписал без подзапроса
SELECT DISTINCT company_name, contact_name
FROM customers
JOIN orders USING(customer_id)
WHERE freight BETWEEN 50 AND 100
## NOT EXIST
SELECT company_name, contact_name
FROM customers
WHERE NOT EXISTS (SELECT customer_id FROM orders
			 WHERE customer_id = customers.customer_id
			 AND order_date BETWEEN '1995-02-01' AND '1995-02-15')

                                                       

SELECT DISTINCT company_name
FROM customers
JOIN orders USING(customer_id)
JOIN order_details USING(order_id)
WHERE quantity > 40

SELECT customer_id
FROM orders
JOIN order_details USING(order_id)
WHERE quantity > 40
## WHERE ANY (любой) если совпадает с записью побочка DISTINCT
SELECT DISTINCT company_name
FROM customers
WHERE customer_id = ANY(SELECT customer_id
                        FROM orders
                        JOIN order_details USING(order_id)
                        WHERE quantity > 40)

SELECT DISTINCT product_name
FROM products
JOIN order_details USING(product_id)
WHERE quantity > (SELECT AVG(quantity)
                  FROM order_details)

SELECT AVG(quantity)
FROM order_details
## WHERE ALL значение "<,>,=.." любого из списка   ALL

SELECT DISTINCT product_name, quantity
FROM products
JOIN order_details USING(product_id)
WHERE quantity > ALL(SELECT AVG(quantity)
                     FROM order_details
                     GROUP BY product_id)
ORDER BY quantity

# DDL (CREATE, ALTER TABLE...) Primary Key , Foreign key,  Счётчики
## DATA DEFINITION LANGUAGE (манипуляции данними)
  CREATE TABLE                                         --  создать таблицу 
  ALTER TABLE ..... name                               --  изменить таблицу
       ADD COLUMN .....name, type                       --  добавить колонку
       RENAME TO ......new_table_name                  --  переименовать таблицу
       RENAME old_column_name TO new_column_name       --  переименовать колонку
       ALTER COLUMN column_name SET DATA TYPE data_type -- задать новый тип данных
       DROP COLUMN column_name                         --  удалить столбец
       DROP CONSTRAINT exam_exam_id_key                --  удалить ограничение уникальности
   DROP TABLE table_name                               --  удаляет таблицу полностью
   TRUNCATE TABLE table_name                           --  удаляет данные в таблице если нет ссылок (чистит без лога)
         RESTART IDENTITY                              --  перезапускает 'serial' и 'identity' после 'TRUNCATE TABLE'
### CREATE TABLE
CREATE TABLE student
(
student_id serial,
first_name varchar,
last_name varchar,
birthday date,
phone varchar
);

CREATE TABLE cathedra
(
	cathedra_id serial,
	cathedra_name varchar,
	dean varchar
);
### ALTER TABLE _ ADD COLUMN _
ALTER TABLE student
ADD COLUMN middle_name varchar;

ALTER TABLE student
ADD COLUMN rating float;

ALTER TABLE student
ADD COLUMN enrolled date;
### ALTER TABLE DROP COLUMN
ALTER TABLE student
DROP COLUMN middle_name;                   
### ALTER TABLE RENAME
ALTER TABLE cathedra
RENAME TO chair;

ALTER TABLE chair
RENAME cathedra_id TO chair_id;

ALTER TABLE chair
RENAME cathedra_name TO chair_name;
### ALTER TABLE ALTER COLUMN
ALTER TABLE student
ALTER COLUMN first_name SET DATA TYPE varchar(64);
ALTER TABLE student
ALTER COLUMN last_name SET DATA TYPE varchar(64);
ALTER TABLE student
ALTER COLUMN phone SET DATA TYPE varchar(30);
### CREATE, INSERT INTO, TRUNCATE , DROP TABLE
CREATE TABLE faculty
(
	faculty_id serial,
	faculty_name varchar
);

INSERT INTO faculty (faculty_name)
VALUES
('faculty 1'),
('faculty 2'),
('faculty 3');

SELECT * FROM faculty;

TRUNCATE TABLE faculty RESTART IDENTITY

DROP TABLE faculty;
## PRIMARY KEY (Первичный ключ) CONSTRAINT (ключ )
UNIQUE      -- только уникальные значения
NOT NULL    -- нельзя оставить пустой
RYMARY KEY  -- уникальный ключ может быть только один в таблице
### DROP CONSTRAINT (удалить ключ)     
       ALTER TABLE chair
       DROP CONSTRAINT chair_chair_id_key (табл_колонка_id_key)
		
  SELECT constraint_name                          -- посмотреть адресс ключа
  FROM information_schema.key_column_usage
  WHERE table_name = 'chair'
  AND table_schema = 'public'                     -- можно и без table_schema
  AND column_name = 'chair_id'
### ADD PRIMARY KEY (задать ключ)
   ALTER TABLE chair                             -- задать ключ
   ADD PRIMARY KEY(chair_id)     
## FOREIGN KEY (внешний ключ) REFERENCES
  ---задать ВНЕШНИЙ (FOREIGN KEY) ключ при создании таблици
  -- установить связь с таблицей publisher
  -- REFERENCES -- ссылка 

CREATE TABLE book
(
	book_id int,
	title text NOT NULL,
	isbn varchar(32) NOT NULL,
	publisher_id int,
	
	CONSTRAINT PK_book_book_id PRIMARY KEY(book_id),
	CONSTRAINT FK_book_publisher FOREIGN KEY (publisher_id) REFERENCES publisher(publisher_id)
)

  задать Внешний ключ

ALTER TABLE book
ADD CONSTRAINT FK_books_publisher FOREIGN KEY(publisher_id) REFERENCES publisher(publisher_id);

       DROP TABLE IF EXIST book;            ---удалить таблицу если она существует

     

ALTER TABLE book
ADD COLUMN price decimal CONSTRAINT CHK_book_price CHECK (price > 0)  --- условие для ячейки CHECK

### DEFAULT значение по умолчанию 
CREATE TABLE customer
(
	customer_id serial,
	full_name text,
	status char DEFAULT 'r',
	
	CONSTRAINT PK_customer_customer_id PRIMARY KEY(customer_id),
	CONSTRAINT CHK_customer_status CHECK (status = 'r' OR status = 'p' )
);
                                                                                  удалить
ALTER TABLE customer
ALTER COLUMN status DROP DEFAULT
                                                                               задать

ALTER TABLE customer
ALTER COLUMN status SET DEFAULT 'r'
## СЧЁТЧИК  ПОСЛЕДОВАТЕЛЬНОСТЬ
### CREATE SEQUENCE DROP SEQUENCE
CREATE SEQUENCE seq1; -- создать последовательность имя seq1
SELECT nextval('seq1'); -- функция ввызывает следуещее значение последовательности
SELECT currval('seq1'); -- возвращает текущее значение(последнее) из последовательности seq1
SELECT lastval();  --    возвращает последнее использованое значение в текущей сессии

CREATE SEQUENCE seq1;                               -- создать

CREATE SEQUENCE IF NOT EXISTS seq2 INCREMENT 16;    -- шаг 16

CREATE SEQUENCE IF NOT EXISTS seq3
INCREMENT 16                                        -- шаг 16
MINVALUE 0                                          -- минималка 
MAXVALUE 128                                        -- максимальная величина
START WITH 0;                                       

SELECT nextval('seq3');                             -- запустить 

ALTER SEQUENCE seq3 RENAME TO seq4;                 -- переименовать

ALTER SEQUENCE seq4 RESTART WITH 16;                -- скинуть счётчик

DROP SEQUENCE seq4;                                 -- удалить счётчик



CREATE SEQUENCE IF NOT EXISTS book_book_id_seq          -- создали счётчик с именем
START WITH 1 OWNED BY book.book_id;  

ALTER TABLE book
ALTER COLUMN book_id SET DEFAULT nextval('book_book_id_seq');  -- связали счётчик с таблицей book.book_id

CREATE TABLE book
(
	book_id int GENERATED ALWAYS AS IDENTITY NOT NULL, -- GENERATED ALWAYS AS IDENTITY вместо serial
	title text NOT NULL,
	isbn varchar(32) NOT NULL,
	publisher_id int NOT NULL,
	
	CONSTRAINT PK_book_book_id PRIMARY KEY(book_id)
)

# INSERT вставить данные в таблицу
## вставить в таблицу (все данные столбцов)
INSERT INTO author 
VALUES (10, 'john Silver', 4.5)
##  вставить в таблицу
INSERT INTO author (author_id, full_name)
VALUES (15, 'Bob Grey'),
       (14, 'Bob Grey1'),
	   (13, 'Bob Grey2'),
	   (12, 'Bob Grey3'),
	   (16, 'Bob Grey4')
	   
### найти в таблице, создать новую и вставить найденные данные
SELECT *
INTO best_authors 
FROM author
WHERE rating >= 4.5

### копировать из таблици где ...
INSERT INTO best_authors
SELECT *
FROM author
WHERE rating < 4.5;

### копировать из таблици всё
INSERT INTO products
SELECT * 
FROM backup_products_discontinued;

### копировать всё из таблицы
CREATE TABLE employees_salary AS 
SELECT * FROM employees
WHERE ....;

## копировать всё из таблицы
SELECT * INTO employees_salary
FROM employees
WHERE ....;

## изменить данные в таблице  http://www.sql-tutorial.ru/ru/book_update_statement.html
UPDATE author
SET full_name = 'Elias', rating = 5
WHERE author_id = 1;

UPDATE author
SET full_name = 'Wolter', rating = 5          -- вносить данные в таблицу можно только по одному значению или ОДНО значение во все строки
WHERE author_id = 1

## удалить данные (строки) с таблицы
DELETE FROM author;

## Очистить таблицу без лога(shift+dell)
TRUNCATE TABLE author;

## Вывод данных после манипуляции RETURNING
INSERT INTO book (title, isbn, publisher_id)
VALUES ('title', 'isbn', 3)
RETURNING *                                   -- выводит данные после манипуляций с таблицами

UPDATE author
SET full_name = 'Wolter', rating = 5          -- вносить данные в таблицу можно только по одному значению или ОДНО значение во все строки
WHERE author_id = 1
RETURNING *

DELETE FROM author
WHERE rating = 5
RETURNING *

# VIEV - сохранённый запрос
## Коротко 
CREATE VIEV ..name.. AS             --  создание 
SELECT

CREATE OR REPLACE VIEV ...name...AS  -- изменение
SELECT

ALTER VIEW ..old_name..RENAME TO ..new_view_name.. -- переименовать

DROP VIEW name      -- можно IF EXISTS -- удалить  

DELETE FROM ..name..                -- удалить строчку
WHERE ...

ПРИ ВСТАВКЕ СТРОЧКИ в VIEW - она попадает в исходную таблицу(материнскую) писец
## примеры
### создать
CREATE VIEW heavi_orders AS
SELECT *
FROM orders
WHERE freight > 50;
### изменить
CREATE OR REPLACE VIEW heavi_orders AS
SELECT *
FROM orders
WHERE freight > 100;
WITH LOCAL CHECK OPTION; -- запрет на вставку строк кот. фильтр = FALSE
WITH CASCADE CHECK OPTION; -- запрет на все виев кот создавались на их основе
### удалить
DROP VIEW IF EXISTS heavi_orders
### переименовать
ALTER VIEW products_suppliers_categories RENAME TO psc_old;
### вставить строчку
INSERT INTO heavi_orders
VALUES (11078, 'VINET', 5, '2019-12-10', '2019-12-15',
        '2019-12-14', 1, 120, 'Hanary Carnes', 
       'Rua de Paco', 'Bern', null, 3012, 'Switzerland');
### удалить строчку
DELETE FROM heavi_orders
WHERE freight < 100.25;

# CASE WHEN
## Select CASE WHEN
SELECT product_name, unit_price, units_in_stock,
    CASE WHEN units_in_stock >= 100 THEN 'Lots of'
	     WHEN units_in_stock >= 50 AND units_in_stock < 100 THEN 'average'
		 WHEN units_in_stock >= 50 THEN 'Low number'
		 ELSE 'unknown'
	END AS amount
FROM products
ORDER BY units_in_stock DESC;

SELECT order_id, order_date,
      CASE WHEN date_part('month', order_date) BETWEEN 3 and 5 THEN 'Spring'
	       WHEN date_part('month', order_date) BETWEEN 6 and 8 THEN 'Summer'
	       WHEN date_part('month', order_date) BETWEEN 9 and 11 THEN 'Autumn'
	  ELSE 'Winter'
	  END AS season
FROM orders
ORDER BY order_date;
## ORDER BY CASE WHEN
SELECT contact_name, country, city
FROM customers
ORDER BY contact_name,
                  (
		   CASE WHEN city IS NULL THEN country
		   ELSE city
		   END
	              );

# ПОЛЕЗНЫЕ ФУНКЦИИ random, coalesce
## random
ALTER TABLE employees ADD COLUMN salary decimal(12,2);

UPDATE employees  SET salary = random()*100;

UPDATE perf_test                             
SET annotation = md5(UPPER(random()::text));

-- заполняем случайными данными
INSERT INTO perf_test(id, reason, annotation) 
SELECT s.id, md5(random()::text), null            -- md5(randoom()) - шифровальная функция генерирует рандомный текст
FROM generate_series(1, 10000000) AS s(id)        -- тоже наверное функция
ORDER BY random();   -- упорядочить случайно      -- тоже функция
## COALESCE -- подменяет значение NULL в таблице на любую другую надпись  
SELECT order_id, order_date, COALESCE(ship_region, 'unknown') AS ship_region
FROM orders;

SELECT DISTINCT contact_name, COALESCE(order_id::text, 'orders') AS order_id -- изменить формат '::text'
FROM customers
LEFT JOIN orders USING(customer_id)
WHERE order_date IS NULL

SELECT contact_name, country, COALESCE(city, 'interes') AS city
FROM customers
## update type 
select contact_name, coalesce(order_id::text, 'no orders') as order_id
from customers
left join orders using (customer_id)
order by order_id desc;
## NULLIF подмена строки
SELECT contact_name, COALESCE(NULLIF(city, ''), ' unknown') AS city  -- переименовать все строки 'unknown' на 'city'
FROM customers           

# Проектирование БД
• Проблема представления предметной области
• Логическое проектирование
• Проблемы плохого проектирования:
  • возможность записи не валидных данных
  • возможность потери информации (нет нужных связей)
  • отсутствие необходимой информации (забыли то, что было нужно)

• Стадии проектирования БД

  • Анализ требования предметной области
  • Логическое моделирование данных предметной области
  • Физическое проектирование и нормализация

• Анализ требований
  • Составление USE CASES
  • Аналитический процесс с участием stakeholders (владельцев, экспертов домена)
  • Концептуальная схема БД

• Логическое проектирование
 • Детализирует концептуальную модель БД
 • Разные источники включают разные компоненты в логическую модель
 • Полностью описывает все ключи
 • Полностью определяет типы данных (безотносительно конкретной СУБД)
 • Полностью описывает все логические ограничения (спорно)
 • Нормализация отношений обычно максимум до формы ЗНФ

• Физическая модель данных
 • Выбирается конкретная СУБД
 • Определяются типы данных
 • Определяются индексы
 • Могут определяться представления (views)
 • Определяются ограничения на доступ (security)

• ER Diagrams
  • Очень много платных инструментов для моделирования
  • MySQL Workbench
  • Oracle SQL Developer Data Modeler
  • pgModeler
  • SQL Power Architect

• Базовые советы по проектированию
   • Таблица: объект, событие, абстракция
   • Поле (колонка): свойство объекта
   • Запись (строка): совокупность полей
   • Значения в каждом поле по отдельности
     не должны содержать не валидных данных
   • Значения в совокупности полей должны быть непротиворечивы

• Плохие практики
   • Игнорирование нормализации - - избыточность данных
   • Отсутствие стандартов именования на проекте
   • Одна таблица для разных по смыслу данных
   • Наплевательское отношение к актуальности репрезентации данных
     (домен меняется - это живой механизм!)

• Плохие практики
   • Поле, содержащее более 1 логической части (full_name)
   • Поле, содержащее более 1 значения (массив, когда не надо)
   • Вычислимое поле (полная зарплата за всё время работы)
   • Неправильно выбранные первичные ключи (ИНН – плохой РК)
   • Избегайте композитных РК
     (может приводить к деградации производительности)
   • В идеале, в таблице кроме суррогатного ключа, должен быть и натуральный
   • Правила иногда можно нарушать!
     Вычислимое поле даёт реrformance boost? Делаем вычислимое поле…

## Формы НФ
 • Нормальная форма (нф) - свойство отношения,
   характеризующее его с точки зрения избыточности

 • Нормализация - процесс минимизации
   избыточности отношения (приведение к НФ)

• 1НФ
   • Нет строк-дубликатов
   • Все атрибуты простых типов данных
   • Все значения скалярные

• 2НФ
  • Таблица удовлетворяет 1НФ
  • Есть первичный ключ
  • Все атрибуты (поля) описывают
    первичный ключ целиком, а не лишь его часть

• ЗНФ
   • Таблица удовлетворяет 2НФ
   • Нет зависимостей одних неключевых атрибутов от других
     (все атрибуты зависят только от первичного ключа)

# FUNC SQL and PgSQL
## ФУНКЦИИ МОГУТ СОДЕРЖАТЬ 

   -- SELECT, INSERT, UPDATE, DELETED -- CRUD

                    НЕ МОГУТ
    -- COMMIT, SAVEPOINT(TCL), VACUUM(utility)
        
                ДЕЛЯТСЯ НА
        -- SQL функции
        -- Процедурные (PL/pgSQL)
    -- Серверные (написанные на С)
    -- Собственные С - функции
```
CREATE FUNCTION ...name..(arg1 type, arg2 type,...) RETURNS type AS$$
--logic
$$ LANGUAGE lang
CREATE OR REPLACE -- создать или изменить

SELECT *
INTO tmp_customers   -- копи в таблицу 
FROM customers;

SELECT *
FROM tmp_customers;

UPDATE tmp_customers        готовим обновлялку для регион IS NULL
SET region = 'unknown'
WHERE region IS NULL;

CREATE OR REPLACE FUNCTION fix_customer_region() RETURNS void AS $$    -- создать функцию
     UPDATE tmp_customers
     SET region = 'unknown'
     WHERE region IS NULL
$$ language SQL;

SELECT fix_customer_region();

SELECT *
FROM tmp_customers;
```
## скалярные функции (линейные 1 цифра) 
```

CREATE OR REPLACE FUNCTION get_total_number_of_goods() RETURNS bigint AS $$
    SELECT SUM(units_in_stock)
	FROM products
$$ LANGUAGE SQL;

SELECT get_total_number_of_goods() AS total_goods;

                         IN входящий (все по умолчанию входящие)
                         OUT исходящий (обязательно писать)
                         INOUT - вход и исход
                         VARIADIC - масив входящих
                         DEFAULT - умолчание
--1 
SELECT * FROM func_minmax_unit_price(0);

CREATE OR REPLACE FUNCTION get_price_boundaries(OUT max_price real, OUT min_price real) AS $$
    SELECT MAX(unit_price), MIN(unit_price)
    FROM products
$$ LANGUAGE SQL;

SELECT * FROM get_price_boundaries();
                                          --возвращает только одну строчку 
--2
CREATE OR REPLACE FUNCTION get_product_price_by_name(prod_name varchar) RETURNS real AS $$
   SELECT unit_price
   FROM products
   WHERE product_name = prod_name
$$ LANGUAGE SQL;
  
SELECT get_product_price_by_name('Chai') AS price;

SELECT *
FROM products
ORDER BY product_name

SELECT get_product_price_by_name('Chai') AS price;

--2
CREATE OR REPLACE FUNCTION get_price_boundaries(OUT max_price real, OUT min_price real) AS $$
    SELECT MAX(unit_price), MIN(unit_price)
	FROM products
$$ LANGUAGE SQL;

SELECT * FROM get_price_boundaries();

--3
CREATE OR REPLACE FUNCTION get_price_boundaries_by_diskontinuoti(is_discontinued int, OUT max_price real, OUT min_price real) AS $$
    SELECT MAX(unit_price), MIN(unit_price)
	FROM products
	WHERE discontinued = is_discontinued
$$ LANGUAGE SQL;

SELECT *
FROM get_price_boundaries_by_diskontinuoti(1);

--4 DEFAULT
CREATE OR REPLACE FUNCTION get_price_boundaries_by_diskontinuoti(is_discontinued int DEFAULT 1, 
						                 OUT max_price real, 
							       OUT min_price real) AS $$
    SELECT MAX(unit_price), MIN(unit_price)
	FROM products
	WHERE discontinued = is_discontinued
$$ LANGUAGE SQL;

SELECT * FROM get_price_boundaries_by_diskontinuoti(); -- по умолчанию default = 1
SELECT * FROM get_price_boundaries_by_diskontinuoti(0); -- discontinued = 0
SELECT * FROM get_price_boundaries_by_diskontinuoti(1); -- discontinued = 1

                                     -- множество строк --

                         -- RETURNS SETOF data_type -- возврат одного столбца
                         -- RETURNS SETOF table     -- возврат всей таблицы типа 
                         -- RETURNS SETOF record    -- возврат всей таблицы неизвестного типа (не использовать , тип - анонимный)
                         -- RETURNS TABLE(column_name data_type,...) -- можно явно указать возвращаемые столбцы
                         -- возврат через out параметры -- одна строка
                   -- RETURNS SETOF record - с OUT параметрами - ВЫДАЁТ ТАБЛИЦУ - так же как RETURNS TABLE(name type,..)

--1 RETURNS SETOF data_type
CREATE OR REPLACE FUNCTION get_average_prices_by_prod_categories() 
                              RETURNS SETOF double precision AS $$
	SELECT AVG(unit_price)
	FROM products
	GROUP BY category_id
$$ LANGUAGE SQL;

SELECT * FROM get_average_prices_by_prod_categories() AS avg_price;

--2 RETURNS SETOF record с указанием OUT
DROP FUNCTION get_average_prices_by_prod_cats;

CREATE OR REPLACE FUNCTION get_average_prices_by_prod_cats(OUT sum_price real, OUT avg_price float8) -- double precision = float8 = AVG
       RETURNS SETOF RECORD AS $$
	SELECT SUM(unit_price), AVG(unit_price)
	FROM products
	GROUP BY category_id
$$ LANGUAGE SQL;

SELECT * FROM get_average_prices_by_prod_cats();
SELECT sum_price FROM get_average_prices_by_prod_cats();
SELECT sum_price, avg_price FROM get_average_prices_by_prod_cats() ;
SELECT sum_price AS sum_of, avg_price AS a_price FROM get_average_prices_by_prod_catDROP FUNCTION get_average_prices_by_prod_cats;

--3 RETURNS SETOF record без указания OUT -- нужно писать параметры, 
                                          -- нельзя вызвать один столбик
DROP FUNCTION get_average_prices_by_prod_cats;

CREATE OR REPLACE FUNCTION get_average_prices_by_prod_cats()
       RETURNS SETOF RECORD 
	   AS $$
	   
	SELECT SUM(unit_price), AVG(unit_price)
	FROM products
	GROUP BY category_id
	
$$ LANGUAGE SQL;

SELECT * FROM get_average_prices_by_prod_cats() AS (sum_price real , avg_price float8);

--4 RETURNS table
SELECT * FROM customers;

CREATE OR REPLACE FUNCTION get_customers_by_country (customer_country varchar)
          RETURNS TABLE(char_cod char, company_name varchar) AS $$
		  
	SELECT customer_id, company_name
	FROM customers
	WHERE country = customer_country
	
$$ LANGUAGE SQL;

SELECT * FROM get_customers_by_country('USA');
SELECT company_name FROM get_customers_by_country('USA');
SELECT char_cod FROM get_customers_by_country('USA');

DROP FUNCTION get_customers_by_country

--4.1
SELECT * FROM customers;

CREATE OR REPLACE FUNCTION get_customers1_by_country ()
          RETURNS TABLE(char_cod char, company_name varchar) AS $$
		  
	SELECT customer_id, company_name
	FROM customers
	
$$ LANGUAGE SQL;

SELECT * FROM get_customers1_by_country();
SELECT company_name FROM get_customers1_by_country();
SELECT char_cod FROM get_customers1_by_country();

DROP FUNCTION get_customers1_by_country

-- 4,2
SELECT * FROM customers;

CREATE OR REPLACE FUNCTION get_customers_by_country (customer_country varchar)
          RETURNS SETOF customers AS $$
		  --won't work SELECT company_name, contact_name
	SELECT *
	FROM customers
	WHERE country = customer_country
	
$$ LANGUAGE SQL;

SELECT * FROM get_customers_by_country('USA');
SELECT company_name, contact_name FROM get_customers_by_country('USA');
SELECT char_cod FROM get_customers_by_country('USA');

DROP FUNCTION get_customers_by_country
```
## --ja
```
--1 OUT параметры

CREATE OR REPLACE FUNCTION get_products_categories_avg_sum(OUT category_name character varying,
                                                    OUT avg_u_i_s numeric, OUT sum_u_i_s bigint) 
                 RETURNS SETOF record AS $$
SELECT category_name, AVG(units_in_stock),SUM(units_in_stock)
FROM products
JOIN categories USING(category_id)
GROUP BY category_name
$$ LANGUAGE SQL;

SELECT * FROM get_products_categories_avg_sum();

--2 RETURNS TABLE(,,)

CREATE OR REPLACE FUNCTION get_products_categories_avg_sum()RETURNS TABLE(category_name character varying,
                                                    avg_u_i_s numeric, sum_u_i_s bigint) AS $$
SELECT category_name, AVG(units_in_stock),SUM(units_in_stock)
FROM products
JOIN categories USING(category_id)
GROUP BY category_name
$$ LANGUAGE SQL;

SELECT * FROM get_products_categories_avg_sum();
```
## PL/pgSQL
```
  -- CREATE FUNCTION func_name('arg1, arg2...') RETURNS data_type AS $$
BEGIN
--logic
END;
$$ LANGUAGE plpgsql;

- BEGIN/END - тело метода
- Создание переменных
- Прогон циклов и развитая логика
- Возврат значения через RETURN (вместо SELECT) для СКАЛЯРНЫХ значений (ОДНО)
  или RETURN QUERY (в дополнение к SELECT) для таблиц
- Для двух и более значений с OUT параметрами return не нужен
```
### 1 Синтаксис
```
-- 1
SELECT SUM(units_in_stock)
FROM products;

CREATE OR REPLACE FUNCTION get_total_number_of_goods() RETURNS bigint AS $$
    BEGIN
	    RETURN SUM(units_in_stock)
		FROM products;
	END;
$$ LANGUAGE plpgsql;

SELECT get_total_number_of_goods();
```
### 2
```
SELECT MAX(unit_price)
FROM products
WHERE discontinued = 1

CREATE OR REPLACE FUNCTION get_max_price_of_frome_discontinued() RETURNS bigint AS $$
    BEGIN
	   RETURN MAX(unit_price)
	   FROM products
	   WHERE discontinued = 1;
	 END;
$$ LANGUAGE plpgsql;

SELECT get_max_price_of_frome_discontinued()
```
### 3  присвоение исходящих в plpgsql
```
SELECT MAX(unit_price), MIN(unit_price)
FROM products

CREATE OR REPLACE FUNCTION get_price_boundaries(OUT max_price real, OUT min_price real) AS $$
BEGIN
        -- max_price = MAX(unit_price) FROM products; 
	-- min_price = MIN(unit_price) FROM products; --присвоение данних
	SELECT MAX(unit_price), MIN(unit_price)       -- или так
	INTO max_price, min_price
	FROM products;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_price_boundaries();
```
### 4
```
CREATE OR REPLACE FUNCTION get_sum(x int, y int, OUT RESULT int) AS $$
BEGIN
    result = x + y;
	--RETURN; если функцию нужно завершить позже
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_sum(2,3);
```
### 5
```
SELECT *
FROM customers
WHERE country = 'USA'

DROP FUNCTION get_customers_by_country

CREATE FUNCTION get_customers_by_country(customer_country varchar) RETURNS SETOF customers AS $$
BEGIN
    RETURN QUERY  -- где QUERY это запрос
	SELECT *
	FROM customers
	WHERE country = customer_country;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_customers_by_country('USA');
```
## DECLARE ДEКЛАРАЦИЯ СТОРОННИХ ДАННЫХ
```
       CREATE FUNCTION func_name(arg1, arg2...) RETURNS data_type AS $$
       DECLARE
              variable type;
       BEGIN
              logic
       END;
       $$ LANGUAGE plpgsql;

--sqrt - корень квадратный, стандартная функция sql

CREATE FUNCTION get_square(ab real, bc real, ac real) RETURNS real AS $$
DECLARE
  perimeter real;
BEGIN
  perimeter = (ab+bc+ac) / 2;
  RETURN sqrt(perimeter * (perimeter - ab) * (perimeter - bc) * (perimeter - ac));
END;
$$ LANGUAGE plpgsql;

SELECT get_square(5,5,5);
```
### 1
```
CREATE FUNCTION get_square(ab real, bc real, ac real) RETURNS real AS $$
DECLARE
  perimeter real;
BEGIN
  perimeter = (ab+bc+ac) / 2;
  RETURN sqrt(perimeter * (perimeter - ab) * (perimeter - bc) * (perimeter - ac));
END;
$$ LANGUAGE plpgsql;

SELECT get_square(5,5,5);
```
### 2
```
CREATE OR REPLACE FUNCTION calc_middle_price() RETURNS SETOF products AS $$
DECLARE
    avg_price real;
	low_price real;
	high_price real;
BEGIN
   SELECT AVG(unit_price) INTO avg_price
   FROM products;
   
   low_price = avg_price * 0.75;
   high_price = avg_price * 1.25;
   
   RETURN QUERY
   SELECT * FROM products
   WHERE unit_price BETWEEN low_price AND high_price;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM calc_middle_price();
```
## IF ELSE ВЕТВЛЕНИЕ
IF expression THEN
  logic
ELSIF expression THEN
  logic
ELSIF expression THEN
  logic
ELSE
  logic
END IF;

Если не выполнено ни одно условие из ELSIF то значение попадает в ELSE
так же возможно написание ELSEIF

### 1
```
CREATE FUNCTION convert_temp_to(temperature real, to_celsius bool DEFAULT true) RETURNS real AS $$
DECLARE
    result_temp real;
BEGIN
    IF to_celsius = true THEN
	                      result_temp = (5.0/9.0)*(temperature-32);
	ELSE 
	     result_temp = (9*temperature+(32*5))/5.0;
	END IF;
	
	RETURN result_temp;
END;
$$ LANGUAGE plpgsql;

SELECT convert_temp_to (80);

SELECT convert_temp_to(26.667, false); --если фолс то ЕЛСЕ, если тру то ИФ (тру можно не указывать)

DROP FUNCTION convert_temp_to;
```
### 2
```
CREATE FUNCTION get_season(month_number int) RETURNS text AS $$
DECLARE
     season text;
BEGIN
     IF month_number BETWEEN 3 AND 5 THEN
	      season = 'Spring';
	 ELSEIF month_number BETWEEN 6 AND 8 THEN
	      season = 'Summer';
	 ELSEIF month_number BETWEEN 9 AND 11 THEN
	      season = 'Autumn';
	 ELSE                                        -- попали все что не входят
	      season = 'Winter';
	 END IF;
	 
	 RETURN season;
END;
$$ LANGUAGE plpgsql;

SELECT get_season(12)
```
## ЦИКЛЫ
### 1
будет исполнятся пока expression true

   WHILE expression
   LOOP
        logic
   END LOOP;

### 2
Бесконечный цикл, выйти когда expression true

    LOOP
        EXIT WHEN expression
        logic
    END LOOP;

### 3
```
Для определенного количества(counter) раз. где a..b интервал, 
      [BYx] x - необязательный шаг, по умолчанию шаг = 1

    FOR counter IN a..b [BYx]
    LOOP
        logic
    END LOOP;

--Возможные условия
CONTINUE WHEN expression - принудительно перейти на следующую итерацию
```
## Примеры
### Пример 1
```
--последовательность Фибоначи 1,2,3,5,8...следующее = суме двух последних

CREATE OR REPLACE FUNCTION fib(n int) RETURNS int AS $$
DECLARE                           -- декларация
   counter int = 0;
   i int = 0;
   j int = 1;
BEGIN
    IF n < 1 THEN                 -- условие для входящих(IN)
	    RETURN 0;
	END IF;
	
	WHILE counter <= n        -- серия повторов, меньше равно n
	LOOP
	    counter = counter + 1; -- прирост серии 
		SELECT j, i+j INTO i, j;   -- рост числа с каждой новой серией
	END LOOP;
		
	RETURN i;
END;
$$ LANGUAGE plpgsql;

SELECT fib(5);
```
### 2
```
CREATE OR REPLACE FUNCTION fib(n int) RETURNS int AS $$
DECLARE
   counter int = 0;
   i int = 0;
   j int = 1;
BEGIN
   IF n < 1 THEN  
     RETURN 0; 
   END IF;
	
	LOOP
	EXIT WHEN counter > n ;
	    counter = counter + 1;
		SELECT j, i+j INTO i, j;
	END LOOP;
		
	RETURN i;
END;
$$ LANGUAGE plpgsql;
```
### 3
```
DO $$
BEGIN
     FOR counter IN 1..5
	 LOOP
	     RAISE NOTICE 'Counter: %', counter;
	 END LOOP;
END $$;

DO $$
BEGIN
     FOR counter IN REVERSE 5..1  -- REVERSE значения тоже нужно менять 5..1
	 LOOP
	     RAISE NOTICE 'Counter: %', counter;
	 END LOOP;
END $$;

--n
DO $$
BEGIN
     FOR counter IN 1..10 BY 2
	 LOOP
	     RAISE NOTICE 'Counter: %', counter;
	 END LOOP;
END $$;
```                    
## RETURN NEXT
```
-- Иногда необходимо накапливать записи в результирующем наборе
   (построчный процессинг)
   *RETURN NEXT expression;

CREATE FUNCTION return_ints() RETURNS SETOF int AS $$
BEGIN
    RETURN NEXT 1;
	RETURN NEXT 2;
	RETURN NEXT 3;
END;
$$ LANGUAGE plpgsql;

SELECT return_ints();
```

# Обработка ошибок       RAISE
       RAISE[level]'message(%)', arg_name;
  Добавить условия при которых будет выдавать ошибку;
## level уровень серьёзности ошибки:
  -- DEBAG - отладка;
  -- LOG -  лог;
  -- INFO -  информация;
  -- NOTICE - замечания;
  -- WARNING - потенциальная опасность;
  -- EXCEPTION - исключение. Выдаёт ошибку с описанием
  ```
- EXCEPTION абортирует текущую транзакцию
- log_min_messages - регулирует уровень сообщений, которые будут писатся в лог сервера(WARNING - по умолчанию)
- client_min_messages - регулирует уровень сообщений которые будут передаваться вызывающей стороне(NOTICE - по умолчанию)
- из самых важных параметров упомянуть HINT подсказка & ERRCODE
- параметры присоединяются с помощью USING:
RAISE 'invalid billing number=%, number USING HINT='Check out the billing number' ERRCODE='12881'.
 - Чтобы поймать исключение нужен блок EXCEPTION WHEN:
   EXCEPTION WHEN condition [others] THEN handing_logic
 - Код в блоке EXCEPTION теряет в производительности
 - Старайтесь избегать исключений
```
## ПРИМЕРЫ
### 1  RAISE EXCEPTION -- выдаёт ошибку с описанием
```
CREATE OR REPLACE FUNCTION get_season(month_number int) RETURNS text AS $$
DECLARE 
   season text;
BEGIN
   IF month_number not BETWEEN 1 AND 12 THEN
      RAISE EXCEPTION 'Invalid month. You passed: (%)', month_number USING HINT='Allowed from 1 up to 12', ERRCODE = 12882;
	END IF;
   IF month_number BETWEEN 3 AND 5 THEN 
       season = 'Spring';
   ELSIF month_number BETWEEN 6 AND 8 THEN
       season = 'Summer';
   ELSIF month_number BETWEEN 9 AND 11 THEN
       season = 'Autumn';
   ELSE 
       season = 'Winter';
   END IF;
   
   RETURN season;
END;
$$ LANGUAGE plpgsql;

SELECT get_season(15);
SELECT get_season(12);
```
### 2 EXCEPTION WHEN ловит ошибку и Заменяем её на NULL
```
CREATE OR REPLACE FUNCTION get_season_caller(month_number int) RETURNS text AS $$
BEGIN
    RETURN get_season(month_number);
EXCEPTION WHEN SQLSTATE '12882' THEN      --функция опирается на функцию get_season и если ошибка ввода данных то выдаёт NULL
    RAISE INFO 'A problem. Nothing special.';
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;

SELECT get_season_caller(15);
```
### 3  Выдаёт в сообщении(message) подробности о ошибке
```
CREATE OR REPLACE FUNCTION get_season_caller(month_number int) RETURNS text AS $$
DECLARE
err_ctx text;
err_msg text;
err_details text;
err_code text;
BEGIN
    RETURN get_season(month_number);
EXCEPTION WHEN SQLSTATE '12882' THEN
   GET STACKED DIAGNOSTICS         
                 err_ctx = PG_EXCEPTION_CONTEXT,  -- в базе ошибок по адрессу - PG_EXCEPTION_CONTEXT
		 err_msg = MESSAGE_TEXT,
	         err_details = PG_EXCEPTION_DETAIL,
		 err_code = RETURNED_SQLSTATE;
    RAISE INFO 'My custom handler:';
	RAISE INFO 'Error msg:%', err_msg;
	RAISE INFO 'Error details:%', err_details;
	RAISE INFO 'Error code:%', err_code;
	RAISE INFO 'Error context:%', err_ctx;
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- ещё пару кодов --
CREATE OR REPLACE FUNCTION get_season_caller1(month_number int) RETURNS text AS $$
BEGIN
    RETURN get_season(month_number);
EXCEPTION WHEN SQLSTATE '12882' THEN
    RAISE INFO 'My custom handler:';
	RAISE INFO 'Error Mame:%', SQLERRM;
	RAISE INFO 'Error details:%', SQLSTATE;
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;

SELECT get_season_caller1(20);
```
### 4 WHEN OTHERS THEN -- ловит все ошибки --сложность в логике - всё без разбора--
```
# CREATE OR REPLACE FUNCTION get_season_caller2(month_number int) RETURNS text AS $$
BEGIN
    RETURN get_season(month_number);
EXCEPTION

WHEN OTHERS THEN
    RAISE INFO 'My custom handler:';
	RAISE INFO 'Error Mame:%', SQLERRM;
	RAISE INFO 'Error details:%', SQLSTATE;
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;

SELECT get_season_caller2(20);
```

# Домен - пользовательские типы данных с ограничениями
## CREATE DOMAIN domain_name AS data_type CONSTRAINTS; где CONSTRAINTS - ограничения 
### Пример
 - CREATE DOMAIN text_no_space_null AS TEXT NOT NULL CHECK(value!~'\s');
   Изменить домен
 - ALTER DOMAIN domain_name ADD CONSTRAINT new_constraint[NOT VALID] - NOT VALID не обязвателен
 - ALTER DOMAIN domain_name VALIDATE CONSTRAINT new constraint
### ПРИМЕР
```
1)
 CREATE DOMAIN text_no_space_null AS TEXT NOT NULL CHECK (value ~ '^(?!\s*$).+');
 --создаём ограничения ~ ['^(?!\s*$).+'] регулярные выражения
2)
CREATE TABLE agent(
	first_name text_no_space_null,
	last_name text_no_space_null
        );
3)
INSERT INTO agent
VALUES ('bob', 'taylor');
INSERT INTO agent
VALUES ('   ', 'taylor'); --не работает
4)
DROP TABLE agent;
DROP DOMAIN IF EXISTS text_no_space_null;
5)
--создать ограничение но не внедрять до валидации - длинна <=32 символа
ALTER DOMAIN text_no_space_null ADD CONSTRAINT text_no_space_null_lenght32 CHECK (length(value)<=32) NOT VALID;
6)
--проверить записи на соответствие ограничению
ALTER DOMAIN text_no_space_null VALIDATE CONSTRAINT text_no_space_null_lenght32;
7)
--удалить (последние добавленые) ограничения
ALTER DOMAIN text_no_space_null DROP CONSTRAINT text_no_space_null_lenght32
```
## Composite Types
### Составной тип - тип обьеденяющий логически взаимосвязанные данные без создания полноценной таблицы
- состоящий из нескольких полей имеющих свой тип
- нельзя задавать ограничения (constraints). Не так часто используется в таблицах.
- часто используется для возврата данных из функций
- как правило таблица лучше
```
CREATE TYPE type_name AS(
  field1 type
  field2 type
  );

      создать экземпляр- вставить данные
'(1,2,"text")' или ROW(1,2,'text')

DROP IF EXIST TYPE type_name[CASCADE];
```
#### пример1
```
DROP FUNCTION get_price_boundaries;

CREATE OR REPLACE FUNCTION get_price_boundaries(OUT max_price real, OUT min_price real)
RETURNS SETOF RECORD AS $$
	SELECT MAX(unit_price), MIN(unit_price) 
	FROM products;
$$ LANGUAGE sql;

SELECT * FROM get_price_boundaries();
```
- создаем композитный тип
CREATE TYPE price_bounds AS (
	max_price real,
	min_price real
);
- указываем его в RETURNS SETOF
```
CREATE OR REPLACE FUNCTION get_price_boundaries()RETURNS SETOF price_bounds AS $$
	SELECT MAX(unit_price), MIN(unit_price) 
	FROM products;
$$ LANGUAGE sql;

SELECT * FROM get_price_boundaries();
```
#### Пример2
CREATE TYPE complex AS(
	r float8,
	f float8
);

CREATE TABLE match_calcs (
	math_id serial,
	val complex
);
--вставляем в ячейку val - 2 значения
INSERT INTO match_calcs(val)
VALUES
(ROW(3.0, 4.0)),
(ROW(2.0, 1.0));

SELECT * FROM match_calcs;
--или с указанием имени при CREATE TYPE
SELECT (val).f FROM match_calcs;
---полный адресс обеих значений
SELECT (val).* FROM match_calcs;
             =
SELECT (match_calcs.val).f, (match_calcs.val).r FROM match_calcs;
### UPDATE--
UPDATE match_calcs
SET val = ROW(5.0, 4.0)
WHERE math_id = 1;
 -- только r
UPDATE match_calcs
SET val.r = 3
WHERE math_id = 1;
   --скобочки
UPDATE match_calcs
SET val.r = (val).r + 1
WHERE math_id = 1;

DROP TABLE match_calcs
### Перечисление - позволяет эмулировать простейшие справочные таблицы
-----служит єлементарной справочной таблицей
-----полезная 
-----ENUM - между собой сравнивать нельзя SQL
 --CREATE TYPE type_name AS ENUM('value1', 'value2',...);
-----Накладывает ограничения (смысл перечисления) на добавление в колонку
             --типа перечисления значения, отсутствующего в перечислении
             --(--Создав тип ENUM - мы указали возможные значения в этой колонке)
-----Значения регистрозависимы
#### 1)
CREATE TABLE chess_title (
	title_id serial PRIMARY KEY,
	title text
);
#### 2)REFERENCES - title_id данные взяли в chess_title
CREATE TABLE chess_player (
	player_id serial PRIMARY KEY,
	first_name text,
	last_name text,
	title_id int REFERENCES chess_title(title_id)
);
#### 3)
INSERT INTO chess_title(title)
VALUES 
	('Candidate Master'),
	('FIDE Master'),
	('International Master'),
	('Grand Master');

SELECT * FROM chess_title;
#### 4)
INSERT INTO chess_player(first_name, last_name, title_id)
VALUES
('Wesley', 'So', 4),
('Vlad', 'Kramnik', 4),
('Vasily', 'Pupkin', 1);

SELECT *
FROM chess_player
JOIN chess_title USING (title_id);
#### 5)
DROP TABLE chess_title CASCADE;
DROP TABLE chess_player;
#### 6) создали 'ENUMEREYSION
CREATE TYPE chess_title AS ENUM
('Candidate Master', 'FIDE Master', 'International Master');
#### 7)
SELECT enum_range(null::chess_title);
#### 8) добавляем значение в єнум (после)
ALTER TYPE chess_title
ADD VALUE 'Grand Master' AFTER 'International Master';
#### 9)title_id chess_title (type!!!)
CREATE TABLE chess_player (
	player_id serial PRIMARY KEY,
	first_name text,
	last_name text,
	title_id chess_title
);
#### 10) 
INSERT INTO chess_player(first_name, last_name, title_id)
VALUES ('Magnus', 'Carlsen', 'Grand Master');

SELECT * FROM chess_player;
#### 11)
DROP TYPE chess_title

# Индексы
## Термины
- database cluster - пространство на диске где находятся Базы Данных и которым управляет одна инстанция сервера
- Файлы хранятся в data -- PGDATA.
- Для каждой БД есть подпапка в PGDATA/base
- Для каждой таблицы выделяется свой отдельный файл
- Таблица состоит из массива страниц (блоков размером 8 кб)
- Файлы таблицы называются HEAP File - содержит списки неупорядоченных записей различной длинны
## Устройство таблиц 
$ Таблица 1ГБ состоит 
    из страниц по 8КБ, 
        каждая из которых содержит 
        --заголовок страницы
        --строки с их заголовками
        --Ссылки на строки 'CTID'
$ Рядом с таблицей лежит файл FSM (free space map),
    который не обновляется при каждом обновлении или удалении строк
        -- VACUUM - очистка дохлых версий строк , не нужен полный доступ - можно работать
        -- VACUUM[FULL] где (FULL) - полный компанктинг таблицы (дефрагментация) - берёт полный доступ
            ---есть тонкая настройка---
        -- VACUUM ANALYZE -- собирает статистику и выдает план - по очистке, можно задать области сканирования. 
        -- Avtovacuum - автоматически каждый день повтор - можно настроить
     $ Рядом лежит файл VM (visitibility map)
## Индексы 

ИНДЕКС - это структура данных ускоряющая выборку данных из таблиц за счёт доп операций записи и 
работы с жёстким диском по средствам  используемых для хранения структуры данных и поддержания 
 её в актуальном состоянии. 

 -- Индекс - обьект бд - можно создать и удалить
 -- можно искать без полного перебора обьекты
 -- оптимизация выборки не большого числа записей
 -- не большое число - число относительно количества записей
 -- PRIMARY KEY и UNIGUE - уникальным столбцам индекс создается автоматически
 -- Индекс не бесплатный
## МЕТОДЫ СКАНИРОВАНИЯ -- 
Оптимизатор выбирает метод сканирования
 -- Индексное (index scan)
 -- Исключительно индексное хранение(index only scan) - файл карта таблицы
 -- Сканирование по битовой карте (bitmap scan) - возвращает карту типов и потом сканирует тоько подходящие адресса - без прыжков по HD (жёсткий диск)
 -- Последовательное сканирование (sequential scan) - сканирует всё по архитектуре - подходит ссд
 -- Может выдать всё кроме не нужного
## SELECT amname FROM pg_am;
 - B-tree(сбалансированное дерево) --btree
      - создается "по умолчанию" в (CREATE INDEX index_name ON table_name (column_name))
      - поддерживает операторы <,>,<=, >=, =.
      - LIKE'abc%'(но не '%abc')
      - NULL
      - сложность поиска O(logN) логарифмическая сложность - с увеличением масива данных - время затраченное растёт в квадратной прогресии
 - Хеш индекс --hash
      - CREATE INDEX index_name ON table_name USING HASH (column_name)
      - только оператор '='
      - не отражается в журнале (WAL  - хранятся все не сохранённые записи (типа кеша)) - можно что то удалить и не вернуть
      - сложность поиска константа 0(1) - нет приблизительных вариантов = 0
## СПЕЦИАЛИЗИРОВАННЫЕ ИНДЕКСЫ
 - GiST (обобщённое дерево поиска) --gist
      - для индексации геометрических типов данных и текста
 - GIN (обобщённый обратный) --gin
      - для индексации массива и ренж типов(диапазон)
      - полнотестовый поиск типа LIKE '%name'
 - SP-GiST (GiST с двоичным разбиением пространства) --spgist
      - поиск наборов данных с естественной упорядоченостью но не являются сбалансированным
 - RIN (блочно-диапазонный)        -- brin
      - полезен на упорядоченных огромных базах данных
## EXPLAIN ANALYZE-- анализатор
   -- поиск проблемы с производительностью
          EXPLAIN query - (EXPLAYN SELECT name) - запрос не исполняется - только аналитика - анализатор
                  - Собирает статистику по данным таблицы
                  - Планировщик смотрит на статистику при построении плана
          EXPLAIN ANALYZE guery  - показывает план и реальность 
              ANALIZE[table_name[(column1, cjlumn2...)]]   -- autovacuum запускает ANALYZE
                 запускать как минимум один раз в день

     -- примечание - индексы не работают на выражение LOWER, SUM ...
## ПРАКТИКА
  ---MD5() функция возвращает 128-битное представление контрольной суммы строки
   --MD5(random) - возвращает случайные числа
   --LOWER(column_name) - приводит данные к нижнему регистру 

### 1 ПРИМЕР 
CREATE TABLE perf_test(
	id int,
	reason text COLLATE "C",                  -- побайтовое сравнение символов латиницы
	annotation text COLLATE "C"
);
                                          -- заполняем случайными данными
INSERT INTO perf_test(id, reason, annotation) 
SELECT s.id, md5(random()::text), null            -- md5(randoom()) - шифровальная функция генерирует рандомный текст
FROM generate_series(1, 10000000) AS s(id)        -- тоже наверное функция
ORDER BY random();   -- упорядочить случайно      -- тоже функция

                                          -- что б записи были точно разными 
UPDATE perf_test                             
SET annotation = md5(UPPER(random()::text));

SELECT *
FROM perf_test
LIMIT 10

SELECT *
FROM perf_test
WHERE id = 3700000; --Successfully run. Total query runtime: 13 secs 782 msec. 1 rows affected.

--запускаем анализатор
EXPLAIN                                           -- анализатор
SELECT *
FROM perf_test
WHERE id = 3700000; -- Seq Scan - последовательный

 -- создаём Индекс для perf_test(id)
CREATE INDEX idx_perf_test_id ON perf_test(id);

 -- запускаем анализатор
EXPLAIN
SELECT *
FROM perf_test
WHERE id = 3700000; -- Index_Scan--теперь есть индекс--и время 39msec

### 2 ПРИМЕР
SELECT *
FROM perf_test
WHERE reason LIKE 'bc%' AND annotation LIKE 'AB%';

--чувак ошибся
UPDATE perf_test                             
SET annotation = UPPER(md5(random()::text));
-- in 12 min 27 secs.

EXPLAIN ANALYZE
SELECT *
FROM perf_test
WHERE reason LIKE 'bc%' AND annotation LIKE 'AB%'; --Paralel Seq--

                                                 -- после запуска анализа создается кеш и работать должно чуть быстрее
--ANALYZE perf_test;

                                                 --Индекс по двум колонкам сразу (Btree - умолчание)
CREATE INDEX idx_perf_test_reason_annotation ON perf_test(reason, annotation);
                                                 -- индекс по двум колонкам даёт возможность делать поиск и ПО ОДНОЙ из колонок.
                                                 --(Когда делаем индекс для двух столбцов то поиск Bit_map будет только 
                                                 --по первой колонке или по обеим одновременно
                                                 --по второй будет paralel seq). ЕСЛИ НУЖЕН btree по второму столбцу то ....ниже 

--запуск анализатора
EXPLAIN
SELECT *
FROM perf_test
WHERE reason LIKE 'bc%' AND annotation LIKE 'AB%'; --index Cond - индексное сканирование

--Запуск поиска
SELECT *
FROM perf_test
WHERE reason LIKE 'bc%' AND annotation LIKE 'AB%'; -- 47msec

--Запуск поиска по одной колонке
EXPLAIN
SELECT *
FROM perf_test
WHERE reason LIKE 'bc%'; --Bitmap Index_scan--382msec

EXPLAIN
SELECT *
FROM perf_test
WHERE annotation LIKE 'AB%'; -- paralel seq
                                                  --(Когда делаем индекс для двух столбцов то поиск Bit_map будет только 
                                                   --по первой колонке или по обеим одновременно
                                                   --по второй будет paralel seq) -- ЕСЛИ НУЖЕН btree по второму столбцу то --

 -- Строим индекс по второй колонке
CREATE INDEX idx_perf_test_annotation ON perf_test(annotation);

EXPLAIN
SELECT *
FROM perf_test
WHERE annotation LIKE 'AB%'; --28msec

EXPLAIN
SELECT *
FROM perf_test
WHERE LOWER(annotation) LIKE('ab%'); --paralel seq - из за функции LOWER() --10seq
-- или избегать выражений в поиске 
-- или строить индекс под 
CREATE INDEX idx_perf_test_annotation_LOWER ON perf_test(LOWER(annotation));
--LOWER(column_name) - приводит данные к нижнему регистру

EXPLAIN
SELECT *
FROM perf_test
WHERE LOWER(annotation) LIKE('ab%'); --30msec--Bitmap Heap Scan
## Сложный индекс для поиска по тексту--
EXPLAIN
SELECT *
FROM perf_test
WHERE reason LIKE '%bc%'; --1) Seg scan (последовательное) 2)--результатов 1146897 штук :)
--Индекс ДЖИН для ускорения поиска основан на правиле (тригран) базе функции
--1
CREATE EXTENSION pg_trgm; -- подключить расширение
--2
CREATE INDEX trgm_idx_perf_test_reason ON perf_test USING gin (reason gin_trgm_ops); --2min 18sec
--где gin - индекс, gin_trgm_ops - место хранения
--3
EXPLAIN ANALYSE
SELECT *
FROM perf_test
WHERE reason LIKE '%bc%'; --seq scan (последовательный) - 3 secs 68 msec -- 1145940 строк
-- Придется изменить условия поиска - слишком много результатов, планировщик отказывается 
--изменять способ поиска
--4 Если изменить условия поиска с (LIKE '%bc%') на (LIKE '%dfe%')
EXPLAIN ANALYSE
SELECT *
FROM perf_test
WHERE reason LIKE '%dfe%'; --Bitmap Scan, 22 msec
-- Изменился способ поиска --SQL сам выбирает метод поиска в зависимости от условий

# Массив, Variadic, foreach
## МАССИВ - колекция данных одного типа
  ```  
   -- Ономерные массивы - 10 чисел типа int, 7 строк типа varchar, 34 числа типа real.
   -- Столбцы и переменные можно обьявить как массив    
   -- Массивы могут быть многомерными
   -- Двумерный массив:
       2  3  6
       4  1  3
       5  2  7
   -- Трёхмерные и больше на практике не встречаются
      (кубы обрабатываются не массивами)

   - SQL - стандарт
     temp int ARRAY      -- без указания размера
     temp int ARRAY[4]   -- явный размер

   - Postgres:
     temp int[]         -- без указания размера
     temp int[8]        -- явный размер
  ----явно указанные размеры не на что не влияют! - 'для документации'.
    
      '{"a","b","c"}' -- инициализация массива (тип char или текстовый)
       '{1,2,3}'      -- инициализация массива (числа не нужно кавичить)
      '{{1,2,3}',{1,2,3},{1,2,3}}' --инициализация ДВУМЕРНОГО массива (тип число)

      -- ARRAY['a','b','c'] --инициализация массива (тип char или текстовый)
      -- ARRAY[1,2,3]        -- инициализация массива (числа не нужно кавичить)
      -- ARRAY[[1,2,3], [1,2,3], [1,2,3]] --инициализация ДВУМЕРНОГО массива (тип число)

        ---------Обращение к массиву
      - tmp[index] - взятие элемента по индексу(от 1)
      - Слайсинг (срезы):
            - tmp[1:3]  с 1го по 3й элемент
            - tmp[:4] с 1го по 4й
            - tmp[2:] со 2го до конца
        
        ---------функции для массивов
      - array_dims(arr_name) - возвращает размерность массива
             если одномерный вернет 1, если двумерный = 2 :)
      - array_length(arr_name, dim) - возвращает длинну массива
            - для ОДНОМЕРНОГО dim = 1; двумерный dim = 2;


    -------МАССИВЫ это не множества. Лучше содержать
            данные в таблице, проще с поиском и доступом---- 

-- создали ТАБЛИЦУ и обьявили поля
CREATE TABLE chess_game(
	white_player text,
	black_player text,
	moves text[],          -- ходы - текстовый массив
	final_state text[][]   -- ДВУМЕРНЫЙ МАССИВ -- состояние доски на конец партии
);

-- вставить записи через{}
INSERT INTO chess_game
       VALUES(
           'Caruana',
          'Nakamura', 
              '{"d4",  "d5", "c4",  "c6"}',
	    '{{"Ra8", "Qe8",  "x",  "x",  "x", "x", "x", "x"},
	     {  "a7",   "x",  "x",  "x",  "x", "x", "x", "x"},
	     { "Kb5", "Bc5", "d5",  "x",  "x", "x", "x", "x"}}'
               );
		 
SELECT *
FROM chess_game;

-- вставить записи через ARRAY[]
INSERT INTO chess_game
            VALUES(
                   'Caruana',
                  'Nakamura',
                ARRAY[  'd4',  'd5', 'c4', 'c6'],
	        ARRAY[['Ra8', 'Qe8',  'x',  'x', 'x', 'x', 'x', 'x'],
	              [ 'a7',   'x',  'x',  'x', 'x', 'x', 'x', 'x'],
	              ['Kb5', 'Bc5', 'd5',  'x', 'x', 'x', 'x', 'x']]
                     );
 -- возможно скобки лишние:)

 запрос на взятие нескольких элементов с 
 -- одномерного массива                                           
SELECT moves[2:3]
FROM chess_game;  

SELECT moves[:3]
FROM chess_game;

SELECT moves[3:]
FROM chess_game;

--найти значение в массиве
SELECT *
FROM chess_game
WHERE 'g6' = ANY(moves)

--возвращает размерность массива
SELECT array_dims(moves)
FROM chess_game

--возвращает длинну массива
SELECT array_length(moves, 1) -- одномерный массив
FROM chess_game;

SELECT array_dims(final_state), array_length(final_state, 1)
FROM chess_game;

-- обновить значения в массиве
UPDATE chess_game
SET moves = ARRAY['e4', 'd6', 'd4', 'kf6'];

 -- обновить  конкретный элемент в массиве
UPDATE chess_game
SET moves[4] = 'g6';
```
## ОПЕРАТОРЫ
```
 --Операторы сравнения
    =true если совпадают значения и последовательность
ARRAY[1, 2, 3, 4] = ARRAY[1, 2, 3, 4]; --true
    >true если в первой паре, єлемент слева больше
    <true если в первой паре, єлемент слева меньше
SELECT ARRAY[1, 2, 3, 4] = ARRAY[2, 1, 3, 4]; --false
SELECT ARRAY[2, 1, 3, 4] > ARRAY[1, 2, 3, 4]; --true
SELECT ARRAY[1, 2, 3, 5] > ARRAY[1, 2, 3, 4]; --true
--сравнение каждого элемента по очереди '5>4=true'
 --Containment Операторы (состава)
    @>true если левый массив включает все элементы правого
SELECT ARRAY[1, 2, 3, 4] @> ARRAY[1, 2]; --true если левый массив включает все элементы правого
SELECT ARRAY[1, 2, 3, 4] @> ARRAY[1, 2, 5]; --false
    @<true если правый массив включает все элементы левого
 -- Оператор пересечения:
   && true если в массивах есть хотя бы один одинаковый элемент
SELECT ARRAY[2, 1, 3, 4] && ARRAY[1, 2];--true
SELECT ARRAY[1, 2, 3, 4] && ARRAY[1, 2, 5]; --true
SELECT ARRAY[3, 3, 3, 4] && ARRAY[1, 2];--false

SELECT *
FROM chess_game
WHERE moves && ARRAY['d4']; - если есть пересечение в одной из записей - true
```
## VARIADIC----FOREACH
```
 -- Что бы передавать н-параметров (одного типа) в функцию,
 нужно обьявить аргумент как VARIADIC arg_name data_type[];
 
 функция выводит только чётные

DROP FUNCTION filter_even(VARIADIC numbers int[]);

CREATE FUNCTION filter_even(VARIADIC numbers int[]) RETURNS SETOF int AS $$
BEGIN
    FOR counter IN 1..array_upper(numbers, 1)
	LOOP
	    CONTINUE WHEN counter % 2 != 0; --когда каунтер не делится на 2 без остатка - переходить к следующей итерации
		RETURN NEXT counter;
	END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM filter_even(1, 2, 3, 4, 5, 6, 7, 8, 9);

---1/2 FOREACH counter IN ARRAY numbers--

CREATE FUNCTION filter_even(VARIADIC numbers int[]) RETURNS SETOF int AS $$
DECLARE
   counter int;
BEGIN
    --FOR counter IN 1..array_upper(numbers, 1)
	FOREACH counter IN ARRAY numbers  
   LOOP
	    CONTINUE WHEN counter % 2 != 0; --когда каунтер не делится на 2 без остатка - переходить к следующей итерации
		RETURN NEXT counter;
	END LOOP;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM filter_even(1, 2, 3, 4, 5, 6, 7, 8, 9);
```

# Импорт .csv
датасеты хранятся в виде .csv

http://www.kaggle.com/datasets

-создаём БД
-создаём таблицу с полями, типы данных которых
  должны соответствовать (колонкам) из csv-файла
-открываем psql и подаём команды:
      --psql db_connection_string
      --\copy table_name(fields) FROM 'csv_file_path' DELIMITR','CSV HEADER;


csv
psql --port 5432 --host=localhost --dbname=testdb --username=postgres

# CTE Иерархический РЕКУРСИВНЫЙ CTE
## CTE запрос Иерархический запрос
 Позволяют строить временные таблицы в рамках большого запроса
 -- Присваивая имя такой временной таблице(подзапросу),
    можно её переиспользовать
----------WITH name AS(
                       SELECT clause
                       )
          SELECT"using WITH part"
 -- Внутри подзапроса можно использовать INSERT, UPDATE, DELETE
 -- CTE выполняется единожды, результат кешируется (не всегда)
 -- Длинные и сложные запросы можно рефакторить в CTE,
    повышая чистоту и читабельность кода

--1)
SELECT company_name
FROM suppliers
WHERE country IN (SELECT country FROM customers);

--CREATE CTE -нет разделения - это всё одно действие
WITH customer_countries AS
(
	SELECT country FROM customers
)
SELECT company_name
FROM suppliers
WHERE country IN (SELECT * FROM customer_countries);

---WHERE NOT EXISTS - подзапрос
SELECT company_name
FROM suppliers
WHERE NOT EXISTS
      (
		  SELECT product_id
		  FROM products
		  JOIN order_details USING(product_id)
		  JOIN orders USING(order_id)
		  WHERE suppliers.supplier_id = products.supplier_id AND 
		        order_date BETWEEN '1998-02-01' AND '1998-02-04'
	  );
-- или вместо подзапроса 
SELECT company_name
FROM products
JOIN order_details USING(product_id)
JOIN orders USING(order_id)
JOIN suppliers USING(supplier_id) -- вместо (suppliers.supplier_id = products.supplier_id)
WHERE order_date BETWEEN '1998-02-01' AND '1998-02-04'

--создаём CTE
WITH filtered AS
(
     SELECT company_name, suppliers.supplier_id
     FROM products
     JOIN order_details USING(product_id)
     JOIN orders USING(order_id)
     JOIN suppliers USING(supplier_id) -- вместо (suppliers.supplier_id = products.supplier_id)
     WHERE order_date BETWEEN '1998-02-01' AND '1998-02-04'
)
SELECT company_name
FROM suppliers
WHERE supplier_id NOT IN (SELECT supplier_id FROM filtered);
--мой вариант - фишка в типах company_name
WITH filtered AS
(
     SELECT company_name
     FROM products
     JOIN order_details USING(product_id)
     JOIN orders USING(order_id)
     JOIN suppliers USING(supplier_id) -- вместо (suppliers.supplier_id = products.supplier_id)
     WHERE order_date BETWEEN '1998-02-01' AND '1998-02-04'
)
SELECT company_name
FROM suppliers
WHERE company_name NOT IN (SELECT * FROM filtered);
## UNION-----------
SELECT company_name
FROM suppliers
WHERE NOT EXISTS
      (
		  SELECT product_id
		  FROM products
		  JOIN order_details USING(product_id)
		  JOIN orders USING(order_id)
		  WHERE suppliers.supplier_id = products.supplier_id AND 
		        order_date BETWEEN '1998-02-01' AND '1998-02-04'
	  )
	  UNION
(WITH filtered AS
(
     SELECT company_name
     FROM products
     JOIN order_details USING(product_id)
     JOIN orders USING(order_id)
     JOIN suppliers USING(supplier_id) -- вместо (suppliers.supplier_id = products.supplier_id)
     WHERE order_date BETWEEN '1998-02-01' AND '1998-02-04'
)
SELECT company_name
FROM suppliers
WHERE company_name NOT IN (SELECT * FROM filtered));

----CTE1, CTE2, CTE3,....
WITH filtered AS
(
     SELECT company_name
     FROM products
     JOIN order_details USING(product_id)
     JOIN orders USING(order_id)
     JOIN suppliers USING(supplier_id) -- вместо (suppliers.supplier_id = products.supplier_id)
     WHERE order_date BETWEEN '1998-02-01' AND '1998-02-04'
),
CTE2 AS(
),
CTE3 AS(
),...
## РЕКУРСИВНЫЙ CTE 
DROP TABLE IF EXISTS employee;
--Создаём таблицу
CREATE TABLE employee (
	employee_id INT PRIMARY KEY,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL,
	manager_id INT,
	FOREIGN KEY (manager_id) REFERENCES employee (employee_id)
    ON DELETE CASCADE
);
INSERT INTO employee (
	employee_id,
	first_name,
	last_name,
	manager_id
)
VALUES
	(1, 'Windy', 'Hays', NULL),
	(2, 'Ava', 'CHristensen', 1),
	(3, 'Hassan', 'Conner', 1),
	(4, 'Anna', 'Reeves', 2),
	(5, 'Sau', 'Norman', 2),
	(6, 'Kelsie', 'Hays', 3),
	(7, 'Tory', 'Goff', 3),
	(8, 'Salley', 'Lester', 3);
## Рекурсивный JOIN - где видно подчинение
SELECT e.first_name || ' ' || e.last_name AS employee,
       m.first_name || ' ' || m.last_name AS manager
FROM employee e
LEFT JOIN employee m ON m.employee_id = e.manager_id
ORDER BY manager;
## рекурсивный СТЕ (WITH RECURSIVE)
WITH RECURSIVE submission(sub_line, employee_id) AS --где sub_line накапливает строчку
(
	SELECT last_name, employee_id FROM employee WHERE manager_id IS NULL --начальник IS NULL - у него нет начальника
	UNION ALL
	SELECT sub_line || '->' || e.last_name, e.employee_id -- sub_line соединяем со следующей записью через->
	FROM employee e, submission s                     -- присвоили псевдонимы e & s
	WHERE e.manager_id = s.employee_id                -- где менеджер из емплойе = employee_id из submission  
)
SELECT * FROM submission

# Оконные функции и ранжирование RANK, LAG, LEAD
## Оконные функции
  -- Позволяют обрабатывать группы строк
     без образования группировок в результирующем наборе (group by - без group by)
  -- Делятся на 
     ---агрегатные - (AVG, SUM, MIN, MAX, COUNT...)
     ---ранжирования - (ROW_NUMBER, RANK, LAG, LEAD..) - (можно добавить ранг--упорядочить)
  -- Отрабатывают после JOIN, WHERE, GROUP BY, HAVING, но перед ORDER BY
  -- "There was SQL before window functions and SQL after window functions" - очень важны в SQL
    Синтаксис
-function OVER(expression)
    Полный синтаксис
-function OVER([PARTITION expression], [ORDER BY expression]) - где [PARTITION expression] - группировка

----------примеры-АГРЕГАТНЫЕ-------
SELECT category_id, AVG(unit_price) AS avg_price
FROM products
GROUP BY category_id
LIMIT 5;

-- OVER (PARTITION BY category_id)  - вводим среднее значение ГРУППИРОВКА по category_id 

SELECT category_id, category_name, product_name, 
       unit_price, AVG(unit_price) OVER (PARTITION BY category_id) AS avg_price
FROM products
JOIN categories USING(category_id);

--- нарастающий результат по product_id (наростает по product_id, считает только по order_id(по заказам))
--- OVER(PARTITION BY order_id ORDER BY product_id) .... ORDER BY order_id

SELECT order_id, order_date, product_id, customer_id, unit_price AS sub_total,
     SUM(unit_price) OVER(PARTITION BY order_id ORDER BY product_id) AS sale_sum
FROM orders
JOIN order_details USING(order_id)
ORDER BY order_id;
## РАНЖИРОВАНИЕ
  -- ROW_NUMBER - присвоение уникального значения строкам
  -- RANK - присвоение ранга(веса) строкам с пропусками
  -- DENSE_RANK - присвоение ранга (веса) строкам без пропусков
  -- LAG - присвоение значений текущей строке, основанное на значении в предидущей
  -- LEAD - LAG наоборот - основанное на значении в следующей
  -- В LEAD & LAG можно передавать смещение (offset)

--- наростающий итог по всем order_id - нужно добавить нумерацию по строкам (row_number)
--- SELECT row_number() OVER(), order_id, order_date FROM orders

SELECT row_id, order_id, order_date, product_id, customer_id, unit_price AS sub_total,
     SUM(unit_price) OVER(ORDER BY row_id ) AS sale_sum
FROM (
	SELECT order_id, order_date, product_id, customer_id, unit_price,
        row_number() OVER() AS row_id
        FROM orders
        JOIN order_details USING(order_id)
      ) AS sabquery

ORDER BY order_id;
--ROW_NUMBER--так же как и RANK -- имитирует id PRIMARY KEY - 
SELECT product_name, units_in_stock,
                           row_number() OVER(ORDER BY product_id) -- упорядочить по product_id и пронумеровать соответственно
FROM products;
      --работает без ORDER BY
SELECT product_name, units_in_stock,
                           row_number() OVER()
FROM products;

-------row_number() OVER(PARTITION BY category_name ORDER BY product_id) -- нумерация по группировке (PARTITION BY category_name)
SELECT category_name, product_id, product_name, units_in_stock,
                           row_number() OVER(PARTITION BY category_name ORDER BY product_id)
FROM products
JOIN categories USING(category_id);
## RANK
           -- так же как и ROW_NUMBER -- имитирует id - 
	   -- без ORDER BY не работает
SELECT * FROM products;
SELECT product_name, units_in_stock,
                           RANK() OVER(ORDER BY product_id)
FROM products;
--Делает промежутки между рангами = количеству строк в предыдущем ранге
SELECT product_name, units_in_stock,
                           RANK() OVER(ORDER BY units_in_stock)
FROM products;
--DENSE_RANK - не делает таких промежутков - так же как и --ROW_NUMBER--
SELECT product_name, units_in_stock,
                           DENSE_RANK() OVER(ORDER BY units_in_stock)
FROM products;
---ранг основанный на логике
SELECT product_name, unit_price,
       DENSE_RANK() OVER(
	        ORDER BY
	           CASE
	               WHEN unit_price > 80 THEN 1
	               WHEN unit_price > 30 AND unit_price < 80 THEN 2
	               ELSE 3
	           END
	   ) AS ranking
FROM products
ORDER BY unit_price DESC;
## LAG ----- ДРУГОЙ принцып РАБОТАЕТ НЕ ТАК КАК RANK-i  !!!

            -- обязателен аргумент---LAG(unit_price)
            -- при вычитании "-" даёт разницу, если убрать минус будет стоять 
            --смещает значение из строки аргумента на строчку ниже
            -- PARTITION делать нет смысла но принимает
SELECT product_name, unit_price,
       LAG(unit_price) OVER(ORDER BY unit_price DESC) - unit_price AS price_lag
FROM products
ORDER BY unit_price DESC;
-- LAG наоборот
SELECT product_name, unit_price,
       LAG(unit_price) OVER(ORDER BY unit_price DESC) AS price_lag
FROM products
ORDER BY unit_price DESC;
## LEAD - LAG наоборот - основанное на значении в следующей
SELECT product_name, unit_price,
       LEAD(unit_price, 2) OVER(ORDER BY unit_price ) - unit_price AS price_lag
FROM products
ORDER BY unit_price ;
## пример1--
SELECT *
FROM products
WHERE product_id = ANY (
	SELECT product_id
	FROM (
	     SELECT product_id, unit_price,
         ROW_NUMBER() OVER(ORDER BY unit_price DESC) AS nth
         FROM products
	     )AS sorted_prices
    WHERE nth < 4
    );

--короче--
SELECT *
FROM (SELECT product_id, product_name, category_id, unit_price, units_in_stock,
	 ROW_NUMBER() OVER(ORDER BY unit_price DESC) AS nth
	 FROM products
	 ) AS sorted_prices
WHERE nth < 4
ORDER BY unit_price

--совсем коротко--
SELECT *
FROM products
WHERE product_id = ANY (
	SELECT product_id
    FROM products
    ORDER BY unit_price DESC
    LIMIT 3
    );
## пример2---
SELECT *
FROM 
(
	SELECT order_id, product_id, unit_price, quantity,
	    RANK() OVER(PARTITION BY order_id ORDER BY (quantity) DESC) AS rank_quant
	FROM orders
	JOIN order_details USING(order_id)
) AS subquery
WHERE rank_quant <=3

# Транзакции
## ТРАНЗАКЦИИ--
• Что если часть данных обновилась, а часть нет?
• Транзакция - логическая группа операций
• Транзакция может быть выполнена только целиком
• Классический пример: банковская транзакция
--- прочесть баланс на счету Х
--- уменьшить баланс на Z денежных единиц
--- сохранить новый баланс счёта Х
--- прочесть баланс на счету Y
--- увеличить баланс на Z денежных единиц
--- сохранить новый баланс счёта Y
## ACID
• Atomiciti (атомарность) -- ROLLBACK-- откат если не удалось завершить полностью
• Consistency (согласованность) -- логика должна быть выдержана
• Isolation_(изолированность) -- данные обрабатываемы одной транзакцией изолируются от 
                              --других транзакций на время требуемое для выполнения данной транзакции
• Durability (долговечность) -- если транзакция выполнена то сбои уже не повлияют на неё
        ----------------TCL---Transaction Control Language-----------
		--START TRANSACTION (pgsql)   OR BEGIN[TRANSACTION] (sql)-- начать транзакцию
		--END[TRANSACTION]            OR COMMIT                  -- закончить
		пример
		BEGIN;
		statement1;
		statement2;
		statement3;
		COMMIT;
		--все операции обёрнуты в транзакции в любом случае (некоторые умолчания просто хранятся в базе)
• Если по ходу транзакции что-то пошло не так, её можно откатить:
--ROLLBACK
• Внутри транзакции можно делать «засечки» в «важных» местах с помощью команды
--SAVEPOINT savepoint_name
• Откат к точке --ROLLBACK TO savepoint_name
  ------------ограничения--------------
• Нельзя создавать транзакции в функциях
• Хотя функции неявно исполняются в рамках транзакции
• Чтобы прервать транзакцию изнутри функции 
-- RAISE EXCEPTION
• Если хотим откатить только часть можно «имитировать» savepoint с помощью 
--BEGIN и EXCEPTION WHEN
• Концепция подразумевает, что уровнем изоляции и откатами управляет внешний
  код (SQL-скрипт или код приложения верхнего уровня)
  ---------ИЗОЛИРОВАННОСТЬ--------------
• Isolation (изолированность)
## Проблемы параллельности:
  -- <<грязное» чTEHИe-ЧTeние частичных изменений
  -- неповторяемое чтение - повторное чтение показывает, что данные были изменены
     после первого чтения
  -- Фантомное чтение - повторное чтение показывает другой результирующий набор
  -- аномалия сериализации - результат параллельно выполняемых транзакций может
     не согласовываться с результатом этих же транзакций, выполняемых по очереди	 
## Уровень изоляции                   «Грязное» чтение           Неповторяемое     Фантомное чтение         Аномалия сериализации
                                                                чтени             
Read uncommited (Чтение              Допускается,               Возможно          Возможно                  Возможно
незафиксированных данных             но не в PG                                            
Read committed (Чтение               Невозможно                 Возможно          Возможно                  Возможно
зафиксированных данных                                                               
Repeatable read (Повторяемое         Невозможно                 Невозможно        Допускается,но нев PG     Возможно
чтение)                                                                                                                                                          
Serializable (Cериализуемость)       Невозможно                 Невозможно        Невозможно                Невозможно

### ПОДХОДЫ к ИЗОЛЯЦИИ--------
 --------------Изоляция транзакций
• MVCC (multiversion control system)
• Изменений версионируются таймштампами (индексы версий по времени сохранения)
• Механизм изоляции ориентируется на снимки данных по таймштампам
• Снижает потребность блокировок	
	------управление
-- В начале транзакции
 BEGIN ISOLATION LEVEL level   -- (level -READ, READ COMMITED, REPEATABLE READ, SERIALIZABLE)
-- Внутри транзакции 
 SET TRANSACTION ISOLATION LEVEL level  -- (level -READ, READ COMMITED, REPEATABLE READ, SERIALIZABLE)
## Пример_1
BEGIN;        -- запускает транзакцию

WITH prod_update AS (  --создать CTE которая обновляет discontinued
	UPDATE products
	SET discontinued = 1
	WHERE units_in_stock < 10
	RETURNING product_id   -- возвращает product_id обновлённых
)
SELECT * INTO last_orders_on_discontinued  -- копировать * в таблицу - паралельно создав таблицу
FROM order_details                         -- из 
WHERE product_id IN (SELECT product_id FROM prod_update); -- где product_id обновлён

DELETE FROM order_details
WHERE product_id IN (SELECT product_id FROM prod_update);

ROLLBACK;  -- если не достигнут COMMIT (транзакция не завершина) делает возврат (в средине транзакции)

COMMIT;  -- завершает транзакцию --или ДЕЛАЕТ ROLLBACK если была ошибка в транзакции

SELECT * FROM last_orders_on_discontinued;

DROP TABLE IF EXISTS last_orders_on_discontinued;
## 2
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;  --уровень изоляции

WITH prod_update AS (  
	UPDATE products
	SET discontinued = 1
	WHERE units_in_stock < 10
	RETURNING product_id   
)
SELECT * INTO last_orders_on_discontinued  
FROM order_details                         
WHERE product_id IN (SELECT product_id FROM prod_update);

SAVEPOINT backup;    -- поставили метку что б вернуть всё до неё (сохранялка)

DELETE FROM order_details
WHERE product_id IN (SELECT product_id FROM last_orders_on_discontinued );

ROLLBACK TO backup; --вернуть SAVEPOINT backup, всё остальное останется выполненным !!!только не завершённые транзакции!!!

UPDATE order_details
SET quantity = 0
WHERE product_id IN (SELECT product_id FROM last_orders_on_discontinued ); 

COMMIT;  

SELECT * 
FROM order_details
WHERE product_id IN (SELECT product_id FROM last_orders_on_discontinued ); 


DROP TABLE IF EXISTS last_orders_on_discontinued;

SELECT COUNT(*) FROM order_details;
SELECT * FROM last_orders_on_discontinued; 
## 3
BEGIN;

ALTER TABLE employees
ADD COLUMN salary decimal(12,2); --с двумя знаками после запятой

UPDATE employees 
SET salary = random()*100;  -- случайное число от 0,1 - 1

UPDATE employees
SET salary = 22
WHERE employee_id = 1
COMMIT;

UPDATE employees
SET salary = 22
WHERE employee_id = 1

SELECT * FROM employees
 
BEGIN;

UPDATE employees
SET salary = salary * 1.5
WHERE salary < 50;

SAVEPOINT inc;

UPDATE employees
SET salary = salary * 0.8
WHERE salary > 80;

ROLLBACK TO inc;

UPDATE employees
SET salary = 0
WHERE employee_id = 7;

COMMIT;
 
SELECT *
from employees
order by salary 

# Событийное программирование в Postgres
•  Что если надо программно отреагировать на изменение данных (событие)?
•  Триггер - объект, который назначает действие на те или иные события
•  Триггеры могут реагировать как на построчное изменение (множественное срабатывание),
   так и единожды на все изменения сразу
  •  Сценарии использования триггеров:
     • Аудит таблиц
     • Дополнительные действия в ответ на изменения
     • Сложные проверки целостности
  • Что необходимо:
     •  Для совершения действия необходима функция
     •  И сам объект триггера
## "Построчные" триггеры
### Создание триггера:
CREATE TRIGGER trigger_name() condition ON table_name
FOR EACH ROW EXECUTE PROCEDURE function_name();
 Создание триггера:

CREATE TRIGGER trigger_name() condition ON table_name
REFERENCING [NEW, OLD] TABLE AS ref_table_name
FOR EACH STATEMENT EXECUTE PROCEDURE function_name();

• Необходимо задекларировать референсную таблицу
• В функции доступна переменная TG_OP, хранящая тип операции
  (INSERT, UPDATE, DELETE)
• Такие триггеры часто используются для настройки аудита
### Управление объектами триггеров
- Удаление триггера:
DROP TRIGGER IF EXISTS trigger_name ON table_name
- Переименование триггера:
ALTER TRIGGER trigger_name ON table_name
RENAME TO new_trigger_name;
- Отключение триггера:
ALTER TABLE table_name
DISABLE TRIGGER trigger_name;
- Отключение всех триггеров на таблице:
ALTER TABLE table_name
DISABLE TRIGGER ALL;
Condition:
[BEFORE, AFTER] [INSERT, UPDATE, DELETE]

Например:
BEFORE INSERT                        -- перед 
AFTER UPDATE                         -- после
BEFORE INSERT OR UPDATE              ---перед  ____или____
DROP TRIGGER IF EXISTS trigger_name  ---Удаление
## функции на построчные триггеры
```
    Вид функции, привязываемой триггером:
CREATE FUNCTION func_name() RETURNS trigger AS $$               ---- trigger
BEGIN;

END;
$$ LANGUAGE plpgsql;

• Должна возвращать NULL или запись,
  соответствующую структуре таблицы, на которую будет вешаться триггер!

• Через аргумент NEW есть доступ к вставленным и модифицированным строкам
(INSERT/UPDATE триггеры)

• Через аргумент OLD есть доступ к вставленным и удалённым строкам
(UPDATE/DELETE триггеры)

-----------Практика по возврату из триггеров
• Должна возвращать NULL или запись,
соответствующую структуре таблицы, на которую будет вешаться триггер!

• Если BEFORE-триггер возвращает NULL, то сама операция и АFTER-триггеры будут отменены
• BEFORE-триггер может изменить строку (INSERT\UPDATE) через NEW и тогда
  операция и AFTER-триггеры будут работать с изменённой строкой
• Если BEFORE-триггер не «хочет>> изменять строку, то надо просто вернуть NEW
• В случае BEFORE-триггера реагирующего на DELETE, возврат не имеет значения
  (кроме NULL: отменит DELETE)
• NEW = null при DELETE, так что если BEFORE-триггер хочет дать ход DELETE, надо вернуть OLD
• Возвращаемое значение из построчного AFTER-триггера
  (или и из BEFORE и из AFTER Tpиггеров на утверждения)
  игнорируется => можно возвращать NULL
• Если построчный AFTER-триггер или триггер на утверждение хочет отменить операцию => raise exception
```
## Триггеры на утверждения
- Создание триггера:

CREATE TRIGGER trigger_name() condition ON table_name
REFERENCING [NEW, OLD] TABLE AS ref_table_name
FOR EACH STATEMENT EXECUTE PROCEDURE function_name();

• Необходимо задекларировать референсную таблицу
• В функции доступна переменная TG_OP, хранящая тип операции
  (INSERT, UPDATE, DELETE)
• Такие триггеры часто используются для настройки аудита
## поСТРОЧНЫЕ триггеры
--функция логирует дату внесения изминений-----
### ПРИМЕР_1 -- TIME--
```
ALTER TABLE customers
ADD COLUMN last_updated timestamp;  -- добавить колонку в таблицу

CREATE OR REPLACE FUNCTION track_changes_on_customers() RETURNS trigger AS $$ -- функция возвращает триггер
BEGIN
    NEW.last_updated = now();                               -- NEW обращаемся к новой колонке 
	RETURN NEW;                                             -- now() - функция которая возвращает текущее время
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS customers_timestamp ON customers; -- удалить триггер, вдруг уже есть такой  

CREATE TRIGGER customers_timestamp BEFORE INSERT OR UPDATE ON customers -- создать триггер___если изменить или обновить  'таблицу'
FOR EACH ROW EXECUTE PROCEDURE track_changes_on_customers();            -- ДЛЯ КАЖДОГО РЯДА ВЫПОЛНИТЬ ПРОЦЕДУРУ track_changes_on_customers
-----------------------
SELECT * FROM customers;

UPDATE customers 
SET city = 'London'
WHERE customer_id = 'BBBBB';

INSERT INTO customers
VALUES ('ABCDE', 'company', 'contact', 'title', 'adress', 'city', null, 'code', 'country', '', '', null);
```
### ПРИМЕР2  ЧЕЛОВЕКА логирует 'person' внёсшего изменения
```
Функция которая логирует 'person' внёсшего изменения
ALTER TABLE employees
ADD COLUMN user_changed text;  --------------------------------------------------1) добавили колонку в таблицу

CREATE OR REPLACE FUNCTION track_changes_on_employees () RETURNS trigger AS $$ --2)создать функцию
BEGIN
    NEW.user_changed = session_user;
	RETURN NEW;
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS employees_user_change ON employees;
CREATE TRIGGER employees_user_change BEFORE INSERT OR UPDATE ON employees  ------3) Создать триггер и привязать колонку к функции
     FOR EACH ROW EXECUTE PROCEDURE track_changes_on_employees ();
	 
SELECT * FROM employees;

UPDATE employees
SET reports_to = 2
WHERE employee_id = 2;

INSERT INTO employees
VALUES (10, ' ', ' ', '', '', null, null, '', '', '', '', ' ', '', '', null, '', null, 0, null);
```
## ТРИГГЕРЫ НА УТВЕРЖДЕНИЯ
```
DROP TABLE IF EXISTS products_audit;

CREATE TABLE products_audit
(
	op char(1) NOT NULL,
	user_changed text NOT NULL,
	time_stamp timestamp NOT NULL,
	
	product_id smallint NOT NULL,
    product_name varchar(40) NOT NULL,
    supplier_id smallint,
    category_id smallint,
    quantity_per_unit varchar(26),
    unit_price real,
    units_in_stock smallint,
    units_on_order smallint,
    reorder_level smallint,
    discontinued integer NOT NULL);

--DROP FUNCTION build_audit_products;

CREATE OR REPLACE FUNCTION build_audit_products() RETURNS trigger AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN                                              -- если TG_OP (функция которая выдаёт действие(insert, update, delete)) = то
	     INSERT INTO products_audit                                    
	     SELECT 'I', session_user, now(), nt.* FROM new_table nt;     -- вставить 'I', session_user, now(), (nt.* FROM new_table nt) = (псевдоним навешен на таблицу)
	ELSEIF TG_OP ='UPDATE' THEN
	     INSERT INTO products_audit
	     SELECT 'U', session_user, now(), nt.* FROM new_table nt;
	ELSEIF TG_OP ='DELETE' THEN
	     INSERT INTO products_audit
	     SELECT 'D', session_user, now(), ot.* FROM old_table ot;
	END IF;
	RETURN NULL;
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS audit_products_insert ON products;

CREATE TRIGGER audit_products_insert AFTER INSERT ON products
REFERENCING NEW TABLE AS new_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_products();

DROP TRIGGER IF EXISTS audit_products_update ON products;
CREATE TRIGGER audit_products_update AFTER UPDATE ON products
REFERENCING NEW TABLE AS new_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_products();

DROP TRIGGER IF EXISTS audit_products_delete ON products;
CREATE TRIGGER audit_products_delete AFTER DELETE ON products
REFERENCING OLD TABLE AS old_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_products();

INSERT INTO products
VALUES (79, 'kozlova mozzarella', 7, 4, '200g per package', 50, 20, 0, 0, 0 );

UPDATE products
SET unit_price = 70
WHERE product_id = 79;

DELETE FROM products
WHERE product_id = 79;

------------------------------------------
SELECT * FROM products ORDER BY product_id DESC;
DELETE FROM products
WHERE product_id = 79;
SELECT * FROM products_audit; 
```

# Безопасность в PostgreSQL
## Теория
• Безопасность - аспект многогранный
• Одна из проблем - управление пользователями и их доступом
• Самый частый подход: создание ролей и назначение ролей пользователям
• Роль - совокупность разрешений и запретов на доступ к БД и её объектам
• Пользователь - учётная запись с логином и паролем
• B PostgreSQL есть только роли
• Пользователи - роли с паролями
• На роль с паролем назначают роль с прописанными доступами
• По сути, можно создавать дерево ролей, поскольку разрешена вложенность
• Также роли называют группами
• Роли не привязываются к конкретной БД (а создаются на уровне экземпляра сервера)

## 6 уровней безопасности
• Экземпляра (кластера/сервера/инсталляции): аутентификация, создание БД,
  управление безопасностью и т.д.
• Базы данных: подключение к конкретной БД, создание в ней ролей и т.д.
• Схемы: управление схемами (создание, удаление)
• Таблицы: CRUD-операции над таблицами
• Колонки: операции над конкретной колонкой конкретной таблицы
• Строки таблицы

## Создание роли и серверные привилегии
CREATE ROLE role_name [PRIVILEGE];
 Privilege:
  • LOGIN [NOLOGIN]
  • SUPERUSER [NOSUPERUSER]
  • CREATEDB [NOCREATEDB]
  • CREATEROLE [NOCREATEROLE]
  • REPLICATION [NOREPLICATION]

-------------------------------SELECT rolname FROM pg_roles - список ролей

• postgres -роль, создаваемая по умолчанию,
  и единственная имеющая привилегии SUPERUSER

### Создание роли и «пользователя»
• CREATE ROLE role_name LOGIN
• CREATE USER user_name
• CREATE USER user_name WITH PASSWORD '*********
• CREATE ROLE role_name создаёт «бесправную» роль (NOLOGIN, NOSUPERUSER и т.д.)

### Лучшие практики
• Не работать с кластером из под SUPERUSER (postgres).
  Создать пользователя с правами CREATEDB и CREATEROLE и работать из под него.
• Создавайте отдельные роли для отдельных приложений
  по сценариям использования: readonly, readwrite
• Давать минимальные права, требующиеся ролями
  Всё запрещено, что не разрешено (а не «всё разрешено, что не запрещено»)
  Т.е., сначала нужно изъять все права, а затем наделять только необходимыми
• Создавайте пользователей по акторам приложения: hr_user, reporting_user, admin_stuff
  Затем назначайте пользователям роли с минимально требуемыми правами (но с достаточными)
• Если на пользователя назначены две роли
  и одна запрещает создание БД, а другая разрешает, то разрешение в приоритете
• При создании БД, создаётся схема public --------------------------------------------------!!!
  и скрытая роль public с доступом на запись в схему public
• Всем новым ролям/пользователям неявно даются привилегии роли public
   REVOKE CREATE ON SCHEMA public FROM public
   REVOKE ALL ON DATABASE db_name FROM public
• Бойтесь роль CREATEROLE. Если «user» имеет привилегию CREATEROLE, но не CREATEDB,
  то он может создать новую роль с правами CREATEDB! CREATEROLE - почти SUPERUSER!
•  GRANT operation ON TABLE table_name TO role
   • SELECT
   • INSERT
   • UPDATE
   • DELETE
   • TRUNCATE
   • REFERENCES - создание внешних ключей
   • TRIGGER
   • GRANT operation ON ALL TABLES IN SCHEMA schema_name TO role

### Доступ к колонкам
• GRANT operation (columns) ON TABLE table_name TO role
  • SELECT
  • INSERT
  • UPDATE
  • REFERENCES - возможность ссылаться на колонки с внешним ключом

### Доступ к строкам
 • Необходимо включать явно:
   ALTER TABLE table ENABLE ROW LEVEL SECURITY;
 • После включения row-level security, доступ к данным запрещён
   Необходимо дать доступ явно через политику:
    • CREATE POLICY policy_name ON table_name
    • FOR operation TO role
    • USING (expression)
 • Две и более политик работают как запросы с UNION между ними, т.е.
   в результирующий набор будут включаться все данные из этих политик,
   даже если они по смыслу ДРУГ ДРУГУ Противоречат
    • DROP POLICY policy_name ON table_name

### Удаление

 • Чтобы удалить роль, необходимо:
   • изъять все выданные ранее права
   • переназначить все объекты на другую роль, которыми владела удаляемая роль
     • ALTER TABLE table_name OWNER TO other_role

 • Незначимые объекты, которыми владела удаляемая роль можно удалить автоматом:
   • DROP OWNED BY deleting_role [CASCADE] [RESTRICT], restrict is default
 • REVOKE ALL PRIVILEGES ON table1, table2... FROM deleting_role
 • REVOKE ALL ON DATABASE db_name FROM deleting_role
 • REVOLE ALL ON SCHEMA schema_name FROM deleting role
 • ...иными словами - удаляем все зависимые объекты и изымаем выданные права
 • и в конче: DROP ROLE [IF EXISTS] deleting role

### Пример
DROP ROLE sales_stuff;                          -- удалить роль
DROP ROLE northwind_admins;

CREATE ROLE sales_stuff;                        -- создать роль -- по умолчанию все привилегии закрыты
CREATE ROLE northwind_admins;
--зайти через роль пока нельзя - нет логина и пароля

DROP USER john_smith;  --удалить пользователя
DROP USER north_admin1;

CREATE USER john_smith WITH PASSWORD 'qwerty';  -- создать пользователя
CREATE USER north_admin1 WITH PASSWORD 'qwerty';

                                                -- теперь можно подключится к >psql northwind
                                                -- хоть мы и не давали прав (проблемы с пабликом)
                                                -- НУЖНО ИЗЬЯТЬ ПРАВА ИЗ роли public - присваивается по умолчанию для новых пользоватилей
REVOKE CREATE ON SCHEMA public FROM public;     -- отменить возможность создания чего либо в схеме public
REVOKE ALL ON DATABASE northwind FROM public;   -- отменить все права схемы public
                                                -- теперь всё закрыто и доступ к базе данных тоже
                                                          
GRANT USAGE ON SCHEMA public TO sales_stuff;             -- дозвіл на дії у схемі (ДОЗВОЛИТИ ВИКОРИСТАННЯ НА СХЕМІ public для ...) 

GRANT USAGE ON SCHEMA public TO northwind_admins;        -- дозвіл на дії у схемі (ДОЗВОЛИТИ ВИКОРИСТАННЯ НА СХЕМІ public для ...)
GRANT CREATE ON SCHEMA public TO northwind_admins;       -- дозвіл на СТВОРЕННЯ СХЕМ (облікові записи) для public для ...адміна
GRANT CREATE ON DATABASE northwind TO northwind_admins;  -- дозвіл на СТВОРЕННЯ у DATABASE (у СХЕМІ(...облікові записи)) ....

GRANT sales_stuff TO john_smith;                         -- связать роль и пользователя
GRANT northwind_admins TO north_admin1;                  -- связать роль и пользователя 

--------------cmd-----------
psql northwind north_admin1
password qwerty

create table admin_demo2 (names varchar); --+
drop table admin_demo2;                                  -- пользователь создавший таблицу имеет к этой таблице полный доступ

>psql northwind john_smit
password qwerty
create table admin_demo2 (names varchar); --'-'          -- не має дозволу на CREATE


SELECT grantee, privilege_type                           -- посмотреть привилегии по таблице (grant ebaut table)
FROM information_schema.role_table_grants
WHERE table_name = 'admin_demo2';              
          
### GRANT FOR TABLE разрешить доступ к КОНКРЕТНЫМ таблицам

#### разрешить доступ к КОНКРЕТНЫМ таблицам---------------
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE            -- разрешение на SELECT,...
public.orders,                                           -- схема public. для таблиц orders...
public.order_details,
public.products
TO sales_stuff;                                          -- для роли sales_stuff;

GRANT SELECT ON TABLE public.employees TO sales_stuff;   -- grant for select on table employees:)

#### разрешить доступ КО ВСЕМ таблицам
 
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES, TRIGGER ON ALL TABLES
IN SCHEMA public
TO northwind_admins;

#### Права на уровне колонок----------------
--но сначала нужно изьять ПРИВИЛЕГИИ на SELECT--
REVOKE SELECT ON employees FROM sales_stuff;
---Нужно явно прописать колонки к которым пользователь будет иметь доступ---
GRANT SELECT (employee_id, last_name, first_name, title, title_of_courtesy, birth_date, hire_date, 
			 city, region, postal_code, country, home_phone, extension, photo, notes, reports_to, photo_path)
ON employees
TO sales_stuff;
SELECT * FROM employees; --работать не будет
SELECT last_name FROM emploees;  -- будет работать
 
#### ПРАВА НА УРОВНЕ СТРОК-----
ALTER TABLE products
ENABLE ROW LEVEL SECURITY;  -- ограничение на все строки для всех пользователей

CREATE POLICY active_products_for_sales_stuff ON products -- разрешение для sales_stuff где discontinued <> 1
FOR SELECT
TO sales_stuff
USING (discontinued <> 1);

SELECT * FROM products;

#### противоречие в политике 

CREATE POLICY reorderer_products_for_sales_stuff ON products -- разрешение для sales_stuff где discontinued <> 1
FOR SELECT
TO sales_stuff
USING (reorder_level > 10);

---------в итоге открылись ещё позиции 

DROP POLICY reorderer_products_for_sales_stuff ON products;  -- удалить политику 

### Удаляем все права и удаляем роли
REVOKE ALL PRIVILEGES ON employees, orders, order_details, products FROM sales_stuff; -- снять все разрешения с таблиц

REVOKE ALL ON DATABASE northwind FROM sales_stuff;    -- снять все разрешения с northwind

REVOKE ALL ON SCHEMA public FROM sales_stuff;  -- снять все разрешения и привязку

DROP POLICY reorderer_products_for_sales_stuff ON products;

DROP POLICY active_products_for_sales_stuff ON products;

DROP ROLE sales_stuff;

### ПОЛЕЗНЫЕ ТАБЛИЧКИ-----
SELECT * FROM pg_roles;

SELECT grantee, privilege_type                           -- посмотреть привилегии по таблице (grant ebaut table)
FROM information_schema.role_table_grants
WHERE table_name = 'admin_demo2'; 

# Home work 
## HOME WORK 2
1
SELECT *
FROM orders
WHERE ship_country IN('France', 'Austria', 'Spain')

2
SELECT *
FROM orders
ORDER BY required_date DESC, shipped_date

3.1
SELECT MIN(unit_price)
FROM products
WHERE units_in_stock > 30
3.2
SELECT MIN(unit_price)
FROM products
WHERE units_in_stock > 30 AND discontinued < 1

4
SELECT MAX(units_in_stock)
FROM products
WHERE unit_price > 30
4.1
SELECT MAX(units_in_stock)
FROM products
WHERE unit_price > 30 OR discontinued > 1

5 
SELECT AVG(shipped_date - order_date)
FROM orders
WHERE ship_country = 'USA'
6
SELECT SUM(unit_price * units_in_stock)
FROM products
WHERE discontinued <> 1
## HOME WORK 3
1 SELECT *
FROM orders
WHERE ship_country LIKE'U%'

2
SELECT order_id, customer_id, freight, ship_country
FROM orders
WHERE ship_country LIKE 'N%'
ORDER BY freight DESC
LIMIT 10

3
SELECT first_name, last_name, home_phone, region
FROM employees
WHERE region IS NULL

4
SELECT COUNT(*)
FROM customers
WHERE region IS NOT NULL

5
SELECT country, COUNT(*)
FROM suppliers
GROUP BY country
ORDER BY COUNT(*) DESC

6
SELECT ship_country, SUM(freight)
FROM orders
WHERE ship_region IS NOT NULL
GROUP BY ship_country
HAVING SUM(freight) > 2750
ORDER BY SUM(freight)

7
SELECT country
FROM suppliers
UNION
SELECT country
FROM customers
ORDER BY country

8
SELECT country
FROM customers
INTERSECT
SELECT country
FROM suppliers
INTERSECT
SELECT country
FROM employees

9
SELECT country
FROM customers
INTERSECT
SELECT country
FROM suppliers
EXCEPT
SELECT country
FROM employees
## HOME WORK  JOIN (обьединения) and As (псевдоним)

### 1.
SELECT c.company_name, CONCAT(e.first_name, ' ', e.last_name)
FROM orders AS o                                 -- из таблицы ордерс, AS О
JOIN customers AS c USING(customer_id)           -- обьединили по ключу AS С
JOIN employees AS e USING(employee_id)           -- обьединили по ключу AS Е
JOIN shippers AS s ON o.ship_via = s.shipper_id  -- обьединили по (разным столбикам) AS S
WHERE c.city = 'London' AND e.city = 'London' AND s.company_name = 'Speedy Express'

### 2
variant1
SELECT p.product_name, p.units_in_stock, s.contact_name, s.phone, c.category_name
FROM products AS p
JOIN suppliers AS s USING(supplier_id)
JOIN categories AS c USING(category_id)
WHERE c.category_name IN('Beverages', 'Seafood') AND discontinued = 0 
      AND p.units_in_stock < 20

       v2
SELECT p.product_name, p.units_in_stock, s.contact_name, s.phone, c.category_name
FROM products AS p
JOIN suppliers AS s USING(supplier_id)
JOIN categories AS c USING(category_id)
WHERE c.category_name = 'Beverages' OR c.category_name = 'Seafood' 
      AND p.discontinued = 0 AND p.units_in_stock < 20
ORDER BY p.units_in_stock

### 3
SELECT contact_name, order_id
FROM customers
LEFT JOIN orders USING(customer_id)
WHERE order_id is NULL

### 4
SELECT contact_name, order_id
FROM orders
RIGHT JOIN customers USING(customer_id)
WHERE order_id is NULL
## HOMEWORK 4 Подзапрос
### 1
SELECT product_name, units_in_stock
FROM products
WHERE units_in_stock < ALL(SELECT AVG(quantity)       -- "< ALL - меньше любого из списка"
                   FROM order_details
		  GROUP BY product_id)
ORDER BY units_in_stock DESC

### 2
SELECT customer_id, SUM(freight) AS freight_sum
FROM orders
INNER JOIN (SELECT customer_id, AVG(freight) AS freight_avg
            FROM orders
            GROUP BY customer_id) AS oa
USING(customer_id)
WHERE freight > freight_avg AND shipped_date BETWEEN '1996-07-15' AND '1996-07-30'
GROUP BY customer_id
ORDER BY freight_SUM

### 3
SELECT customer_id, ship_country, order_price
FROM orders
JOIN (SELECT order_id, SUM(unit_price * quantity - unit_price * quantity * discount) AS order_price
      FROM order_details
      GROUP BY order_id) AS od
  USING(order_id)
WHERE ship_country IN ('Argentina', 'Bolivia', 'Brazil', 'Chile', 'Colombia',
					  'Ecuador', 'Guyana', 'Paraguay', 'Peru', 'Suriname', 'Uruguay', 'Venezuela')
AND order_date > '1997-09-01'
ORDER BY order_price DESC
LIMIT 3

### 4
SELECT product_name
FROM products
WHERE product_id = ANY (SELECT product_id FROM order_details WHERE quantity = 10) 
## HOMEWORK 5 DDL + INSERT INTO
CREATE TABLE teacher
(
	teacher_id serial,
	first_name varchar,
	last_name varchar,
	birthday date
);

ALTER TABLE teacher
ADD COLUMN middle_name varchar;

ALTER TABLE teacher
ADD COLUMN phone varchar;

ALTER TABLE teacher
DROP COLUMN middle_name;

ALTER TABLE teacher
RENAME birthday TO birth_day;

ALTER TABLE teacher
ALTER COLUMN phone SET DATA TYPE varchar(32);

CREATE TABLE exam
(
	exam_id serial,
	exam_name varchar(256),
	exam_date date
);

INSERT INTO exam (exam_name, exam_date)
VALUES
('Vasja1', '1985-02-17'),
('Vitja2', '1987-03-12'),
('Olja3', '1980-11-02');

SELECT *
FROM exam

TRUNCATE TABLE exam RESTART IDENTITY
## HOME WORK 6 Primary key
--1
CREATE TABLE exam
(
	exam_id serial UNIQUE NOT NULL,
	exam_name varchar,
	exam_date date
);
--2
ALTER TABLE exam
DROP CONSTRAINT exam_exam_id_key;

--3
ALTER TABLE exem
ADD PRIMARY KEY(exam_id);

--3.1
SELECT constraint_name
FROM information_schema.key_column_usage
WHERE table_name = 'exam'
AND column_name = 'exam_id';

--4
DROP TABLE person;
DROP TABLE passport;

CREATE TABLE person
(
	person_id int NOT NULL,
	first_name varchar(64) NOT NULL,
	last_name varchar(64) NOT NULL,
	
	CONSTRAINT pk_person_person_id PRIMARY KEY(person_id)
);
--5
CREATE TABLE passport
(
	passport_id int NOT NULL,
	isbn int NOT NULL,
	registration text NOT NULL,
	person_id int NOT NULL,
	
	CONSTRAINT pk_passport_passport_id PRIMARY KEY(passport_id),
	CONSTRAINT pk_passport_person_id FOREIGN KEY(person_id) REFERENCES person(person_id)
);

--6
SELECT *
FROM book;

ALTER TABLE book
ADD COLUMN weight decimal CONSTRAINT CHK_book_weight CHECK (weight > 0 AND weight < 100 );

--7
UPDATE book
SET weight = 101
WHERE book_id = 1;

INSERT INTO book (weight)
VALUES (1, 1, 'dertfd', 121, 99 );

--8
DROP TABLE student;

CREATE TABLE student
(
	identiti serial NOT NULL,
	fuul_name text NOT NULL,
	kurs decimal DEFAULT 1
);

--9
INSERT INTO student (fuul_name)
VALUES ('Edvard Ruki nozznicu'),
        ('Putin Xyilo');
		
INSERT INTO student (fuul_name, kurs)
VALUES ('Lavro kon', 2);

SELECT *
FROM student;

--10
ALTER TABLE student
ALTER COLUMN kurs DROP DEFAULT;
--11
SELECT * FROM products;

ALTER TABLE products
ADD CONSTRAINT chk_products_price CHECK(unit_price > 0);

--12
SELECT MAX(product_id) FROM products;

CREATE SEQUENCE IF NOT EXISTS products_product_id_seq
START WITH 78 OWNED BY products.product_id;

ALTER TABLE products
ALTER COLUMN product_id SET DEFAULT nextval('products_product_id_seq');
--13
INSERT INTO products (product_name, supplier_id, category_id, quantity_per_unit,
					 unit_price, units_in_stock, units_on_order, reorder_level, discontinued)
VALUES
('prod', 1, 1, 10, 20, 20, 10, 1, 0)
RETURNING product_id;
## HOME WORK 7 view
### 1
CREATE VIEW or_cu_em AS
SELECT order_date, required_date, shipped_date, ship_postal_code, 
        company_name, contact_name, phone, last_name, first_name,
		title
FROM orders
JOIN customers USING(customer_id)
JOIN employees USING(employee_id);

SELECT *
FROM or_cu_em
WHERE order_date > '1997-01-01'
ORDER BY order_date;

### 2
CREATE VIEW ord_cus_emp AS
SELECT order_date, required_date, shipped_date, ship_postal_code,
       ship_country, company_name, contact_name, phone, last_name,
	   first_name
FROM orders
JOIN customers USING(customer_id)
JOIN employees USING(employee_id);

CREATE OR REPLACE VIEW ord_cus_emp AS
SELECT postal_code
FROM customers

CREATE OR REPLACE VIEW ord_cus_emp AS
SELECT order_date, required_date, shipped_date, ship_postal_code,
       ship_country, company_name, contact_name, phone, last_name,
	   first_name, customers.postal_code, employees.reports_to
FROM orders
JOIN customers USING(customer_id)
JOIN employees USING(employee_id);

ALTER VIEW ord_cus_emp RENAME TO oce

SELECT *
FROM oce
ORDER BY ship_country

SELECT reports_to
FROM employees

DROP VIEW oce

### 3
CREATE VIEW product_continued AS
SELECT *
FROM products
WHERE discontinued = 0
WITH LOCAL CHECK OPTION;

SELECT *
FROM product_continued;

INSERT INTO product_continued
VALUES
(79, 'kaka', 50, 50, 25, 15, 140, 1, 5, 1);
## HOME WORK 8 CASE WHEN and COALESCE and CONCAT
### 1 CASE WHEN in ORDER BY
SELECT contact_name, country, city
FROM customers
ORDER BY contact_name,
                  (
		   CASE WHEN city IS NULL THEN country
		   ELSE city
		   END
	              );

SELECT contact_name, country, COALESCE(city, 'interes') AS city
FROM customers
ORDER BY contact_name,
       (
		   CASE WHEN city IS NULL THEN country
		   ELSE city
		   END
	   );
### 2 CASE WHEN in SELECT
SELECT product_name, unit_price, 
       CASE WHEN unit_price >= 100 THEN 'too expensive'
	        WHEN unit_price >= 50 AND unit_price <= 100 THEN 'average'
			ELSE 'low price'
			END AS curva_mat
FROM products
ORDER BY curva_mat;
### 3 COALESCE 
SELECT DISTINCT contact_name, COALESCE(order_id::text, 'orders') AS order_id
FROM customers
LEFT JOIN orders USING(customer_id)
WHERE order_date IS NULL

### 4 CONCAT обьеденение названий 
SELECT CONCAT(first_name,' ', last_name) AS FIO,
       COALESCE(NULLIF(title, 'Sales Representative'), 'Sales Stuff') AS title
FROM employees
## HOME WORK Func
### 1
```
CREATE OR REPLACE FUNCTION backup_customers() RETURNS void AS $$
DROP TABLE IF EXISTS b_up_customers;
SELECT *                                   -- (CREATE TABLE b_up_customers AS SELECT * FROM customers)
INTO b_up_customers
FROM customers
$$ LANGUAGE SQL;

SELECT backup_customers();
```
### 2
```
double precision = float8;

CREATE OR REPLACE FUNCTION avg_freight_from_orders() RETURNS float8 AS $$
SELECT avg(freight)
FROM orders
$$ LANGUAGE SQL;

SELECT avg_freight_from_orders();


SELECT *
FROM orders
```
### 3
```
CREATE OR REPLACE FUNCTION get_random(a1 real, b1 real) RETURNS float8 AS $$
DECLARE 
    ran float8 ;
BEGIN
    ran = RANDOM();
	RETURN FLOOR(((a1 - b1 + 1) * ran) + b1);   -- начудил можно без деклараций RETURN FLOOR((RANDOM()*(a1 - b1 + 1) ) + b1);
END;
$$ LANGUAGE plpgsql;

SELECT get_random(5, 2)
FROM generate_series(1, 5);
```
### 4
```
CREATE OR REPLACE FUNCTION salary_employee_city(employee_city varchar , OUT max_salary numeric , OUT min_salary numeric) AS $$
SELECT MAX(salary), MIN(salary)
FROM employee_salary
WHERE city = employee_city
$$ LANGUAGE SQL;											 
SELECT *
FROM salary_employee_city('London');
```
### 5
```
SELECT *
FROM employee_salary;

create or replace function read_salary_employees() returns void AS $$
UPDATE employee_salary
SET salary = salary * 1.15
WHERE salary < 6000
$$ LANGUAGE SQL

SELECT read_salary_employees();
```
### 6
```
DROP FUNCTION reading_salary_employees(numeric,numeric);

CREATE OR REPLACE FUNCTION reading_salary_employees(upper_boundary numeric DEFAULT 6000, correction_rate numeric DEFAULT 0.15)
                                                    RETURNS SETOF employee_salary AS $$
UPDATE employee_salary
SET salary = FLOOR(salary + (salary * correction_rate))
WHERE salary < upper_boundary
RETURNING *
$$ LANGUAGE SQL;

SELECT * FROM reading_salary_employees();
```
### 7
```
DROP FUNCTION reading_salary_employees(numeric,numeric);

CREATE OR REPLACE FUNCTION reading_salary_employees(upper_boundary numeric DEFAULT 6000, correction_rate numeric DEFAULT 0.15)

               RETURNS TABLE (last_name text, first_name  text, title text, salary numeric) AS $$
UPDATE employee_salary
SET salary = FLOOR(salary + (salary * correction_rate))
WHERE salary < upper_boundary
RETURNING last_name, first_name, title, salary;
$$ LANGUAGE SQL;

SELECT * FROM reading_salary_employees();
```
### 8
```
DROP FUNCTION get_avg_max_freight_orders_where_ship_wia(integer);

CREATE OR REPLACE FUNCTION get_avg_max_freight_orders_where_ship_wia(shipp_via int ) RETURNS SETOF orders AS $$
DECLARE
average real;
maximum real;
where_pox real;
BEGIN
 SELECT MAX(freight) INTO maximum FROM orders WHERE ship_via = shipp_via ;
 
 SELECT AVG(freight) INTO average FROM orders WHERE ship_via = shipp_via ;
 
 maximum = maximum * 0.7 ;
 
where_pox = (maximum + average) / 2 ;

RETURN QUERY
SELECT *
FROM orders
WHERE freight < where_pox ;
END
$$ LANGUAGE plpgsql;

SELECT * FROM get_avg_max_freight_orders_where_ship_wia(1);
```
### 9
```
CREATE OR REPLACE FUNCTION cur_salory(
	salary numeric,
	max_salary numeric DEFAULT 80,
	min_salary numeric DEFAULT 30,
	kof_rost numeric DEFAULT 0.2
) RETURNS bool AS $$                -- если тру или фолс то 'bool'
DECLARE
   new_salary numeric;
BEGIN
   IF salary >= max_salary OR salary >= min_salary THEN
      RETURN false;
   END IF;
   
   IF salary < min_salary THEN
      new_salary = salary + (salary * kof_rost);
   END IF;
   
   IF new_salary > max_salary THEN
      RETURN false;
   ELSE RETURN true;
   END IF;
END
$$ LANGUAGE plpgsql;

SELECT cur_salory(40, 80, 30, 0.2);
SELECT cur_salory(79, 81, 80, 0.2);
SELECT cur_salory(79, 95, 80, 0.2);
```
## HOME WORK Raise (ошибки)
```
create or replace function should_increase_salary(
	cur_salary numeric,
	max_salary numeric DEFAULT 80, 
	min_salary numeric DEFAULT 30, -- dont min_salary > max_salary
	increase_rate numeric DEFAULT 0.2
	) returns bool AS $$
declare
	new_salary numeric;
begin
   IF min_salary > max_salary THEN
     RAISE EXCEPTION 'Invalid salary. You passed:(%)', min_salary USING HINT = 'Allowed min_salary<max_salary', ERRCODE = 12887;
	END IF;
   IF min_salary < 0 OR max_salary < 0 THEN
     RAISE EXCEPTION 'Min and max salary >= 0. Min is %', min_salary USING HINT = 'zaebal';
	END IF;
   IF increase_rate < 0.05 THEN
     RAISE EXCEPTION 'Invalid salary. You passed;(%)', increase_rate USING HINT='Allowed increase_rate<5%', ERRCODE = 12889;
	END IF;
	if cur_salary >= max_salary or cur_salary >= min_salary then 		
		return false;
	end if;
	
	if cur_salary < min_salary then
		new_salary = cur_salary + (cur_salary * increase_rate);
	end if;
	
	if new_salary > max_salary then
		return false;
	else
		return true;
	end if;	
end;
$$ language plpgsql;

SELECT should_increase_salary(79,10,80,02);
SELECT should_increase_salary(79,10,-1,02);
SELECT should_increase_salary(79,10,10,0.04);
```
## HOME WORK композит тип
```
--1. Переписать функцию, которую мы разработали ранее в одном из ДЗ таким образом,
--чтобы функция возвращала экземпляр композитного типа. Вот та самая функция:
create or replace function get_salary_boundaries_by_city(
	emp_city varchar, out min_salary numeric, out max_salary numeric) 
AS 
$$
	SELECT MIN(salary) AS min_salary,
	   	   MAX(salary) AS max_salary
  	FROM employees
	WHERE city = emp_city
$$ language sql;
--1)
CREATE TYPE get_salary_boudaries AS (
	min_salary numeric,
	max_salary numeric
);
--
CREATE FUNCTION get_boundaries(emp_city varchar) RETURNS SETOF get_salary_boudaries AS $$
    SELECT MIN(salary) AS min_salary,
	   	   MAX(salary) AS max_salary
  	FROM employee_salary
	WHERE city = emp_city
$$ LANGUAGE SQL;
--
SELECT * FROM get_boundaries('London');
--
DROP TYPE get_salary_boudaries;

--2. Задание состоит из пунктов:
--2.1 Создать перечисление армейских званий США, включающее следующие значения: Private, Corporal, Sergeant
--2.2 Вывести все значения из перечисления.
--2.3 Добавить значение Major после Sergeant в перечисление
--2.4 Создать таблицу личного состава с колонками: person_id, first_name, last_name, person_rank (типа перечисления)
--2.5 Добавить несколько записей, вывести все записи из таблицы

--2.1
CREATE TYPE usa_range AS ENUM
('Private', 'Corporal', 'Sergeant');
--2.2
SELECT enum_range(NULL::usa_range); 
--2.3
ALTER TYPE usa_range
ADD VALUE 'Major' AFTER 'Sergeant';
--2.4
CREATE TABLE person_list 
(person_id serial PRIMARY KEY, first_name text, last_name text, person_rank usa_range);
--2.5
INSERT INTO person_list(first_name, last_name, person_rank)
VALUES
('Ivan', 'Hto', 'Private'),
('Sergiy', 'Kym', 'Major'),
('Muron', 'Drug', 'Sergeant');

SELECT *
FROM person_list
```
## HOME WORK Масив
```
--Создать функцию, которая вычисляет средний фрахт по заданным странам (функция принимает список стран).

CREATE OR REPLACE FUNCTION avg_freight_orders(VARIADIC country varchar[]) RETURNS float8 AS $$ 	
   	  SELECT  AVG(freight)
      FROM orders
      WHERE ship_country = ANY(country) ---попробуй догадатся
$$ LANGUAGE sql;

SELECT * FROM avg_freight_orders('Brazil', 'USA');
---или так
SELECT * FROM avg_freight_orders(VARIADIC ARRAY['Brazil', 'USA']);
## HOME WORK Масив 2
--Написать функцию, которая фильтрует телефонные номера по коду оператора.
--Принимает 3-х значный код мобильного оператора и список телефонных номеров в формате +1(234)5678901 (variadic)
--Функция возвращает только те номера, код оператора которых соответствует значению соответствующего аргумента.
DROP FUNCTION Phone_number([]);

CREATE FUNCTION Phone_number(oper int, VARIADIC numbers text[]) RETURNS SETOF text AS $$
DECLARE
     cur_val text;
BEGIN
    FOREACH cur_val IN ARRAY numbers
	LOOP
	    RAISE NOTICE 'cur_val is %', cur_val;
	    CONTINUE WHEN cur_val NOT LIKE CONCAT('__(', oper, ')%');
		RETURN NEXT cur_val;
	END LOOP;
END
$$ LANGUAGE plpgsql;

SELECT * FROM Phone_number(903, '+7(903)1901235', '+7(926)8567589', '+7(903)1532476');
SELECT * FROM Phone_number(903, VARIADIC ARRAY ['+7(903)1901235', '+7(926)8567589', '+7(903)1532476']);
-- +Проверить функцию передав следующие аргументы:
-- +903, +7(903)1901235, +7(926)8567589, +7(903)1532476
---Попробовать передать аргументы с созданием массива и без.
--Подсказка: чтобы передать массив в VARIADIC-аргумент, надо перед массивом прописать, собственно, ключевое слово variadic.
```
## HOME WORK GROUPING SET, ROLLUP, CUBE Подитоги и общие итоги
### 1 Вывести сумму продаж (цена * кол-во) по каждому сотруднику с подсчётом полного итога
----(полной суммы по всем сотрудникам) отсортировав по сумме продаж (по убыванию).
SELECT employee_id, SUM(unit_price * quantity)
FROM orders
LEFT JOIN order_details USING(order_id)
GROUP BY ROLLUP(employee_id)
ORDER BY SUM(unit_price * quantity);

### 2 Вывести отчёт показывающий сумму продаж по сотрудникам и странам отгрузки
----с подытогами по сотрудникам и общим итогом.
SELECT ship_country, employee_id, SUM(unit_price * quantity)
FROM orders
LEFT JOIN order_details USING(order_id)
GROUP BY ROLLUP(employee_id, ship_country)
ORDER BY employee_id, SUM(unit_price * quantity) DESC

### 3 Вывести отчёт показывающий сумму продаж по сотрудникам, странам отгрузки, сотрудникам
----и странам отгрузки с подытогами по сотрудникам и общим итогом.
SELECT ship_country, employee_id, SUM(unit_price * quantity)
FROM orders
LEFT JOIN order_details USING(order_id)
GROUP BY CUBE(ship_country, employee_id)
ORDER BY employee_id, SUM(unit_price * quantity) DESC
## HOME WORK Оконные функции и ранжирование
--Вывести отчёт показывающий по сотрудникам суммы продаж SUM(unit_price*quantity),
--и сопоставляющий их со средним значением суммы продаж по сотрудникам 
--(AVG по SUM(unit_price*quantity)) сортированный по сумме продаж по убыванию.

SELECT DISTINCT employee_id, sum_total, AVG(sum_total) OVER (), SUM(sum_total) OVER ()
FROM (SELECT DISTINCT employee_id, SUM(unit_price*quantity) OVER (PARTITION BY employee_id) AS sum_total
      FROM orders
      JOIN order_details USING(order_id)) AS total
ORDER BY sum_total DESC;

SELECT employee_id,  total_sum, AVG(total_sum) OVER (), SUM(total_sum) OVER ()
FROM (SELECT employee_id, SUM(unit_price * quantity) AS total_sum
      FROM order_details
      JOIN orders USING(order_id)
      GROUP BY employee_id) as d
ORDER BY total_sum DESC;

SELECT SUM(unit_price * quantity)
FROM order_details
JOIN orders USING(order_id)

--2--Вывести ранг сотрудников по их зарплате, без пропусков. 
-- Также вывести имя, фамилию и должность.

SELECT CONCAT(first_name,' ', last_name) as FIO, salary, title, 
           DENSE_RANK() OVER (ORDER BY salary DESC)
FROM employee_salary;
## ## HOME WORK Transactions
 --  В рамках транзакции с уровнем изоляции Repeatable Read выполнить следующие операции:
--  Заархивировать (SELECT INTO или CREATE TABLE AS) заказчиков(customers), которые сделали покупок менее чем на 2000 у.е.
-- удалить из таблицы заказчиков (customers) всех заказчиков (customers) , которые были предварительно заархивированы
(подсказка: для этого придётся удалить данные из связанных таблиц);

 SELECT * INTO backup_customers_correct FROM customers;

BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ;

DROP TABLE IF EXISTS backup_customers_less_2000;

CREATE TABLE backup_customers_company_less_2000 AS
                      SELECT customer_id, company_name, SUM(unit_price * quantity)
                      FROM orders
                      JOIN order_details USING(order_id)
					  JOIN customers USING(customer_id)
                      GROUP BY customer_id, company_name
                      HAVING SUM(unit_price * quantity) < 2000;

DELETE FROM order_details
WHERE order_id IN (SELECT order_id 
				   FROM orders 
				   WHERE customer_id IN (SELECT customer_id 
										 FROM backup_customers_company_less_2000 ));

DELETE FROM orders
WHERE customer_id IN (SELECT customer_id FROM backup_customers_company_less_2000);

DELETE FROM customers
WHERE customer_id IN (SELECT customer_id FROM backup_customers_company_less_2000);
COMMIT;
### 2
SELECT * FROM products
BEGIN;
DROP TABLE IF EXISTS backup_products;

SELECT * INTO backup_products FROM products;
-- заархивировать все продукты, снятые с продажи (см. колонку discontinued)
DROP TABLE IF EXISTS backup_products_discontinued;

SELECT * 
INTO backup_products_discontinued FROM products WHERE discontinued = 1;
--  поставить savepoint после архивации
SAVEPOINT backup;
-- удалить из таблицы продуктов все продукты, которые были заархивированы
SELECT * 
INTO backup_order_details_in_discontinued 
FROM order_details
WHERE product_id IN (SELECT product_id 
					FROM backup_products_discontinued);
					
DELETE FROM order_details WHERE order_id IN (SELECT order_id 
											  FROM backup_order_details_in_discontinued);
DELETE FROM products 
WHERE product_id IN (SELECT product_id FROM backup_products_discontinued);

-- откатиться к savepoint
ROLLBACK TO backup;

COMMIT;
### 3
вернуть удалённые данные 2 задание

BEGIN;

INSERT INTO products
SELECT * FROM backup_products_discontinued;

INSERT INTO order_details 
SELECT * FROM backup_order_details_in_discontinued;

COMMIT;

SELECT * 
FROM products
JOIN order_details USING(product_id)
WHERE discontinued = 0;
## HOME WORK Тригеры
### 1. Автоматизировать логирование времени последнего изменения в таблице products.
--   Добавить в products соответствующую колонку и реализовать построчный триггер.
```
ALTER TABLE products
ADD COLUMN last_updated timestamp;

CREATE OR REPLACE FUNCTION track_changes_products() RETURNS trigger AS $$
BEGIN
    NEW.last_updated = now();
	RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER products_timestamp BEFORE INSERT OR UPDATE ON products
FOR EACH ROW EXECUTE PROCEDURE track_changes_products();
--------cheking--------
UPDATE products
SET unit_price = 11
WHERE product_id = 3;

select * from products where product_id = 3
----------a new attempt-----
ALTER TABLE products
DROP COLUMN IF EXISTS last_updated;
DROP TRIGGER IF EXISTS products_timestamp ON products;
DROP FUNCTION IF EXISTS track_changes_products;
-----------clear history-----
DROP TABLE products_audit;
DROP TRIGGER IF EXISTS audit_products_insert ON products;
DROP TRIGGER IF EXISTS audit_products_update ON products;
DROP TRIGGER IF EXISTS audit_products_delete ON products;
DROP FUNCTION build_audit_products;
```
### 2. Автоматизировать аудит операций в таблице order_details.
```
--   Создайте отдельную таблицу для аудита, добавьте туда колонки для хранения наименования операций, 
--   имени пользователя и временного штампа. Реализуйте триггеры на утверждения.
DROP TABLE IF EXISTS order_details_audit;

CREATE TABLE order_details_audit (
	op char(2) NOT NULL,
	user_changed text NOT NULL,
	time_stamp timestamp NOT NULL,
	
    order_id smallint NOT NULL,
    product_id smallint NOT NULL,
    unit_price real NOT NULL,
    quantity smallint NOT NULL,
    discount real NOT NULL
);
--
CREATE OR REPLACE FUNCTION build_audit_order_details() RETURNS trigger AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN 
	    INSERT INTO order_details_audit 
	    SELECT 'I', session_user, now(), nt.* FROM new_table nt;
	ELSEIF TG_OP ='UPDATE'
	    THEN INSERT INTO order_details_audit
		SELECT 'U', session_user, now(), nt.* FROM new_table nt;
	ELSEIF TG_OP ='DELETE'
	    THEN INSERT INTO order_details_audit
		SELECT 'D', session_user, now(), ot.* FROM old_table ot;
	END IF;
	RETURN NULL;
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS audit_order_details_insert ON order_details;

CREATE TRIGGER audit_order_details_insert AFTER INSERT ON order_details
REFERENCING NEW TABLE AS new_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_order_details();

DROP TRIGGER IF EXISTS audit_order_details_update ON order_details;
CREATE TRIGGER audit_order_details_update AFTER UPDATE ON order_details
REFERENCING NEW TABLE AS new_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_order_details();

DROP TRIGGER IF EXISTS audit_order_details_delete ON order_details;
CREATE TRIGGER audit_order_details_delete AFTER DELETE ON order_details
REFERENCING OLD TABLE AS old_table
FOR EACH STATEMENT EXECUTE PROCEDURE build_audit_order_details();
--------------------------------------------------
--INSERT INTO orders 
--VALUES (11078, 'VINET', 5, '1996-07-04', '1996-08-01', '1996-07-16', 3, 32.3800011, 'Vins et alcools Chevalier', '59 rue de l''Abbaye', 'Reims', NULL, '51100', 'France');
--------------------------------------
INSERT INTO order_details
VALUES( 11078, 77, 5, 1, 0);

UPDATE order_details
SET discount = 1
where order_id = 11078;

DELETE FROM order_details
WHERE order_id = 11078;
--------------
select * from orders order by order_id desc;
select * from order_details_audit;
SELECT * FROM order_details ORDER BY order_id DESC;
---clear -----
DROP TRIGGER IF EXISTS audit_order_details_updete ON order_details;

DROP TRIGGER IF EXISTS audit_order_details_insert ON order_details;
DROP TRIGGER IF EXISTS audit_order_details_update ON order_details;
DROP TRIGGER IF EXISTS audit_order_details_delete ON order_details;
drop function build_audit_order_details;
DROP TABLE IF EXISTS order_details_audit;
```
#