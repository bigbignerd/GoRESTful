package main

import (
	"github.com/bigbignerd/GoRESTful/config"
	"github.com/bigbignerd/GoRESTful/model"
	"github.com/bigbignerd/GoRESTful/router"
	"github.com/lexkong/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apisever cofig file path")
)

func main() {
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
	//init db
	model.DB.Init()
	defer model.DB.Close()

	log.Infof("start to listening the incoming requests on http address %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
