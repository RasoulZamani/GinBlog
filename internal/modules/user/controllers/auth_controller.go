package controllers

import (
	"log"
	"net/http"

	"github.com/RasoulZamani/hiGin/internal/modules/user/requests/auth"
	UserService "github.com/RasoulZamani/hiGin/internal/modules/user/services"
	"github.com/RasoulZamani/hiGin/pkg/converters"
	"github.com/RasoulZamani/hiGin/pkg/errors"
	"github.com/RasoulZamani/hiGin/pkg/html"
	"github.com/RasoulZamani/hiGin/pkg/old"
	"github.com/RasoulZamani/hiGin/pkg/sessions"
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

		// save errors in session
		sessions.Set(c, "registrationFormErrors", converters.MapToString(errors.Get()))
		// session then send to template by WithGlobalData function

		// save form data to session in order to re-present them if errors happened
		old.Init()
		old.Set(c)
		sessions.Set(c, "oldRegisterForm", converters.ListMapToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
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
