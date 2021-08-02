package models

type Review struct {
	Title  string `json:"title" binding:"required"`
	Review string `json:"review,omitempty"`
	Score  uint16 `json:"score,omitempty"`
}

type Movie struct {
	Title string `json: "title" binding:"required"`
	Year  int    `json: "year,omitempty"`
}
