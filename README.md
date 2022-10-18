# Project (Scalable Web Service with Golang - DTS Kominfo) - MyGram
Aplikasi untuk menyimpan foto dan membuat comment foto orang lain

# Init Project
go mod tidy

# Install Packages
go get github.com/asaskevich/govalidator
go get golang.org/x/crypto/bcrypt  
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/gin-gonic/gin 

## Register
Path: [http://localhost:8080/users/register ](http://localhost:8080//users/register)

Method: Post
## Login
Path: [http://localhost:8080/users/login ](http://localhost:8080//users/login)

Method: Post

## Update User
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Put
## Delete
Path: [http://localhost:8080/users ](http://localhost:8080//users)

Method: Delete
