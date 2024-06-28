package v1

import (
	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type OrgDashboardMembersRequest struct {
	*request.BaseRequest
	Org       string  `json:"org"`
	Dashboard string  `json:"dashboard"`
	User      *string `json:"user,omitempty"`
	Current   *bool   `json:"-"`
}

func NewOrgDashboardMembersRequest(org, dashboard string) *OrgDashboardMembersRequest {
	return &OrgDashboardMembersRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Org:       org,
		Dashboard: dashboard,
		Current:   sdk.Bool(true),
	}
}

func (c *Client) OrgDashboardMembers(req *OrgDashboardMembersRequest) (*OrgDashboardMembersResponse, error) {
	if req == nil {
		req = NewOrgDashboardMembersRequest("", "")
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}"
	}

	req.BaseRequest.URL += "%s/orgs/{org}/dashboards/{dashboard}/members"

	resp := NewOrgDashboardMembersResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewOrgDashboardMembersResponse() *OrgDashboardMembersResponse {
	return &OrgDashboardMembersResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type OrgDashboardMembersResponse struct {
	*response.BaseResponse

	Data       []DashboardMember `json:"data"`
	NextPage   int               `json:"next_page"`
	Page       int               `json:"page"`
	PrevPage   int               `json:"prev_page"`
	Total      int               `json:"total"`
	TotalPages int               `json:"total_pages"`
}
