# Register

Used to register and obtain a JWT Token for a new User.

**URL** : `/api/v1/register`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "username": "[valid username]",
    "password": "[password in plain text]",
    "name": "[name of user]"
}
```

**Data example**

```json
{
    "username": "this_is_username",
    "password": "this_is_passowrd",
    "name": "this_is_name_of_user"
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