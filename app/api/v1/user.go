package v1

import (
	"blog-api/app/middlewares"
	"blog-api/app/models"
	"blog-api/app/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){  //登陆
	var data models.User
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
	}
	user,err := models.FindUser(data.Account,data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "查询出错",
			"err":err.Error(),
		})
	} else {
		if len(user) > 0 {
			t, _ := time.ParseDuration("1h")
			claims :=middlewares.CustomClaims{
				user[0].ID,
				user[0].Account,
				user[0].Password,
				jwt.StandardClaims{
					NotBefore: int64(time.Now().Unix() - 60),
					ExpiresAt: int64(time.Now().Add(t).Unix()),
					IssuedAt: int64(time.Now().Unix()),
					Issuer:    "U.amazing",
				},
			}
			jwt := middlewares.NewJWT()
			token,err := jwt.CreateToken(claims);
			if err!=nil{
				c.JSON(http.StatusOK,gin.H{
					"msg": "token创建失败",
					"error": err,
				})
			}
			c.JSON(http.StatusOK,gin.H{
				"msg": "验证通过",
				"userInfo":user[0],
				"token": token,
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg": "用户名或密码错误",
			})
		}
	}
}
func UpdateUser(c *gin.Context){
	var data models.User
	if err :=c.BindJSON(&data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"绑定出错",
			"err":err.Error(),
		})
		return
	}
	if data.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"id不能为空",
		})
		return
	}
	//if queryResult,err :=models.FindNew(data.ID);err != nil{
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
	if data.Avater != "" {
		if urlResult,err := utils.MoveFileToS(data.Avater);err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"文件移动出错",
				"err":err.Error(),
			})
		} else{
			data.Avater = urlResult
		}
	}
	if err :=models.UpdateUser(data);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "更新出错",
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"msg": "更新成功",
		})
	}
}
func Verify(c *gin.Context){ //token验证登陆,同时返回user信息
	token := c.Request.Header.Get("token")
	claims,err := middlewares.NewJWT().ParseToken(token)
	if err != nil {
		c.JSON(401,gin.H{
			"msg": "验证失败",
		})
	}
	user,err := models.FindUser(claims.Account,claims.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "查询出错",
			"err":err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"msg": "验证通过",
		"userInfo":user[0],
	})
}

func Refresh(c *gin.Context){ //token刷新
	token := c.Request.Header.Get("token")
	d,_ := time.ParseDuration("72h")
	reToken,err := middlewares.NewJWT().RefreshToken(token,d);
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"msg": "token刷新成功",
		"re_token": reToken,
	})
}