# Go Webhook Service - STAGE 1
Source code: app/

This is a HTTPS webhook server written in Go.

* HTTPS server Listens on port **8443**
* Provides a single endpoint: POST /webhook
* Accepts payloads in **JSON format** with the following structure:

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

**Example of successful response**
```json
{
  "received": true,
  "build_id": "abc123",
  "pass_rate": 66.67
}
```
# Prerequisites

- Go installed (version ≥ 1.21 recommended)

- TLS certificate and key for HTTPS

- Setting environment variables for server
```bash
export TLS_CERT_PATH=/path/to/tls.crt

export TLS_KEY_PATH=/path/to/tls.key

export SERVICE_HOST="0.0.0.0"
```
# Local development

## Starting server:

Run the command.

```bash
cd app
go run main.go 
```
Port can be easily changed by simply adding parameter.

**ex:** ```go run main.go [port]```

## Generate self-signed certificates 
Project contains a script for self-signed certificates generation. The script can be found in scripts directory.

```bash 
bash scripts/generate_certs.sh
```
Script should create directory named 'certs' and self-signed certificates.

## Error Handling
| HTTP Status | Error Type            | Message Example                             | Description                                                      |
| ----------- | --------------------- | ------------------------------------------- | ---------------------------------------------------------------- |
| 400         | Invalid JSON          | `Bad request`                               | Request body is not a valid JSON or cannot be decoded            |
| 400         | Invalid Payload       | `Invalid payload: Missing required fields`  | Required JSON fields are missing or payload structure is invalid |
| 400         | Invalid Test Data     | `Error in tests: No tests provided`         | `tests` array is empty                                           |
| 400         | Invalid Test Data     | `Error in tests: Test name cannot be empty` | Test case contains an empty `name` field                         |
| 405         | Method Not Allowed    | `Method not allowed`                        | HTTP method other than `POST` was used                           |
| 500         | Internal Server Error | `Failed to encode response`                 | Server failed to encode JSON response                            |
