package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yeric17/edcomments/controllers"
)

//SetCommentRouter crea un comentario usando el controlador
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
