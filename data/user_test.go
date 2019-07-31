package data

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	fmt.Println(Db)
	Db.Delete(&User{})
	Db.Delete(&Role{})
	InsertRoles()
	fmt.Println("----------test begin------------")
}

func shutdown() {
	fmt.Println("----------test end------------")
}

// func TestCreateFriendShip(t *testing.T) {
// 	friend := Friendships{}

// 	friend.FollowUser(nil, nil)
// }

func TestFollowUserDB(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()

	if err := user.FollowedUser(user2); err != nil {
		t.Error(err.Error())
	}

	if user.FollowedCount() != 1 {
		t.Error("count error")
	}

	user.FollowedUser(user2)
	if user.FollowedCount() != 1 {
		t.Error("count error")
	}

}

func TestUnFollowedUser(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()

	user.FollowedUser(user2)
	if err := user.UnFollowUser(user2); err != nil {
		t.Error(err)
	}
	if user.FollowedCount() != 0 {
		t.Error("count error")
	}
	// if err := user.UnFollowUser(user2); err != nil {
	// 	t.Error(err)
	// }
}

func TestGetFollowedCount(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()
	user.FollowedUser(user2)

	if user.FollowedCount() != 1 {
		t.Error("count error")
	}

	if user2.FollowedCount() != 0 {
		t.Error("count error")
	}

}

func TestGetFollowersCount(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()
	user.FollowedUser(user2)

	if user.FollowerCount() != 0 {
		t.Error("count error")
	}

	if user2.FollowerCount() != 1 {
		t.Error("count error")
	}
}

func TestGetFollowersUser(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()
	user.FollowedUser(user2)

	users, err := user2.GetFollowers()
	if err != nil {
		t.Error("get followers error")
	}

	if len(users) != 1 {
		t.Error("get followers count error")
	}

	if users[0].ID != user.ID {
		t.Error("get follower user not match")
	}
}

func TestGetFollowedUser(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Follows{})
	userLg := UserLogin{Username: "soojin206", PasswordHash: "lemon"}
	userRg := UserRegister{Email: "soojin@qq.com", UserLogin: userLg}
	user2 := User{UserRegister: userRg}
	user2.CreateUserDB()
	user.FollowedUser(user2)

	users, err := user.GetFollowed()
	if err != nil {
		t.Error("get followers error")
	}

	if len(users) != 1 {
		t.Error("get followers count error")
	}

	if users[0].ID != user2.ID {
		t.Error("get follower user not match")
	}
}

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
		t.Errorf("Db create exist user, %s", err.Error())
	}

	userLg1 := UserLogin{Username: "lemon1"}
	userRg1 := UserRegister{Email: "lemon1@qq.com", UserLogin: userLg1}
	user3 := User{UserRegister: userRg1}
	err = user3.CreateUserDB()
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
	user := createTestUser()
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
	user := createTestUser()
	user1 := User{UserRegister: UserRegister{UserLogin: UserLogin{Username: "lemonlee", PasswordHash: "123456"}}}
	user1.GeneratePasswordHash()
	user1.GetUserInfo()

	if user1.Email != user.Email {
		t.Errorf("get userinfo fail,\n user1: %s, user2:%s", user1.RoleID, user.RoleID)
	}

}

func createTestUser() User {
	Db.Delete(&User{})
	user := CreateUser("lemon@qq.com", "123456", "lemonlee")
	user.GeneratePasswordHash()
	err := user.CreateUserDB()
	if err != nil {
		fmt.Println(err)
	}
	return user
}
