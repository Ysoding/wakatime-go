package v1

import (
	"fmt"
	"time"

	"github.com/Ysoding/wakatime-go/sdk"
	"github.com/Ysoding/wakatime-go/sdk/request"
	"github.com/Ysoding/wakatime-go/sdk/response"
)

type StatsAggregatedRequest struct {
	*request.BaseRequest
	Range *string `json:"range"`
}

func NewStatsAggregatedRequest(rng string) *StatsAggregatedRequest {
	return &StatsAggregatedRequest{
		BaseRequest: &request.BaseRequest{
			Version: "v1",
			Method:  sdk.MethodGet,
			URL:     "stats/{range}",
		},
		Range: sdk.String(rng),
	}
}

func (c *Client) StatsAggregated(req *StatsAggregatedRequest) (*StatsAggregatedResponse, error) {
	if req == nil {
		req = NewStatsAggregatedRequest(fmt.Sprintf("%d", time.Now().AddDate(-1, 0, 0).Year()))
	}

	resp := NewStatsAggregatedResponse()

	err := c.Send(req, resp)
	return resp, err
}

func NewStatsAggregatedResponse() *StatsAggregatedResponse {
	return &StatsAggregatedResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

type StatsAggregatedResponse struct {
	*response.BaseResponse
	Data StatsData `json:"data"`
}

type StatsData struct {
	Categories       []CategoryStats        `json:"categories"`
	DailyAverage     DailyAverageStats      `json:"daily_average"`
	Editors          []EditorStats          `json:"editors"`
	Languages        []LanguageStats        `json:"languages"`
	OperatingSystems []OperatingSystemStats `json:"operating_systems"`
	Total            TotalStats             `json:"total"`
	Range            SummaryRangeStats      `json:"range"`
	Timeout          int                    `json:"timeout"`
	WritesOnly       bool                   `json:"writes_only"`
}

// CategoryStats represents the 'categories' array item.
type CategoryStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	Average    struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
	Sum struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"sum"`
}

// DailyAverageStats represents the 'daily_average' object.
type DailyAverageStats struct {
	Average struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
}

// EditorStats represents the 'editors' array item.
type EditorStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	Average    struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
	Sum struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"sum"`
}

// LanguageStats represents the 'languages' array item.
type LanguageStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	Average    struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
	Sum struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"sum"`
}

// OperatingSystemStats represents the 'operating_systems' array item.
type OperatingSystemStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	Average    struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
	Sum struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"sum"`
}

// TotalStats represents the 'total' object.
type TotalStats struct {
	Average struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"average"`
	Count struct {
		Text string `json:"text"`
	} `json:"count"`
	Max struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"max"`
	Median struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"median"`
	Sum struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
	} `json:"sum"`
}

// SummaryRangeStats represents the 'range' object.
type SummaryRangeStats struct {
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	Text      string `json:"text"`
}
