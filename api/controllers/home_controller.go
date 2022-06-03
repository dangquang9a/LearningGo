package controllers

import (
	"net/http"

	"github.com/dangquang9a/LearningGo/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to API")
}
