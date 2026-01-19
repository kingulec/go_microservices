# Go Full-Lifecycle App Deployment – ​​from Code to Kubernetes with Helm
## Table of Contents
- [Project Overview](#project-overview)
- [Clone the Repository](#clone-the-repository)
- [Project structure](#project-structure)
- [Documentation](#documentation)

## Project Overview
Project contains complete production ready webhook microservice application.
The service works as CI/CD webhook handler, receiving build results, calculating test pass rates, and returning a structured JSON response.

The project reflects the full application lifecycle: from local development, through containerization, to Kubernetes deployment.

## Requirements
Lista narzędzi potrzebnych do uruchomienia projektu:

* Go 
* Docker
* Kubernetes 
* Helm
* TLS certificates

pytest (opcjonalnie)

## Clone the repository
```bash
git clone https://github.com/<githubuser>/go_microservices.git

cd go_microservices
```
## Project structure 
```
GO_MICROSERVICES
├── app/                # Go application source code
│   ├── models/         # Data models
│   ├── validators/     
│   ├── wbhandler/      # Webhook handler logic
│   └── main.go         # main logic
├── docs/ 
├── k8s/              # kubernetes mainfests.yaml
│
├── my-webhook/      # Helm chart for Kubernetes deployment
│
├── scripts/     
│   └── generate_certs.sh  #script to generate self-signed certs
├── test_data/   # example of test payloads json
├── tests/               
│   ├── pytest.ini         # pytest tests configuration file
│   ├── service_tests.sh    
│   ├── test_service.py      
│   └── main.go         # main logic
├── Dockerfile          
└── README.md
```
## Documentation
For usage instruction follow the links:
- [Go Webhook Service - STAGE 1](docs/go_service.md) - running the service locally.
- [Testing](docs/testing.md) - Test setup an usage.
- [Docker-STAGE 2](docs/docker.md) - Building and running the service in a container.
- [Kubernetes-STAGE 3](docs/kubernetes.md) - deployment the service to Kubernetes using manifests.
- [Helm Chart-STAGE 4](docs/helm.md) -Automated deployment using Helm.


