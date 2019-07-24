package data

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
)

//User Model
type User struct {
	Model
	UserRegister
	UserInfo
}

//UserLogin 用户登录信息
type UserLogin struct {
	Username     string `gorm:"type:varchar(64);unique;not null" json:"username" binding:"required,alphanum,min=6,max=12"`
	PasswordHash string `gorm:"not null;type:varchar(128)" json:"password" binding:"required"`
}

//UserRegister 用户注册信息
type UserRegister struct {
	UserLogin
	Email string `gorm:"unique;not null;type:varchar(64)" json:"email" binding:"omitempty,email"`
}

//UserInfo 用户数据信息
type UserInfo struct {
	RoleID     string
	Confirmed  bool   `gorm:"default:false"`
	Name       string `gorm:"type:varchar(64)"`
	Location   string `gorm:"type:varchar(64)"`
	AboutMe    string
	AvatarHash string `gorm:"type:varchar(64)"`
}

const salt string = "kvdevsalt"

//CreateUserDB 创建用户
func (userRG *UserRegister) CreateUserDB() error {
	Db.CreateTable(&User{})
	if len(userRG.PasswordHash) < 1 {
		return errors.New("passwordHash length error")
	}
	role := Role{}
	Db.Where("name = ?", "User").First(&role)
	user := User{UserRegister: *userRG}
	user.RoleID = role.ID
	db := Db.Create(&user)

	return db.Error
}

//FindUser 通过username查询用户
func FindUser(username string) (User, error) {
	user := User{}
	db := Db.Where("username = ?", username).First(&user)
	return user, db.Error
}

//CreateUser 创建用户封装方法
func CreateUser(email, password, username string) User {
	return User{UserRegister: UserRegister{UserLogin: UserLogin{Username: username, PasswordHash: password},
		Email: email}}
}

//GeneratePasswordHash 创建MD5+salt
func (user *UserLogin) GeneratePasswordHash() {
	h := md5.New()
	io.WriteString(h, user.PasswordHash+salt)
	user.PasswordHash = hex.EncodeToString(h.Sum(nil))
}

//ValidateUser 验证用户名和密码
func (user *UserLogin) ValidateUser() bool {
	user.GeneratePasswordHash()
	user1 := User{}
	Db.Where("username = ?", user.Username).First(&user1)
	// fmt.Printf("user1: %s user: %s\n", user1.PasswordHash, user.PasswordHash)
	return user.PasswordHash == user1.PasswordHash
}

//GetUserInfo 获取用户信息
func (user *User) GetUserInfo() {
	Db.Where("username = ?", user.Username).First(&user)
}

//UpdateUserInfo 更新用户信息
func (user *User) UpdateUserInfo() {
	Db.Save(&user)
}
