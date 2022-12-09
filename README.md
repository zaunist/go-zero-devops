# go-zero-devops

## deploy with docker-compose

```
cd deploy && docker compose up -d
```

## API example

### Default Response

```json
{
  "ok": true,
  "error": ""
}
```

#### Login

+ Api: /login
+ Method: POST
+ RequestBody:

```json
{
  "username": "",
  "password": ""
}
```

#### Add user

+ Api: /users
+ Method: POST
+ RequestBody:

```json
{
  "username": "",
  "password": ""
}
```

#### Delete user

+ Api: /users/id/:id
+ Method: DELETE

#### Update user

+ Api: /users/id/:id
+ Method: PUT

```json
{
  "username": "",
  "password": ""
}
```

#### Get user

+ Api: /users/id/:id
+ Method: GET
+ Response:

```json
{
  "id": 0,
  "username": "",
  "password": ""
  }
```

#### Get all user

+ Api: /users
+ Method: GET
+ Response:

```json
[
{
  "id": 0,
  "username": "",
  "password": ""
  }
]
```
