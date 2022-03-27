package main

type ServerInterface interface {
	//保存数据
	SaveData(cs *CircleState) error

	//更新数据
	UpdateData(cs *CircleState) error

	//查询所有数据
	GetDataList() ([]*CircleState, error)

	//查询一条数据
	GetOneData(userId int64) (*CircleState, error)

	//删除数据
	DeleteDataByDeleted(Id int64) error
}
