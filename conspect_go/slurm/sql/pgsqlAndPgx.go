package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

//**очень скупой вариант от gpt**
// в slurm много больше функций

func pgsqlAndPgx() {
	// Установите соединение с базой данных
	conn, err := pgx.Connect(context.Background(), "your_connection_string")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}
	defer conn.Close(context.Background())

	// Выполните запрос
	rows, err := conn.Query(context.Background(), "SELECT * FROM user")
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer rows.Close()

	// Обработайте результаты запроса
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Ошибка при сканировании результата:", err)
			return
		}
		fmt.Println("ID:", id, "Name:", name)
	}

	if rows.Err() != nil {
		fmt.Println("Ошибка при получении результатов:", rows.Err())
		return
	}
}
