package entity

type Index struct {
	Id   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}
