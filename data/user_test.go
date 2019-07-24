package data

import (
	"testing"
)

func TestCreateUserDB(t *testing.T) {
	db := Db.Delete(&User{})
	if db.Error != nil {
		t.Errorf(db.Error.Error())
	}
	userLg := UserLogin{Username: "lemon", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "lemon@qq.com", UserLogin: userLg}
	user := User{UserRegister: userRg}
	err := user.CreateUserDB()

	if err != nil {
		t.Errorf(err.Error())
	}

	err = user.CreateUserDB()
	if err == nil {
		t.Errorf("Db create exist user")
	}

	userLg1 := UserLogin{Username: "lemon1"}
	userRg1 := UserRegister{Email: "lemon1@qq.com", UserLogin: userLg1}
	user1 := User{UserRegister: userRg1}
	err = user1.CreateUserDB()
	if err == nil {
		t.Errorf("pasword should not null")
	}
}

func TestFindUser(t *testing.T) {
	Db.Delete(&User{})
	userLg := UserLogin{Username: "lemon", PasswordHash: "lemon"}
	_, err := FindUser(userLg.Username)
	if err == nil {
		t.Error("should not found user")
	}

	userRg := UserRegister{Email: "lemon@qq.com", UserLogin: userLg}
	user := User{UserRegister: userRg}
	user.CreateUserDB()

	_, err = FindUser(user.Username)
	if err != nil {
		t.Errorf("find user error, %s", err.Error())
	}
}

func TestGeneratePasswordHash(t *testing.T) {
	user := CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	if len(user.PasswordHash) != 32 {
		t.Errorf("generatePassword fail, should length 32, password:%s", user.PasswordHash)
	}
}

func TestValidateUser(t *testing.T) {
	user := createUser()
	user1 := UserLogin{Username: user.Username, PasswordHash: "123456"}

	if !user1.ValidateUser() {
		t.Errorf("user should validate true")
	}

	user2 := UserLogin{Username: user.Username, PasswordHash: "12345"}
	if user2.ValidateUser() {
		t.Errorf("user should validate fail")
	}

	user3 := UserLogin{Username: "lemon", PasswordHash: "123456"}
	if user3.ValidateUser() {
		t.Errorf("user should validate fail")
	}
}

func TestGetUserInfo(t *testing.T) {
	user := createUser()
	user1 := User{UserRegister: UserRegister{UserLogin: UserLogin{Username: "lemonlee", PasswordHash: "123456"}}}
	user1.GeneratePasswordHash()
	user1.GetUserInfo()
	if user1.Email != user.Email || user1.RoleID != user.RoleID {
		t.Errorf("get userinfo fail")
	}

}

func createUser() User {
	Db.Delete(&User{})
	user := CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	user.CreateUserDB()
	return user
}
