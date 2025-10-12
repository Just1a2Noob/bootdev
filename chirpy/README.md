
# Goal

This small project is meant to showcase and learn how HTTP Servers work in go.

The practicality of this project is creating a REST API with some basic functionality:

- Reading from database.
- Getting data from database.
- Data from database is sent to the client in JSON format.
- Managing database migration and changes.
- Authentication and Authorization using JWT Tokens.


# API Documentation

## Admin (group)
### GET /admin/metrics

Gets the total amount of page visits in HTML format:

```html
<html>

<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited 0 times!</p>
</body>

</html>
```

### POST /admin/reset

Deletes all users within the database. Does not return a body

### DELETE /admin/users

Deletes a user given the specified JSON format:

```json
{
    "email": "user@example.com",
    "password": "04225",
}
```

## User (group)


### POST /api/users

Creates a new user given JSON format:

```json
{
    "email": "user@example.com",
    "password": "04225",
}
```
HTTP Response:

```json
{
    "id":"2b3916f0-56c3-49ac-b8ef-3d69ad017a94",
    "created_at":"2025-10-12T08:03:59.573532Z",
    "updated_at":"2025-10-12T08:03:59.573533Z",
    "email":"test@gmail.com",
    "is_chirpy_red":"false"
}
```
- `id` = The identifier for this resource.
- `created_at` = The time and date of when this resource is created.
- `updated_at` = The time and date of when this resource is last updated.
- `email` = The email of the user.
- `is_chirpy_red` = Is the user paying customer or not.


### POST /api/login

Login a user given a JSON format:

```json
{
    "email": "user@example.com",
    "password": "04225",
    "expire_in_seconds": 3600
}
```

`expire_in_seconds` means how long the user stays logged in before it requires to re-login. If `expire_in_seconds` is not given, it has a default value of 3600.

HTTP Response:

```json
{
  "id": "2b3916f0-56c3-49ac-b8ef-3d69ad017a94",
  "created_at": "2025-10-12T08:03:59.573532Z",
  "updated_at": "2025-10-12T08:03:59.573533Z",
  "email": "test@gmail.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiIyYjM5MTZmMC01NmMzLTQ5YWMtYjhlZi0zZDY5YWQwMTdhOTQiLCJleHAiOjE3NjAyNjAwNTQsImlhdCI6MTc2MDI1NjQ1NH0.Pq9KmT8-6WK6zq_MEwWCfllje1mwbK3GzlSG5dI80QM",
  "refresh_token": "bfb1543e695a2927a3044bfafb7b0b6d614df8e0033e02faf72ec0382d7e9c17",
  "is_chirpy_red": "false"
}
```

- `id` = The identifier for this resource.
- `created_at` = The time and date of when this resource is created.
- `updated_at` = The time and date of when this resource is last updated.
- `email` = The email of the user.
- `token` = The JWT token for the user, used for accessing certain endpoints.
- `refresh_token` = Token used for refreshing `expires_in_seconds` avoiding re-login user.
- `is_chirpy_red` = Is the user paying customer or not.

### PUT /api/users

Updates a user email and password given JSON and `token` for authentication:

```json
{
    "email": "user@example.com",
    "password": "04225"
}
```

Output:

```json
{
    "id": "2b3916f0-56c3-49ac-b8ef-3d69ad017a94",
    "email": "user@example.com",
    "updated_at": "2025-10-12T08:03:59.573533Z"
}
```
- `id` = The identifier for this resource.
- `email` = The email of the user.
- `updated_at` = The time and date of when this resource is last updated.

## Chirps (group)

The following group requires to have the following HTTP header:

```json
{
    "Authorization" : "Bearer <token>"
}
```
The `token` is generated from logging in as a user.

### POST /api/chirps

Creates a new chirp given JSON format:

```json
{
  "body": "Hello, world!",
  "user_id": "275cebdc-3cc0-4e86-927d-09cafe887fab"
}
```

Server response:

```json
{
  "id": "33ffab2b-e359-4366-8fb7-a722d2057e35",
  "created_at": "2025-10-11T09:03:13.747292Z",
  "updated_at": "2025-10-11T09:03:13.747292Z",
  "body": "Yo fam this feast is lit ong",
  "user_id": "275cebdc-3cc0-4e86-927d-09cafe887fab"
}
```

