package controller






















if err == nil && errorResponse["error"] == "Model is currently loading" {
			retryCount++
			time.Sleep(retryDelay)
			continue
		}
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
		return
	}
}

if response.StatusCode() != 200 {
	helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error from Hugging Face API "+string(response.Body()))
	return
}

var data []map[string]interface{}

	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		if len(data) > 0 {
			generatedText, ok := data[0]["generated_text"].(string)
			if !ok {
				helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "error extracting generated text")
				return
			}
			helper.WriteJSON(respw, http.StatusOK, map[string]string{"answer": generatedText})
	} else {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server: response")
	}
}
