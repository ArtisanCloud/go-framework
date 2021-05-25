package SOAService

import (
	. "github.com/ArtisanCloud/go-framework/config"
	"github.com/gin-gonic/gin"
)

type Meta struct {
	ResultCode    int    `json:"result_code"`
	ResultMessage string `json:"result_message"`
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_message"`
	Locale        string `json:"locale"`
	Timezone      string `json:"timezone"`
}

type APIResponse struct {
	Context *gin.Context `json:"-"`

	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func NewAPIResponse(ctx *gin.Context) (rs *APIResponse) {
	if ctx == nil {
		ctx = &gin.Context{}
	}

	rs = &APIResponse{
		Context: ctx,
		Meta: Meta{
			ReturnCode:    API_RETURN_CODE_INIT,
			ReturnMessage: "",
			ResultCode:    API_RESULT_CODE_INIT,
			ResultMessage: "",
		},
		Data: nil,
	}
	return rs
}

