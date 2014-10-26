package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func V0GetLists(res http.ResponseWriter, req *http.Request) {
	return "Success"
}
