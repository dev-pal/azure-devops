package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-pal/azure-devops-pullrequest/pkg/api/rest"
	"github.com/dev-pal/azure-devops-pullrequest/pkg/listing"
)

func main() {
	var lister listing.Service = listing.NewService()

	router := rest.Handler(lister)

	fmt.Println("The Azure DevOps PullRequest server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
