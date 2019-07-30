package data

import "testing"

func TestCreateComment(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	err := comment.CreateComment(user, post)
	if err != nil {
		t.Errorf("create comment fail")
	}
	comment.Body = "he"
	err = comment.CreateComment(user, post)
	if err == nil {
		t.Error("comment should not create, len must be > 6")
	}
}

func TestUpdateComment(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	comment.CreateComment(user, post)

	comment.Body = "hello world"
	err := comment.UpdateComment()
	if err != nil {
		t.Error("comment cannot update")
	}

	comment.Body = "he"
	err = comment.UpdateComment()
	if err == nil {
		t.Error("comment should not update")
	}
}

func TestGetCommentByID(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	comment.CreateComment(user, post)

	comment1, err := GetCommentByID(comment.ID)
	if err != nil {
		t.Error("get comment error")
	}

	if comment1.ID != comment.ID {
		t.Error("get comment not equal comment")
	}
}

func TestGetAllPostComment(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	for i := 0; i < 10; i++ {
		comment.CreateComment(user, post)
	}

	comments, err := GetAllPostComment(post)
	if err != nil {
		t.Error("get all post comment error")
	}
	if len(comments) != 10 {
		t.Error("get all post comment nums not equal 10")
	}

}

func TestGetAllUserComment(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	for i := 0; i < 10; i++ {
		comment.CreateComment(user, post)
	}

	comments, err := GetAllUserComment(user)
	if err != nil {
		t.Error("get all user comment error")
	}
	if len(comments) != 10 {
		t.Error("get all user comment nums not equal 10")
	}
}

func TestDeleteComment(t *testing.T) {
	Db.Delete(&Comment{})
	user := createTestUser()
	post := CreateTestPost(user)
	comment := CreateTestComment()
	comment.CreateComment(user, post)

	err := comment.DeleteComment()
	if err != nil {
		t.Error("delete comment error")
	}

	comment1, _ := GetCommentByID(comment.ID)
	if comment1.ID == comment.ID {
		t.Error("delete comment error")
	}

}
