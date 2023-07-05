package store

import (
	"fmt"
	"log"

	"github.com/Wolfxxxz/Dictionary/internal/app/models"
)

type WordRepository struct {
	store *Store
}

var (
	tableWord string = "words"
)

// For Post request
func (ar *WordRepository) Create(a *models.Words) (*models.Words, error) {
	//Вставить данные в таблицу
	query := fmt.Sprintf("INSERT INTO %s (english, russian, theme) VALUES ($1, $2, $3) RETURNING id", tableWord)
	if err := ar.store.db.QueryRow(query, a.English, a.Russian, a.Theme).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// For DELETE request
func (ar *WordRepository) DeleteById(id int) (*models.Words, error) {
	article, ok, err := ar.FindWordById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tableWord)
		_, err = ar.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return article, nil
}

// Helper for Delete by id and GET by id request
func (ar *WordRepository) FindWordById(id int) (*models.Words, bool, error) {
	articles, err := ar.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var articleFinded *models.Words
	for _, a := range articles {
		if a.ID == id {
			articleFinded = a
			founded = true
		}
	}

	return articleFinded, founded, nil

}

// Get all request and helper for FindByID
func (ar *WordRepository) SelectAll() ([]*models.Words, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableWord)
	rows, err := ar.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := make([]*models.Words, 0)
	for rows.Next() {
		a := models.Words{}
		err := rows.Scan(&a.ID, &a.English, &a.Russian, &a.Theme)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}
