package models

type Words struct {
	ID      int    `json:"id"`
	English string `json:"english"`
	Russian string `json:"russian"`
	Theme   string `json:"theme"`
}
