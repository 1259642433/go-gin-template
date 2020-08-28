package utils

import (
	"io"
	"log"
	"os"
	"path"
)

func MoveFileToS(url string)(result string,err error){
	cache := "./public/file/cache/"
	storage := "/public/file/storage/"
	name := path.Base(url)
	if url == "" {
		return
	}
	fileC, err := os.Open(cache+name)
	if err != nil{
		log.Println("源文件读取错误:%v",err.Error())
		return "",err
	}
	defer fileC.Close()
	fileS,err := os.Create(storage+name)
	if err != nil {
		log.Println("目标文件创建错误:%v",err.Error())
		return "",err
	}
	defer fileS.Close()
	_, err = io.Copy(fileS,fileC)
	if err != nil{
		log.Println("文件复制错误:%v",err.Error())
		return "",err
	}
	return "/public/file/storage/"+name,err
}

func RemoveFile(url string)(err error){
	if url == "" {
		return
	}
	err = os.Remove(url)
	if err != nil {
		log.Println("文件删除失败")
		return err
	}
	return
}