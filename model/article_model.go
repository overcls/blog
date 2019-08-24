package model

type ArticleModel struct {
	Title string `json:"title"`
	Content string `json:"content"`
	ViewCount int `json:"view_count"`
	CategoryId int `json:"category_id"`

}
