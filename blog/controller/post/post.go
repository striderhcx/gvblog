package post

import (
	utils "blog/common"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"strings"
)

var JsonResponse = utils.JsonResponse

func GetPostDetailView(c *gin.Context) {
	postId := c.DefaultQuery("post_id", "-1")
	detailId, err := strconv.Atoi(postId)
	if err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "post_id参数不是整数", c)
		return
	}
	if postId != "-1" {
		category, post := model.Category{}, model.Post{}
		model.DB.Where("id = ?", detailId).First(&post)
		model.DB.Where("id = ?", post.CategoryID).First(&category)

		result := map[string]interface{}{}
		result["post"] = map[string]interface{}{
			"id":            post.ID,
			"title":         post.Title,
			"content":       post.Content,
			"category_id":   post.CategoryID,
			"category_name": category.Name,
			"created_time":  post.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_time":  post.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		c.JSON(http.StatusOK, gin.H{
			"code": utils.Success,
			"msg":  "ok",
			"data": result["post"],
		})
	}
}

func PostListView(c *gin.Context) {
	type PostListParam struct {
		Page         int   `form:"page"`
		Limit        int   `form:"limit" binding:"required"`
		CategoryId  int   `form:"category_id"`
		TagId		int	  `form:"tag_id"`
		Search      string 	`form:"search"`
	}
	params := PostListParam{}
	if err := c.ShouldBindQuery(&params); err != nil {
		JsonResponse(utils.ParamError, "请求参数有误", c)
		fmt.Println(err)
		return
	}
	pagePosts, categories := []model.Post{}, []model.Category{}
	fmt.Println("params.CategoryId", params.CategoryId)

	postQuery := model.DB
	if params.CategoryId != 0 {
		postQuery = postQuery.Where(" category_id= ? ", params.CategoryId)
	}
	if params.Search != "" {
		postQuery = postQuery.Where("title LIKE ?", "%" + fmt.Sprintf("%s", params.Search) + "%")
	}
	if params.TagId != 0 {
		postTags := []model.PostTag{}
		model.DB.Where(" tag_id= ? ", params.TagId).Find(&postTags)
		postIds := []int{}
		for _, post_tag := range postTags {
			postIds = append(postIds, post_tag.PostID)
		}
		postQuery = postQuery.Where("id IN (?)", postIds)
	}
	postQuery.Where("pstatus = ? ", model.DataStatusNormal).Order("created_at desc, id desc").
	Offset(params.Page * params.Limit).Limit(params.Limit).
	Find(&pagePosts)

	count := 0
	model.DB.Table("posts").Where("pstatus= ? ", model.DataStatusNormal).Count(&count)
	model.DB.Order("name").Where("pstatus= ? ", model.DataStatusNormal).Find(&categories)
	for _, post := range pagePosts{
		p := &post
		limit_content := []rune(p.Content)
		// 只显示固定的前面200个字符作为简介
		if len(limit_content) > 200 {
			p.Content = string(limit_content[:200])
		}
	}
	respPosts := []map[string]interface{}{}
	categoryMap := map[int]interface{}{}
	for _, category := range categories {
		categoryMap[category.ID] = category.Name
	}
	for _, post := range pagePosts {
		category_name, ok := categoryMap[post.CategoryID]
		if !ok {
			category_name = ""
		}
		postItem := map[string]interface{}{
			"id":            post.ID,
			"title":         post.Title,
			"content":       post.Content,
			"category_id":   post.CategoryID,
			"category_name": category_name,
			"created_time":  post.CreatedAt.Format("2006-01-02 15:04:05"),
			"updated_time":  post.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		respPosts = append(respPosts, postItem)
	}
	pagination := utils.Pagination(count, params.Page, params.Limit)
	c.JSON(http.StatusOK, gin.H{
		"code":       utils.Success,
		"msg":        "ok",
		"data":       respPosts,
		"pagination": pagination,
	})
}

func Create(c *gin.Context) {
	// todo 自定义错误提示返回
	type PostParams struct {
		CategoryID   int       `json:"category_id" binding:"required,numeric,min=1"`
		Title        string     `json:"title" binding:"required,min=8,max=40"`
		Content      string     `json:"content" binding:"required"`
		TagIds		 []int      `json:"tag_ids" binding:"required"`
	}
	postForm := PostParams{}
	if err := c.ShouldBindWith(&postForm, binding.JSON); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	if postForm.Title == "" {
		JsonResponse(utils.ParamError, "文章title不能为空",  c)
		return
	}
	if postForm.Content == "" {
		JsonResponse(utils.ParamError, "文章内容不能为空", c)
		return
	}

	newPost := model.Post{}
	newPost.Title = strings.TrimSpace(postForm.Title)
	newPost.Content = strings.TrimSpace(postForm.Content)
	newPost.CategoryID = postForm.CategoryID

	tx := model.DB.Begin()
	category := model.Category{}
	if tx.First(&category, postForm.CategoryID).Error != nil {
		JsonResponse(utils.ParamError, fmt.Sprintf("不存在ID为%v的分类", postForm.CategoryID), c)
		return
	}
	saveErr := tx.Create(&newPost).Error
	fmt.Println("新创建的post id：", newPost.ID)
	if saveErr != nil {
		fmt.Println(saveErr.Error())
		JsonResponse(utils.MySqlError, "error",  c)
		return
	}

	// Todo 这里解析这个表单数组可以写得更通用
	//tag_ids_strs := c.PostForm("tag_ids")
	//tag_ids_strs = strings.Trim(strings.Trim(tag_ids_strs, "["), "]")
	//tag_ids_str_slice := strings.Split(tag_ids_strs, ",")
	//tag_ids := []int{}
	//for _, tag_id_str := range tag_ids_str_slice {
	//	num_str := strings.TrimSpace(tag_id_str)
	//	num, c_err := strconv.Atoi(num_str)
	//	if c_err == nil {
	//		tag_ids = append(tag_ids, int(num))
	//	}
	//}

	tag_ids := postForm.TagIds
	fmt.Println("tag_ids", tag_ids)
	tags := []model.Tag{}
	if len(tag_ids) != 0 {
		tx.Where("id IN (?)", tag_ids).Find(&tags)
		if len(tags) != len(tag_ids) {
			JsonResponse(utils.ParamError, "存在重复添加或者不存在的tag", c)
			tx.Rollback()
			return
		}
		for _, tagId := range tag_ids {
			newPostTag := model.PostTag{}
			newPostTag.PostID = newPost.ID
			newPostTag.TagId = tagId
			createTagErr := tx.Create(&newPostTag).Error
			if createTagErr != nil {
				fmt.Println("创建标签失败了", createTagErr)
				JsonResponse(utils.ParamError, "内部错误：创建标签失败了", c)
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Update(c *gin.Context) {
	type PostUpdateForm struct {
		PostId      int		`form:"post_id" binding:"required"`
		Title 		string  `form:"title,omitempty"`
		Content 	string  `form:"content,omitempty"`
	}
	postForm := PostUpdateForm{}
	if err := c.ShouldBindWith(&postForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	// 记录前端非空请求字段
	updateFieldMap := map[string]interface{}{}
	postForm.Title = strings.TrimSpace(postForm.Title)
	postForm.Content = strings.TrimSpace(postForm.Content)
	if postForm.Title != "" {
		updateFieldMap["title"] = postForm.Title
	}
	if postForm.Content != "" {
		updateFieldMap["content"] = postForm.Content
	}
	post := model.Post{}
	if model.DB.First(&post, postForm.PostId).Error != nil {
		JsonResponse(utils.ParamError, "不存在ID为"+fmt.Sprintf("%v", postForm.PostId)+"的博文",  c)
		return
	}
	model.DB.Model(&post).Updates(updateFieldMap)
	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

func Delete(c *gin.Context) {
	type PostDeleteForm struct {
		PostId      int		`form:"post_id" binding:"required"`
	}
	postDeleteForm := PostDeleteForm{}
	if err := c.ShouldBindWith(&postDeleteForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError, "参数无效", c)
		return
	}
	post := model.Post{}
	if model.DB.First(&post, postDeleteForm.PostId).Error != nil {
		JsonResponse(utils.ParamError, "不存在ID为"+fmt.Sprintf("%v", postDeleteForm.PostId)+"的博文", c)
		return
	}
	delete_err := model.DB.Delete(&post).Error
	if delete_err != nil {
		fmt.Println(delete_err.Error())
		JsonResponse(utils.MySqlError, "error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

// 给某篇文章添加tag
func PostAddTag(c *gin.Context) {
	type AddTagForm struct {
		PostId      int		`form:"post_id" binding:"required"`
		TagId       int     `form:"tag_id" binding:"required"`
	}
	addTagForm := AddTagForm{}
	if err := c.ShouldBindWith(&addTagForm, binding.Form); err != nil {
		fmt.Println(err.Error())
		JsonResponse(utils.ParamError,"参数无效", c)
		return
	}
	post, tag := model.Post{}, model.Tag{}
	if model.DB.First(&post, addTagForm.PostId).Error != nil {
		JsonResponse(utils.ParamError, "不存在ID为"+fmt.Sprintf("%v", addTagForm.PostId)+"的博文", c)
		return
	}

	if model.DB.First(&post, addTagForm.PostId).Error != nil {
		JsonResponse(utils.ParamError, "不存在ID为"+fmt.Sprintf("%v", addTagForm.PostId)+"的博文", c)
		return
	}

	if model.DB.First(&tag, addTagForm.TagId).Error != nil {
		JsonResponse(utils.ParamError, "不存在ID为"+fmt.Sprintf("%v", addTagForm.TagId)+"的博文", c)
		return
	}

	newPostTag := model.PostTag{}
	newPostTag.PostID = addTagForm.PostId
	newPostTag.TagId = addTagForm.TagId

	create_err := model.DB.Create(&newPostTag).Error
	if create_err != nil {
		fmt.Println(create_err.Error())
		JsonResponse(utils.MySqlError, "error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": utils.Success,
		"msg":  "ok",
		"data": nil,
	})
}

