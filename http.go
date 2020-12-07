package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/get", HandleGet)
	r.POST("/post", HandlePost)
	r.PUT("/put", HandlePut)
	r.DELETE("/delete", HandleDelete)
	r.PATCH("/patch", HandlePatch)
	r.HEAD("/head", HandleHead)
	r.OPTIONS("/options", HandleOptions)
	r.Run(":9000")

}

func HandleGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"receive": "65536",
	})

}

func HandlePut(C *gin.Context) {
	body, _ := ioutil.ReadAll(C.Request.Body)
	C.String(http.StatusOK, "Put,%s", body)

}

func HandleDelete(C *gin.Context) {
	body, _ := ioutil.ReadAll(C.Request.Body)
	C.String(http.StatusOK, "Delete,%s", body)

}

func HandleOptions(C *gin.Context) {
	body, _ := ioutil.ReadAll(C.Request.Body)
	C.String(http.StatusOK, "Options,%s", body)

}

func HandlePatch(C *gin.Context) {
	body, _ := ioutil.ReadAll(C.Request.Body)
	C.String(http.StatusOK, "patch,%s", body)

}

func HandlePost(C *gin.Context) {
	body, _ := ioutil.ReadAll(C.Request.Body)
	C.String(http.StatusOK, "post,%s", body)

}

func HandleHead(C *gin.Context) {
	// http head only response  header
	fmt.Printf("head http \r\n")

}
