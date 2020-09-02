package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gin-template/config"
	"go-gin-template/controllers/v1"
	"go-gin-template/middlewares/cors"
	"net/http"
)

func init() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则会返回404
	router.Use(cors.Cors())
	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	//gin.default()创建路由时其实已经默认调用Logger,Recovery这两个中间件了
	//router.Use(gin.Logger(),gin.Recovery())

	// token鉴权
	//router.Use(middlewares.Auth())

	gin.SetMode(config.RunMode)

	// var jwt = middlewares.Auth()
	base := router.Group("/api/v1")
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
				//user.GET("/refresh", v1.Refresh)
			}
			admin.POST("/file", v1.UploadFile)
		}

		//cases := v1.Group("/case")
		//{
		//	cases.GET("",controllers.GetCases)
		//	cases.GET("/:regionId",controllers.GetRegionCases)
		//	cases.POST("",controllers.CreateCase)
		//	cases.PUT("",controllers.UpdateCase)
		//	cases.DELETE("",controllers.DeleteCase)
		//}
		//intelligence := v1.Group("/intelligence")
		//{
		//	intelligence.GET("",controllers.GetIntelligences)
		//	intelligence.POST("",controllers.CreateIntelligence)
		//	intelligence.PUT("",controllers.UpdateIntelligence)
		//	intelligence.DELETE("",controllers.DeleteIntelligence)
		//}
		//link := v1.Group("/link")
		//{
		//	link.GET("",controllers.GetLinks)
		//	link.POST("",controllers.CreateLink)
		//	link.PUT("",controllers.UpdateLink)
		//	link.DELETE("",controllers.DeleteLink)
		//}
		//maps := v1.Group("/map")
		//{
		//	maps.GET("",controllers.GetMap)
		//	maps.PUT("",controllers.UpdateMap)
		//}
		//news := v1.Group("/new")
		//{
		//	news.GET("",controllers.GetNews)
		//	news.GET("/:id",controllers.GetNewsDetail)
		//	news.POST("",controllers.CreateNew)
		//	news.PUT("",controllers.UpdateNew)
		//	news.DELETE("",controllers.DeleteNew)
		//}
	}
	router.StaticFS("/public/file", http.Dir("/file"))
	//router.StaticFile("/static/file", "../static/file/cacheArea/ZQEUANJFFSPVOHLINGOE.jpg")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//router.Run(":12577")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
