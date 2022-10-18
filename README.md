# Project (Scalable Web Service with Golang - DTS Kominfo) - MyGram
Aplikasi untuk menyimpan foto dan membuat comment foto orang lain

## Register
Path: [http://localhost:8080/users/register ](http://localhost:8080//users/register)'

Request 

1. body:
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "password":"password",
    
    "username":"vika"
    
}
```

Response

1. status 201

2. Data
```
{

    "age":20,
    
    "email":"vika@gmail.com",
    
    "id":1,
    
    "username":"vika"
    
}
```

Method: Post
## Login
Path: [http://localhost:8080/users/login ](http://localhost:8080//users/login)

Method: Post

Request 

1. body:
```
{
  
    "email":"vika@gmail.com",
    
    "password":"password",
    
}
```

Response

1. status 200

2. Data
```
{

    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZpa2FAZ21haWwuY29tIiwiaWQiOjV9.LCpNoDe29cp6KUjclllQ0EPNwqqQbol8ibvaYC9GVMQ",
    
}
```

## Update User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Put
## Delete User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Delete
