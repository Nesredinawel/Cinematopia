package models

type Director struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
