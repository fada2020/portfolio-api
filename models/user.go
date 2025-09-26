package models

import "time"

// User represents a user in the system
type User struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe" binding:"required"`
	Email     string    `json:"email" example:"john@example.com" binding:"required,email"`
	Role      string    `json:"role" example:"developer" binding:"required"`
	Avatar    string    `json:"avatar,omitempty" example:"https://example.com/avatar.jpg"`
	Bio       string    `json:"bio,omitempty" example:"Full-stack developer with 5+ years experience"`
	Website   string    `json:"website,omitempty" example:"https://johndoe.dev"`
	Location  string    `json:"location,omitempty" example:"Seoul, Korea"`
	Skills    []string  `json:"skills,omitempty" example:"Go,JavaScript,React"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name     string   `json:"name" binding:"required" example:"John Doe"`
	Email    string   `json:"email" binding:"required,email" example:"john@example.com"`
	Role     string   `json:"role" binding:"required" example:"developer"`
	Avatar   string   `json:"avatar,omitempty" example:"https://example.com/avatar.jpg"`
	Bio      string   `json:"bio,omitempty" example:"Full-stack developer"`
	Website  string   `json:"website,omitempty" example:"https://johndoe.dev"`
	Location string   `json:"location,omitempty" example:"Seoul, Korea"`
	Skills   []string `json:"skills,omitempty" example:"Go,JavaScript"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name     *string   `json:"name,omitempty" example:"Jane Doe"`
	Email    *string   `json:"email,omitempty" example:"jane@example.com"`
	Role     *string   `json:"role,omitempty" example:"senior-developer"`
	Avatar   *string   `json:"avatar,omitempty" example:"https://example.com/new-avatar.jpg"`
	Bio      *string   `json:"bio,omitempty" example:"Senior full-stack developer"`
	Website  *string   `json:"website,omitempty" example:"https://janedoe.dev"`
	Location *string   `json:"location,omitempty" example:"Busan, Korea"`
	Skills   *[]string `json:"skills,omitempty" example:"Go,JavaScript,React,Vue"`
}