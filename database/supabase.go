package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var SupabaseDB *sql.DB

// InitSupabase initializes Supabase connection
func InitSupabase() error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY")
	supabaseDBURL := os.Getenv("SUPABASE_DB_URL")

	if supabaseDBURL == "" {
		return fmt.Errorf("SUPABASE_DB_URL environment variable is required")
	}

	var err error
	SupabaseDB, err = sql.Open("postgres", supabaseDBURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Supabase: %v", err)
	}

	// Test connection
	if err = SupabaseDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping Supabase: %v", err)
	}

	log.Printf("Successfully connected to Supabase: %s", supabaseURL)

	// Create tables
	if err = createSupabaseTables(); err != nil {
		return fmt.Errorf("failed to create Supabase tables: %v", err)
	}

	return nil
}

// createSupabaseTables creates all tables in Supabase
func createSupabaseTables() error {
	// Enable Row Level Security (RLS) by default
	tables := []string{
		// Users table with RLS
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			email VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			role VARCHAR(100) NOT NULL DEFAULT 'user',
			avatar_url TEXT,
			bio TEXT,
			website TEXT,
			location VARCHAR(255),
			skills JSONB DEFAULT '[]',
			is_public BOOLEAN DEFAULT true,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,

		// Enable RLS on users
		`ALTER TABLE users ENABLE ROW LEVEL SECURITY;`,

		// Policy: Users can read public profiles
		`CREATE POLICY "Public profiles are viewable by everyone"
		 ON users FOR SELECT
		 USING (is_public = true);`,

		// Policy: Users can update their own profile
		`CREATE POLICY "Users can update own profile"
		 ON users FOR UPDATE
		 USING (auth.uid() = id);`,

		// Projects table
		`CREATE TABLE IF NOT EXISTS projects (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			tech_stack JSONB DEFAULT '[]',
			status VARCHAR(50) NOT NULL DEFAULT 'in-progress',
			featured BOOLEAN DEFAULT false,
			live_url TEXT,
			github_url TEXT,
			image_url TEXT,
			start_date DATE,
			end_date DATE,
			is_public BOOLEAN DEFAULT true,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,

		// Enable RLS on projects
		`ALTER TABLE projects ENABLE ROW LEVEL SECURITY;`,

		// Policy: Anyone can read public projects
		`CREATE POLICY "Public projects are viewable by everyone"
		 ON projects FOR SELECT
		 USING (is_public = true);`,

		// Skills table
		`CREATE TABLE IF NOT EXISTS skills (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			category VARCHAR(100) NOT NULL,
			level VARCHAR(50) NOT NULL,
			years_exp INTEGER DEFAULT 0,
			featured BOOLEAN DEFAULT false,
			icon_url TEXT,
			color VARCHAR(20),
			description TEXT,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,

		// Contact messages table
		`CREATE TABLE IF NOT EXISTS contact_messages (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			subject VARCHAR(255) NOT NULL,
			message TEXT NOT NULL,
			status VARCHAR(50) DEFAULT 'unread',
			ip_address INET,
			user_agent TEXT,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			read_at TIMESTAMP WITH TIME ZONE
		);`,

		// Analytics table for tracking visits
		`CREATE TABLE IF NOT EXISTS analytics (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			page VARCHAR(255) NOT NULL,
			referrer TEXT,
			user_agent TEXT,
			ip_address INET,
			country VARCHAR(10),
			city VARCHAR(100),
			device_type VARCHAR(50),
			browser VARCHAR(100),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,

		// Create updated_at trigger function
		`CREATE OR REPLACE FUNCTION update_updated_at_column()
		 RETURNS TRIGGER AS $$
		 BEGIN
			 NEW.updated_at = NOW();
			 RETURN NEW;
		 END;
		 $$ language 'plpgsql';`,

		// Add triggers for updated_at
		`CREATE TRIGGER update_users_updated_at
		 BEFORE UPDATE ON users
		 FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();`,

		`CREATE TRIGGER update_projects_updated_at
		 BEFORE UPDATE ON projects
		 FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();`,

		// Create indexes for better performance
		`CREATE INDEX IF NOT EXISTS idx_projects_user_id ON projects(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects(featured);`,
		`CREATE INDEX IF NOT EXISTS idx_projects_status ON projects(status);`,
		`CREATE INDEX IF NOT EXISTS idx_skills_user_id ON skills(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_skills_category ON skills(category);`,
		`CREATE INDEX IF NOT EXISTS idx_analytics_page ON analytics(page);`,
		`CREATE INDEX IF NOT EXISTS idx_analytics_created_at ON analytics(created_at);`,
	}

	for _, table := range tables {
		if _, err := SupabaseDB.Exec(table); err != nil {
			// Ignore errors for policies that might already exist
			log.Printf("Warning: %v", err)
		}
	}

	log.Println("Supabase tables and policies created successfully")
	return nil
}

// InsertSampleData inserts sample portfolio data
func InsertSampleData() error {
	// Insert sample user
	userSQL := `
		INSERT INTO users (name, email, role, bio, website, location, skills, is_public)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (email) DO NOTHING
		RETURNING id;`

	var userID string
	err := SupabaseDB.QueryRow(
		userSQL,
		"이혁주",
		"hyoukjoo@example.com",
		"Full-Stack Developer",
		"Backend engineer specializing in Go, TypeScript, and cloud architecture",
		"https://hyoukjoolee.github.io/portfolio",
		"Seoul, Korea",
		`["Go", "TypeScript", "Flutter", "AWS", "Docker", "PostgreSQL"]`,
		true,
	).Scan(&userID)

	if err != nil {
		log.Printf("Sample user might already exist: %v", err)
		return nil // Continue even if user exists
	}

	// Insert sample projects
	projectsSQL := `
		INSERT INTO projects (user_id, title, description, tech_stack, status, featured, live_url, github_url)
		VALUES
		($1, 'Portfolio API Server', 'Go REST API with Supabase integration', '["Go", "Supabase", "Railway"]', 'completed', true, 'https://portfolio-api.up.railway.app', 'https://github.com/hyoukjoolee/portfolio-api'),
		($1, 'Flutter Portfolio', 'Responsive portfolio website with i18n', '["Flutter", "Dart", "GitHub Pages"]', 'completed', true, 'https://hyoukjoolee.github.io/portfolio', 'https://github.com/hyoukjoolee/portfolio');`

	if _, err := SupabaseDB.Exec(projectsSQL, userID); err != nil {
		log.Printf("Warning: Sample projects: %v", err)
	}

	// Insert sample skills
	skillsSQL := `
		INSERT INTO skills (user_id, name, category, level, years_exp, featured, color)
		VALUES
		($1, 'Go', 'backend', 'expert', 3, true, '#00ADD8'),
		($1, 'TypeScript', 'frontend', 'advanced', 4, true, '#3178C6'),
		($1, 'Flutter', 'mobile', 'advanced', 2, true, '#02569B'),
		($1, 'Supabase', 'database', 'intermediate', 1, true, '#3ECF8E');`

	if _, err := SupabaseDB.Exec(skillsSQL, userID); err != nil {
		log.Printf("Warning: Sample skills: %v", err)
	}

	log.Println("Sample data inserted successfully")
	return nil
}

// CloseSupabase closes the Supabase connection
func CloseSupabase() error {
	if SupabaseDB != nil {
		return SupabaseDB.Close()
	}
	return nil
}