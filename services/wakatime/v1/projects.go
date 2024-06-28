package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type ProjectsRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`

	Q *string `json:"q,omitempty"`
}

func NewProjectsRequest() *ProjectsRequest {
	return &ProjectsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) Projects(req *ProjectsRequest) (*ProjectsResponse, error) {
	if req == nil {
		req = NewProjectsRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/projects"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/projects"
	}

	resp := NewProjectsResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewProjectsResponse() *ProjectsResponse {
	return &ProjectsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type ProjectsResponse struct {
	*response.BaseResponse
	Data []ProjectDetail `json:"data"`
}
