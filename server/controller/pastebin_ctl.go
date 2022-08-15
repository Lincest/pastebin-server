package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/model"
)

/**
    controller
    @author: roccoshi
    @desc: pastebin 存储和返回
**/

func PastebinGetAction(c *gin.Context) {
	resp := model.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	id := c.Param("id")
	log.Printf("pastebin get id = %s", id)
}
