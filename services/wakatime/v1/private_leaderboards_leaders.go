package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type PrivateLeaderBoardLeadersRequest struct {
	*request.BaseRequest
	Board   *string `json:"board,omitempty"`
	User    *string `json:"user,omitempty"`
	Current *bool   `json:"-"`

	Language    *string `json:"language,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Page        *int    `json:"page,omitempty"`
}

func NewPrivateLeaderBoardLeadersRequest(board string) *PrivateLeaderBoardLeadersRequest {
	return &PrivateLeaderBoardLeadersRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Board:   sdk.String(board),
		Current: sdk.Bool(true),
	}
}

func (c *Client) PrivateLeaderBoardLeaders(req *PrivateLeaderBoardLeadersRequest) (*PrivateLeaderBoardLeadersResponse, error) {
	if req == nil {
		req = NewPrivateLeaderBoardLeadersRequest("")
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/leaderboards"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}/leaderboards"
	}

	if req.Board != nil {
		req.BaseRequest.URL += "/{board}"
	}

	resp := NewPrivateLeaderBoardLeadersResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewPrivateLeaderBoardLeadersResponse() *PrivateLeaderBoardLeadersResponse {
	return &PrivateLeaderBoardLeadersResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type PrivateLeaderBoardLeadersResponse struct {
	*response.BaseResponse

	Data        []LeaderboardMember `json:"data"`
	Language    string              `json:"language"`
	CountryCode string              `json:"country_code"`
	ModifiedAt  time.Time           `json:"modified_at"`
	Page        int                 `json:"page"`
	TotalPages  int                 `json:"total_pages"`
	Range       LeaderboardRange    `json:"range"`
	Timeout     int                 `json:"timeout"`
	WritesOnly  bool                `json:"writes_only"`
}
