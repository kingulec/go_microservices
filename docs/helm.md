# Helm Chart -Stage 4

## Helm installation 
To install helm go to https://helm.sh/docs/intro/install/.

## 
Run the following and copy it to my-webhook/values_template.yaml.
```bash
base64 -w 0 tls.crt
base64 -w 0 tls.key
``` 

values_template.yaml change te following and rename file to values.yaml.
```bash
tls:
  cert: "<base64-encoded-cert-string>"
  key: "<base64-encoded-key-string>"
```

## Installation
Run the following command to install Helm chart
```bash
helm install my-webhook ./my-webhook
```

To upgrade
```bash
helm upgrade my-webhook ./my-webhook
```
## Testing app
Forward the service port to local machine:
```bash
kubectl port-forward svc/my-webhook-service 8443:443
```
You can run the tests to verify if application is working correctly.
```bash
cd tests
python3 -m pytest test_service.py
```
or shell test scripts:
```bash
'cd tests'
'bash service_tests.sh  https://localhost:8443/webhook ../test_data/test_data01.json'
```
Check section
- [Testing](testing.md)

## Verification
You can verify te results with the following commands

To check the pods:
```bash
kubectl get pods -l app=my-webhook
```

TLS secret created:
```bash
kubectl get secret my-webhook-tls
```
Chack logs:
```bash
kubectl logs pod_name
```
Check service:

```bash
kubectl get svc my-webhook-service
```

## Uninstallation
To uninstall run the following:
```bash
helm upgrade my-webhook ./my-webhook
```