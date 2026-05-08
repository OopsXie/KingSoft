package model

type Question struct {
	ID         uint    `json:"id"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
	DeletedAt  *string `json:"deletedAt,omitempty"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	Options    string  `json:"options"`
	Answer     string  `json:"answer"`
	Difficulty string  `json:"difficulty"`
}
