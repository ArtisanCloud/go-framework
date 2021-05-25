package en_TW

import (
	fmt2 "fmt"
	"github.com/ArtisanCloud/go-framework/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.Chinese

func LoadLang() {
		message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_INIT) , "")
		message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_WARNING) , "警告消息")
		message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_ERROR) , "错误消息")

		message.SetString(lang, fmt2.Sprintf("%d",config.API_RESULT_CODE_INIT) , "")

		message.SetString(lang, fmt2.Sprintf("%d",config.API_RESULT_CODE_SUCCESS_RESET_PASSWORD) , "密码修改成功")

		message.SetString(lang, fmt2.Sprintf("%d",config.API_WARNING_CODE_IN_MAINTENANCE) , "系统维护中")
		message.SetString(lang, fmt2.Sprintf("%d",config.API_WARNING_CODE_NEED_UPDATE) , "推出新版本，请更新最新的版本")

		message.SetString(lang, fmt2.Sprintf("%d",config.API_ERR_CODE_REQUEST_PARAMETER) , "请求参数错误")

}
