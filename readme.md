
## Golang Rest API Starter

**Clone Environment Repository  :**

    go get -u "github.com/jinzhu/gorm"
    go get -u "github.com/gin-gonic/gin"
    go get -u "github.com/gin-gonic/gin"
    go get -u "github.com/jinzhu/gorm/dialects/postgres"

**Clone our Repository  :**
 https://github.com/fascalsj/golang-rest-starter.git 

**Start Application**
To run the application you must provide postgres and adjust the database settings, then run the command:

    go run main.go

**Technology :**

> Rest Framework : Gin-Gonic
> ORM : GORM

**File Structure**
We are use using controller, model, and config. Controller handle url request, Model used for mapping json body and database,  Configuration used for save all configuration of project
```
Golang Starter
├── config
│ ├── dbconfig.go
│ └── response.go
├── controllers
│ ├── gorm.go
│ └── usercontroller.go
├── model
│ └── usermodel.go
├── main.go
└── readme.md

```
