package model

import (
	"strconv"
)

type Model interface {
	ToString() string
}

var UserDatas map[string]Model

type User struct {
	username string
	password string
	age      int
	sex      string
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}
func (u *User) SetAge(age int) {
	u.age = age
}
func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetAge() int {
	return u.age
}
func (u *User) GetSex() string {
	return u.sex
}
func GetUser(username string) *User {
	if UserDatas == nil {
		return nil
	}
	return UserDatas[username].(*User)
}

func (u *User) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}
