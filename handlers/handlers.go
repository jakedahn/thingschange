package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jakedahn/thingschange/models"
)

func V0GetCheckList(w http.ResponseWriter, req *http.Request) {
	api_key := req.URL.Query().Get("api_key")
	if api_key != "" {
		log.Printf("Getting list of checks for %s", api_key)
	}

	// checks := models.GetUsersChecks(api_key) => models.CheckList
	models.GetChecksByOwner(api_key)
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

func V0PostCheck(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	api_key := req.PostForm.Get("api_key")
	check_url := req.PostForm.Get("url")

	if api_key == "" {
		http.Error(w, "forbdden", http.StatusForbidden)
		return
	}

	if check_url == "" {
		http.Error(w, "Incorrect url", http.StatusNotAcceptable)
		return
	}

	check := models.Check{
		Owner:      api_key,
		Url:        check_url,
		Created_at: time.Now().UnixNano(),
		Updated_at: time.Now().UnixNano(),
	}

	check.Save()

	res, err := json.Marshal(check)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return
}
