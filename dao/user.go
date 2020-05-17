package dao

import (
	"fmt"
	"github.com/SnDragon/gin_restful_demo/model"
	"github.com/SnDragon/gin_restful_demo/utils"

)

func CreateUser(user *model.User) {
	db := NewDB()
	db.Create(user)
}

func GetUserById(id int) (user *model.User) {
	if userId, err := utils.GetInt(id); err != nil {
		fmt.Println(err)
		return nil
	}else {
		db := NewDB()
		user = &model.User{}
		db.First(user, userId)
		if user.ID == 0 {
			return nil
		}
		return
	}
}

func UpdateUserById(id int, user *model.User) *model.User {
	dbUser := GetUserById(id)
	if dbUser == nil {
		return nil
	}
	db := NewDB()
	if err := db.Model(dbUser).Updates(*user).Error;err != nil {
		fmt.Println(err)
		return nil
	}
	return dbUser
}

func DeleteUserById(id interface{}) bool {
	if userId, err := utils.GetInt(id); err != nil {
		fmt.Println(err)
		return false
	} else {
		user := model.User{ID: uint(userId)}
		db := NewDB()
		if err := db.Delete(&user).Error; err != nil {
			fmt.Println(err)
			return false
		} else {
			return true
		}
	}
}

func GetUsers(page, limit int) (total int, users []model.User) {
	// 获取总数
	db := NewDB()
	//users = make([]model.User, 10)
	db.Model(&model.User{}).Count(&total)
	// 分页
	if total > 0 {
		offset := (page - 1) * limit
		db.Offset(offset).Limit(limit).Find(&users)
	}
	return
}
