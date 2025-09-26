package models

import "time"

// ViewStats represents portfolio view statistics
type ViewStats struct {
	TotalViews    int            `json:"total_views" example:"1250"`
	UniqueVisitors int           `json:"unique_visitors" example:"950"`
	ViewsToday    int            `json:"views_today" example:"45"`
	ViewsThisWeek int            `json:"views_this_week" example:"320"`
	ViewsThisMonth int           `json:"views_this_month" example:"1100"`
	TopPages      []PageStat     `json:"top_pages"`
	ViewsByCountry []CountryStat `json:"views_by_country"`
	ViewsOverTime []TimeStat     `json:"views_over_time"`
}

// PageStat represents statistics for a specific page
type PageStat struct {
	Page  string `json:"page" example:"/projects"`
	Views int    `json:"views" example:"450"`
}

// CountryStat represents statistics by country
type CountryStat struct {
	Country string `json:"country" example:"South Korea"`
	Views   int    `json:"views" example:"680"`
}

// TimeStat represents views over time
type TimeStat struct {
	Date  string `json:"date" example:"2024-01-01"`
	Views int    `json:"views" example:"85"`
}

// ProjectStats represents project-related statistics
type ProjectStats struct {
	TotalProjects     int                `json:"total_projects" example:"12"`
	CompletedProjects int                `json:"completed_projects" example:"10"`
	FeaturedProjects  int                `json:"featured_projects" example:"5"`
	TechStackStats    []TechStackStat    `json:"tech_stack_stats"`
	ProjectsByStatus  []ProjectStatusStat `json:"projects_by_status"`
}

// TechStackStat represents statistics for technology usage
type TechStackStat struct {
	Technology string `json:"technology" example:"Go"`
	Count      int    `json:"count" example:"8"`
	Percentage float64 `json:"percentage" example:"66.7"`
}

// ProjectStatusStat represents project statistics by status
type ProjectStatusStat struct {
	Status string `json:"status" example:"completed"`
	Count  int    `json:"count" example:"10"`
}

// VisitRequest represents a visit tracking request
type VisitRequest struct {
	Page      string `json:"page" binding:"required" example:"/projects"`
	UserAgent string `json:"user_agent,omitempty" example:"Mozilla/5.0..."`
	Country   string `json:"country,omitempty" example:"KR"`
	Referrer  string `json:"referrer,omitempty" example:"https://google.com"`
}

// Visit represents a recorded visit
type Visit struct {
	ID        int       `json:"id" example:"1"`
	Page      string    `json:"page" example:"/projects"`
	UserAgent string    `json:"user_agent,omitempty"`
	Country   string    `json:"country,omitempty" example:"KR"`
	Referrer  string    `json:"referrer,omitempty"`
	IP        string    `json:"ip,omitempty"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
}