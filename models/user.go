package models

import (
	"time"
)


type User struct {
	Model
	Username     string    `json:"username"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Avater      string    `json:"avater"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	LoginTime   time.Time `json:"login_time"`
	Type        int       `json:"type"`
}

func FindUserByID(id int)(user []*User,err error){
	if err= db.Where("id = ?", id).Find(&user).Error;err != nil{
		return nil,err
	}
	db.Where("id = ?", id).Model(&user).Update("login_time",time.Time(time.Now()))
	return
}

func CheckUser(account, password string) (user User,err error) {
	err = db.Where(User{Username : account, Password : password}).First(&user).Error
	return
}

func UpdateUser(data User)(err error){
	if err= db.Model(&User{}).Updates(&data).Error;err != nil{
		return err
	}
	return
}