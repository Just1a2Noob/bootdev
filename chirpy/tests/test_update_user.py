import json

import requests

ENDPOINT = "http://localhost:8080/api"


def test_update_user():
    create_user(
        "user@example.com",
        "04225",
    )

    login = login_user(
        "user@example.com",
        "04225",
    )
    token = login["token"]

    update_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    update_payload = {
        "email": "user@gmail.com",
        "password": "112412",
    }

    update_response = requests.put(
        ENDPOINT + "/users", headers=update_headers, data=json.dumps(update_payload)
    )

    update_data = json.loads(update_response.text)
    print(update_data)

    assert update_response.status_code == 200
    assert update_data["email"] == update_payload["email"]

    reset_user("user@gmail.com")


def login_user(email, password):
    login_header = {"Content-Type": "application/json"}

    login_payload = {
        "email": email,
        "password": password,
    }

    login_response = requests.post(
        ENDPOINT + "/login", headers=login_header, data=json.dumps(login_payload)
    )

    login_data = json.loads(login_response.text)

    return login_data


def reset_user(email):
    if type(email) != str:
        return 400

    user_payload = {
        "email": email,
    }

    reset_response = requests.delete(
        "http://localhost:8080/admin/users",
        data=json.dumps(user_payload),
    )

    return reset_response.status_code


def create_user(email, password):
    user_header = {"Content-Type": "application/json"}

    user_payload = {
        "email": email,
        "password": password,
    }

    user_response = requests.post(
        ENDPOINT + "/users", headers=user_header, data=json.dumps(user_payload)
    )

    user_data = json.loads(user_response.text)

    return user_data
