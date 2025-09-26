package models

import "time"

// ContactMessage represents a contact form submission
type ContactMessage struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe" binding:"required"`
	Email     string    `json:"email" example:"john@example.com" binding:"required,email"`
	Subject   string    `json:"subject" example:"Project Inquiry" binding:"required"`
	Message   string    `json:"message" example:"I would like to discuss a project opportunity" binding:"required"`
	Status    string    `json:"status" example:"unread"` // unread, read, replied
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	ReadAt    *time.Time `json:"read_at,omitempty" example:"2024-01-01T01:00:00Z"`
}

// ContactFormRequest represents the request body for contact form submission
type ContactFormRequest struct {
	Name    string `json:"name" binding:"required" example:"John Doe"`
	Email   string `json:"email" binding:"required,email" example:"john@example.com"`
	Subject string `json:"subject" binding:"required" example:"Project Inquiry"`
	Message string `json:"message" binding:"required" example:"I would like to discuss a project opportunity"`
}

// Skill represents a technical skill
type Skill struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Go" binding:"required"`
	Category    string `json:"category" example:"backend" binding:"required"`
	Level       string `json:"level" example:"expert"` // beginner, intermediate, advanced, expert
	YearsExp    int    `json:"years_exp" example:"3"`
	Featured    bool   `json:"featured" example:"true"`
	Icon        string `json:"icon,omitempty" example:"https://example.com/go-icon.svg"`
	Color       string `json:"color,omitempty" example:"#00ADD8"`
	Description string `json:"description,omitempty" example:"Backend development and microservices"`
}

// AddSkillRequest represents the request body for adding a skill
type AddSkillRequest struct {
	Name        string `json:"name" binding:"required" example:"Python"`
	Category    string `json:"category" binding:"required" example:"backend"`
	Level       string `json:"level" binding:"required" example:"advanced"`
	YearsExp    int    `json:"years_exp" example:"2"`
	Featured    bool   `json:"featured" example:"false"`
	Icon        string `json:"icon,omitempty" example:"https://example.com/python-icon.svg"`
	Color       string `json:"color,omitempty" example:"#3776AB"`
	Description string `json:"description,omitempty" example:"Data analysis and web development"`
}