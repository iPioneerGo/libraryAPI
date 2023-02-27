package entity

import (
	"encoding/json"
	"strconv"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name   string `json:"name"`
		Author string `json:"author"`
		Year   string `json:"year"`
	}{
		Name:   b.Name,
		Author: b.Author,
		Year:   strconv.Itoa(b.Year),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Name   string `json:"name"`
		Author string `json:"author"`
		Year   string `json:"year"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	b.Name = aux.Name
	b.Author = aux.Author
	b.Year, _ = strconv.Atoi(aux.Year)

	return nil
}
