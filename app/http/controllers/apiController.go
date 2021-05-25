package controllers

import (
	"errors"
	"fmt"
	. "github.com/ArtisanCloud/go-framework/app/http"
	service "github.com/ArtisanCloud/go-framework/app/services"
	"github.com/ArtisanCloud/go-framework/config"
	"github.com/gin-gonic/gin"
	"runtime"
	"runtime/debug"
)

type APIController struct {
	ServiceUser *service.UserService
	Context     *gin.Context
	RS          *APIResponse
}

func NewAPIController(context *gin.Context) *APIController {
	return &APIController{
		ServiceUser: service.NewUserService(context),
		Context:     context,
		RS:          NewAPIResponse(context),
	}
}

func RecoverResponse(context *gin.Context, action string) {

	if p := recover(); p != nil {
		var err error
		apiResponse := &APIResponse{}
		apiResponse.Context = context
		switch rs := p.(type) {

		// 获取业务流程中的异常错误
		case *APIResponse:
			rs.ThrowJSONResponse(context)
			break

		case runtime.Error:
			err = p.(runtime.Error)
		case string:
			err = errors.New(p.(string))
		// 若非APIResponse，也许默认抛出一个若非APIResponse
		default:
		}

		if err != nil {
			fmt.Printf("Unknown panic: %v \r\n", err.Error())
			fmt.Printf("err stack: %v \r\n", string(debug.Stack()))

			apiResponse.SetReturnCode(config.API_RETURN_CODE_ERROR, "Inner Error")
			apiResponse.ThrowJSONResponse(context)
		}

	}
}
