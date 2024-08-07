package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// function for set key-val in session
func Set(c *gin.Context, key string, value string) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

// function for get and then deleting key-val from session
func Flash(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Get(key)
	session.Save()

	session.Delete(key)
	session.Save()

	if response != nil {
		return response.(string)
	}
	return ""

}

// function for get val of key from session
func Get(c *gin.Context, key string) string {
	session := sessions.Default(c)

	response := session.Get(key)
	session.Save()

	if response != nil {
		return response.(string)
	}
	return ""
}

// function for deleting key-val from session
func Remove(c *gin.Context, key string) {
	session := sessions.Default(c)

	session.Delete(key)
	session.Save()
}
