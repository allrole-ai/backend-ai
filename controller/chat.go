package controller

import (
	"encoding/json"
	"log"
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
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body: "+err.Error())
		return
	}

	if chat.Prompt == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}
	// if chat.Query == "" {
	// 	helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
	// 	return
	// }

	client := resty.New()

	// URL API Hugging Face dan token
	apiUrl := config.GetEnv("HUGGINGFACE_API_KEY")
	apiToken := "Bearer " + tokenmodel

	var response *resty.Response
	var retryCount int
	maxRetries := 5
	retryDelay := 20 * time.Second

	

	if response.StatusCode() != http.StatusOK {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API: "+string(response.Body()))
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body: "+err.Error())
		return
	}

	// Logging untuk memeriksa struktur data yang diterima
	log.Printf("Response data: %v", data)

	if len(data) > 0 {
		if answer, ok := data[0]["answer"].(string); ok {
			helper.WriteJSON(respw, http.StatusOK, map[string]string{"answer": answer})
			return
		}
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting answer")
	} else {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}
