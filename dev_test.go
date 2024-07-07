package test

import (
	"testing"

	"github.com/allrole-ai/backend-ai/config"
	helper "github.com/allrole-ai/backend-ai/helper"
	module "github.com/allrole-ai/backend-ai/module"
)

var db = module.MongoConnect("MONGOSTRING", "allrole-ai",)

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := helper.GenerateKey()
	t.Logf("PrivateKey : %v", privateKey)
	t.Logf("PublicKey : %v", publicKey)
}

// TestInsertOneDoc
func TestInsertOneDoc(t *testing.T) {
	var data = map[string]interface{}{
		"username": "allrole",
		"password": "ganteng123",
	}
	insertedDoc, err := helper.InsertOneDoc(config.Mongoconn, "users", data)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	t.Logf("InsertedDoc : %v", insertedDoc)
}

func TestRegister(t *testing.T) {
	var doc model.User
	doc.NamaLengkap = "Fahad Abdul Aziz"
	doc.Email = "fahad@gmail.com"
	doc.Password = "fahad#123"
	doc.Confirmpassword = "fahad#123"
	email, err := module.Register(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan email:", email)
	}
}
