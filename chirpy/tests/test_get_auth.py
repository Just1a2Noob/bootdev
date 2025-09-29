import json

import requests


def test_login_api_call():
    endpoint = "http://localhost:8080/api/login"
    header = {"Content-Type": "application/json"}

    payload = {
        "email": "user@example.com",
        "password": "04225",
    }

    raw_json = json.dumps(payload)

    response = requests.post(endpoint, headers=header, data=raw_json)

    assert response.status_code == 200

    data = json.loads(response.text)
    add_chirp(data["token"])


def add_chirp(token):
    endpoint = "http://localhost:8080/api/chirps"
    headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    payload = {
        "body": "Hello, world!",
        "user_id": "a64515aa-2d04-4f4b-818a-8ade2d9a2924",
    }
    raw_json = json.dumps(payload)

    response = requests.post(endpoint, headers=headers, data=raw_json)

    assert response.status_code == 201

    data = json.loads(response.text)
    print(data)


test_login_api_call()
