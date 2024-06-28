package v1

import "time"

type StatsData struct {
	TotalSeconds                       float64           `json:"total_seconds"`
	TotalSecondsIncludingOtherLanguage float64           `json:"total_seconds_including_other_language"`
	HumanReadableTotal                 string            `json:"human_readable_total"`
	HumanReadableTotalIncludingOther   string            `json:"human_readable_total_including_other_language"`
	DailyAverage                       float64           `json:"daily_average"`
	DailyAverageIncludingOtherLanguage float64           `json:"daily_average_including_other_language"`
	HumanReadableDailyAverage          string            `json:"human_readable_daily_average"`
	HumanReadableDailyAverageIncluding string            `json:"human_readable_daily_average_including_other_language"`
	Categories                         []Category        `json:"categories"`
	Projects                           []Project         `json:"projects"`
	Languages                          []Language        `json:"languages"`
	Editors                            []Editor          `json:"editors"`
	OperatingSystems                   []OperatingSystem `json:"operating_systems"`
	Dependencies                       []Dependency      `json:"dependencies"`
	Machines                           []Machine         `json:"machines"`
	BestDay                            BestDay           `json:"best_day"`
	Range                              string            `json:"range"`
	HumanReadableRange                 string            `json:"human_readable_range"`
	Holidays                           int               `json:"holidays"`
	DaysIncludingHolidays              int               `json:"days_including_holidays"`
	DaysMinusHolidays                  int               `json:"days_minus_holidays"`
	Status                             string            `json:"status"`
	PercentCalculated                  int               `json:"percent_calculated"`
	IsAlreadyUpdating                  bool              `json:"is_already_updating"`
	IsCodingActivityVisible            bool              `json:"is_coding_activity_visible"`
	IsOtherUsageVisible                bool              `json:"is_other_usage_visible"`
	IsStuck                            bool              `json:"is_stuck"`
	IsIncludingToday                   bool              `json:"is_including_today"`
	IsUpToDate                         bool              `json:"is_up_to_date"`
	Start                              string            `json:"start"`
	End                                string            `json:"end"`
	Timezone                           string            `json:"timezone"`
	Timeout                            int               `json:"timeout"`
	WritesOnly                         bool              `json:"writes_only"`
	UserID                             string            `json:"user_id"`
	Username                           string            `json:"username"`
	CreatedAt                          time.Time         `json:"created_at"`
	ModifiedAt                         time.Time         `json:"modified_at"`
}

type BestDay struct {
	Date         string  `json:"date"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

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

type AverageData struct {
	Seconds float64 `json:"seconds"`
	Text    string  `json:"text"`
}

type CountData struct {
	Text string `json:"text"`
}

type StatsBase struct {
	Average AverageData `json:"average"`
	Count   CountData   `json:"count"`
	Max     AverageData `json:"max"`
	Median  AverageData `json:"median"`
	Sum     AverageData `json:"sum"`
}

type CategoryStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type DailyAverageStats struct {
	StatsBase
}

type EditorStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type LanguageStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type OperatingSystemStats struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
	StatsBase
}

type TotalStats struct {
	StatsBase
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

type SummaryBase struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
	Digital      string  `json:"digital"`
	Text         string  `json:"text"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Seconds      int     `json:"seconds"`
}

type Category SummaryBase
type Project SummaryBase
type Language SummaryBase
type Editor SummaryBase
type OperatingSystem SummaryBase
type Dependency SummaryBase
type Machine struct {
	SummaryBase
	MachineNameID string `json:"machine_name_id"`
}

type Branch SummaryBase

type Entity SummaryBase

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
