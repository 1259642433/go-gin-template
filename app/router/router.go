package router

import (
	"blog-api/app/api/v1"
	"blog-api/app/middlewares"
	"blog-api/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则会返回404

	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	//router.Use(middlewares.Recover)
	router.Use(middlewares.Cors())

	// var jwt = middlewares.Auth()
	base := router.Group("/app/v1")
	{
		admin := base.Group("/admin")
		{
			banner := admin.Group("/banner")
			{
				banner.GET("", v1.GetBanners)
				banner.POST("", v1.CreateBanner)
				banner.PUT("", v1.UpdateBanner)
				banner.DELETE("", v1.DeleteBanner)
			}
			user := admin.Group("/user")
			{
				user.GET("", v1.Verify)
				user.POST("", v1.UpdateUser)
				user.POST("/login", v1.Login)
				user.GET("/refresh", v1.Refresh)
			}
			admin.POST("/file", v1.UploadFile)
		}

		//cases := v1.Group("/case")
		//{
		//	cases.GET("",api.GetCases)
		//	cases.GET("/:regionId",api.GetRegionCases)
		//	cases.POST("",api.CreateCase)
		//	cases.PUT("",api.UpdateCase)
		//	cases.DELETE("",api.DeleteCase)
		//}
		//intelligence := v1.Group("/intelligence")
		//{
		//	intelligence.GET("",api.GetIntelligences)
		//	intelligence.POST("",api.CreateIntelligence)
		//	intelligence.PUT("",api.UpdateIntelligence)
		//	intelligence.DELETE("",api.DeleteIntelligence)
		//}
		//link := v1.Group("/link")
		//{
		//	link.GET("",api.GetLinks)
		//	link.POST("",api.CreateLink)
		//	link.PUT("",api.UpdateLink)
		//	link.DELETE("",api.DeleteLink)
		//}
		//maps := v1.Group("/map")
		//{
		//	maps.GET("",api.GetMap)
		//	maps.PUT("",api.UpdateMap)
		//}
		//news := v1.Group("/new")
		//{
		//	news.GET("",api.GetNews)
		//	news.GET("/:id",api.GetNewsDetail)
		//	news.POST("",api.CreateNew)
		//	news.PUT("",api.UpdateNew)
		//	news.DELETE("",api.DeleteNew)
		//}
	}
	router.StaticFS("/public/file", http.Dir("/file"))
	//router.StaticFile("/static/file", "../static/file/cacheArea/ZQEUANJFFSPVOHLINGOE.jpg")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Run(":12577")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
