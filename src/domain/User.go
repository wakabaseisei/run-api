package domain

import "time"

// ユーザー
type User struct {
	Id           int
	Sex          int
	Introduction string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 返すユーザー情報
type UserGet struct {
	Id           int    `json:"id"`
	Sex          int    `json:"sex"`
	Introduction string `json:"introduction"`
}

func (u *User) BuildForGet() UserGet {
	user := UserGet{}
	user.Id = u.Id
	user.Sex = u.Sex
	user.Introduction = u.Introduction
	return user
}
