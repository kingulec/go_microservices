
# Docker
The Docker image is built using a multi-stage build to minimize the final image size.
Only the compiled Go binary is included in the runtime image.
### TLS Certificates
TLS certificates are not included in the Docker image.
They are mounted at runtime using a read-only volume and provided to the application via environment variables.

## Build image

Build a Docker image for the webhook server using the following command: 
```bash
docker build -t webhook-server:latest .
```
-t webhook-server → assigns the name webhook-server
. uses the current directory

## Run the container
Run the following command to run the container.
```bash
docker run -p 8443:8443 \
  -v $(pwd)/certs:/certs:ro \
  -e TLS_CERT_PATH=/certs/tls.crt \
  -e TLS_KEY_PATH=/certs/tls.key \
  webhook-server
```
* -p 8443:8443 → maps port 8443 in the container to port 8443 on your host machine.
* -v $(pwd)/certs:/certs:ro → mounts the certs directory from your host into the container as read-only
* -e TLS_CERT_PATH=/certs/tls.crt and -e TLS_KEY_PATH=/certs/tls.key → provide the paths to the TLS certificate and private key for HTTPS.

* webhook-server → the name of the Docker image to run.

After starting the container, the HTTPS server will be available at:

https://localhost:8443

You can run tests to verify that the server is working correctly:
- [Testing](testing.md)