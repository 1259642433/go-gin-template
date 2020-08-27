package models

import (
	"blog-api/app/utils"
	"blog-api/pkg/util"
	"log"
)

type Banner struct {
	Model
	Title string `json:"title"`
	Url string `json:"url"`
	Description string `json:"description"`
	Link string `json:"link"`
}

//获取多个banner
func GetBanners(pageVar util.PageVar, maps interface {})(data []Banner,err error){
	err = db.Where(maps).Offset(pageVar.Page).Limit(pageVar.Size).Find(&data).Error
	return
}

// 查询指定banner详情
func FindBanner(id int)(data []*Banner,err error){
	err =db.Where("id = ?",id).Find(&data).Error
	return
}

//// 查询banner是否存在
//func ExistBannerByTitle(title string)bool{
//	if err := db.Where("title = ?",title).Error;err !=nil {
//		return false
//	}
//	return true
//}

func  ExistBannerByID(id int) bool {
	var banner Banner
	db.Select("id").Where("id = ?", id).First(&banner)
	if banner.ID > 0 {
		return true
	}
	return false
}


// 创建banner记录
func CreateBanner(data Banner)(err error){
	err = db.Create(&data).Error
	return
}

//修改指定banner数据
func UpdateBanner(id int,data Banner)(err error){
	if data.Url != "" {
		if urlResult,err := utils.MoveFileToS(data.Url);err != nil{
			log.Printf("文件删除失败,id:%s",err.Error())
		} else{
			data.Url = urlResult
		}
	}
	err = db.Model(&Banner{}).Where("id=?",id).Updates(&data).Error
	return
}

//删除指定banner
func DeleteBanner(id int)(err error){
	var banner Banner
	err = db.Where("id = ?",id).Delete(&Banner{}).First(&banner).Error
	if banner.Url != "" {
		if err := utils.RemoveFile(banner.Url);err != nil{
			// 列入消息队列
			// 删除达到最大次数记录到日志当中
			log.Printf("文件删除失败,id:%d",banner.ID)
		}
	}
	return
}

func GetBannerTotal(maps interface {}) (count int,err error){
	err = db.Model(&Banner{}).Where(maps).Count(&count).Error
	return
}