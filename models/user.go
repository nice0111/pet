package models

type User struct {
	Id        int `gorm:"primarykey:autoIncrement"`
	Username  string
	Password  string
	Phone     string
	Email     string
	Level     string
	Gender    string
	Logintime int64
	Loginip   string
	City      string
	Token     string
}

// 默认情况表名是结构体名称复数形式，结构体名字默认就是操作与结构体同名的表
// 使用结构体中自定义方法TableName 改变结构体的默认表名称
func (User) TableName() string {
	return "User"
}
