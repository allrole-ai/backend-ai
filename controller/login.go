package controller




	hash := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		helper.ErrorResponse(respw, req, http.StatusUnauthorized, "Unauthorized", "password salah")
		return
	}
	tokenstring, err := helper.Encode(user.ID, user.Email, privatekey)
	if err != nil {
		helper.ErrorResponse(respw, req, http.StatusInternalServerError, "Internal Server Error", "kesalahan server : token")
		return
	}
	resp := map[string]any{
		"status":  "success",
		"message": "login berhasil",
		"token":   tokenstring,
		"data": map[string]string{
			"email":       existsDoc.Email,
			"namalengkap": existsDoc.NamaLengkap,
		},
	}
	helper.WriteJSON(respw, http.StatusOK, resp)
}
