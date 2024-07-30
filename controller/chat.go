package controller

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func Login(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {}

