#! /bin/bash
set -e
CERT_DIR="certs"
CERT_FILE="$CERT_DIR/tls.crt"
KEY_FILE="$CERT_DIR/tls.key"
DAYS_VALID=365

mkdir -p "$CERT_DIR"

openssl req -x509 -newkey rsa:4096 \
  -keyout "$KEY_FILE" \
  -out "$CERT_FILE" \
  -days "$DAYS_VALID" \
  -nodes \
  -subj "/CN=localhost"

echo "TLS certificates generated:"
echo "Certificate: $CERT_FILE"
echo "Private key: $KEY_FILE"