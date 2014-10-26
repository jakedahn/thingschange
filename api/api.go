package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jakedahn/thingschange/models"
)

func V0GetLists(w http.ResponseWriter, req *http.Request) {
	api_key := req.URL.Query().Get("api_key")
	if api_key != "" {
		log.Printf("Getting list of checks for %s", api_key)
	}

	// checks := models.GetUsersChecks(api_key) => models.CheckList
	checks := models.CheckList{}

	res, err := json.Marshal(checks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
