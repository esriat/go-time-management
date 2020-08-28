package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
	"golang.org/x/crypto/bcrypt"
)

//	GetTokenHandler
/*	The handler called by the following endpoint : POST /get-token.
	This method takes the email adress and the password of the user in order to connect them.
	Uses bcrypt to compare the given password and the crypted database password.
	Is this method works, it sets up a cookie that contains the token.
*/
func (env *Env) GetTokenHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err          error
		databaseUser model.User
		formUser     model.User
	)

	// Parsing the form
	if err = r.ParseForm(); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when parsing the form",
			Code:    http.StatusBadRequest,
		}
	}

	// Decoding the form
	if err = json.NewDecoder(r.Body).Decode(&formUser); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.WithFields(logrus.Fields{"Form User : ": formUser}).Debug("GetTokenHandler")

	// Trying to identify the User
	if databaseUser, err = env.DB.GetUserFromEmail(formUser.Mail); err != nil {
		if err == sql.ErrNoRows {
			return &AppError{
				Code:    http.StatusUnauthorized,
				Error:   err,
				Message: "User does not exist",
			}
		}
		return &AppError{
			Code:    http.StatusInternalServerError,
			Error:   err,
			Message: "Error when getting the user",
		}
	}

	globals.Log.WithFields(logrus.Fields{"Databse User : ": databaseUser}).Debug("GetTokenHandler")

	// Trying to authenticate the User with his password
	if err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(formUser.Password)); err != nil {
		return &AppError{
			Code:    http.StatusUnauthorized,
			Error:   err,
			Message: "The password is incorrect",
		}
	}

	// Creating the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Initializing the claims and creating them
	claims := token.Claims.(jwt.MapClaims)
	claims["mail"] = databaseUser.Mail
	claims["user_id"] = databaseUser.UserId
	claims["role_id"] = databaseUser.RoleId
	claims["expiration"] = time.Now().Add(time.Hour * 8)

	// Sign the token with the Globals secret key
	tokenString, _ := token.SignedString(globals.TokenSignKey)

	// Write the token to the browser
	cookieToken := http.Cookie{
		Name:  "token",
		Value: tokenString,
	}

	http.SetCookie(w, &cookieToken)

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(tokenString)); err != nil {
		return &AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

//	DeleteTokenHandler
/*	The handler called by the following endpoint : POST /get-token.
	This method deletes the token cookie.
*/
func (env *Env) DeleteTokenHandler(w http.ResponseWriter, r *http.Request) *AppError {
	globals.Log.Debug("DeleteTokenHandler called")

	// Erase the cookies
	cookieToken := http.Cookie{
		Name:  "token",
		Value: "",
	}

	http.SetCookie(w, &cookieToken)

	w.WriteHeader(http.StatusOK)
	return nil
}
