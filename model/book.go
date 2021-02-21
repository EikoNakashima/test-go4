package model

type Book struct {
	Id      int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Title   string `xorm:"varchar(40)" json:"title" form:"title"`
	Content string `xorm:"varchar(40)" json:"content" form:"content"`
}
