package controllers

import (
	"net/http"

	"github.com/mozarik/testMajoo/api_model/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hihi selamat datang")
}
