package models

type Articles struct {
	TableId
	Name     string `json:"name"`
	CateName string `json:"cate_name"`
	Content  string `json:"content"`
	TableTime
}
