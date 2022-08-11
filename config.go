package main

import "fmt"

var (
	scope               = "https://mail.google.com/"    // 阅读、撰写、发送及永久删除您的全部 Gmail 电子邮件
	response_type       = "code"                        // 授权码模式
	state               = "userID"                      // 授权服务器会原封不动的返回,可以作为定位用户的字段
	code_redirect_uri   = "http://localhost:5000/code"  // 重定向网址，用于接收授权code响应
	token_redirect_uri  = "http://localhost:5000/token" // 重定向网址，用于接收授权token响应
	client_id           = ""                            // 客户端ID
	client_secret       = ""                            // 客户端密钥
	authUrlFormat       = "https://accounts.google.com/o/oauth2/v2/auth?scope=%s&access_type=offline&response_type=%s&state=%s&redirect_uri=%s&client_id=%s"
	tokenUri            = "https://oauth2.googleapis.com/token"
	googleOAuthEndpoint = fmt.Sprintf(authUrlFormat, scope, response_type, state, code_redirect_uri, client_id)
)
