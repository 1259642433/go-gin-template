package v1

import (
	"github.com/gin-gonic/gin"
	"blog-api/app/utils"
	"net/http"
	"path"
)

func UploadFile(c *gin.Context) {
	f, _:=c.FormFile("file")
	suffix := path.Ext(f.Filename)
	name := utils.RandString(20) + suffix
	dst := path.Join("/public/file/cache/", name)
	err := c.SaveUploadedFile(f,dst)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传出错",
			"err":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "上传完成",
			"url": "/public/file/cache/"+name,
		})
	}
}