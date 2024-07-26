package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
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

	if chat.Query == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "dimohon untuk melengkapi data")
		return
	}

	client := resty.New()

	// URL API Hugging Face dan token
	apiUrl := config.GetEnv("HUGGINGFACE_API_KEY")
	apiToken := "Bearer " + tokenmodel

	parsedURL, err := url.Parse(apiUrl)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model Hugging Face: "+err.Error())
		return
	}

	segments := strings.Split(parsedURL.Path, "/")
	modelName := strings.Join(segments[2:], "/")

	// Request ke API Hugging Face dengan retry logic
	var response *resty.Response
	retryCount := 0
	maxRetries := 5
	retryDelay := 20 * time.Second

	for retryCount < maxRetries {
		response, err = client.R().
			SetHeader("Authorization", apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(`{"inputs": "` + chat.Query + `"}`).
			Post(apiUrl)

		if err != nil {
			log.Printf("Error making request: %v", err)
			retryCount++
			time.Sleep(retryDelay)
			continue
		}

		if response.StatusCode() == http.StatusOK {
			break
		} else {
			var errorResponse map[string]interface{}
			err = json.Unmarshal(response.Body(), &errorResponse)
			if err == nil && errorResponse["error"] == "Model "+modelName+" is currently loading" {
				retryCount++
				time.Sleep(retryDelay)
				continue
			}
			helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API: "+string(response.Body()))
			return
		}
	}

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
