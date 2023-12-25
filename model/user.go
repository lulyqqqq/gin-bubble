package model

type User struct {
	Id       int    `gorm:"type:int" json:"id"`
	Name     string `gorm:"varchar(11);not null;unique" json:"name"`
	Password string `gorm:"varchar(12);not null" json:"password"`
	Number   string `gorm:"varchar(11);not null;unique" json:"number"`
	Address  string `gorm:"varchar(256)" json:"address"`
	Tag      string `gorm:"varchar(5)" json:"tag"`
	Role     string `gorm:"varchar(2)" json:"role"` // 0-管理员 1-正常用户 2-禁止使用的用户
}
