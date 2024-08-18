package dao

type InterfaceDao interface {
	IAdminUsers
}

var IDao InterfaceDao

type Dao struct {
}

func New() {
	IDao = &Dao{}
}
