package service

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	Context *gin.Context
}

/**
 ** 初始化构造函数
 */

// 模块初始化函数 import 包时被调用
func init() {
}

func NewService(ctx *gin.Context) (r *Service) {
	r = &Service{
		Context: ctx,
	}
	return r
}
