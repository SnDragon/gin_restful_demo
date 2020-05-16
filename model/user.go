package model

import (
	"time"
)

type User struct {
	// 这里想忽略请求的参数绑定,所以不直接使用gorm.Model
	// TODO 忽略请求参数绑定
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name string `gorm:"type:varchar(255);column:name;default:''" binding:"required" json:"name"`
	Age int	`gorm:"column:age;default:0" binding:"required" json:"age"`
}