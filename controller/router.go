package controller

import "github.com/gin-gonic/gin"

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体w
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
}
