package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"net/url"
	"os"
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
