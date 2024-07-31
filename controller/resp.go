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

		if generatedText == "" {
			WriteJSON(respw, http.StatusNoContent, map[string]string{"message": "No content generated"})
		} else if generatedText == "special condition" { // Contoh kondisi khusus
			WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText, "note": "This is a special response"})
		} else {
			WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
		}
	} else {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
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

		if generatedText == "" {
			WriteJSON(respw, http.StatusNoContent, map[string]string{"message": "No content generated"})
		} else if generatedText == "special condition" { // Contoh kondisi khusus
			WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText, "note": "This is a special response"})
		} else {
			WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
		}
	} else {
		ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}
