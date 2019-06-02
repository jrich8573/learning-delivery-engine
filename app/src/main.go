package main

import (
        //"encoding/json"
        //"errors"
        //"fmt"
        //"log"
        "net/http"
        //"os"
        "strconv"
        "github.com/gin-gonic/contrib/static"
        "github.com/gin-gonic/gin"
)


type User struct {
    ID int `json: "id" binding:"required"`
    FirstName string `json: "firstname"`
    MiddleName string `json: "middlename"`
    LastName string `json: "lastname"`
    Organization []string `json: "organization"`
    Logins int `json: "logins"` //count the number if logins
}

var users = []User{
    // ToDo: build a scalable users arrary
}



func main(){
    // Set the router as the default one shipped with Gin
    router := gin.Default()
    // Serve frontend static files
    router.Use(static.Serve("/", static.LocalFile("./views", true)))

    // Setup route group for the API
    api := router.Group("/api")
    api.GET("/", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
              "message": "pong",
        })
    })

    // Start and run the server
    router.Run(":3000")

     // Our API will consit of just two and possibly more routes
     api.GET("/users", userHandler)
     api.POST("/users/:userID", getUsers)
}
     // JokeHandler retrieves a list of available
func userHandler(c *gin.Context){
        c.Header("Content-Type", "application/json")
        c.JSON(http.StatusOK, gin.H{
            "message":"User handler is not implemented yet",
        })
}

// user handler to retrieve valid users
func getUsers(c *gin.Context){
    //check for a valid user-id
    if userid, err := strconv.Atoi(c.Param("userID")); err == nil {
        //find user and increment logins
        for i :=0; i < len(users); i++ {
            if users[i].ID == userid {
              users[i].Logins = users[i].Logins + 1
            }
        }
        c.JSON(http.StatusOK, &users)
    } else {
        // new user
        c.AbortWithStatus(http.StatusNotFound)
      }
}
