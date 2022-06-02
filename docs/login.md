# Login

Used to get a JWT Token for a registered User.

**URL** : `/api/v1/login/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "username": "[valid username]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "username": "this_is_username",
    "password": "this_is_passowrd"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{   
    "status": 200,
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im15bmFtZWlzIiwiZXhwIjoxNjU0NDcyNjY3fQ.6F74-GCleoBb3_dukjanGetVweW5VcF9vy13CcCyyKk"
    }
    
}
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "status": 400,
    "message": "Cannot login with provided username and password",
}
```