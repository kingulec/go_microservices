# Go Full-Lifecycle App Deployment – ​​from Code to Kubernetes with Helm

## Project Overview
Project contains complete production ready webhook microservice application.
The service works as CI/CD webhook handler, receiving build results, calculating test pass rates, and returning a structured JSON response.

The project reflects the full application lifecycle: from local development, through containerization, to Kubernetes deployment.

## Go microservice
Source code: app/
* HTTPS server Listens on port 8443
* Provides a single endpoint: POST /webhook
* Request handling:
The endpoint accepts payloads in JSON format, in the following structure:

```json
{
"build_id": "abc123",
"status": "success",
"timestamp": "2025-06-13T16:00:00Z",
"tests": [
{"name": "TestLogin", "passed": true},
{"name": "TestSignup", "passed": false},
{"name": "TestCheckout", "passed": true}
]
}
```

**Example of successfull response**
```json
{
  "received": true,
  "build_id": "abc123",
  "pass_rate": 66.67
}
```

## Error Handling
| HTTP Status | Error Type            | Message Example                             | Description                                                      |
| ----------- | --------------------- | ------------------------------------------- | ---------------------------------------------------------------- |
| 400         | Invalid JSON          | `Bad request`                               | Request body is not a valid JSON or cannot be decoded            |
| 400         | Invalid Payload       | `Invalid payload: Missing required fields`  | Required JSON fields are missing or payload structure is invalid |
| 400         | Invalid Test Data     | `Error in tests: No tests provided`         | `tests` array is empty                                           |
| 400         | Invalid Test Data     | `Error in tests: Test name cannot be empty` | Test case contains an empty `name` field                         |
| 405         | Method Not Allowed    | `Method not allowed`                        | HTTP method other than `POST` was used                           |
| 500         | Internal Server Error | `Failed to encode response`                 | Server failed to encode JSON response                            |






## Local development
App source code directory: app/

**Starting server:**

Setting environment variables for server.
```bash
export TLS_CERT_PATH=/path/to/tls.crt

export TLS_KEY_PATH=/path/to/tls.key
```

```bash
go run main.go 
```
Port can be easly changed by simply adding parameter.

**ex:** ```go run main.go [port]```

## Generate self-signed certificates 
```bash 
bash scripts/generate_certs.sh
```
Script should create directory named 'certs' and self-signed certificates.

## Testing locally 
test configuration file: tests/pytest.ini
automated tests: tests/test_service.py
**Running pytest tests**
Server should be running.

Setting environment variable for tests:
```bash 
export WEBHOOK_URL=https://localhost:8443/webhook
```
```bash
pytest test_service.py
```

**Curl tests**
```bash
bash service_tests.sh https://localhost:8443/webhook  path/to/test_data.json [METHOD]
```

Method parameter is optional by default it is set to POST. Can be change to check the proper error handling of other methods.

Sample files with valid and invalid json payloads are in test_data directory.

## Docker
**build image** :
```bash
docker build -t webhook-server:latest .
```
**run container**:
```bash
docker run -p 8443:8443 \
  -v $(pwd)/certs:/certs:ro \
  -e TLS_CERT_PATH=/certs/tls.crt \
  -e TLS_KEY_PATH=/certs/tls.key \
  webhook-server
```

