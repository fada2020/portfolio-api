package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"portfolio-api/database"
	"portfolio-api/models"
)

// Handler function aliases for better organization
var (
	// Project handlers
	GetProjects   = getProjects
	GetProject    = getProject
	CreateProject = createProject
	UpdateProject = updateProject
	DeleteProject = deleteProject

	// Skill handlers
	GetSkills   = getSkills
	AddSkill    = addSkill
	RemoveSkill = removeSkill

	// Contact handlers
	SubmitContactForm = submitContactForm

	// Stats handlers
	GetViewStats    = getViewStats
	GetProjectStats = getProjectStats
	RecordVisit     = recordVisit
)

// Mock data for projects
var projects = []models.Project{
	{
		ID:          1,
		Title:       "Portfolio Website",
		Description: "A responsive portfolio website built with Flutter Web, featuring i18n support and GitHub Pages deployment",
		TechStack:   []string{"Flutter", "Dart", "GitHub Actions", "GitHub Pages"},
		Status:      "completed",
		Featured:    true,
		LiveURL:     "https://hyoukjoolee.github.io/portfolio",
		GithubURL:   "https://github.com/hyoukjoolee/portfolio",
		ImageURL:    "https://via.placeholder.com/600x400",
		StartDate:   time.Now().Add(-time.Hour * 24 * 60),
		EndDate:     timePtr(time.Now().Add(-time.Hour * 24 * 30)),
		CreatedAt:   time.Now().Add(-time.Hour * 24 * 30),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "Go REST API",
		Description: "RESTful API server built with Go and Gin framework, featuring Swagger documentation and Docker deployment",
		TechStack:   []string{"Go", "Gin", "Docker", "AWS", "Swagger"},
		Status:      "in-progress",
		Featured:    true,
		GithubURL:   "https://github.com/hyoukjoolee/portfolio-api",
		ImageURL:    "https://via.placeholder.com/600x400",
		StartDate:   time.Now().Add(-time.Hour * 24 * 7),
		CreatedAt:   time.Now().Add(-time.Hour * 24 * 7),
		UpdatedAt:   time.Now(),
	},
}

// Mock data for skills
var skills = []models.Skill{
	{ID: 1, Name: "Go", Category: "backend", Level: "expert", YearsExp: 3, Featured: true, Color: "#00ADD8"},
	{ID: 2, Name: "JavaScript", Category: "frontend", Level: "expert", YearsExp: 5, Featured: true, Color: "#F7DF1E"},
	{ID: 3, Name: "TypeScript", Category: "frontend", Level: "advanced", YearsExp: 3, Featured: true, Color: "#3178C6"},
	{ID: 4, Name: "Flutter", Category: "mobile", Level: "advanced", YearsExp: 2, Featured: true, Color: "#02569B"},
	{ID: 5, Name: "Docker", Category: "devops", Level: "advanced", YearsExp: 3, Featured: false, Color: "#2496ED"},
	{ID: 6, Name: "AWS", Category: "cloud", Level: "intermediate", YearsExp: 2, Featured: false, Color: "#FF9900"},
}

// Mock data for contacts
var contacts = []models.ContactMessage{}

// Mock data for visits
var visits = []models.Visit{}

// Helper function to create time pointer
func timePtr(t time.Time) *time.Time {
	return &t
}

// Project handlers
func getProjects(c *gin.Context) {
	status := c.Query("status")
	featured := c.Query("featured")

	filteredProjects := projects

	if status != "" {
		var filtered []models.Project
		for _, p := range projects {
			if p.Status == status {
				filtered = append(filtered, p)
			}
		}
		filteredProjects = filtered
	}

	if featured == "true" {
		var filtered []models.Project
		for _, p := range filteredProjects {
			if p.Featured {
				filtered = append(filtered, p)
			}
		}
		filteredProjects = filtered
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  filteredProjects,
		"count": len(filteredProjects),
	})
}

func getProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	for _, project := range projects {
		if project.ID == id {
			c.JSON(http.StatusOK, project)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func createProject(c *gin.Context) {
	var req models.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProject := models.Project{
		ID:          len(projects) + 1,
		Title:       req.Title,
		Description: req.Description,
		TechStack:   req.TechStack,
		Status:      req.Status,
		Featured:    req.Featured,
		LiveURL:     req.LiveURL,
		GithubURL:   req.GithubURL,
		ImageURL:    req.ImageURL,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	projects = append(projects, newProject)
	c.JSON(http.StatusCreated, newProject)
}

func updateProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req models.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, project := range projects {
		if project.ID == id {
			if req.Title != nil {
				projects[i].Title = *req.Title
			}
			if req.Description != nil {
				projects[i].Description = *req.Description
			}
			if req.TechStack != nil {
				projects[i].TechStack = *req.TechStack
			}
			if req.Status != nil {
				projects[i].Status = *req.Status
			}
			if req.Featured != nil {
				projects[i].Featured = *req.Featured
			}
			if req.LiveURL != nil {
				projects[i].LiveURL = *req.LiveURL
			}
			if req.GithubURL != nil {
				projects[i].GithubURL = *req.GithubURL
			}
			if req.ImageURL != nil {
				projects[i].ImageURL = *req.ImageURL
			}
			if req.StartDate != nil {
				projects[i].StartDate = *req.StartDate
			}
			if req.EndDate != nil {
				projects[i].EndDate = req.EndDate
			}
			projects[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, projects[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

func deleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	for i, project := range projects {
		if project.ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
}

// Skill handlers
func getSkills(c *gin.Context) {
	category := c.Query("category")
	featured := c.Query("featured")

	filteredSkills := skills

	if category != "" {
		var filtered []models.Skill
		for _, s := range skills {
			if s.Category == category {
				filtered = append(filtered, s)
			}
		}
		filteredSkills = filtered
	}

	if featured == "true" {
		var filtered []models.Skill
		for _, s := range filteredSkills {
			if s.Featured {
				filtered = append(filtered, s)
			}
		}
		filteredSkills = filtered
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  filteredSkills,
		"count": len(filteredSkills),
	})
}

func addSkill(c *gin.Context) {
	var req models.AddSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSkill := models.Skill{
		ID:          len(skills) + 1,
		Name:        req.Name,
		Category:    req.Category,
		Level:       req.Level,
		YearsExp:    req.YearsExp,
		Featured:    req.Featured,
		Icon:        req.Icon,
		Color:       req.Color,
		Description: req.Description,
	}

	skills = append(skills, newSkill)
	c.JSON(http.StatusCreated, newSkill)
}

func removeSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	for i, skill := range skills {
		if skill.ID == id {
			skills = append(skills[:i], skills[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
}

// Contact handler
func submitContactForm(c *gin.Context) {
	var req models.ContactFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newContact := models.ContactMessage{
		ID:        len(contacts) + 1,
		Name:      req.Name,
		Email:     req.Email,
		Subject:   req.Subject,
		Message:   req.Message,
		Status:    "unread",
		CreatedAt: time.Now(),
	}

	contacts = append(contacts, newContact)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact form submitted successfully",
		"id":      newContact.ID,
	})
}

// Stats handlers
func getViewStats(c *gin.Context) {
	// Generate some mock statistics
	stats := models.ViewStats{
		TotalViews:     1250 + rand.Intn(100),
		UniqueVisitors: 950 + rand.Intn(50),
		ViewsToday:     45 + rand.Intn(20),
		ViewsThisWeek:  320 + rand.Intn(50),
		ViewsThisMonth: 1100 + rand.Intn(100),
		TopPages: []models.PageStat{
			{Page: "/projects", Views: 450},
			{Page: "/", Views: 320},
			{Page: "/resume", Views: 280},
			{Page: "/contact", Views: 200},
		},
		ViewsByCountry: []models.CountryStat{
			{Country: "South Korea", Views: 680},
			{Country: "United States", Views: 320},
			{Country: "Japan", Views: 180},
			{Country: "Others", Views: 70},
		},
		ViewsOverTime: generateTimeStats(),
	}

	c.JSON(http.StatusOK, stats)
}

func getProjectStats(c *gin.Context) {
	completed := 0
	featured := 0
	techCount := make(map[string]int)
	statusCount := make(map[string]int)

	for _, project := range projects {
		if project.Status == "completed" {
			completed++
		}
		if project.Featured {
			featured++
		}

		statusCount[project.Status]++

		for _, tech := range project.TechStack {
			techCount[tech]++
		}
	}

	var techStats []models.TechStackStat
	for tech, count := range techCount {
		percentage := float64(count) / float64(len(projects)) * 100
		techStats = append(techStats, models.TechStackStat{
			Technology: tech,
			Count:      count,
			Percentage: percentage,
		})
	}

	var statusStats []models.ProjectStatusStat
	for status, count := range statusCount {
		statusStats = append(statusStats, models.ProjectStatusStat{
			Status: status,
			Count:  count,
		})
	}

	stats := models.ProjectStats{
		TotalProjects:     len(projects),
		CompletedProjects: completed,
		FeaturedProjects:  featured,
		TechStackStats:    techStats,
		ProjectsByStatus:  statusStats,
	}

	c.JSON(http.StatusOK, stats)
}

func recordVisit(c *gin.Context) {
	var req models.VisitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newVisit := models.Visit{
		ID:        len(visits) + 1,
		Page:      req.Page,
		UserAgent: req.UserAgent,
		Country:   req.Country,
		Referrer:  req.Referrer,
		IP:        c.ClientIP(),
		CreatedAt: time.Now(),
	}

	visits = append(visits, newVisit)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Visit recorded successfully",
		"id":      newVisit.ID,
	})
}

// Helper function to generate time-based statistics
func generateTimeStats() []models.TimeStat {
	var stats []models.TimeStat
	now := time.Now()

	for i := 6; i >= 0; i-- {
		date := now.AddDate(0, 0, -i)
		views := 40 + rand.Intn(60)
		stats = append(stats, models.TimeStat{
			Date:  date.Format("2006-01-02"),
			Views: views,
		})
	}

	return stats
}