package main

import (
	"server/router"
	"server/service"
)

/**
    server
    @author: roccoshi
    @desc: main
**/

func main() {
	//logFile, err := os.OpenFile("./log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Fatal("open log file failed, err:", err)
	//	return
	//}
	//log.SetOutput(logFile)
	router.LoadConfig()
	service.InitRedis()
	router.InitRoutes()
}
