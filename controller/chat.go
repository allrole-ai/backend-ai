package controller





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
