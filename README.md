# Project (Scalable Web Service with Golang - DTS Kominfo) - MyGram
Aplikasi untuk menyimpan foto dan membuat comment foto orang lain

## Register
Path: [http://localhost:8080/users/register ](http://localhost:8080//users/register)

Method: Post
## Login
Path: [http://localhost:8080/users/login ](http://localhost:8080//users/login)

Method: Post

Request 

body:
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "password":"password",
    
    "username":"vika"
    
}
```

Response

status 201

data
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "id":1,
    
    "username":"vika"
    
}
```

## Update User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Put
## Delete User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Delete
