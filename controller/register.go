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
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email tidak valid")
		return
	}
	userExists, _ := helper.GetUserFromEmail(user.Email, db)
	if user.Email == userExists.Email {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "email sudah terdaftar")
		return
	}
	if len(user.Password) < 8 {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password minimal 8 karakter")
		return
	}
	if strings.Contains(user.Password, " ") {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password tidak boleh mengandung spasi")
		return
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
