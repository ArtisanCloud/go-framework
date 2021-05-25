package service

import (
	. "github.com/ArtisanCloud/go-framework/app/models"
	"github.com/ArtisanCloud/go-framework/config"
)

type WelcomeService struct {
	User *User
}


/**
 ** 初始化构造函数
*/

// 模块初始化函数 import 包时被调用
func init() {
}

func NewWelcomeService() (r * WelcomeService)  {
	r = &WelcomeService{
		User: NewUser(),
	}
	return r
}


/**
 ** 实例函数
 */

func (srv *WelcomeService) GetWelcome() string {

	return "Welcome! " + config.APP_NAME + " version:" + config.APP_VERSION

}


func (srv *WelcomeService) GetWelcomeAPI() string {

	return "Welcome API!"

}
