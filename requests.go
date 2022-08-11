package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getToken(code string) (err error) {

	log.Println("get token request")

	bodyFormat := "code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code"
	var body = strings.NewReader(fmt.Sprintf(bodyFormat, code, client_id, client_secret, code_redirect_uri))
	response, err := http.Post(tokenUri, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response error code:%v", response.StatusCode)
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println("get token resp:", string(respBody))
	return nil
}

func refreshToken(refresh_token string) error {

	log.Println("refresh token request")

	bodyFormat := "client_id=%s&client_secret=%s&refresh_token=%s&grant_type=refresh_token"
	var body = strings.NewReader(fmt.Sprintf(bodyFormat, client_id, client_secret, refresh_token))
	response, err := http.Post(tokenUri, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response error code:%v", response.StatusCode)
	}

	return nil
}
