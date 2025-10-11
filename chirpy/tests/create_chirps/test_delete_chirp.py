import requests
from fill_chirps import login_user

ENDPOINT = "http://localhost:8080/api"


def test_delete_chirp():
    login = login_user(
        "gislaine@gmail.com",
        "441201",
    )
    token = login["token"]

    chirp_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    # NOTE: After reupdating the chirps, don't forget to change the chirpID
    chirpID = "cb24717a-6d8e-4f67-940a-565924a01e03"

    chirp_response = requests.delete(
        ENDPOINT + f"/chirps/{chirpID}", headers=chirp_headers
    )

    assert chirp_response.status_code == 204


def test_delete_chirp_err():
    login = login_user(
        "gislaine@gmail.com",
        "441201",
    )
    token = login["token"]

    chirp_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    chirpID = "33ffab2b-e359-4366-8fb7-a722d2057e35"

    chirp_response = requests.delete(
        ENDPOINT + f"/chirps/{chirpID}", headers=chirp_headers
    )

    assert chirp_response.status_code == 400
