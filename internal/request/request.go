package request

import (
	"net/http"
)

type StressRequest struct {
	request *http.Request
}

func NewStressRequest(url string) (*StressRequest, error) {
	req, error := http.NewRequest(http.MethodGet, url, nil)
	if error != nil {
		return nil, error
	}
	return &StressRequest{
		request: req,
	}, nil
}

func (sr *StressRequest) Execute() (int, error) {
	res, error := http.DefaultClient.Do(sr.request)
	if error != nil {
		return 0, error
	}
	return res.StatusCode, nil
}
