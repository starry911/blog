package dao

type InterfaceDao interface {
	ITestDao
}

var IDao InterfaceDao

type Dao struct {
}

func New() {
	IDao = &Dao{}
}
