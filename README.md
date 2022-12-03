# koham
Role, Access and Token Management Service
## Setup and Run
  ### Using Dockerfile - (Has a bug, refrain using while development)
  <!-- 1. Build docker - `docker build -t koham .`
  2. Run Docker - `docker run -it -p 8080:8080 koham` -->
  ### Run using go
  <!-- 1. Using docker compose to run mysql and redis - `docker compose up -d` -->
  2. Download dependencies - `go mod download`
  3. Run main - `go run main.go`
  ### Setup and run pre-commit
  <!-- 1. Install golangci-lint - `brew install golangci-lint` or follow this - `https://golangci-lint.run/usage/install/
  2. Install pre-commit for enabling hooks - `brew install pre-commit`
  3. Install pre-commit `pre-commit install`
  4. Try adding some changes, and check if on `git commit` pre-commit hooks are running or not. -->