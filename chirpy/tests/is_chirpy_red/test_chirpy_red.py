import json
import os

import requests
from dotenv import load_dotenv

ENDPOINT = "http://localhost:8080/api"

# Load the .env file
load_dotenv()

API_Key = os.getenv("POLKA_KEY")


def test_is_chirpy_red():
    file_path = "gislaine.json"
    with open(file_path, "r") as file:
        data = json.load(file)

    chirpy_headers = {"Authorization": f"ApiKey {API_Key}"}

    chirpy_response = requests.post(
        ENDPOINT + "/polka/webhooks", headers=chirpy_headers, json=data
    )

    assert chirpy_response.status_code == 204


def test_is_chirpy_red_no_api_key():
    file_path = "gislaine.json"
    with open(file_path, "r") as file:
        data = json.load(file)

    chirpy_response = requests.post(ENDPOINT + "/polka/webhooks", json=data)

    assert chirpy_response.status_code == 400


def test_is_chirpy_red_invalid_key():
    file_path = "gislaine.json"
    with open(file_path, "r") as file:
        data = json.load(file)

    chirpy_headers = {"Authorization": "ApiKey 012h9buh9hfdw2h891hb11"}

    chirpy_response = requests.post(
        ENDPOINT + "/polka/webhooks", headers=chirpy_headers, json=data
    )

    assert chirpy_response.status_code == 401
