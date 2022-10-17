package basic

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	Id   int
	Name string
}

type UserModel struct {
	BaseModel
}

func (m *UserModel) Model() *gorm.DB {
	return m.DB().Table("user")
}

func (m *UserModel) GetByCondition() (user *User) {
	err := m.Model().Find(&user).Error
	fmt.Println(err)
	return
}
