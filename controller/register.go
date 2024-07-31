package controller



func Register(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User
