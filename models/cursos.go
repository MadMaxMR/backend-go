package models

type Cursos struct {
	ID          uint   `json:"id"`
	Nombre      string `json:"nombre" `
	Description string `json:"description" `
	Image       string `json:"image" `
}
