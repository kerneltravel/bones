package handlers

import (
	"bones/entities"
	"bones/repositories"
	"bones/web/actions"
	"bones/web/forms"
	"net/http"
	"strconv"
)

type ProfileContext struct {
	*BaseContext
	User *entities.User
}

func LoadLoginPage(res http.ResponseWriter, req *http.Request) {
	actions.RenderPage(res, newLoginContext())
}

func CreateNewSession(res http.ResponseWriter, req *http.Request) {
	form := &forms.LoginForm{Request: req}
	err := actions.ProcessForm(req, form)

	if err != nil {
		actions.RenderPageWithErrors(res, newLoginContext(), err)

		return
	}

	url := routeURL("userProfile", "id", strconv.Itoa(form.User.Id))
	http.Redirect(res, req, url, 302)
}

func LoadUserProfilePage(res http.ResponseWriter, req *http.Request) {
	id, err := pathParamInt(req, "id")

	if err != nil {
		actions.Render404(res, req)

		return
	}

	user, err := repositories.Users.FindById(id)

	if err != nil {
		actions.Render404(res, req)

		return
	}

	ctx := ProfileContext{newBaseContext("profile.html"), user}
	actions.RenderPage(res, &ctx)
}

func newLoginContext() *BaseContext {
	return newBaseContext("login.html")
}
