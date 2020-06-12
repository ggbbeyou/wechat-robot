package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanyways/wechat-robot/config"
	"github.com/kanyways/wechat-robot/controller/receive"
	"net/http"
)

// Route 路由
func Route(router *gin.Engine) {

	// 配置资源的文件目录
	router.Static("/static", "static")
	// 配置网站的图标
	router.StaticFile("/favicon.ico", "static/favicon.ico")
	// 加载模板
	router.LoadHTMLGlob("views/*")
	// 默认的欢迎页面
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   config.ServerConfig.SiteName,
			"message": "你来了呀？",
		})
	})

	apiPrefix := config.ServerConfig.ApiPrefix
	api := router.Group(apiPrefix)
	{
		// 收到Post请求
		api.POST("/receive/:id", receive.PostMessage)
		// 收到Get请求
		api.GET("/receive/:id", receive.GetMessage)
	}
}
