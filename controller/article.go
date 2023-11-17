package controller

import (
	"TreeHole/dao"
	"TreeHole/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Article 实例化chat类型对象，首字母大写用于跨包调用
var Article article

// 声明article结构体
type article struct{}

// CreateArticle 创建帖子
func (a article) CreateArticle(ctx *gin.Context) {
	// 拿到身份
	claims, _ := ctx.Get("claims")
	role := claims.(map[string]interface{})["role"]
	id := claims.(map[string]interface{})["id"]
	//参数绑定
	params := new(struct {
		Title string `form:"title" binding:"required"`
		Text  string `form:"text" binding:"required"`
	})
	if role.(int) >= 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "权限不够！",
			"data": nil,
		})
		return
	}
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	createArticle, err := dao.Dao.CreateArticle(utils.GetUint(fmt.Sprint(id)), params.Text, params.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "创建成功",
		"data": createArticle,
	})
}

// GetArticleList 获取帖子列表
func (a article) GetArticleList(ctx *gin.Context) {
	// 拿到身份
	//claims, _ := ctx.Get("claims")
	//role := claims.(map[string]interface{})["role"]
	articles, err := dao.Dao.GetArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取帖子列表失败",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取帖子列表成功",
		"data": articles,
	})
}
