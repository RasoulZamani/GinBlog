package controllers

import (
	"log"
	"net/http"

	"github.com/RasoulZamani/hiGin/internal/modules/user/requests/auth"
	UserService "github.com/RasoulZamani/hiGin/internal/modules/user/services"
	"github.com/RasoulZamani/hiGin/pkg/errors"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Registration Form",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {

	// validate request
	var request auth.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		c.JSON(http.StatusOK, gin.H{
			"errors": errors.Get(),
		})
		return
	}

	// create user
	user, err := controller.userService.Create(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	log.Printf("user created successfully with username: %s", user.UserName)
	c.Redirect(http.StatusFound, "/")
}
