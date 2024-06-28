package v1

import (
	"fmt"
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type OrgDashboardMemberSummariesRequest struct {
	*request.BaseRequest
	Org       string  `json:"org"`
	Dashboard string  `json:"dashboard"`
	Member    string  `json:"member"`
	User      *string `json:"user,omitempty"`
	Current   *bool   `json:"-"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	Project  *string               `json:"project,omitempty"`
	Branches *string               `json:"branches,omitempty"`
	Range    *SummariesRangeOption `json:"range,omitempty"`
}

func NewOrgDashboardMemberSummariesRequest(org, dashboard, member string, start, end time.Time) *OrgDashboardMemberSummariesRequest {
	return &OrgDashboardMemberSummariesRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Org:       org,
		Dashboard: dashboard,
		Member:    member,
		Current:   sdk.Bool(true),
		Start:     start,
		End:       end,
	}
}

func (c *Client) OrgDashboardMemberSummaries(req *OrgDashboardMemberSummariesRequest) (*OrgDashboardMemberSummariesResponse, error) {
	if req == nil {
		req = NewOrgDashboardMemberSummariesRequest("", "", "", time.Now().AddDate(0, 0, -7), time.Now())
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user}"
	}

	req.BaseRequest.URL = fmt.Sprintf("%s/orgs/%s/dashboards/%s/members/%s/summaries",
		req.BaseRequest.URL,
		req.Org,
		req.Dashboard,
		req.Member,
	)

	resp := NewOrgDashboardMemberSummariesResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewOrgDashboardMemberSummariesResponse() *OrgDashboardMemberSummariesResponse {
	return &OrgDashboardMemberSummariesResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type OrgDashboardMemberSummariesResponse struct {
	*response.BaseResponse

	Data                   []OrgDashboardMemberSummariesData `json:"data"`
	Start                  string                            `json:"start"`
	End                    string                            `json:"end"`
	DefaultPersonalPrivacy string                            `json:"default_personal_privacy"`
	CumulativeTotal        CumulativeTotal                   `json:"cumulative_total"`
	DailyAverage           DailyAverage                      `json:"daily_average"`
}
