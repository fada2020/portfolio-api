package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"portfolio-api/database"
	"portfolio-api/models"
)

// GetUsers retrieves all users from Supabase
// @Summary Get all users
// @Description Get list of all users with optional filtering
// @Tags users
// @Produce json
// @Param is_public query boolean false "Filter by public profile"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetUsers(c *gin.Context) {
	isPublic := c.Query("is_public")

	query := "SELECT id, email, name, role, avatar_url, bio, website, location, skills, is_public, created_at, updated_at FROM users WHERE 1=1"
	args := []interface{}{}

	if isPublic == "true" {
		query += " AND is_public = $1"
		args = append(args, true)
	}

	query += " ORDER BY created_at DESC"

	rows, err := database.SupabaseDB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var skills string

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Role,
			&user.AvatarURL,
			&user.Bio,
			&user.Website,
			&user.Location,
			&skills,
			&user.IsPublic,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			continue
		}

		// Parse skills JSON array - simplified for now
		user.Skills = []string{}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"count": len(users),
	})
}

// GetUserByID retrieves a specific user by ID
// @Summary Get user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	query := "SELECT id, email, name, role, avatar_url, bio, website, location, skills, is_public, created_at, updated_at FROM users WHERE id = $1"

	var user models.User
	var skills string

	err := database.SupabaseDB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.Website,
		&user.Location,
		&skills,
		&user.IsPublic,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse skills JSON array - simplified for now
	user.Skills = []string{}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO users (email, name, role, avatar_url, bio, website, location, skills, is_public)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at`

	var user models.User
	err := database.SupabaseDB.QueryRow(
		query,
		req.Email,
		req.Name,
		req.Role,
		req.AvatarURL,
		req.Bio,
		req.Website,
		req.Location,
		"[]", // Empty skills array for now
		req.IsPublic,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Fill in the rest of the user data
	user.Email = req.Email
	user.Name = req.Name
	user.Role = req.Role
	user.AvatarURL = req.AvatarURL
	user.Bio = req.Bio
	user.Website = req.Website
	user.Location = req.Location
	user.Skills = []string{}
	user.IsPublic = req.IsPublic

	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user
// @Summary Update a user
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUserRequest true "User update data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build dynamic update query
	query := "UPDATE users SET updated_at = NOW()"
	args := []interface{}{}
	argIndex := 1

	if req.Name != nil {
		query += ", name = $" + strconv.Itoa(argIndex)
		args = append(args, *req.Name)
		argIndex++
	}
	if req.Role != nil {
		query += ", role = $" + strconv.Itoa(argIndex)
		args = append(args, *req.Role)
		argIndex++
	}
	if req.AvatarURL != nil {
		query += ", avatar_url = $" + strconv.Itoa(argIndex)
		args = append(args, *req.AvatarURL)
		argIndex++
	}
	if req.Bio != nil {
		query += ", bio = $" + strconv.Itoa(argIndex)
		args = append(args, *req.Bio)
		argIndex++
	}
	if req.Website != nil {
		query += ", website = $" + strconv.Itoa(argIndex)
		args = append(args, *req.Website)
		argIndex++
	}
	if req.Location != nil {
		query += ", location = $" + strconv.Itoa(argIndex)
		args = append(args, *req.Location)
		argIndex++
	}
	if req.IsPublic != nil {
		query += ", is_public = $" + strconv.Itoa(argIndex)
		args = append(args, *req.IsPublic)
		argIndex++
	}

	query += " WHERE id = $" + strconv.Itoa(argIndex) + " RETURNING id, email, name, role, avatar_url, bio, website, location, skills, is_public, created_at, updated_at"
	args = append(args, id)

	var user models.User
	var skills string

	err := database.SupabaseDB.QueryRow(query, args...).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.Website,
		&user.Location,
		&skills,
		&user.IsPublic,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse skills JSON array - simplified for now
	user.Skills = []string{}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user from the system
// @Tags users
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM users WHERE id = $1"
	result, err := database.SupabaseDB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.Status(http.StatusNoContent)
}