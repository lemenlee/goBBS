package data

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

//Post 文章模型
type Post struct {
	Model
	PostModel
}

//PostModel 具体内容
type PostModel struct {
	Body   string `gorm:"type:text;not null" json:"body" validate:"required,min=6" binding:"required,min=6"`
	UserID string `gorm:"not null" json:"user_id" validate:"required,min=6"`
}

var validate *validator.Validate

func (post *PostModel) validatePost() error {
	validate = validator.New()
	err := validate.Struct(post)
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

//GeneratePost 生成Post
func (post *Post) GeneratePost(user User) error {
	// post := Post{}
	Db.CreateTable(&Post{})
	post.UserID = user.ID
	err := post.validatePost()
	if err != nil {
		return err
	}
	// post.Body = postModel.Body
	db := Db.Create(&post)
	return db.Error
}

//UpdatePost 更新Post
func (post *Post) UpdatePost() error {
	if err := post.validatePost(); err != nil {
		return err
	}
	db := Db.Save(&post)
	return db.Error
}

//DeletePost delete post by id
func (post *Post) DeletePost() error {
	err := post.validatePost()
	if err != nil {
		return err
	}
	db := Db.Delete(&post)
	return db.Error
}

func FindUserAllPost(user User) ([]Post, error) {
	posts := []Post{}
	err := Db.Model(&user).Related(&posts).Error
	return posts, err
}

func FindUserPostCount(user User) (int, error) {
	count := 0
	db := Db.Model(&Post{}).Where("user_id = ?", user.ID).Count(&count)
	return count, db.Error
}

//FindPostByID 通过post ID 查找
func FindPostByID(id string) (Post, error) {
	post := Post{}
	db := Db.Where("id = ?", id).First(&post)
	return post, db.Error
}

func CreateTestPost(user User) Post {
	post := Post{}
	post.Body = "12342342342342234"
	post.UserID = user.ID
	return post
}
