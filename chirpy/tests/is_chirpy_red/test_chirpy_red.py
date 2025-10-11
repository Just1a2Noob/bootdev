import json

import requests

ENDPOINT = "http://localhost:8080/api"


def test_is_chirpy_red():
    file_path = "gislaine.json"
    with open(file_path, "r") as file:
        data = json.load(file)

    chirpy_response = requests.post(ENDPOINT + "/polka/webhooks", json=data)

    assert chirpy_response.status_code == 204
