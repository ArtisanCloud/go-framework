package controllers

import (
	"fmt"
	"github.com/ArtisanCloud/go-framework/app/http/request"
	. "github.com/ArtisanCloud/go-framework/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelcomeAPIController struct {
	ServiceWelcomeAPI *WelcomeService
}

// 模块初始化函数 import 包时被调用
func init() {
}

func NewWelcomeAPIController() (ctl *WelcomeAPIController) {

	return &WelcomeAPIController{
		ServiceWelcomeAPI: NewWelcomeService(),
	}
}

func APIGetHome(context *gin.Context) {
	ctl := NewWelcomeAPIController()

	params, _ := context.Get("params")
	para := params.(request.ParaLogin)
	fmt.Printf("user: %s",para.UserName)
	//fmt.Printf("password: %s",para.Password)

	r := ctl.ServiceWelcomeAPI.GetWelcomeAPI()

	context.JSON(http.StatusOK, r)
}