- `id` = The identifier for this resource.
- `created_at` = The time and date of when this resource is created.
- `updated_at` = The time and date of when this resource is last updated.
- `body` = The body text within the chirp.
- `user_id` = The id of the user.

### GET /api/chirps

Gets all the chirps within the database.

Server Response:

```json
[
  {
    "id": "3aca5e8d-36e6-4f34-b482-af34f03afc15",
    "created_at": "2025-10-10T08:08:16.700779Z",
    "updated_at": "2025-10-10T08:08:16.70078Z",
    "body": "This CLI kinda lit ngl",
    "user_id": "22630bcf-37e5-43bf-9282-a7748892b635"
  },
  {
    "id": "33ffab2b-e359-4366-8fb7-a722d2057e35",
    "created_at": "2025-10-11T09:03:13.747292Z",
    "updated_at": "2025-10-11T09:03:13.747292Z",
    "body": "Yo fam this feast is lit ong",
    "user_id": "f1900ef8-c33d-4601-b40b-937f371b5dbf"
  },
  {
    "id": "c3fad2f6-d01f-4eec-a5ba-dad46c91eabf",
    "created_at": "2025-10-11T09:29:49.107503Z",
    "updated_at": "2025-10-11T09:29:49.107503Z",
    "body": "Yo fam this feast is lit ong",
    "user_id": "f1900ef8-c33d-4601-b40b-937f371b5dbf"
  }
]
```

- `id` = The identifier for this resource.
- `created_at` = The time and date of when this resource is created.
- `updated_at` = The time and date of when this resource is last updated.
- `body` = The body text within the chirp.
- `user_id` = The id of the user.

### GET /api/chirps?author_id={user_id}

Gets chirps that is published by the specified user.

Server Response:

```json
[
  {
    "ID": "3aca5e8d-36e6-4f34-b482-af34f03afc15",
    "CreatedAt": "2025-10-10T08:08:16.700779Z",
    "UpdatedAt": "2025-10-10T08:08:16.70078Z",
    "Body": "This CLI kinda lit ngl",
    "UserID": "22630bcf-37e5-43bf-9282-a7748892b635"
  },
  {
    "ID": "33ffab2b-e359-4366-8fb7-a722d2057e35",
    "CreatedAt": "2025-10-11T09:03:13.747292Z",
    "UpdatedAt": "2025-10-11T09:03:13.747292Z",
    "Body": "Yo fam this feast is lit ong",
    "UserID": "f1900ef8-c33d-4601-b40b-937f371b5dbf"
  },
  {
    "ID": "c3fad2f6-d01f-4eec-a5ba-dad46c91eabf",
    "CreatedAt": "2025-10-11T09:29:49.107503Z",
    "UpdatedAt": "2025-10-11T09:29:49.107503Z",
    "Body": "Yo fam this feast is lit ong",
    "UserID": "f1900ef8-c33d-4601-b40b-937f371b5dbf"
  }
]
```

### DELETE /api/chirps/{chirpsID}

Deletes the chirp given specified chirps id. There is no body in the HTTP response but has the HTTP status of `204 No Content`.

## Tokens (groups)

This groups contains endpoints that changes the `token` or `refresh_token` of a user.


The following group requires to have the following HTTP header:

```json
{
    "Authorization" : "Bearer <token>"
}
```

### POST /api/refresh

Refreshed the `expire_in_seconds` (avoiding re-login) to 3600 seconds. It will give back a new `token` and checks if user is *revoked* (revoke, means the user cannot refresh the token). 

Server Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiIyYjM5MTZmMC01NmMzLTQ5YWMtYjhlZi0zZDY5YWQwMTdhOTQiLCJleHAiOjE3NjAyNjAwNTQsImlhdCI6MTc2MDI1NjQ1NH0.Pq9KmT8-6WK6zq_MEwWCfllje1mwbK3GzlSG5dI80QM",
}
```

### POST /api/revoke

Revokes the user, in other words the user cannot refresh the token, and is forced to re-login.

Does not have server response body, but will give HTTP status `204 No Content`.

## POST /api/polka/webhooks

Sets the user status of `is_chirpy_red` to be true. Requires the following JSON body:

```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "275cebdc-3cc0-4e86-927d-09cafe887fab"
  }
}
```

Additionally needs `polka` API Key as the header, to be able to acces this endpoint:

```json
{
    "Authorization": "ApiKey <API_Key>"
}
```

Server response will return an empty body and the HTTP status of `204 No Content`.
