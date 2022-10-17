package controllers

import (
	"net/http"

	"github.com/abdulhanananwari/golang_example/src/infrastructure/utils/wrapper"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	wrapper.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
