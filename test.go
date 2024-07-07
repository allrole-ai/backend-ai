package test

import (
	"testing"

	"github.com/allrole-ai/backend-ai/config"
	helper "github.com/allrole-ai/backend-ai/helper"
)

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