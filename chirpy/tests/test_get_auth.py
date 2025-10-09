import json

import requests

ENDPOINT = "http://localhost:8080/api"


def test_login_and_add_chirp():
    # Checks login handling
    login_header = {"Content-Type": "application/json"}

    login_payload = {
        "email": "gislaine@gmail.com",
        "password": "441201",
    }

    login_response = requests.post(
        ENDPOINT + "/login", headers=login_header, data=json.dumps(login_payload)
    )

    assert login_response.status_code == 200

    login_data = json.loads(login_response.text)
    # Gets the token for adding chirp
    token = login_data["token"]

    # Check add chirp
    add_chirp_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    add_chirp_payload = {
        "body": "Hello, world!",
        "user_id": "275cebdc-3cc0-4e86-927d-09cafe887fab",
    }

    add_chirp_response = requests.post(
        ENDPOINT + "/chirps",
        headers=add_chirp_headers,
        data=json.dumps(add_chirp_payload),
    )

    assert add_chirp_response.status_code == 201

    add_chirp_data = json.loads(add_chirp_response.text)

    # Checks if the sent payload is the same with received payload
    assert add_chirp_data["Body"] == add_chirp_payload["body"]


def test_login_and_add_chirp_fail():
    # Test checks if login and adding chirp user is different it should fail

    # Checks login handling
    login_header = {"Content-Type": "application/json"}

    login_payload = {"email": "gislaine@gmail.com", "password": "441201"}

    login_response = requests.post(
        ENDPOINT + "/login", headers=login_header, data=json.dumps(login_payload)
    )

    assert login_response.status_code == 200

    login_data = json.loads(login_response.text)
    # Gets the token for adding chirp
    token = login_data["token"]

    # Check add chirp
    add_chirp_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    add_chirp_payload = {
        "body": "Chirpy is pretty cool",
        # Different user_id
        "user_id": "a64515aa-2d04-4f4b-818a-8ade2d9a2924",
    }

    add_chirp_response = requests.post(
        ENDPOINT + "/chirps",
        headers=add_chirp_headers,
        data=json.dumps(add_chirp_payload),
    )

    assert add_chirp_response.status_code == 401
