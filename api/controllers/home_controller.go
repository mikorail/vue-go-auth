package controllers

import (
	"net/http"

	"github.com/mikorail/go-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API OK")
}
