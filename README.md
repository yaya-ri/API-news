# API-news
Learn to build API POST and GET News using SQL, GIN, GORM, Elasticsearch and RabbitMQ

# Setup for local
1. Run docker-compose.yml:
```
docker-compose up -d
```

2. copy `.env.example` to `.env` and remove command or you can copy 
```
APP_PORT=8070
APP_ENV=development

SQL_HOST=localhost
SQL_PORT=3307
SQL_DATABASE=news
SQL_USER=root
SQL_PASSWORD=admin

ES_HOST=localhost
ES_PORT=9207

RABBITMQ_HOST=localhost
RABBITMQ_PORT=5677
RABBITMQ_USER=admin
RABBITMQ_PASSWORD=admin
```
in your `.env`

3. run API with:
```
refresh run
```
or
```
go run .
```
or
```
go build
./API-news
```

4. Unit test {i am sorry, the unit test is not complete :( }, bu you can check with:
```
go test -v ./...
```

# Documentation API-news
`POST /news`

request body:
```
{
    "author": "yaya",
    "body": "body yaya"
}
```

response:
```
{
    "code": 200,
    "error": null,
    "data": {},
    "message": "Success store news",
    "status": "success"
}
```

`GET /news` or `GET /news?page=1&limit=3` 
result with cacher 5 second

response
```
{
    "code": 200,
    "error": null,
    "data": {
        "current_page": 1,
        "data": [
            {
                "id": 4,
                "author": "yayaaaaa",
                "body": "Lorem ipsum dolor",
                "created": "2020-03-09 11:01:23"
            },
            {
                "id": 3,
                "author": "yayaaaaa",
                "body": "Lorem ipsum dolor",
                "created": "2020-03-09 11:01:21"
            },
            {
                "id": 2,
                "author": "yayaaaaa",
                "body": "Lorem ipsum dolor",
                "created": "2020-03-09 11:01:19"
            }
        ]
    },
    "message": "Success get news",
    "status": "success"
}
```

NOTE:
if there is any problem with GET, please:
```
docker-compose down
```

then start again from number 1 (`docker-compose up`), 

i thinks the problem is on elasticsearch service test. cause i am still learning and i don't know how to mock elasticsearch test. so this test direct access to elasticsearch.
