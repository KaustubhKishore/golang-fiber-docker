# golang-fiber-docker
A sample API built using Golang Fiber Framework which is then containerised through Docker

> docker build -t golangfiber .

> docker run -d -p3000:3000 golangfiber

##### Get all books
`GET`
> localhost:3000/all

##### Get book through id
`GET`
> localhost:3000/api/getonebook/1

##### Add Book
`POST - JSON in raw body (postman)`
> localhost:3000/api/addonebook

```JSON
{
    "BookId": 3,
    "BookName": "Kaustubh",
    "BookAuthor": "Kishore"
}
```

##### Delete book through ID
`DELETE`
> localhost:3000/api/deletebook/1
