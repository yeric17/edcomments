package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yeric17/edcomments/commons"
	"github.com/yeric17/edcomments/configuration"
	"github.com/yeric17/edcomments/models"
)

//Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))

	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)

	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})

		if err != nil {
			log.Fatalf("Error al convertir el token a json %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message: "Usuario o clave no válido",
			Code:    http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}
