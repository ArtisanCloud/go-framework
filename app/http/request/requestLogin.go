package request

import (
	. "github.com/ArtisanCloud/go-framework/app/http"
	. "github.com/ArtisanCloud/go-framework/config"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/gin-gonic/gin"
)

type ParaLogin struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func ValidateRequestLogin(context *gin.Context) {
	var form ParaLogin

	if err := context.ShouldBind(&form); err != nil {
		if err := context.ShouldBindJSON(&form); err != nil {
			//Error("Error occurs",gin.H{"error": err.Error()})
			//js, _ := json.Marshal(gin.H{"error": err.Error()})
			//context.Writer.Header().Set("Content-Type", "application/json")
			//context.Writer.Write(js)
			apiResponse := &APIResponse{}
			apiResponse.Context = context
			apiResponse.SetCode(
				API_ERR_CODE_REQUEST_PARAM_ERROR,
				API_RETURN_CODE_ERROR,
				"", err.Error()).SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(context)
		}
	}

	context.Set("params", form)
	context.Next()
}
