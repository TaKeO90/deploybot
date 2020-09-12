package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const RESPFILENAME string = "webhook.json"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var mtx *sync.Mutex = &sync.Mutex{}

func write(c *gin.Context) {
	mtx.Lock()
	defer mtx.Unlock()
	reqBody := c.Request.Body
	data, err := ioutil.ReadAll(reqBody)
	checkError(err)
	if len(data) != 0 {
		err := ioutil.WriteFile(RESPFILENAME, data, 0600)
		checkError(err)
	}
}

func checkFileInDir() bool {
	dir, err := ioutil.ReadDir(".")
	checkError(err)
	for _, f := range dir {
		if f.Name() == RESPFILENAME {
			return true
		}
	}
	return false
}

func getReq(c *gin.Context) {
	mtx.Lock()
	defer mtx.Unlock()
	interSrvErr := struct {
		ErrorMsg string `json:"ErrorMsg"`
	}{
		ErrorMsg: "No response yet from telegram, or webhook setup is not correct",
	}
	if checkFileInDir() {
		data, err := ioutil.ReadFile(RESPFILENAME)
		if err != nil {
			log.Fatal(err)
		}
		if len(data) > 0 {
			c.Data(http.StatusOK, "application/json", data)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	} else {
		c.JSON(http.StatusInternalServerError, interSrvErr)
	}
}
