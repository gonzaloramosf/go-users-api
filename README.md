## Go Users Api

Go CRUD with Users and Tasks Association  
  
- **Gorilla Mux**
- **Gorm**
- **PostgreSQL**
  
### Installation  
Requisites: 
- Go
- Docker
- PostgreSQL  
  
Clone repository:
```
https://github.com/gonzaloramosf/go-users-api.git
```

To use PostgreSQL with docker:
```
docker run --name go-users-api -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres
```

Then create the database:  
```
docker ps // list docker containers
docker exec -it go-users-api bash; // connect to container
psql -u root --password
// type your password (root in this case)
CREATE DATABSE gorm;
\c gorm // to connect to database
// type your password (root in this case)
```

In db/connection.go:
```
var DSN = "host=localhost user=root password=root dbname=gorm port=5432"
```


### Basic usage:
**Users:**  

Get all:
```
GET /users
```
Get by id:
```
GET /users/{id}
```
Create:
```
{
    name: "Jack",
    surname: "Musk",
    email: "jask.musk@gmail.com"
}
```
Delete:
```
DELETE /users/{id}
```

**Tasks:**  

Get all:
```
GET /tasks
```
Get by id:
```
GET /tasks/{id}
```
Create:
```
{
    title: "My task",
    description: "My task description",
    status: false,
    userId: 1
}
```
Delete:
```
DELETE /tasks/{id}
```



