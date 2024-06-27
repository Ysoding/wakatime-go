package v1

import (
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

//

type RangeOption string

const (
	RangeToday                  RangeOption = "Today"
	RangeYesterday              RangeOption = "Yesterday"
	RangeLast7Days              RangeOption = "Last 7 Days"
	RangeLast7DaysFromYesterday RangeOption = "Last 7 Days from Yesterday"
	RangeLast14Days             RangeOption = "Last 14 Days"
	RangeLast30Days             RangeOption = "Last 30 Days"
	RangeThisWeek               RangeOption = "This Week"
	RangeLastWeek               RangeOption = "Last Week"
	RangeThisMonth              RangeOption = "This Month"
	RangeLastMonth              RangeOption = "Last Month"
)

type SummariesRequest struct {
	*request.BaseRequest

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	Current    *bool        `json:"-"`
	User       *string      `json:"user_id,omitempty"`
	Project    *string      `json:"project,omitempty"`
	Branches   *string      `json:"branches,omitempty"`
	Timeout    *int         `json:"timeout,omitempty"`
	WritesOnly *bool        `json:"writes_only,omitempty"`
	Timezone   *string      `json:"timezone,omitempty"`
	Range      *RangeOption `json:"range,omitempty"`
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

// SummaryData represents each item in the 'data' array.
type SummaryData struct {
	GrandTotal       GrandTotal        `json:"grand_total"`
	Categories       []Category        `json:"categories"`
	Projects         []Project         `json:"projects"`
	Languages        []Language        `json:"languages"`
	Editors          []Editor          `json:"editors"`
	OperatingSystems []OperatingSystem `json:"operating_systems"`
	Dependencies     []Dependency      `json:"dependencies"`
	Machines         []Machine         `json:"machines"`
	Branches         []Branch          `json:"branches,omitempty"`
	Entities         []Entity          `json:"entities,omitempty"`
	Range            SummaryRange      `json:"range"`
}

// GrandTotal represents the 'grand_total' object.
type GrandTotal struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

// Category represents the 'categories' array item.
type Category struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
}

// Project represents the 'projects' array item.
type Project struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
}

// Language represents the 'languages' array item.
type Language struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// Editor represents the 'editors' array item.
type Editor struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// OperatingSystem represents the 'operating_systems' array item.
type OperatingSystem struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// Dependency represents the 'dependencies' array item.
type Dependency struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// Machine represents the 'machines' array item.
type Machine struct {
	Name          string  `json:"name"`
	MachineNameID string  `json:"machine_name_id"`
	TotalSeconds  float64 `json:"total_seconds"`
	Percent       float64 `json:"percent"`
	Digital       string  `json:"digital"`
	Text          string  `json:"text"`
	Hours         int     `json:"hours"`
	Minutes       int     `json:"minutes"`
	Seconds       int     `json:"seconds"`
}

// Branch represents the 'branches' array item.
type Branch struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// Entity represents the 'entities' array item.
type Entity struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

// SummaryRange represents the 'range' object.
type SummaryRange struct {
	Date     string `json:"date"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

// CumulativeTotal represents the 'cumulative_total' object.
type CumulativeTotal struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
	Decimal string  `json:"decimal"`
	Digital string  `json:"digital"`
}

// DailyAverage represents the 'daily_average' object.
type DailyAverage struct {
	Holidays                  int     `json:"holidays"`
	DaysIncludingHolidays     int     `json:"days_including_holidays"`
	DaysMinusHolidays         int     `json:"days_minus_holidays"`
	Seconds                   float64 `json:"seconds"`
	Text                      string  `json:"text"`
	SecondsIncludingOtherLang float64 `json:"seconds_including_other_language"`
	TextIncludingOtherLang    string  `json:"text_including_other_language"`
}
