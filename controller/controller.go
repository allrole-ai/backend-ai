package controller

func GetAllUser(mongoconn *mongo.Database, collection string) []UserNew {
	user := atdb.GetAllDoc[[]UserNew](mongoconn, collection)
	return user
}

func FindUserUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func Deleteuser(mongoconn *mongo.Database, collection string, userdata2 UserNew) interface{} {
	filter := bson.M{"username": userdata2.Username}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedUser(mongoconn *mongo.Database, collection string, filter bson.M, userdata2 UserNew) interface{} {
	updatedFilter := bson.M{"username": userdata2.Username}
	return atdb.ReplaceOneDoc(mongoconn, collection, updatedFilter, userdata2)
}

func EditUser(mongoenv *mongo.Database, collname string, datauser UserNew) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, datauser)
}

func GetAllUser(mongoconn *mongo.Database, collection string) []UserNew {
	user := atdb.GetAllDoc[[]UserNew](mongoconn, collection)
	return user
}

func FindUserUser(mongoconn *mongo.Database, collection string, userdata User) User {
	filter := bson.M{
		"username": userdata.Username,
	}
	return atdb.GetOneDoc[User](mongoconn, collection, filter)
}

func Deleteuser(mongoconn *mongo.Database, collection string, userdata2 UserNew) interface{} {
	filter := bson.M{"username": userdata2.Username}
	return atdb.DeleteOneDoc(mongoconn, collection, filter)
}

func UpdatedUser(mongoconn *mongo.Database, collection string, filter bson.M, userdata2 UserNew) interface{} {
	updatedFilter := bson.M{"username": userdata2.Username}
	return atdb.ReplaceOneDoc(mongoconn, collection, updatedFilter, userdata2)
}

func EditUser(mongoenv *mongo.Database, collname string, datauser UserNew) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mongoenv, collname, filter, datauser)
}


func GetUserFromID(db *mongo.Database, col string, _id primitive.ObjectID) (*UserNew, error) {
	cols := db.Collection(col)
	filter := bson.M{"_id": _id}

	userlist := new(UserNew)

	err := cols.FindOne(context.Background(), filter).Decode(userlist)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return nil, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}

	return userlist, nil
}

//login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("Dimohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("Email tidak valid")
	} 
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
	if err != nil {
		return 
	}
	