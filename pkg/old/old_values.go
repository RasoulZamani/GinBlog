package old

import "github.com/gin-gonic/gin"

var oldValues = make(map[string][]string)

func Init() {
	oldValues = map[string][]string{}
}

func Set(c *gin.Context) {
	c.Request.ParseForm()
	oldValues = c.Request.PostForm
}

func Get() map[string][]string {
	return oldValues

}
