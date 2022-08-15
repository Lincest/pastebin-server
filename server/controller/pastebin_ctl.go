package controller

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/model"
	"server/service"
	"time"
	"unsafe"
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
	log.Printf("pastebin get id = %s\n", id)
	res, err := service.Rdb.Get(id).Result()
	if err != nil {
		log.Printf("error with id %s, err = %s\n", id, err)
		resp.Msg = "unexpected id: " + id
		resp.Code = model.CodeErr
		return
	}
	resp.Data = res
}

func PastebinSetAction(c *gin.Context) {
	resp := model.NewBasicResp()
	defer c.JSON(http.StatusOK, resp)
	var data model.Paste
	if err := c.ShouldBind(&data); err != nil {
		resp.Msg = err.Error()
		resp.Code = model.CodeErr
		return
	}
	log.Printf("data = %#v\n", data)
	// check size
	if unsafe.Sizeof(data) > 1024*1024 {
		resp.Msg = "data too big, should <= 1MB!"
		resp.Code = model.CodeErr
		return
	}
	b := make([]byte, 4)
	rand.Read(b)
	key := hex.EncodeToString(b)
	expired := data.Expired
	if expired == 0 {
		expired = int8(7) // default set 7 days
	}
	if expired > 30 {
		resp.Msg = "最多给你保存30天"
		resp.Code = model.CodeErr
		return
	}
	_, err := service.Rdb.Set(key, data.Data, time.Duration(expired)*time.Hour*24).Result()
	if err != nil {
		resp.Msg = "redis保存失败: " + err.Error()
		resp.Code = model.CodeErr
		return
	}
	log.Printf("保存成功, 保存的key为: %s, 保存的天数为: %d\n", key, expired)
	resp.Msg = "保存成功"
	resp.Data = key
}
