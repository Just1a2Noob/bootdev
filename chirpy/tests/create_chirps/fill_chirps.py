import json

import requests

ENDPOINT = "http://localhost:8080/api"

users = {
    "gislaine@gmail.com": "441201",
    "rudeus_greyrat@gmail.com": "1701823",
    "teej@gmail.com": "123409",
}


def main():
    """
    filepath, user -> sends HTTP POST to chirps

    User -> a valid user that is in the list
    Filepath -> JSON file containing standard chirp format
                (user ID must match with login ID)
    """

    email = input("User: ")
    login = login_user(email, users[email])
    token = login["token"]

    file_path = input("filepath: ")

    with open(file_path, "r") as file:
        data = json.load(file)

    chirp_headers = {
        "Content-Type": "application/json",
        "Authorization": "Bearer " + token,
    }

    chirp_response = requests.post(
        ENDPOINT + "/chirps", headers=chirp_headers, json=data
    )

    if chirp_response.status_code != 201:
        print(chirp_response.text)
        raise Exception("Sending chirp is unsuccessful")

    print(email, " Has sent chirp")
    print_colored_response(chirp_response)
    return


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


if __name__ == "__main__":
    main()
