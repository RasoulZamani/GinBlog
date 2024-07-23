package html

import "github.com/gin-gonic/gin"

func LoadHtml(router *gin.Engine) {
	//internal/modules/home//html//tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
