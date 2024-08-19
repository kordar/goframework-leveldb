package goframework_leveldb

import (
	"github.com/kordar/goleveldb"
)

type LevelDbIns struct {
	name string
	ins  *goleveldb.LevelDB
}

func NewLevelDbIns(name string, filepath string) *LevelDbIns {
	levelDB := goleveldb.NewLevelDB(filepath)
	return &LevelDbIns{name: name, ins: levelDB}
}

func (c LevelDbIns) GetName() string {
	return c.name
}

func (c LevelDbIns) GetInstance() interface{} {
	return c.ins
}

func (c LevelDbIns) Close() error {
	c.ins.Close()
	return nil
}
