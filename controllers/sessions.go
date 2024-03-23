package controllers

import (
	"beego_todos/models"

	beego "github.com/beego/beego/v2/server/web"
)

type SessionsController struct {
	beego.Controller
}

func (c *SessionsController) SignupPage() {
	c.TplName = "sessions/signup.tpl"
}

func (c *SessionsController) LoginPage() {
	c.TplName = "sessions/login.tpl"
}

func (c *SessionsController) Signup() {
	// TODO:
	// Step 1: Fetch the email, password, and confirm password from the request parameters.
	// Step 2: Check if the email already exists in the database.
	//         This can be checked by calling models.UserCheckAvailability method.
	// Step 3: If the email already exists, set an alert message accordingly in c.Data["alert"]
	//		   and render the "sessions/signup.tpl" template.
	// Step 4: If the password and confirm password do not match,
	//         set an alert message indicating this mismatch
	//         and render the "sessions/signup.tpl" template.
	// Step 5: If the email is available and password and confirm password match,
	//         create a new user by calling models.UserCreate.
	// Step 6: If the user creation is successful, set the session with name "user_id" and the user's ID (user.Id),
	//         redirect the user to the home page ("/").
	//         If user creation fails, set an alert message and render the "sessions/signup.tpl" template.
}

func (c *SessionsController) Login() {
	// TODO:
	// Step 1: Get the email and password from the request parameters.
	// Step 2: Authenticate the user by calling models.UserCheck(). This function returns the user if the user exists and password matches.
	//         If the user exists, set the session with name "user_id" and the user's ID (user.Id),
	//         redirect the user to the home page ("/").
	//         If authentication fails, set an alert message indicating "Email and/or password mismatch!" and render the "sessions/login.tpl" template.
}

func (c *SessionsController) Logout() {
	// TODO:
	// Clear the session by calling c.DelSession("user_id").
	// Redirect the user to the login page ("/login").
}

func GetUserFromSession(c *beego.Controller) *models.User {
	sessionID := c.GetSession("user_id")
	if sessionID == nil {
		return &models.User{}
	}
	userId := sessionID.(uint64)
	return models.UserFind(userId)
}
