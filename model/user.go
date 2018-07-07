package model

import (
    "github.com/bigbignerd/GoRESTful/pkg/auth"
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

//update user
func (u *UserModel) Update() error {
    return DB.Self.Save(u).Error
}


func (u *UserModel) Encrypt() (err error) {
    u.Password, err = auth.Encrypt(u.Password)
    return
}

func (u *UserModel) Validate() error {
    validate := validator.New()
    return validate.Struct(u)
}
