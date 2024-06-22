package goframework_leveldb_test

import (
	goframeworkleveldb "github.com/kordar/goframework-leveldb"
	logger "github.com/kordar/gologger"
	"testing"
)

func TestAddLevelDbInstance(t *testing.T) {
	goframeworkleveldb.AddLevelDbInstance("aa", "./2.db")
	db := goframeworkleveldb.GetLevelDB("aa")
	db.Write("c", []byte("ccccc"))
	read, _ := db.Read("c")
	logger.Infof("========%s", string(read))
	db.Write("c", []byte("211111111111111111"))
	read, _ = db.Read("c")
	logger.Infof("========%s", string(read))

}
