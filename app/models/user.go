package models

import (
	"time"
)

type User struct {
	Model
	Account     string    `json:"account"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Avater      string    `json:"avater"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	LoginTime   time.Time `json:"login_time"`
	Type        int       `json:"type"`
}

func FindUser(account string,pwd string)(user []*User,err error){
	if err:=db.Where("account = ? AND password = ?", account, pwd).Find(&user).Error;err != nil{
		return nil,err
	}
	db.Where("account = ? AND password = ?", account, pwd).Model(&user).Update("login_time",time.Time(time.Now()).Format(TimeFormat))
	return
}
func UpdateUser(data User)(err error){
	if err:=db.Model(&User{}).Updates(&data).Error;err != nil{
		return err
	}
	return
}