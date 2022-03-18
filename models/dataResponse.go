package models

type Data struct {
	Success bool              `json:"success"`
	Data    interface{}       `json:"data"`
	Message map[string]string `json:"message"`
}
