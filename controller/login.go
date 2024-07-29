package controller



// Login handles user login
func Login(db *mongo.Database, respw http.ResponseWriter, req *http.Request, privatekey string) {
	var user model.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "error parsing request body "+err.Error())
		return
	}

	if user.Email == "" || user.Password == "" {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "mohon untuk melengkapi data")
		return
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}

	existsDoc, err := helper.GetUserFromEmail(user.Email, db)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : get email "+err.Error())
		return
	}

	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}