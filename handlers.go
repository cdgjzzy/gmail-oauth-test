package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirectHandler(ctx *gin.Context) {
	// 重定向到OAuth同意屏幕，一定要使用302，301会被浏览器缓存
	ctx.Redirect(http.StatusFound, googleOAuthEndpoint)
}

func code(ctx *gin.Context) {
	log.Println("code callback server is being requested.")

	errMsg, errExist := ctx.GetQuery("error")
	if errExist {
		log.Fatalln("code callback error: ", errMsg)
	}

	code, codeExist := ctx.GetQuery("code")
	if !codeExist {
		log.Fatalln("code not exist.")
	}
	log.Printf("code: %v", code)

	err := getToken(fmt.Sprintf("%v", code))
	if err != nil {
		log.Fatalln("get token err: ", err)
	}
}
