// package test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/allrole-ai/backend-ai/config"
// 	helper "github.com/allrole-ai/backend-ai/helper"
// 	"github.com/allrole-ai/backend-ai/model"
// 	module "github.com/allrole-ai/backend-ai/module"
// )

// var db = module.MongoConnect("MONGOSTRING", "all")

// func TestGenerateKey(t *testing.T) {
// 	privateKey, publicKey := helper.GenerateKey()
// 	t.Logf("PrivateKey : %v", privateKey)
// 	t.Logf("PublicKey : %v", publicKey)
// }

// // TestInsertOneDoc
// func TestInsertOneDoc(t *testing.T) {
// 	var data = map[string]interface{}{
// 		"username": "allrole",
// 		"password": "ganteng123",
// 	}
// 	insertedDoc, err := helper.InsertOneDoc(config.Mongoconn, "users", data)
// 	if err != nil {
// 		t.Errorf("Error : %v", err)
// 	}
// 	t.Logf("InsertedDoc : %v", insertedDoc)
// }

// func TestRegister(t *testing.T) {
// 	var doc model.User
// 	doc.NamaLengkap = "Riziq"
// 	doc.Email = "riziq@gmai.com"
// 	doc.Password = "qobel123"
// 	doc.Confirmpassword = "riziq123"
// 	email, err := module.Register(db, "user", doc)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 		fmt.Println("Data berhasil disimpan dengan email:", email)
// 	}
// }

// func TestLogin(t *testing.T) {
// 	var user model.User
// 	user.Email = "riziq@gmail.com"
// 	user.Password = "riziq123"
// 	user, err := module.Login(db, user)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Berhasil LogIn : ", user.Email)
//  	}
//  }