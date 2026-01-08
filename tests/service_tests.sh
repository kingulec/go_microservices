#!/bin/bash
#USAGE:
# bash service_tests.sh [WEBHOOK_URL] [PAYLOAD_FILE] [NEGATIVE_TEST]
WEBHOOK_URL="$1" # URL jako pierwszy argument lub domyślny
PAYLOAD_FILE="$2"
METHOD="${3:-POST}"

echo -e "Using Webhook URL: $WEBHOOK_URL.\nPayload File: $PAYLOAD_FILE.\nMethod: $METHOD"

if [ -z "$WEBHOOK_URL" ] || [ -z "$PAYLOAD_FILE" ]; then
    echo -e "Error: WEBHOOK_URL and PAYLOAD_FILE parameters are required. \n USAGE: bash service_tests.sh [WEBHOOK_URL] [PAYLOAD_FILE] [METHOD]"
    exit 1
fi

if [ ! -f "$PAYLOAD_FILE" ]; then
    echo "Payload file not found: $PAYLOAD_FILE"
    exit 1
fi
if [ "$METHOD" != "POST" ]; then
    echo "WARNING: Only POST method is supported."
fi

# -k to ignore SSL certificate validation
response=$(curl -s -k -X "$METHOD" "$WEBHOOK_URL" -d @"$PAYLOAD_FILE" -w "\n%{http_code}")
status_code=$(echo "$response" | tail -n1)
json_response=$(echo "$response" | head -n-1)

echo "HTTP Code: $status_code"
if [[ "$status_code" == 200 ]]; then
    echo "Request was successful."
    echo "Response: $json_response"
else
    echo "Request failed with status code $status_code."
    echo "ERROR: $json_response"
    exit 1
fi