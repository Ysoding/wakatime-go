package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type StatusBarRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`
}

func NewStatusBarRequest() *StatusBarRequest {
	return &StatusBarRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) StatusBar(req *StatusBarRequest) (*StatusBarResponse, error) {
	if req == nil {
		req = NewStatusBarRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/status_bar/today"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/status_bar/today"
	}

	resp := NewStatusBarResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewStatusBarResponse() *StatusBarResponse {
	return &StatusBarResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type StatusBarResponse struct {
	*response.BaseResponse
	Data            SummaryData `json:"data"`
	CachedAt        time.Time   `json:"cached_at"`
	HasTeamFeatures bool        `json:"has_team_features"`
}
