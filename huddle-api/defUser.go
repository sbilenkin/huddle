package main

import (
	"gorm.io/gorm"
)

func findOrCreateUser(db *gorm.DB, userInfo UserInfo) (UserInfo, error) {
	var user UserInfo
	result := db.FirstOrCreate(&user, UserInfo{Email: userInfo.Email}, userInfo)
}