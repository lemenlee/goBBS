package data

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

//Model 数据库基础数据类
type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time
}

//Db Create
var Db *gorm.DB

//BeforeCreate 写入之前初始化UUID
func (base *Model) BeforeCreate(scope *gorm.Scope) error {
	fmt.Println("------------before insert-------------")
	uuid := uuid.NewV4().String()
	return scope.SetColumn("ID", uuid)
}

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:soojin206@tcp(db:3306)/bbs_test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("open data base")
	if err != nil {
		fmt.Printf("data base erorr: %s", err.Error())
	}

	Db.AutoMigrate(&User{}, &Follows{}, &Role{}, &Post{}, &Comment{})
	Db.LogMode(true)
	return
}
