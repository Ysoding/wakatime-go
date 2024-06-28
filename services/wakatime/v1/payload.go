package v1

import "time"

type TimeData struct {
	DailyAverage      float64   `json:"daily_average"`
	Decimal           string    `json:"decimal"`
	Digital           string    `json:"digital"`
	IsUpToDate        bool      `json:"is_up_to_date"`
	PercentCalculated int       `json:"percent_calculated"`
	Range             TimeRange `json:"range"`
	Text              string    `json:"text"`
	Timeout           int       `json:"timeout"`
	TotalSeconds      float64   `json:"total_seconds"`
}

type TimeRange struct {
	End       string `json:"end"`
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Start     string `json:"start"`
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	Timezone  string `json:"timezone"`
}

type StatsAggregatedData struct {
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

type SummaryRangeStats struct {
	EndDate   string `json:"end_date"`
	EndText   string `json:"end_text"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	StartText string `json:"start_text"`
	Text      string `json:"text"`
}

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

type GrandTotal struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Category struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
}

type Project struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
}

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

type SummaryRange struct {
	Date     string `json:"date"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Text     string `json:"text"`
	Timezone string `json:"timezone"`
}

type CumulativeTotal struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
	Decimal string  `json:"decimal"`
	Digital string  `json:"digital"`
}

type DailyAverage struct {
	Holidays                  int     `json:"holidays"`
	DaysIncludingHolidays     int     `json:"days_including_holidays"`
	DaysMinusHolidays         int     `json:"days_minus_holidays"`
	Seconds                   float64 `json:"seconds"`
	Text                      string  `json:"text"`
	SecondsIncludingOtherLang float64 `json:"seconds_including_other_language"`
	TextIncludingOtherLang    string  `json:"text_including_other_language"`
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

type User struct {
	ID                   string    `json:"id"`
	HasPremiumFeatures   bool      `json:"has_premium_features"`
	DisplayName          string    `json:"display_name"`
	FullName             string    `json:"full_name"`
	Email                string    `json:"email"`
	Photo                string    `json:"photo"`
	IsEmailPublic        bool      `json:"is_email_public"`
	IsEmailConfirmed     bool      `json:"is_email_confirmed"`
	PublicEmail          *string   `json:"public_email"` // Nullable
	PhotoPublic          bool      `json:"photo_public"`
	Timezone             string    `json:"timezone"`
	LastHeartbeatAt      time.Time `json:"last_heartbeat_at"`
	LastPlugin           string    `json:"last_plugin"`
	LastPluginName       string    `json:"last_plugin_name"`
	LastProject          string    `json:"last_project"`
	LastBranch           string    `json:"last_branch"`
	Plan                 string    `json:"plan"`
	Username             string    `json:"username"`
	Website              string    `json:"website"`
	HumanReadableWebsite string    `json:"human_readable_website"`
	WonderfulDevUsername string    `json:"wonderfuldev_username"`
	GitHubUsername       string    `json:"github_username"`
	TwitterUsername      string    `json:"twitter_username"`
	LinkedInUsername     string    `json:"linkedin_username"`
	City                 City      `json:"city"`
	LoggedTimePublic     bool      `json:"logged_time_public"`
	LanguagesUsedPublic  bool      `json:"languages_used_public"`
	IsHireable           bool      `json:"is_hireable"`
	CreatedAt            time.Time `json:"created_at"`
	ModifiedAt           time.Time `json:"modified_at"`
}

type City struct {
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Title       string `json:"title"`
}
