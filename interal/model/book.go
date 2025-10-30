package model

type Libro struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Autor string `json:"autor"`
}