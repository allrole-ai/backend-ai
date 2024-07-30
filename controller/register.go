

	





	if strings.Contains(user.Password, " ") {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password tidak boleh mengandung spasi")
		return
	}
	if len(user.Password) < 8 {
		helper.ErrorResponse(respw, req, http.StatusBadRequest, "Bad Request", "password minimal 8 karakter")
		return
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : salt")
		return
	}
	hashedPassword := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	userData := bson.M{
		"namalengkap": user.NamaLengkap,
		"email":       user.Email,
		"password":    hex.EncodeToString(hashedPassword),
		"salt":        hex.EncodeToString(salt),
	}
	insertedID, err := helper.InsertOneDoc(db, col, userData)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : insert data, "+err.Error())
		return
	}
	resp := map[string]any{
		"message":    "berhasil mendaftar",
		"insertedID": insertedID,
		"data": map[string]string{
			"email": user.Email,
		},
	}
	helper.WriteJSON(respw, http.StatusCreated, resp)
}
