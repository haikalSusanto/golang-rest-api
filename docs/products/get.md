# Get All Products

Show all Products that available in database.

**URL** : `/api/v1/products/`

**Method** : `GET`

**Auth required** : YES

**Permissions required** : None

**Data constraints** : `{}`

## Success Responses

**Condition** : User can see products

**Code** : `200 OK`

**Content** :

```json
[
    {
        "status": 200,
        "message": "success",
        "data": [
            {
                "ID": 1,
                "Name": "sepeda",
                "Category": "olahraga",
                "Price": 25000000
            },
            {
                "ID": 2,
                "Name": "bola basket",
                "Category": "olahraga",
                "Price": 200000
            },
            {
                "ID": 3,
                "Name": "panci",
                "Category": "dapur",
                "Price": 500000
            }
        ]
    }
]
```