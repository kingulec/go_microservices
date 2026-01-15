
## Testing locally 
test configuration file: tests/pytest.ini
tests directory: tests/

The project includes automated integration tests that validate application behavior.
The tests assume that the HTTPS server is already running and accessible at the URL specified by the `WEBHOOK_URL` environment variable.
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
To run python test run the following command:

```bash
python3 -m pytest test_service.py
```

Tags:
Running test from test_service.py file with tag 'smoke'. 
```bash
 python3 -m pytest -m smoke test_service.py
```
Running all the tests with tag 'smoke'. 
```bash
 python3 -m pytest -m smoke
```
### Curl tests**
The `service_tests.sh` script can be used to manually send test payloads via curl.
It allows testing different HTTP methods and payloads to verify server behavior outside of Python tests.
```bash
bash service_tests.sh https://localhost:8443/webhook  path/to/test_data.json [METHOD]
```

Method parameter is optional by default it is set to POST. Can be change to check the proper error handling of other methods.

Sample files with valid and invalid json payloads are in test_data directory. Files can be used for testing purposes.
