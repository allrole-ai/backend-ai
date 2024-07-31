package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
)

// Config dan helper fungsi dapat diimplementasikan sesuai kebutuhan.

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/chat", Chat).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Chat(respw http.ResponseWriter, req *http.Request) {
	// Mengganti "tokenmodel" dengan token API yang sebenarnya
	tokenmodel := "your_api_token_here"

	var chat AIRequest
	err := json.NewDecoder(req.Body).Decode(&chat)
	if err != nil {
		ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if chat.Query == "" {
		ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	client := resty.New()
	apiUrl := "https://api-inference.huggingface.co/models/your_model_here" // Ganti dengan URL model Hugging Face yang sebenarnya
	apiToken := "Bearer " + tokenmodel

	var response *resty.Response
	var retryCount int
	maxRetries := 5
	retryDelay := 20 * time.Second

	parsedURL, err := url.Parse(apiUrl)
	if err != nil {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing URL model hugging face"+err.Error())
		return
	}

	segments := strings.Split(parsedURL.Path, "/")
	modelName := strings.Join(segments[2:], "/")

	for retryCount < maxRetries {
		response, err = client.R().
			SetHeader("Authorization", apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(`{"inputs": "` + chat.Query + `"}`).
			Post(apiUrl)

		if err != nil {
			log.Fatalf("Error making request: %v", err)
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
			ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
			return
		}
	}

	if response.StatusCode() != 200 {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
		return
	}

	var data []map[string]interface{}
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error parsing response body "+err.Error())
		return
	}

	if len(data) > 0 {
		generatedText, ok := data[0]["generated_text"].(string)
		if !ok {
			ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
			return
		}
		WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
	} else {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}

// Fungsi untuk menangani respon JSON
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Fungsi untuk menangani respon error
func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, errorType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": errorType, "message": message})
}

// Struktur data untuk request
type AIRequest struct {
	Query string `json:"query"`
}

// Fungsi untuk menangani respon error
func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, errorType, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": errorType, "message": message})
}

// Struktur data untuk request
type AIRequest struct {
	Query string `json:"query"`
}
