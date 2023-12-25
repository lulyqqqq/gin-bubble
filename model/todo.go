package model

type Todo struct {
	Id      int    `gorm:"type:int" json:"id"`
	UserId  int    `gorm:"type:int" json:"userId"`
	Content string `gorm:"type:varchar(512)" json:"content"`
	Status  int    `gorm:"type:int" json:"status"`
	Time    string `gorm:"type:varchar(24)" json:"time"`
}
