# avengers
Its a simple golang server which stores Avengers record in mongodb. You can perform CRUD operations on it. 


## Build locally
Run the following command to create application binary locally

```shell
$ make install
```

## Sample Request-Response
POST: http://localhost:8000/avengers/createNewAvenger
- Request Body:
```shell
{
    "name" : "sourav patnaik",
    "alias" : "sikan",
    "weapon": "hammer"
}
```
- Response:
```shell
{
    "InsertedID": "61de7b19c51fa632a454a881"
}
```

GET: http://localhost:8000/avengers/getAllAvengers
- Response:
```shell
[
    {
        "_id": "61de7b19c51fa632a454a881",
        "name": "sourav patnaik",
        "alias": "sikan",
        "weapon": "hammer"
    },
    {
        "_id": "61de88b8c51fa632a454a896",
        "name": "jack danniel",
        "alias": "jd",
        "weapon": "sword "
    }
]
```


PUT: http://localhost:8000/avengers/updateAvengerByName
- Request Body:
```shell
{
    "name": "sourav patnaik",
    "alias": "sikan, goodboy",
    "weapon": "hammer, axe"
}
```
- Response
```shell
{
    "MatchedCount": 1,
    "ModifiedCount": 1,
    "UpsertedCount": 0,
    "UpsertedID": null
}
```

DELETE: http://localhost:8000/avengers/deleteAvengerByName?name=sourav patnaik
- Response:
```shell
{
    "DeletedCount": 1
}
```
