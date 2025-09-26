package models

import "time"

// Project represents a portfolio project
type Project struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Portfolio Website" binding:"required"`
	Description string    `json:"description" example:"A responsive portfolio website built with Flutter" binding:"required"`
	TechStack   []string  `json:"tech_stack" example:"Flutter,Dart,GitHub Pages"`
	Status      string    `json:"status" example:"completed" binding:"required"`
	Featured    bool      `json:"featured" example:"true"`
	LiveURL     string    `json:"live_url,omitempty" example:"https://johndoe.github.io/portfolio"`
	GithubURL   string    `json:"github_url,omitempty" example:"https://github.com/johndoe/portfolio"`
	ImageURL    string    `json:"image_url,omitempty" example:"https://example.com/project-image.jpg"`
	StartDate   time.Time `json:"start_date" example:"2024-01-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-02-01T00:00:00Z"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// CreateProjectRequest represents the request body for creating a project
type CreateProjectRequest struct {
	Title       string     `json:"title" binding:"required" example:"New Project"`
	Description string     `json:"description" binding:"required" example:"Project description"`
	TechStack   []string   `json:"tech_stack" example:"Go,React,PostgreSQL"`
	Status      string     `json:"status" binding:"required" example:"in-progress"`
	Featured    bool       `json:"featured" example:"false"`
	LiveURL     string     `json:"live_url,omitempty" example:"https://example.com"`
	GithubURL   string     `json:"github_url,omitempty" example:"https://github.com/user/repo"`
	ImageURL    string     `json:"image_url,omitempty" example:"https://example.com/image.jpg"`
	StartDate   time.Time  `json:"start_date" example:"2024-01-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-02-01T00:00:00Z"`
}

// UpdateProjectRequest represents the request body for updating a project
type UpdateProjectRequest struct {
	Title       *string    `json:"title,omitempty" example:"Updated Project Title"`
	Description *string    `json:"description,omitempty" example:"Updated description"`
	TechStack   *[]string  `json:"tech_stack,omitempty" example:"Go,React,PostgreSQL,Docker"`
	Status      *string    `json:"status,omitempty" example:"completed"`
	Featured    *bool      `json:"featured,omitempty" example:"true"`
	LiveURL     *string    `json:"live_url,omitempty" example:"https://updated.example.com"`
	GithubURL   *string    `json:"github_url,omitempty" example:"https://github.com/user/updated-repo"`
	ImageURL    *string    `json:"image_url,omitempty" example:"https://example.com/updated-image.jpg"`
	StartDate   *time.Time `json:"start_date,omitempty" example:"2024-01-15T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2024-03-01T00:00:00Z"`
}