
# Docker
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
