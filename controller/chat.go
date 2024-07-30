package controller

import (
	"net/http"

	"github.com/allrole-ai/backend-ai/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {}
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}
	

