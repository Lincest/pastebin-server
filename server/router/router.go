package router

/**
    server
    @author: roccoshi
    @desc: router init
**/

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	ctl "server/controller"
	"server/model"
)

func LoadConfig() {
	// load config
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, &model.Config)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
	log.Printf("Init Config Successfully with conf = %#v", model.Config)
}

func InitRoutes() {
	router := gin.New()
	// 中间件
	// 允许跨域: https://github.com/gin-contrib/cors
	router.Use(gin.Logger(), gin.Recovery(), cors.Default())
	Conf := router.Group("pastebin")
	{
		Conf.GET("/:id", ctl.PastebinGetAction)
		Conf.POST("", ctl.PastebinSetAction)
	}
	_ = router.Run(":" + model.Config.Port)
}
