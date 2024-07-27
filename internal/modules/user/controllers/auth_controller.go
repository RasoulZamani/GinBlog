package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		"title": "Registration",
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

	// check if user is already exist with email

	if controller.userService.CheckEmailExist(request.Email) {
		errors.Init()
		errors.Add("Email", "The user with this email is already exists")

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
	//store created user id in sessions
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("user created successfully with username: %s", user.UserName)
	c.Redirect(http.StatusFound, "/")
}

// ******************************** Login ****************************************
func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {

	// validate request
	var request auth.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)

		// save errors in session
		sessions.Set(c, "loginFormErrors", converters.MapToString(errors.Get()))
		// session then send to template by WithGlobalData function

		// save form data to session in order to re-present them if errors happened
		old.Init()
		old.Set(c)
		sessions.Set(c, "oldLoginForm", converters.ListMapToString(old.Get()))
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// check email and password
	user, err := controller.userService.LoginUser(request)
	if err != nil {
		errors.Init()
		errors.Add("Email", err.Error())

		// save errors in session
		sessions.Set(c, "loginFormErrors", converters.MapToString(errors.Get()))
		// session then send to template by WithGlobalData function

		// save form data to session in order to re-present them if errors happened
		old.Init()
		old.Set(c)
		sessions.Set(c, "oldLoginForm", converters.ListMapToString(old.Get()))
		c.Redirect(http.StatusFound, "/login")
		return

	}

	//store login user id in sessions
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("user logged successfully with username: %s", user.UserName)
	c.Redirect(http.StatusFound, "/")
}
