# Simple Web app

## Endpoints

BASE_URL = "http://localhost:8080"

### Public Endpoints

URL | Method | Description
/ping | GET | Check status
/login | POST |returns a jwt string

### Private Endpoints /api/\*

Use jwt authentication token from /login as Bearer Token.

```json
{
  "userID": 4
}
```

| URL                             | Method | Description                                                                          |
| ------------------------------- | ------ | ------------------------------------------------------------------------------------ |
| /api/renew                      | GET    | Returns new token with the currently logged in user                                  |
| /api/v1/appointment?page=1      | GET    | returns all appointments, paginated with &page=:int query                            |
| /api/v1/appointment/:id         | GET    | returns the appointment with :id, used for frontend to get data                      |
| /api/v1/appointment/comment/:id | GET    | returns all comments associated with the appointment id (decoupled from appointment) |
| /api/v1/appointment/archive/:id | PATCH  | Archive the appointment                                                              |
| /api/v1/appointment             | POST   | Create new appointment                                                               |

---

POST /api/v1/appointment

```json
{
  "title": "title01",
  "description": "lorem 01"
}
```

---

PATCH /api/v1/appointment/:id

```json
{
  "id": 7,
  "title": "title007",
  "description": "desc 07",
  "status": 2
}
```
