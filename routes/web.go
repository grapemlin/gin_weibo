package routes

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers/sessions"
	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/user"
	"gin_weibo/middleware"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{
		g.GET("/", staticpage.Home)
		g.GET("/help", staticpage.Help)
		g.GET("/about", staticpage.About)
	}

	// ------------------------------ user ------------------------------
	{
		g.GET("/signup", middleware.Guest(user.Create))
		userRouter := g.Group("/users")
		{
			// 用户列表页面
			userRouter.GET("", user.Index)
			// 展示具体用户页面
			userRouter.GET("/show/:id", user.Show)

			// 创建用户页面
			userRouter.GET("/create", middleware.Guest(user.Create))
			// 保存新用户
			userRouter.POST("", user.Store)

			// 编辑用户页面
			userRouter.GET("/edit/:id", middleware.Auth(user.Edit))
			// 修改用户
			userRouter.POST("/update/:id", middleware.Auth(user.Update))

			// 删除用户
			userRouter.POST("/destory/:id", user.Destroy)
		}
	}

	// ------------------------------ sessions ------------------------------
	{
		// 登录页面
		g.GET("/login", sessions.Create)
		// 登录
		g.POST("/login", sessions.Store)
		// 登出
		g.POST("/logout", sessions.Destroy)
	}
}
