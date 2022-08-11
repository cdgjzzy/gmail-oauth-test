package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	err := getToken(ctx, fmt.Sprintf("%v", code))
	if err != nil {
		log.Fatalln("get token err: ", err)
	}

	time.Sleep(2000)

	if tokenResponse != "" {
		ctx.Redirect(302, "http://localhost:5000/display")
	}
}

func display(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tokenResponse)
}
