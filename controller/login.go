package controller



// Login handles user login
func Login(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
	var user model.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}