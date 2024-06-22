package goframework_leveldb

import (
	"github.com/kordar/godb"
	"github.com/kordar/goleveldb"
	"github.com/kordar/gologger"
)

var (
	leveldbpool = godb.NewDbPool()
)

func GetLevelDB(db string) *goleveldb.LevelDB {
	return leveldbpool.Handle(db).(*goleveldb.LevelDB)
}

// InitLevelDbHandle 初始化LevelDb句柄
func InitLevelDbHandle(dbs map[string]string) {
	for db, filepath := range dbs {
		ins := NewLevelDbIns(db, filepath)
		if ins == nil {
			continue
		}
		err := leveldbpool.Add(ins)
		if err != nil {
			logger.Warnf("初始化异常，err=%v", err)
		}
	}

}

// AddLevelDbInstance 添加LevelDb句柄
func AddLevelDbInstance(db string, filepath string) error {
	ins := NewLevelDbIns(db, filepath)
	return leveldbpool.Add(ins)
}

// RemoveLevelDbInstance 移除LevelDb句柄
func RemoveLevelDbInstance(db string) {
	leveldbpool.Remove(db)
}

// HasLevelDbInstance LevelDb句柄是否存在
func HasLevelDbInstance(db string) bool {
	return leveldbpool != nil && leveldbpool.Has(db)
}
