package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type CreateDB struct {
	conn *gorm.DB
}

func DBConnect(conn *gorm.DB) ServerInterface {
	if conn == nil {
		log.Fatal("数据库连接为空")
	}
	return &CreateDB{
		conn: conn,
	}
}

// 添加数据
func (db CreateDB) SaveData(cs *CircleState) error {
	res := db.conn.Create(cs)
	if err := res.Error; err != nil {
		fmt.Printf("添加数据失败：%v\n", err)
		return err
	}
	fmt.Printf("添加数据成功！\n")
	return nil
}

// 更新数据
func (db CreateDB) UpdateData(cs *CircleState) error {
	res := db.conn.Updates(cs)
	if err := res.Error; err != nil {
		fmt.Printf("更新数据失败：%v\n", err)
		return err
	}
	fmt.Printf("更新数据成功！\n")
	return nil
}

// 查询所有数据
func (db CreateDB) GetDataList() ([]*CircleState, error) {
	circleState := make([]*CircleState, 0, 10)
	res := db.conn.Where(&CircleState{Deleted: false}).Find(&circleState)
	if res.Error != nil {
		fmt.Println("查询列表失败：", res.Error)
		return nil, res.Error
	}

	data, err := json.Marshal(circleState)
	if err != nil {
		fmt.Println("编码出错：", err)
	}
	fmt.Printf("查询列表数据：%v\n", string(data))

	return circleState, nil
}

// 获取单条数据
func (db CreateDB) GetOneData(userId int64) (*CircleState, error) {
	cs := &CircleState{}
	res := db.conn.Where(&CircleState{UserId: userId}).Find(&cs)
	if res.Error != nil {
		fmt.Println("查询失败：", res.Error)
		return nil, res.Error
	}

	data, err := json.Marshal(cs)
	if err != nil {
		fmt.Println("编码出错：", err)
	}
	fmt.Printf("查询的数据：%v\n", string(data))
	return cs, nil
}

// 删除数据
func (db CreateDB) DeleteDataByDeleted(Id int64) error {
	upD := &CircleState{
		Id:      Id,
		Deleted: false,
	}
	res := db.conn.Model(upD).Select("deleted").Updates(upD)
	if res.Error != nil {
		fmt.Println("删除失败：", res.Error)
		return res.Error
	}
	fmt.Printf("删除成功！\n")
	return nil
}
