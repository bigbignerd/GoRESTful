package model

import (
    "fmt"
    validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
    BaseModel
    Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
    Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

//create new user
func (u *UserModel) Create() error {
    return DB.Self.Create(&u).Error
}

//delete user
func DeleteUser(id uint64) error {
    user := UserModel{}
    user.BaseModel.Id = id
    return DB.Self.Delete(&user).Error
}
