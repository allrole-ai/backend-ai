package controller

// "net/url"
// "os"
// "strings"

func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
	var chat model.AIRequest