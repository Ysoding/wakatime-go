package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type SummariesRequest struct {
	*request.BaseRequest

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	Current    *bool                 `json:"-"`
	User       *string               `json:"user_id,omitempty"`
	Project    *string               `json:"project,omitempty"`
	Branches   *string               `json:"branches,omitempty"`
	Timeout    *int                  `json:"timeout,omitempty"`
	WritesOnly *bool                 `json:"writes_only,omitempty"`
	Timezone   *string               `json:"timezone,omitempty"`
	Range      *SummariesRangeOption `json:"range,omitempty"`
}

func NewSummariesRequest(start time.Time, end time.Time) *SummariesRequest {
	return &SummariesRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
		},
		Current: sdk.Bool(true),
		Start:   start,
		End:     end,
	}
}

func (c *Client) Summaries(req *SummariesRequest) (*SummariesResponse, error) {
	if req == nil {
		req = NewSummariesRequest(time.Now().AddDate(0, 0, -7), time.Now())
	}

	if req.Current != nil && *req.Current {
		req.BaseRequest.URL = "users/current/summaries"
	}

	if req.User != nil {
		req.BaseRequest.URL = "users/{user_id}/summaries"
	}

	resp := NewSummariesResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewSummariesResponse() *SummariesResponse {
	return &SummariesResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

// SummaryResponse represents the JSON structure for the summary data.
type SummariesResponse struct {
	*response.BaseResponse
	Data            []SummaryData   `json:"data"`
	CumulativeTotal CumulativeTotal `json:"cumulative_total"`
	DailyAverage    DailyAverage    `json:"daily_average"`
	Start           time.Time       `json:"start"`
	End             time.Time       `json:"end"`
}
