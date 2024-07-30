package controller

import (
	"encoding/json"
	"net/http"

	"github.com/allrole-ai/backend-ai/model"
)

func Chat(respw http.ResponseWriter, req *http.Request, tokenmodel string) {
	var chat model.AIRequest
	err := json.NewDecoder(req.Body).Decode(&chat)

}
