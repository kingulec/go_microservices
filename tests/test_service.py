import os
import json
import pytest
import requests
import subprocess

BASE_URL = os.getenv("WEBHOOK_URL", f"https://localhost:8443/webhook")

# test payloads
VALID_PAYLOAD_TEMPLATE = {
    "build_id": "abc123",
    "status": "success",
    "timestamp": "2025-06-13T16:00:00Z",
    "tests": [
        {"name": "TestLogin", "passed": True},
        {"name": "TestSignup", "passed": False},
        {"name": "TestCheckout", "passed": True}
        ]
}


MISSING_BUILD_ID = {
    "status": "success",
    "timestamp": "2025-06-13T16:00:00Z",
    "tests": [
        {"name": "TestLogin", "passed": True}
        ]
}

NO_TESTS = {
    "build_id": "abc123",
    "status": "success",
    "timestamp": "2025-06-13T16:00:00Z",
    "tests": []
}


# expected results
EXPECTED_MISSING_BUILD_ID = "Invalid payload: Missing required fields\n"
EXPECTED_NO_TESTS = "Error in tests: No tests provided\n"
INV_DATA = "This is not a valid JSON payload"

def prep_payload(tests_list=None):
    """
    Preapare payload based on template 

    @param tests_list: list of test dicts 
    """
    payload = VALID_PAYLOAD_TEMPLATE.copy()
    if tests_list is not None:
        payload['tests'] = tests_list
    return payload

@pytest.mark.smoke
def test_service_running():
    response = requests.post(BASE_URL, json=VALID_PAYLOAD_TEMPLATE, verify=False, timeout=5)
    
    assert response.status_code == 200

@pytest.mark.parametrize(
    "payload, expected_result",
    [
        (MISSING_BUILD_ID, EXPECTED_MISSING_BUILD_ID),
        (NO_TESTS, EXPECTED_NO_TESTS),
    ],
)
def test_invalid_payload(payload, expected_result):
    response = requests.post(BASE_URL, json=payload, verify=False, timeout=5)
    print(response.text)
    assert response.status_code == 400
    assert response.text == expected_result

def test_invalid_json():
    response = requests.post(BASE_URL, data=INV_DATA, verify=False, timeout=5)
    assert response.status_code == 400


def test_method_not_allowed():
    response = requests.get(BASE_URL, verify=False, timeout=5)
    assert response.status_code == 405

@pytest.mark.parametrize(
    "payload, expected_result",
    [   
        (VALID_PAYLOAD_TEMPLATE, {'build_id': 'abc123', 'pass_rate': 66.67, 'received': True}),
        (prep_payload([{"name": "T1", "passed": True}]), {'build_id': 'abc123', 'pass_rate': 100, 'received': True}),
        (prep_payload([{"name": "T1", "passed": False}]), {'build_id': 'abc123', 'pass_rate': 0, 'received': True}),
        (prep_payload([{"name": "T1", "passed": False}, {"name": "T1", "passed": True}]),
         {'build_id': 'abc123', 'pass_rate': 50, 'received': True})
    ],
)
def test_valid_payload(payload, expected_result):
    response = requests.post(BASE_URL, json=payload, verify=False, timeout=5)
    assert response.status_code == 200
    assert response.json() == expected_result