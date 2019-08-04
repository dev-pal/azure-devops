package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dev-pal/azure-devops-pullrequest/pkg/listing"
	"github.com/julienschmidt/httprouter"
)

func Handler(l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/pullrequests", getPullRequests(l))
	router.GET("/pullrequests/:id", getPullRequest(l))

	return router
}

// getPullRequests returns a handler for GET /pullrequests requests
func getPullRequests(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetPullRequests()
		json.NewEncoder(w).Encode(list) // TODO: err handling
	}
}

// getPullRequest returns a handler for GET /pullrequests/:id requests
func getPullRequest(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ID, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s is not a valid pullrequest ID, it must be a number.", p.ByName("id")), http.StatusBadRequest)
			return
		}

		pullRequest, err := s.GetPullRequest(ID)
		if err != nil {
			http.Error(w, "failed to get pullrequest.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pullRequest) // TODO: err handling
	}
}
