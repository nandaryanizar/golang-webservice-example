package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/helpers"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/logging"
	"github.com/nandaryanizar/golang-webservice-example/services"
	"github.com/sarulabs/di"
)

// FindUser find specific user in database
func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logging.Logger.Error(err.Error())
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Bad parameters",
		})
		return
	}

	user, err := di.Get(r, "user-service").(services.UserService).FindUserByID(idInt)
	if err != nil || user.Email == "" {
		if err != nil {
			logging.Logger.Error(err.Error())
		}

		helpers.JSONResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error": "User not found",
		})
		return
	}

	helpers.JSONResponse(w, http.StatusOK, user)
}
