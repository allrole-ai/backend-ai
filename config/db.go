package config

import (
	"github.com/allrole-ai/backend-ai/helper"
)

var MongoString string = GetEnv("MONGOSTRING")

var mongoinfo = helper.DBInfo{
	DBString: MongoString,
	DBName:   "db_allrole",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
