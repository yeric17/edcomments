package migration

import (
	"github.com/yeric17/edcomments/configuration"
	"github.com/yeric17/edcomments/models"
)

//Migrate es una funcion para migrar la base de datos
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
