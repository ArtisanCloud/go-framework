package models

import (
	"database/sql"
	"github.com/ArtisanCloud/go-framework/config"
	"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (*User) TableName() string {
	return config.DatabaseConn.Schemas["default"] + "." + "users"
}

type User struct {
	Id                    int    `gorm:"json:id column:id"`
	Password              string `gorm:"column:password json:password"`
	Locale                string `gorm:"column:locale json:locale"`
	UUID                  string `gorm:"column:uuid" json:"uuid" primaryKey;autoIncrement:"false" sql:"index"`
}

// 模块初始化函数 import 包时被调用
func init() {
	//fmt.Println("user model module init function")
}

func NewUser() *User {
	return &User{}
}

/**
 * Scope Where Conditions
 */
func (user *User) WhereUserName(uuidOrPhone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid=@value OR mobile=@value", sql.Named("value", uuidOrPhone))
	}
}

func (user *User) WhereIsValid(db *gorm.DB) *gorm.DB {
	return db.Where("account_uuid != null")
}
