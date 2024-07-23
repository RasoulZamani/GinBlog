package view

import (
	"time"

	"github.com/gin-gonic/gin"
)

func WithGlobalData(data gin.H) gin.H {
	data["date"] = time.Now()
	return data
}
