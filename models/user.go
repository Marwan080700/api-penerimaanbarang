package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey:autoIncrement"`
	UserName string `json:"username" form:"username" gorm:"column:username;type:varchar(255)"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" form:"role" gorm:"type:varchar(255);default:'user'"`
	Status   string `json:"status" form:"status" gorm:"type:varchar(255);default:'inactive'"`
}
