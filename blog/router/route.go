package router

import (
	"blog/controller/category"
	"blog/controller/comment"
	"blog/controller/post"
	"blog/controller/tag"
	"blog/controller/test"
	"blog/controller/user"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

// Route 路由
func Route(router *gin.Engine) {
	apiPrefix := "/"

	api := router.Group(apiPrefix)
	{
		api.Static("/static", "static/")
		// 文章
		api.GET("/post/list", post.PostListView)
		api.GET("/post/detail", post.GetPostDetailView)
		api.POST("/post/create", middleware.LoginRequired, post.Create)
		api.POST("/post/update", middleware.LoginRequired, post.Update)
		api.POST("/post/delete", middleware.LoginRequired, post.Delete)
		api.POST("/post/add_tag", middleware.LoginRequired, post.PostAddTag)

		// 评论
		api.GET("/comment/list", comment.GetCommentListView)
		api.POST("/comment/create", middleware.LoginRequired, comment.Create)
		api.POST("/comment/update", middleware.LoginRequired, comment.Update)
		api.POST("/comment/delete", middleware.LoginRequired, comment.Delete)

		// 文章归类
		api.GET("/category/list", category.CategoryListView)
		api.POST("/category/create", middleware.LoginRequired, category.Create)
		api.POST("/category/update", middleware.LoginRequired, category.Update)
		api.POST("/category/delete", middleware.LoginRequired, category.Delete)


		// 文章tag
		api.GET("/tag/list", tag.TagListView)
		api.POST("/tag/create", middleware.LoginRequired, tag.Create)

		// 用户
		api.POST("/user/create", user.Signup)
		api.POST("/login", user.Signin)
		//api.GET("/logout", middleware.LoginRequired, user.Signout)

		// 测试用
		api.GET("/test", test.TestGet)
	}

}
