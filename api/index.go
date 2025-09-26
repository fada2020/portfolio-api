package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handler is the main entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Route handling
	switch r.URL.Path {
	case "/health":
		handleHealth(w, r)
	case "/api/v1/users":
		handleUsers(w, r)
	case "/api/v1/projects":
		handleProjects(w, r)
	case "/api/v1/skills":
		handleSkills(w, r)
	default:
		handleNotFound(w, r)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"version":   "1.0.0",
		"platform":  "vercel",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	users := []map[string]interface{}{
		{
			"id":       "1",
			"name":     "이혁주",
			"email":    "hyoukjoo@example.com",
			"role":     "Full-Stack Developer",
			"bio":      "Backend engineer specializing in Go, TypeScript, and cloud architecture",
			"website":  "https://hyoukjoolee.github.io/portfolio",
			"location": "Seoul, Korea",
			"skills":   []string{"Go", "TypeScript", "Flutter", "AWS", "Docker", "PostgreSQL"},
			"is_public": true,
		},
	}

	response := map[string]interface{}{
		"data":  users,
		"count": len(users),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleProjects(w http.ResponseWriter, r *http.Request) {
	projects := []map[string]interface{}{
		{
			"id":          "1",
			"title":       "Portfolio API Server",
			"description": "Go REST API with Supabase integration deployed on Vercel",
			"tech_stack":  []string{"Go", "Supabase", "Vercel"},
			"status":      "completed",
			"featured":    true,
			"live_url":    "https://portfolio-api.vercel.app",
			"github_url":  "https://github.com/hyoukjoolee/portfolio-api",
			"is_public":   true,
		},
		{
			"id":          "2",
			"title":       "Flutter Portfolio",
			"description": "Responsive portfolio website with i18n support",
			"tech_stack":  []string{"Flutter", "Dart", "GitHub Pages"},
			"status":      "completed",
			"featured":    true,
			"live_url":    "https://hyoukjoolee.github.io/portfolio",
			"github_url":  "https://github.com/hyoukjoolee/portfolio",
			"is_public":   true,
		},
	}

	response := map[string]interface{}{
		"data":  projects,
		"count": len(projects),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleSkills(w http.ResponseWriter, r *http.Request) {
	skills := []map[string]interface{}{
		{"id": "1", "name": "Go", "category": "backend", "level": "expert", "years_exp": 3, "featured": true, "color": "#00ADD8"},
		{"id": "2", "name": "TypeScript", "category": "frontend", "level": "advanced", "years_exp": 4, "featured": true, "color": "#3178C6"},
		{"id": "3", "name": "Flutter", "category": "mobile", "level": "advanced", "years_exp": 2, "featured": true, "color": "#02569B"},
		{"id": "4", "name": "Supabase", "category": "database", "level": "intermediate", "years_exp": 1, "featured": true, "color": "#3ECF8E"},
	}

	response := map[string]interface{}{
		"data":  skills,
		"count": len(skills),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	response := map[string]interface{}{
		"error": "Endpoint not found",
		"path":  r.URL.Path,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}