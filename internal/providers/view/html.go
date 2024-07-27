package view

import (
	"time"

	"github.com/RasoulZamani/hiGin/pkg/converters"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["date"] = time.Now()

	data["registrationFormErrors"] = converters.StringToMap(sessions.Flash(c, "registrationFormErrors"))

	return data
}
