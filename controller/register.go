package controller




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
