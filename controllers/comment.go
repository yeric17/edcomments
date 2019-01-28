package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yeric17/edcomments/commons"
	"github.com/yeric17/edcomments/configuration"
	"github.com/yeric17/edcomments/models"
)

//CommentCreate permite registrar un comentario
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con éxito"
	commons.DisplayMessage(w, m)
}
