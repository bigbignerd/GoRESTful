package main

import (
    "net/http"
    "log"
    "github.com/bigbignerd/GoRESTful/router"
    "github.com/gin-gonic/gin"
)

func main(){
    g := gin.New()
    middlewares := []gin.HandlerFunc{}

    router.Load(
        g,
        middlewares...,
    )
    log.Print("start to listening the incoming requests on http address %s", ":8080")
    log.Print(http.ListenAndServe(":8080", g).Error())
}
