package model

type MetaData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
	Script      string `json:"script"`
	Primiera    string `json:"primiera"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
}
