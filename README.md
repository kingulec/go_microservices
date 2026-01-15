# Go Full-Lifecycle App Deployment – ​​from Code to Kubernetes with Helm
## Table of Contents
- [Project Overview](#project-overview)
- [Clone the Repository](#clone-the-repository)
- [Documentation](#documentation)

## Project Overview
Project contains complete production ready webhook microservice application.
The service works as CI/CD webhook handler, receiving build results, calculating test pass rates, and returning a structured JSON response.

The project reflects the full application lifecycle: from local development, through containerization, to Kubernetes deployment.



## Clone the repository
```bash
git clone https://github.com/<githubuser>/go_microservices.git

cd go_microservices
```
Download dependencies 
```bash 
go mod tidy
```

## Documentation

- [Go Webhook Service](docs/go_service.md)
- [Testing](docs/testing.md)
- [Docker](docs/docker.md)
- [Kubernetes](docs/kubernetes.md)
- [Error Handling](docs/error_handling.md)


