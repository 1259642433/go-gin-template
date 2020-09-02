package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-template/config"
)

type PageVar struct{
	Page int
	Size int
}

func GetPageVar(c *gin.Context) (result PageVar) {
	result.Page, _ = com.StrTo(c.DefaultQuery("page",config.CurrentPage)).Int()
	result.Page--
	result.Size, _ = com.StrTo(c.DefaultQuery("size",config.PageSize)).Int()
	return
}

//func GetPageSize(c *gin.Context) int {
//	result := 0
//	page, _ := com.StrTo(c.Query("size")).Int()
//	if page > 0 {
//		result = (page - 1) * setting.PageSize
//	}
//
//	return result
//}