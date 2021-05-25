package en_US

import (
	fmt2 "fmt"
	"github.com/ArtisanCloud/go-framework/config"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var lang = language.English

func LoadLang() {

	message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_INIT), "")
	message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_WARNING), "return warning")
	message.SetString(lang, fmt2.Sprintf("%d",config.API_RETURN_CODE_ERROR), "return error")

	message.SetString(lang, fmt2.Sprintf("%d",config.API_RESULT_CODE_INIT), "")

	message.SetString(lang, fmt2.Sprintf("%d",config.API_RESULT_CODE_SUCCESS_RESET_PASSWORD), "Success")

	message.SetString(lang, fmt2.Sprintf("%d",config.API_WARNING_CODE_IN_MAINTENANCE), "System in maintenance")
	message.SetString(lang, fmt2.Sprintf("%d",config.API_WARNING_CODE_NEED_UPDATE), "New version found, please update the MiniProgram")

	message.SetString(lang, fmt2.Sprintf("%d",config.API_ERR_CODE_REQUEST_PARAMETER), "invalid request parameter")

}
