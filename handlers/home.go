package handlers

import (
	"net/http"

	"github.com/kamran0812/gobase/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())

}
