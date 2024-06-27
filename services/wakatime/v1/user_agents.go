package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type UserAgentsRequest struct {
	*request.BaseRequest
	User    *string `json:"user_id"`
	Current *bool   `json:"-"`
}

func NewUserAgentsRequest() *UserAgentsRequest {
	return &UserAgentsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) UserAgents(req *UserAgentsRequest) (*UserAgentsResponse, error) {
	if req == nil {
		req = NewUserAgentsRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/user_agents"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user_id}/user_agents"
	}

	resp := NewUserAgentResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewUserAgentResponse() *UserAgentsResponse {
	return &UserAgentsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type UserAgentsResponse struct {
	*response.BaseResponse
	Data []UserAgent `json:"data"`
}

type UserAgent struct {
	ID                 string    `json:"id"`
	Value              string    `json:"value"`
	Editor             string    `json:"editor"`
	Version            string    `json:"version"`
	OS                 string    `json:"os"`
	LastSeenAt         time.Time `json:"last_seen_at"`
	IsBrowserExtension bool      `json:"is_browser_extension"`
	IsDesktopApp       bool      `json:"is_desktop_app"`
	CreatedAt          time.Time `json:"created_at"`
}
