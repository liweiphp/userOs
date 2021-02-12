package model

import "strconv"

type Model interface {
	ToString() string
}

var Datas map[string]Model

type User struct {
	username string
	password string
	age      int
	sex      string
}

func NewUser() *User {
	return &User{}
}

func (u User) SetUsername(username string)  {
	u.username = username
}
func (u User) SetPassword(password string)  {
	u.password  = password
}
func (u User) SetAge(age int)  {
	u.age = age
}
func (u User) SetSex(sex string)  {
	u.sex = sex
}

func GetUsername(u User) string {
	return u.username
}
func GetPassword(u User) string {
	return u.password
}
func GetAge(u User) int {
	return u.age
}
func GetSex(u User) string {
	return u.sex
}

func (u User)ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

