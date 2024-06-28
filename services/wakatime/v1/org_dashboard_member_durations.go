package v1

import (
	"fmt"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type OrgDashboardMemberDurationsRequest struct {
	*request.BaseRequest
	Org       string  `json:"org"`
	Dashboard string  `json:"dashboard"`
	Member    string  `json:"member"`
	User      *string `json:"user,omitempty"`
	Current   *bool   `json:"-"`

	Date     string  `json:"date"`
	Project  *string `json:"project,omitempty"`
	Branches *string `json:"branches,omitempty"`
}

func NewOrgDashboardMemberDurationsRequest(org, dashboard, member, date string) *OrgDashboardMemberDurationsRequest {
	return &OrgDashboardMemberDurationsRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Org:       org,
		Dashboard: dashboard,
		Member:    member,
		Current:   sdk.Bool(true),
		Date:      date,
	}
}

func (c *Client) OrgDashboardMemberDurations(req *OrgDashboardMemberDurationsRequest) (*OrgDashboardMemberDurationsResponse, error) {
	if req == nil {
		req = NewOrgDashboardMemberDurationsRequest("", "", "", "2024-06-28")
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}"
	}

	req.BaseRequest.URL = fmt.Sprintf("%s/orgs/%s/dashboards/%s/members/%s/durations",
		req.BaseRequest.URL,
		req.Org,
		req.Dashboard,
		req.Member,
	)

	resp := NewOrgDashboardMemberDurationsResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewOrgDashboardMemberDurationsResponse() *OrgDashboardMemberDurationsResponse {
	return &OrgDashboardMemberDurationsResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type OrgDashboardMemberDurationsResponse struct {
	*response.BaseResponse

	Data     []OrgDashboardMemberDurationsData `json:"data"`
	Start    string                            `json:"start"`
	End      string                            `json:"end"`
	Timezone string                            `json:"timezone"`
}
