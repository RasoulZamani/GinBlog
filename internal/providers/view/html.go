package view

import (
	"time"

	"github.com/RasoulZamani/hiGin/internal/modules/user/helper"
	"github.com/RasoulZamani/hiGin/pkg/converters"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["date"] = time.Now()

	data["registrationFormErrors"] = converters.StringToMap(sessions.Flash(c, "registrationFormErrors"))
	data["loginFormErrors"] = converters.StringToMap(sessions.Flash(c, "loginFormErrors"))
	data["articleFormErrors"] = converters.StringToMap(sessions.Flash(c, "articleFormErrors"))

	data["oldRegisterForm"] = converters.StringToListMap(sessions.Flash(c, "oldRegisterForm"))
	data["oldLoginForm"] = converters.StringToListMap(sessions.Flash(c, "oldLoginForm"))
	data["oldArticleForm"] = converters.StringToListMap(sessions.Flash(c, "oldArticleForm"))

	data["auth"] = helper.Auth(c)

	return data
}
