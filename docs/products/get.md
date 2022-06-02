# Get All Products

Show all Products that available in system.

**URL** : `/api/v1/products/`

**Method** : `GET`

**Auth required** : YES

**Permissions required** : None

**Data constraints** : `{}`

## Success Responses

**Condition** : User can products

**Code** : `200 OK`

**Content** :

```json
[
    {
        "status": 200,
        "message": "success",
        "data": {
        "Products": [
            {
                "ID": 1,
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "Name": "sepeda",
                "Category": "olahraga",
                "Price": 25000000
            },
            {
                "ID": 2,
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "Name": "bola basket",
                "Category": "olahraga",
                "Price": 200000
            },
            {
                "ID": 3,
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "Name": "panci",
                "Category": "dapur",
                "Price": 500000
            }
        ]
    }
}
]
```