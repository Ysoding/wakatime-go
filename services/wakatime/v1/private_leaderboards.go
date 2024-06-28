package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type PrivateLeaderBoardRequest struct {
	*request.BaseRequest
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`
}

func NewPrivateLeaderBoardRequest() *PrivateLeaderBoardRequest {
	return &PrivateLeaderBoardRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
	}
}

func (c *Client) PrivateLeaderBoard(req *PrivateLeaderBoardRequest) (*PrivateLeaderBoardResponse, error) {
	if req == nil {
		req = NewPrivateLeaderBoardRequest()
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/leaderboards"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/leaderboards"
	}

	resp := NewPrivateLeaderBoardResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewPrivateLeaderBoardResponse() *PrivateLeaderBoardResponse {
	return &PrivateLeaderBoardResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type PrivateLeaderBoardResponse struct {
	*response.BaseResponse

	Data       []PrivateLeaderboard `json:"data"`
	Total      int                  `json:"total"`
	TotalPages int                  `json:"total_pages"`
}
