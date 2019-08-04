package listing

import "errors"

// Service provides pullrequest listing operations
type Service interface {
	GetPullRequest(int) (PullRequest, error)
	GetPullRequests() []PullRequest
}

type service struct {
}

// NewService creates a listing service with necessary dependencies
func NewService() Service {
	return &service{}
}

// GetPullRequest returns a pullrequest
func (s *service) GetPullRequest(id int) (PullRequest, error) {
	// TODO: implement GetPullRequest(id int)
	return PullRequest{}, errors.New("not implemented yet")
}

// GetPullRequests returns all pullrequests
func (s *service) GetPullRequests() []PullRequest {
	// TODO: implement GetPullRequest(id int)
	return nil
}
