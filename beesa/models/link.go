package models

// Link models
type Link struct {
	UserID       string `json:"user_id"`
	OriginalLink string `json:"original_link"`
	ShortLink    string `json:"short_link"`
}
