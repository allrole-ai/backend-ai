package controller

import (
	"net/http"

	"github.com/allrole-ai/backend-ai/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {}
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)


