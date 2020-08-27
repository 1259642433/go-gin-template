package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type PageVar struct{
	Page int
	Size int
}

func GetPageVar(c *gin.Context) (result PageVar) {
	if result.Page, _ = com.StrTo(c.Query("page")).Int();result.Page == 0{
		result.Page = 1
	}
	if result.Size, _ = com.StrTo(c.Query("size")).Int();result.Size == 0{
		result.Size = 10
	}
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