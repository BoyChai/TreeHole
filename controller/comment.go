package controller

import (
	"TreeHole/dao"
	"TreeHole/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Comment 实例化comment类型对象，首字母大写用于跨包调用
var Comment comment

// 声明comment结构体
type comment struct{}

// GetComment 获取对应帖子评论
func (c comment) GetComment(ctx *gin.Context) {
	// 拿到身份
	//claims, _ := ctx.Get("claims")
	//id := claims.(map[string]interface{})["id"]
	params := new(struct {
		Article int `form:"article" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	Comments, err := dao.Dao.GetComment(fmt.Sprint(params.Article))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "获取成功",
		"data": Comments,
	})
}

func (c comment) GetHomeComment(ctx *gin.Context) {
	homeComment, err := dao.Dao.GetHomeComment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": homeComment,
	})
}

// SendComment 发送评论
func (c comment) SendComment(ctx *gin.Context) {
	// 拿到身份
	claims, ok := ctx.Get("claims")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "未找到身份信息",
			"data": nil,
		})
		return
	}

	claimsMap, ok := claims.(map[string]interface{})
	if !ok || claimsMap == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "身份信息格式错误",
			"data": nil,
		})
		return
	}
	id, idExists := claimsMap["id"].(string)
	if !idExists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "未找到用户ID",
			"data": nil,
		})
		return
	}

	params := new(struct {
		Article string `form:"article" binding:"required"`
		Text    string `form:"text" binding:"required"`
		Parent  int    `form:"parent"`
	})

	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	sendComment, err := dao.Dao.SendComment(utils.GetUint(fmt.Sprint(id)), utils.GetUint(fmt.Sprint(params.Article)), params.Text, utils.GetUint(fmt.Sprint(params.Parent)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "发送成功",
		"data": sendComment,
	})
}
