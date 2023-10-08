package view

type RequestUpdate struct {
	Id              int64  `json:"id" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Author          string `json:"author" binding:"required"`
	YearPublication int    `json:"year_publication" binding:"required"`
	Summary         string `json:"summary" binding:"required"`
}
