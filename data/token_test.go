package data

import (
	"testing"
)

func TestGenerateValidateToken(t *testing.T) {
	rf := RefreshToken{}
	userLg := UserLogin{Username: "lemon", PasswordHash: "123456"}
	rf.GenerateToken(userLg)
	ac := AccessToken{}
	ac.GenerateToken(userLg)
	if len(ac.AccessTokenStr) == 0 || len(rf.RefreshTokenStr) == 0 {
		t.Errorf("token length should be > 0")
	}
	if ac.AccessTokenStr == rf.RefreshTokenStr {
		t.Error("token should not equal")
	}
	if ac.ValidToken() != userLg.Username || rf.ValidToken() != userLg.Username {
		t.Errorf("token error")
	}
}
