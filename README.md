# Portfolio API

Backend API server for portfolio showcase built with Go and Gin framework.

## Features

- 🚀 RESTful API with comprehensive endpoints
- 📚 Swagger/OpenAPI documentation
- 🐳 Docker containerization
- ☁️ AWS Free Tier deployment ready
- 🔄 CI/CD with GitHub Actions
- 🛡️ CORS configured for portfolio frontend
- 📊 Analytics and statistics endpoints
- 💌 Contact form handling

## API Endpoints

### System
- `GET /health` - Health check

### Users
- `GET /api/v1/users` - Get all users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user

### Projects
- `GET /api/v1/projects` - Get projects (filterable)
- `POST /api/v1/projects` - Create project
- `GET /api/v1/projects/{id}` - Get project by ID
- `PUT /api/v1/projects/{id}` - Update project
- `DELETE /api/v1/projects/{id}` - Delete project

### Skills
- `GET /api/v1/skills` - Get skills (filterable)
- `POST /api/v1/skills` - Add skill
- `DELETE /api/v1/skills/{id}` - Remove skill

### Contact
- `POST /api/v1/contact` - Submit contact form

### Analytics
- `GET /api/v1/stats/views` - View statistics
- `GET /api/v1/stats/projects` - Project statistics
- `POST /api/v1/stats/visit` - Record visit

## Quick Start

### Local Development

1. **Clone and setup**
   ```bash
   git clone <repository>
   cd portfolio-api
   go mod download
   ```

2. **Run locally**
   ```bash
   go run main.go
   ```

3. **Access API**
   - API: http://localhost:8080
   - Swagger docs: http://localhost:8080/swagger/index.html
   - Health check: http://localhost:8080/health

### Docker Development

1. **Build and run with Docker**
   ```bash
   docker build -t portfolio-api .
   docker run -p 8080:8080 portfolio-api
   ```

2. **Or use Docker Compose**
   ```bash
   docker-compose up --build
   ```

## Deployment

### AWS Free Tier Deployment

This project is configured for AWS Free Tier deployment using:
- **ECS Fargate** (1 vCPU, 512MB memory)
- **ECR** for container registry
- **CloudWatch** for logging
- **Default VPC** to minimize costs

#### Prerequisites

1. **AWS Account** with programmatic access
2. **AWS CLI** configured
3. **Terraform** installed (optional, for infrastructure)

#### Setup Steps

1. **Configure AWS credentials**
   ```bash
   aws configure
   ```

2. **Deploy infrastructure (using Terraform)**
   ```bash
   cd terraform
   terraform init
   terraform plan
   terraform apply
   ```

3. **Push initial image to ECR**
   ```bash
   # Get login token
   aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin <account-id>.dkr.ecr.ap-northeast-2.amazonaws.com

   # Build and push
   docker build -t portfolio-api .
   docker tag portfolio-api:latest <account-id>.dkr.ecr.ap-northeast-2.amazonaws.com/portfolio-api:latest
   docker push <account-id>.dkr.ecr.ap-northeast-2.amazonaws.com/portfolio-api:latest
   ```

4. **Set up GitHub Actions secrets**
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`
   - `DOCKER_USERNAME` (optional)
   - `DOCKER_PASSWORD` (optional)

### CI/CD Pipeline

GitHub Actions automatically:
1. **Tests** code on PR/push
2. **Builds** Docker image
3. **Pushes** to ECR
4. **Deploys** to ECS Fargate
5. **Updates** service

## Configuration

### Environment Variables

- `GIN_MODE` - Gin mode (debug/release)
- `PORT` - Server port (default: 8080)

### CORS Configuration

Pre-configured for:
- `http://localhost:3000` (local Flutter dev)
- `https://hyoukjoolee.github.io` (GitHub Pages)

## Project Structure

```
portfolio-api/
├── main.go                 # Application entry point
├── go.mod                  # Go dependencies
├── handlers/               # HTTP handlers
│   ├── handlers.go         # Main handlers
│   └── users.go           # User-specific handlers
├── models/                 # Data models
│   ├── user.go
│   ├── project.go
│   ├── contact.go
│   └── stats.go
├── docs/                   # API documentation
│   └── swagger.json        # OpenAPI specification
├── terraform/              # Infrastructure as code
│   └── main.tf
├── .github/workflows/      # CI/CD pipelines
│   ├── deploy.yml          # AWS deployment
│   └── docker.yml          # Docker build/test
├── Dockerfile              # Container definition
├── docker-compose.yml      # Local development
└── nginx.conf             # Reverse proxy config
```

## Portfolio Integration

The API is designed to work with the Flutter portfolio frontend:

1. **OpenAPI Spec** is copied to `portfolio/assets/openapi/portfolio-api.json`
2. **CORS** is configured for the portfolio domain
3. **cURL generation** works with the existing portfolio API explorer
4. **Analytics endpoints** provide data for portfolio statistics

### Adding to Portfolio

Update the portfolio's OpenAPI loader to include the new API:

```dart
// In portfolio project
// lib/services/openapi_loader.dart
const apiSpecs = [
  'assets/openapi/openapi.json',        // existing demo API
  'assets/openapi/portfolio-api.json',  // new portfolio API
];
```

## Development

### Adding New Endpoints

1. **Define model** in `models/`
2. **Create handler** in `handlers/`
3. **Add route** in `main.go`
4. **Update OpenAPI spec** in `docs/swagger.json`
5. **Add tests** (optional but recommended)

### Testing

```bash
# Run tests
go test -v ./...

# Test with Docker
docker-compose up --build
curl http://localhost:8080/health
```

## Monitoring

- **Health endpoint**: `/health`
- **CloudWatch logs**: ECS service logs
- **ECS metrics**: Container insights enabled

## Cost Optimization

Configured for AWS Free Tier:
- Single Fargate task (0.25 vCPU, 512MB)
- CloudWatch logs with 7-day retention
- ECR with lifecycle policy (keep 10 images)
- No load balancer (direct ECS access)

## Security

- CORS properly configured
- Security headers in nginx
- Rate limiting (10 req/s)
- Container security scanning enabled

## License

MIT License