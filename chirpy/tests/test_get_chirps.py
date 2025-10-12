import json

import requests

ENDPOINT = "http://localhost:8080/api"


def test_get_chirps_no_author():
    chirp_response = requests.get(
        ENDPOINT + "/chirps",
    )

    print_colored_response(chirp_response)
    assert chirp_response.status_code == 200


def test_get_chirps_author():
    author_id = "f1900ef8-c33d-4601-b40b-937f371b5dbf"
    chirp_response = requests.get(
        ENDPOINT + f"/chirps?author_id={author_id}",
    )
    print_colored_response(chirp_response)
    assert chirp_response.status_code == 200


class Colors:
    HEADER = "\033[95m"
    BLUE = "\033[94m"
    GREEN = "\033[92m"
    YELLOW = "\033[93m"
    RED = "\033[91m"
    END = "\033[0m"
    BOLD = "\033[1m"


def print_colored_response(response):
    print(f"{Colors.BOLD}Status: {Colors.END}", end="")

    if response.status_code == 200:
        print(f"{Colors.GREEN}{response.status_code}{Colors.END}")
    else:
        print(f"{Colors.RED}{response.status_code}{Colors.END}")

    print(f"{Colors.BLUE}Response:{Colors.END}")
    print(json.dumps(response.json(), indent=2))


test_get_chirps_no_author()
print("=======================================")
test_get_chirps_author()
