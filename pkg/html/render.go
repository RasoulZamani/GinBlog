package html

import (
	"github.com/RasoulZamani/hiGin/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)
	c.HTML(code, name, data)
}
