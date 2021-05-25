package controllers

import (
	. "github.com/ArtisanCloud/go-framework/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeController struct {
	ServiceWelcome *WelcomeService
}

// 模块初始化函数 import 包时被调用
func init() {
}

func NewWelcomeController() (ctl *WelcomeController) {

	return &WelcomeController{
		ServiceWelcome: NewWelcomeService(),
	}
}

func WebGetHome(context *gin.Context) {
	ctl := NewWelcomeController()

	r := ctl.ServiceWelcome.GetWelcome()

	context.JSON(http.StatusOK, r)
}
