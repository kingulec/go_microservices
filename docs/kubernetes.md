# Kubernetes - Stage 3

## CERTIFICATES
You can generate self-signed certificates for testing purposes:
```bash
bash scripts/generate_certs.sh
```
Kubernetes Secrets require the data to be base64-encoded. Run the following and copy it to secret_template.yaml.
```bash
base64 -w 0 tls.crt
base64 -w 0 tls.key
```

secret_template.yaml: Modify the following lines
```yaml
data:
  tls.crt: aadddfFKQ1dUd...=
  tls.key: Lddd4...=
  ...
```
You can also change the name to service.yaml.

```bash
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/deployment.yaml
```

For details go to 

Kubernetes documentation: https://kubernetes.io