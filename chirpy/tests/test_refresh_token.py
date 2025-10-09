import json

import requests

ENDPOINT = "http://localhost:8080/api"


def test_refresh_token():
    login = login_user()
    refresh_token = login["refresh_token"]

    refresh_header = {
        "Authorization": "Bearer " + refresh_token,
    }

    refresh_response = requests.post(
        ENDPOINT + "/refresh",
        headers=refresh_header,
    )

    assert refresh_response.status_code == 200


def test_revoke_token():
    login = login_user()

    refresh_token = login["refresh_token"]

    refresh_header = {
        "Authorization": "Bearer " + refresh_token,
    }

    revoke_response = requests.post(
        ENDPOINT + "/revoke",
        headers=refresh_header,
    )

    assert revoke_response.status_code == 204

    # Failed refreshing due to revoke

    refresh_response = requests.post(
        ENDPOINT + "/refresh",
        headers=refresh_header,
    )

    assert refresh_response.status_code == 400


def login_user():
    login_header = {"Content-Type": "application/json"}

    login_payload = {
        "email": "gislaine@gmail.com",
        "password": "441201",
    }

    login_response = requests.post(
        ENDPOINT + "/login", headers=login_header, data=json.dumps(login_payload)
    )

    login_data = json.loads(login_response.text)

    return login_data
