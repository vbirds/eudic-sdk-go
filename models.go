package dict

// Category represents a study list category
type Category struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Name     string `json:"name"`
}

// Word represents a word entry in study list
type Word struct {
	Word string `json:"word"`
	Exp  string `json:"exp"`
}

// Response represents the API response structure
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// CategoryResponse represents category list response
type CategoryResponse struct {
	Data    []Category `json:"data"`
	Message string     `json:"message"`
}

// WordResponse represents word list response
type WordResponse struct {
	Data    []Word `json:"word"`
	Message string `json:"message"`
}
