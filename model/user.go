package model

type User struct {
	Id       int    `json:"id" gorm:"column:Id;auto_increment;primary_key"` //用户编号
	Name     string `json:"name" gorm:"column:name"`
	Status   int    `json:"status" gorm:"column:status"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
}

func (user *User) TableName() string {
	return "user"
}
