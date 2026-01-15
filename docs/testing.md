
## Testing locally 
test configuration file: tests/pytest.ini
automated tests: tests/test_service.py

### Python virtual environment (Optional)
Create venv
```bash
python3 -m venv venv 
source venv/bin/activate
```
All the packages required for test can be installed inside of virtual environment.

To deactivate the venv run command.
```bash
deactivate
```
### Test requirements
Required packages: requests, pytest.
Packages can be installed with pip:
```bash 
pip install <package_name>
```

Setting environment variable for tests:
```bash 
export WEBHOOK_URL=https://localhost:8443/webhook
```
If not set the default value will be used https://localhost:8443/webhook.

### Running pytest tests
Tests does not start thew server, the server should be running before running the tests.

```bash
python3 -m pytest test_service.py
```

### Curl tests**
```bash
bash service_tests.sh https://localhost:8443/webhook  path/to/test_data.json [METHOD]
```

Method parameter is optional by default it is set to POST. Can be change to check the proper error handling of other methods.

Sample files with valid and invalid json payloads are in test_data directory. Files can be used for testing purposes.
