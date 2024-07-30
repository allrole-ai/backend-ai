package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/allrole-ai/backend-ai/config"
	"github.com/allrole-ai/backend-ai/helper"
	"github.com/allrole-ai/backend-ai/model"
	"github.com/go-resty/resty/v2"
)

func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
	var chat model.AIRequest
	err := json.NewDecoder(req.Body).Decode(&chat)

	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}
	
	if chat.Query == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}
	
	client := resty.New()

	apiUrl := config.GetEnv("HUGGINGFACE_API_URL")

	apiToken := "Bearer " + tokenmodel

	var response *resty.Response
	var retryCount int

	maxRetries := 5
	retryDelay := 20 * time.Second

	for retryCount < maxRetries {
	response, err = client.R().
    SetHeader("Authorization", apiToken).
    SetHeader("Content-Type", "application/json").
    SetBody({"inputs": " + chat.Query + "}).
    Post(apiUrl)

	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	

}
