package handlers

import (
	"context"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

func (env *Env) HeadersMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, req)
	})
}

func (env *Env) AppMiddleware(h AppHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if e := h(w, r); e != nil {
			http.Error(w, e.Message, e.Code)
		}
	})
}

func (env *Env) AuthenticateMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err             error
			claimMail       interface{}
			userMail        string
			claims          jwt.MapClaims
			ok              bool
			claimId         interface{}
			userId          string
			claimRoleId     interface{}
			roleId          string
			requestToken    *http.Cookie
			requestTokenStr string
			token           *jwt.Token
		)

		// Token regex
		tokenRegex := regexp.MustCompile("token=.+")

		// Extracting the token from the header
		if requestToken, err = r.Cookie("token"); err != nil {
			globals.Log.Debug("Token was not found in cookies")
			http.Error(w, "Token not found in cookies", http.StatusUnauthorized)
			return
		}

		// Verifying the token
		if !tokenRegex.MatchString(requestToken.String()) {
			globals.Log.Debug("The token has an invalid format")
			http.Error(w, "Token has an invalid format", http.StatusUnauthorized)
		}

		splitToken := strings.Split(requestToken.String(), "token=")
		requestTokenStr = splitToken[1]
		if token, err = jwt.Parse(requestTokenStr, func(token *jwt.Token) (interface{}, error) {
			return globals.TokenSignKey, nil
		}); err != nil {
			http.Error(w, "Token signature is invalid", http.StatusUnauthorized)
			return
		}

		// Extracting the claims
		if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Claim the email
			if claimMail, ok = claims["mail"]; !ok {
				globals.Log.Debug("Did not find Mail in claims")
				http.Error(w, "Mail not found in claims", http.StatusBadRequest)
				return
			}
			userMail = claimMail.(string)

			// Claim the id
			if claimId, ok = claims["user_id"]; !ok {
				globals.Log.Debug("Did not find UserId in claims")
				http.Error(w, "UserId not found in claims", http.StatusBadRequest)
				return
			}
			userId = strconv.FormatFloat(claimId.(float64), 'f', 0, 64)

			// Claim the role id
			if claimRoleId, ok = claims["role_id"]; !ok {
				globals.Log.Debug("Did not find RoleId in claims")
				http.Error(w, "RoleId not found in claims", http.StatusBadRequest)
				return
			}
			roleId = strconv.FormatFloat(claimRoleId.(float64), 'f', 0, 64)

		} else {
			globals.Log.Debug("Can not extract claims")
			http.Error(w, "Could not extract claims", http.StatusBadRequest)
			return
		}

		// Setting up context data
		values := map[string]string{
			"user_id": userId,
			"role_id": roleId,
			"mail":    userMail,
		}

		ctx := context.WithValue(r.Context(), "UserData", values)

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (env *Env) AuthorizeMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			err      error
			roleId   int
			userRole model.Role
			item     string
			goal     string
			itemId   int
		)

		// Extracting data from the context
		ctx := r.Context()
		userData := ctx.Value("UserData").(map[string]string)

		// Extracting context data
		if roleId, err = strconv.Atoi(userData["role_id"]); err != nil {
			globals.Log.Debug("Could not convert RoleId from string to int")
			http.Error(w, "Atoi conversion error", http.StatusBadRequest)
			return
		}

		// Getting permissions from database
		if userRole, err = env.DB.GetRole(int64(roleId)); err != nil {
			globals.Log.Debug("Could not retrieved role information")
			http.Error(w, "GetRole error", http.StatusBadRequest)
			return
		}

		// Extracting request variables
		vars := mux.Vars(r)

		item = vars["item"]
		if itemId, err = strconv.Atoi(vars["id"]); err != nil {
			globals.Log.Debug("No id in the request")
		}
		goal = vars["goal"]

		// And now, checking for authorizations
		switch r.Method {
		case "DELETE":
			switch item {
			case "users":
				// Can't Delete a user if you don't have the right to
				if !userRole.CanAddAndModifyUsers {
					globals.Log.Debug("Current user can't add or modify users")
					http.Error(w, "Deleting a user is forbidden", http.StatusForbidden)
					return
				}
			case "projects":
				// Can't delete the vacation project
				if itemId == 1 {
					globals.Log.Debug("Can't Delete the vacation project")
					http.Error(w, "Can't delete the vacation project", http.StatusUnauthorized)
					return
				}

				// Can't delete a project if you don't have the right to add one
				if !userRole.CanAddProjects {
					globals.Log.Debug("Current user can't delete project")
					http.Error(w, "Deleting a project is forbidden", http.StatusForbidden)
					return
				}
			case "roles":
				// Can't delete one of the basic roles
				if itemId == 1 || itemId == 2 || itemId == 3 {
					globals.Log.Debug("Can't delete one of the basic roles")
					http.Error(w, "Can't delete one of the basic roles", http.StatusUnauthorized)
					return
				}
			}
		case "PATCH":
			switch item {
			case "users":
				// Can't Modify a user if you don't have the right to
				if !userRole.CanAddAndModifyUsers {
					globals.Log.Debug("Current user can't add or modify users")
					http.Error(w, "Updating a user is forbidden", http.StatusForbidden)
					return
				}
			case "projects":
				// Can't modify the vacation project
				if itemId == 1 {
					globals.Log.Debug("Can't Modify the vacation project")
					http.Error(w, "Can't Modify the vacation project", http.StatusUnauthorized)
					return
				}

			}
		case "PUT":
			switch item {
			case "users":
				// Can't Add a user if you don't have the right to
				if !userRole.CanAddAndModifyUsers {
					globals.Log.Debug("Current user can't add or modify users")
					http.Error(w, "Creating a user is forbidden", http.StatusForbidden)
					return
				}
			case "projects":
				// Can't add a project if you don't have the right to
				if !userRole.CanAddProjects {
					globals.Log.Debug("Current user can't add new projects")
					http.Error(w, "Creating a new project is forbidden", http.StatusForbidden)
					return
				}
			}
		case "GET":
			switch item {
			case "users":
				// Can't see other scedules if you don't have the right
				if goal == "schedules" && !userRole.CanSeeOtherSchedules {
					globals.Log.Debug("Current user can't see other schedules")
					http.Error(w, "Getting other schedules is forbidden", http.StatusForbidden)
				}
			}
		}

		h.ServeHTTP(w, r)
	})
}
