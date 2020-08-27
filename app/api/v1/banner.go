package v1

import (
	"blog-api/app/models"
	"blog-api/app/utils"
	"blog-api/pkg/e"
	"blog-api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strconv"
)

// @Summary 获取banner列表
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success {"data": } *modelss.Banner{}
// @Router /api/v1/GetBanners [get]
func GetBanners(c *gin.Context){
	//maps := make(map[string]interface{})
	//data,err := models.GetBanners(10,1,maps)
	//count,err1 := models.GetBannerTotal(maps)
	//if err != nil||err1 != nil {
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": '',
	//		"msg": "查询出错",
	//		"data": err,
	//	})
	//} else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"data": data,
	//	})
	//}
	title := c.Query("title")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})



	if title != "" {
		maps["title"] = title
	}

	//var state int = -1
	//if arg := c.Query("state"); arg != "" {
	//	state = com.StrTo(arg).MustInt()
	//	maps["state"] = state
	//}

	code := e.SUCCESS

	pagingAttri := util.GetPageVar(c)

	data["lists"],_ = models.GetBanners(pagingAttri, maps)
	data["total"],_ = models.GetBannerTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func CreateBanner(c *gin.Context){
	//title,_ := c.GetQuery("title")
	//url,_ := c.GetQuery("url")
	//description,_ := c.GetQuery("description")
	//link,_ := c.GetQuery("link")

	var data models.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "绑定出错",
			"err": err.Error(),
		})
		return
	}
	if data.Url != "" {
		if urlResult,err := utils.MoveFileToS(data.Url);err != nil{
			log.Printf("文件移动出错：%+v",err)
		} else {
			data.Url = urlResult
		}
	}
	valid := validation.Validation{}
	valid.Required(data.Title, "title").Message("名称不能为空")
	valid.MaxSize(data.Title, 20, "title").Message("名称最长为20字符")
	valid.Required(data.Url, "url").Message("图片不能为空")
	code := e.INVALID_PARAMS

	if ! valid.HasErrors() {
		code = e.SUCCESS
		models.CreateBanner(data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
	//if err :=models.CreateBanner(data);err != nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"msg":"创建出错",
	//		"err":err.Error(),
	//	})
	//} else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg": "创建成功",
	//	})
	//}
}

func UpdateBanner(c *gin.Context){
	id,_ := strconv.Atoi(c.Query("id"))
	var data models.Banner
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"id不能为空",
		})
		return
	}
	if queryResult,err :=models.FindBanner(id);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"根据id索引查询失败",
			"err":err.Error(),
		})
		return
	} else if len(queryResult) == 0 {
		c.JSON(http.StatusOK,gin.H{
			"msg":"没有找到该条记录",
		})
		return
	}
	if err :=models.UpdateBanner(id,data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"更新出错",
			"err":err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg": "更新成功",
		})
	}
}

func DeleteBanner(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistBannerByID(id) {
			models.DeleteBanner(id)
		} else {
			code = e.ERROR_NOT_EXIST_BANNER
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
	//var data models.Banner
	//if err :=c.BindJSON(&data);err != nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"msg":"绑定出错",
	//		"err":err.Error(),
	//	})
	//	return
	//}
	//if queryResult,err :=models.FindBanner(data.ID);err != nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"msg":"根据id索引查询失败",
	//		"err":err.Error(),
	//	})
	//	return
	//} else if len(queryResult) == 0 {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg":"没有找到该条记录",
	//	})
	//	return
	//}
	//if data.ID ==0 {
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"msg":"id不能为空",
	//	})
	//	return
	//}
	//if err :=models.DeleteBanner(data);err != nil{
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"msg":"删除成功",
	//		"err":err.Error(),
	//	})
	//} else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"msg": "删除成功",
	//	})
	//}
}