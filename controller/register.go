package controller



func Register(db *mongo.Database, col string, respw http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}
	if user.NamaLengkap == "" || user.Email == "" || user.Password == "" || user.Confirmpassword == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}
