package data

import (
	"testing"
)

func TestGeneratePost(t *testing.T) {
	user := createTestUser()
	post := Post{PostModel: PostModel{Body: "12342342342342234", UserID: user.ID}}
	err := post.GeneratePost(user)
	if err != nil {
		t.Errorf("\nerr: %s\n user: %s\n", err.Error(), user.Username)
	}

	post = Post{}
	err = post.GeneratePost(user)
	if err == nil {
		t.Errorf("post should not success, post is nil")
	}

	post = Post{PostModel: PostModel{Body: "12342342342342234"}}
	err = post.GeneratePost(User{})

	if err == nil {
		t.Errorf("post should not success, userid is nil")
	}

	post = Post{PostModel: PostModel{UserID: user.ID}}
	err = post.GeneratePost(user)
	if err == nil {
		t.Errorf("body is nil, should not success")
	}
}

func TestUpdatePost(t *testing.T) {
	user := createTestUser()
	Db.Delete(&Post{})
	post := CreateTestPost(user)
	post.GeneratePost(user)
	post.Body = "hello world"
	post.UpdatePost()

	post1, err := FindPostByID(post.ID)
	if err != nil || post1.Body != post.Body {
		t.Errorf("find post by id error")
	}

	post.Body = ""
	err = post.UpdatePost()
	if err == nil {
		t.Errorf("should not update null post")
	}
}

func TestFindUserAllPost(t *testing.T) {
	Db.Delete(&Post{})
	user := createTestUser()
	post := Post{}
	post.Body = "1111111"
	post.GeneratePost(user)
	post.GeneratePost(user)
	post.GeneratePost(user)
	post.GeneratePost(user)
	post.GeneratePost(user)
	posts, err := FindUserAllPost(user)
	if err != nil || len(posts) != 5 {
		t.Errorf("find user all posts error")
	}
}

func TestDeletePost(t *testing.T) {
	Db.Delete(&Post{})
	user := createTestUser()
	post := CreateTestPost(user)
	post.GeneratePost(user)

	err := post.DeletePost()
	if err != nil {
		t.Error("delete post fail")
	}

	post = Post{}
	err = post.DeletePost()
	if err == nil {
		t.Error("should not delete null post")
	}
}

func TestFindPostByID(t *testing.T) {
	user := createTestUser()
	post := CreateTestPost(user)
	post.GeneratePost(user)

	post1, err := FindPostByID(post.ID)
	if err != nil || post1.ID != post.ID {
		t.Error("can not find post by id")
	}

	// post2 := Post{}
	post2, _ := FindPostByID("")
	if len(post2.ID) > 0 {
		t.Error("should not find null post")
	}

	post3, _ := FindPostByID("123")
	if len(post3.ID) > 0 {
		t.Error("should not find null post")
	}

}
