package modelos

type Data struct {
	Message map[string]string `json:"message"`
}

type Resultado struct {
	Page  int         `json:"page"`
	Prev  bool        `json:"prev"`
	Next  bool        `json:"next"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}
