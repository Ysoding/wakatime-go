package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type StatsRequest struct {
	*request.BaseRequest
	Range   *string `json:"range,omitempty"`
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`

	Timeout    *int  `json:"timeout,omitempty"`
	WritesOnly *bool `json:"writes_only,omitempty"`
}

func NewStatsRequest(rng string) *StatsRequest {
	return &StatsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Range:   sdk.String(rng),
		Current: sdk.Bool(true),
	}
}

func (c *Client) Stats(req *StatsRequest) (*StatsResponse, error) {
	if req == nil {
		req = NewStatsRequest("last_7_days")
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/stats"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/stats"
	}

	if req.Range != nil {
		req.BaseRequest.URL += "/{range}"
	}

	resp := NewStatsResponse()
	err := c.Send(req, resp)
	return resp, err
}

func NewStatsResponse() *StatsResponse {
	return &StatsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type StatsResponse struct {
	*response.BaseResponse
	Data StatsData `json:"data"`
}
