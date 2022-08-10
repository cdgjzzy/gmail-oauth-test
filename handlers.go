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

	errMsg, errExist := ctx.Get("error")
	if errExist {
		log.Fatalln("code callback error: ", errMsg)
	}

	code, codeExist := ctx.Get("code")
	if !codeExist {
		log.Fatalln("code not exist.")
	}
	log.Printf("code: %v", code)

	err := getToken(fmt.Sprintf("%v", code))
	if err != nil {
		log.Fatalln("get token err: ", err)
	}
}

func token(ctx *gin.Context) {

	log.Println("token callback server is being requested.")

	access_token, access_token_exist := ctx.Get("access_token")
	if !access_token_exist {
		log.Println("access_token not exist.")
	} else {
		log.Println("access_token: ", access_token)
	}

	state, state_exist := ctx.Get("state")
	if !state_exist {
		log.Println("state not exist.")
	} else {
		log.Println("state: ", state)
	}

	expires_in, expires_in_exist := ctx.Get("expires_in")
	if !expires_in_exist {
		log.Println("expires_in not exist.")
	} else {
		log.Println("expires_in: ", expires_in)
	}

	scope, scope_exist := ctx.Get("scope")
	if !scope_exist {
		log.Println("scope not exist.")
	} else {
		log.Println("scope: ", scope)
	}

	token_type, token_type_exist := ctx.Get("token_type")
	if !token_type_exist {
		log.Println("token_type not exist.")
	} else {
		log.Println("token_type: ", token_type)
	}

	refresh_token, refresh_token_exist := ctx.Get("refresh_token")
	if !refresh_token_exist {
		log.Println("refresh_token not exist.")
	} else {
		log.Println("refresh_token: ", refresh_token)
	}

}
