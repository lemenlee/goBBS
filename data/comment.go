package data

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Comment struct {
	Model
	CommentModel
}

type CommentModel struct {
	UserID string
	PostID string
	Body   string `gorm:"type:text;not null" validate:"required,min=6"`
}

func (comment *CommentModel) validateComment() error {
	validate = validator.New()
	err := validate.Struct(comment)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Printf("error: %s", err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("\nerr type: %s\n", err.StructField()) // by passing alt name to ReportError like below
		}
		return err
	}
	return nil
}

func (comment *Comment) CreateComment(user User, post Post) error {
	Db.CreateTable(&Comment{})
	comment.PostID = post.ID
	comment.UserID = user.ID
	err := comment.validateComment()
	if err != nil {
		return err
	}
	db := Db.Create(&comment)
	return db.Error
}

func (comment *Comment) UpdateComment() error {
	err := comment.validateComment()
	if err != nil {
		return err
	}
	db := Db.Save(&comment)
	return db.Error
}

func (comment *Comment) DeleteComment() error {
	err := comment.validateComment()
	if err != nil {
		return err
	}
	db := Db.Delete(&comment)
	return db.Error
}

func GetCommentByID(commentID string) (Comment, error) {
	comment := Comment{}
	db := Db.Where("id = ?", commentID).First(&comment)
	return comment, db.Error
}

func GetAllPostComment(post Post) ([]Comment, error) {
	comments := []Comment{}
	db := Db.Model(&post).Related(&comments)
	return comments, db.Error
}

func GetAllUserComment(user User) ([]Comment, error) {
	comments := []Comment{}
	db := Db.Model(&user).Related(&comments)
	return comments, db.Error
}
