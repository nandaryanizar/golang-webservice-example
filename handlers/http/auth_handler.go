package http

import (
	"net/http"

	"github.com/nandaryanizar/golang-webservice-example/internal/app/logging"

	"github.com/sarulabs/di"

	"github.com/nandaryanizar/golang-webservice-example/internal/app/provider"
	"github.com/nandaryanizar/golang-webservice-example/services"

	"github.com/nandaryanizar/golang-webservice-example/entities"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/helpers"
)

// Token authenticate the user
func Token(w http.ResponseWriter, r *http.Request) {
	var input entities.User

	err := helpers.ReadJSONBody(r, &input)
	if err != nil {
		logging.Logger.Error(err.Error())
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Bad parameters",
		})
		return
	}

	user, err := di.Get(r, "user-service").(services.UserService).AuthenticateUser(input.Email, input.Password)
	if err != nil || user.ID == 0 {
		if err != nil {
			logging.Logger.Error(err.Error())
		}
		
		helpers.JSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error": "Invalid email and/or password",
		})
		return
	}

	if user.ID != 0 {
		token, _ := provider.GenerateToken(user.ID)
		helpers.JSONResponse(w, http.StatusOK, token)
	}
}
