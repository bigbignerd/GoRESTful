package main

import (
    "net/http"
    "log"
    "github.com/bigbignerd/GoRESTful/router"
    "github.com/bigbignerd/GoRESTful/config"

    "github.com/gin-gonic/gin"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
)

var (
    cfg = pflag.StringP("config", "c", "", "apisever cofig file path")
)
func main(){
    pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

    g := gin.New()
    middlewares := []gin.HandlerFunc{}

    router.Load(
        g,
        middlewares...,
    )
    log.Printf("start to listening the incoming requests on http address %s",viper.GetString("addr")) 
    log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
