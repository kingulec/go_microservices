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
Check
```bash
kubectl get pods
```

# Testing 
Run the following 
```bash
kubectl port-forward svc/webhook-service 8443:443
```
and in differnet terminal run tests

```bash
export WEBHOOK_URL=https://localhost:8443/webhook
python3 -m pytest test_service.py
```

For details go to 

Kubernetes documentation: https://kubernetes.io