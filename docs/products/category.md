# Get Products by category

Show all products by category that available in system.

**URL** : `/api/v1/products/?category=`

**Method** : `GET`

**Auth required** : YES

**Permissions required** : None

**Parameters**: 
```
{
    name:category
    type:string 
}
```
**Data constraints** : `{}`

## Success Responses

**Condition** : User can see products by category

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
            }
        ]
    }
]
```