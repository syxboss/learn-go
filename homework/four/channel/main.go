package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Person struct {
	Name    string
	FatRate float64
}

type FatRateScope struct {
	min float64
	max float64
}

type FatRateRank struct {
	Persons []Person
	total   int
	Scope   *FatRateScope
	group   sync.WaitGroup

	initNameCh   chan string
	readPersonCh chan Person
}

func main() {
	f := &FatRateRank{
		total: 1000,
		Scope: &FatRateScope{
			min: 0.0,
			max: 0.4,
		},
		initNameCh:   make(chan string, 1000),
		readPersonCh: make(chan Person, 1000),
	}

	f.group.Add(f.total)
	for i := 0; i < f.total; i++ {
		go func(n int) {
			defer f.group.Done()
			f.pushChannel(n)
		}(i)
	}
	go f.Register()          // channel获取数据,注册用户
	go f.UpdateUserFatRate() // 刷新当前排行

	for i := 0; i < f.total; i++ {
		go f.pushPersonsChannel(i)
		go f.UpdateUserInfo()      //更新信息
		go f.UpdateUserFatRate()   // 刷新排行
		go f.GetUserFatRateRank(i) // 获取当前用户排行
	}

	f.group.Wait()
	close(f.initNameCh)

	time.Sleep(5 * time.Second)
}

func (f *FatRateRank) pushChannel(i int) {
	f.initNameCh <- fmt.Sprintf("User-%v", i)
}

func (f *FatRateRank) UpdateUserFatRate() {
	sort.Slice(f.Persons, func(i, j int) bool {
		return f.Persons[i].FatRate < f.Persons[j].FatRate
	})
}

func (f *FatRateRank) Register() {
	for {
		if UserName, ok := <-f.initNameCh; ok {
			FatRate := rand.Float64() * (f.Scope.max - f.Scope.min)
			f.Persons = append(f.Persons, Person{
				Name:    UserName,
				FatRate: FatRate,
			})
		}
	}
}

func (f *FatRateRank) pushPersonsChannel(i int) {
	currentName := fmt.Sprintf("User-%v", i)
	newFatRate := rand.Float64() * (f.Scope.max - f.Scope.min)
	person := Person{
		Name:    currentName,
		FatRate: newFatRate,
	}
	f.readPersonCh <- person
}

func (f *FatRateRank) GetUserFatRateRank(i int) {
	currentName := fmt.Sprintf("User-%v", i)
	for index, item := range f.Persons {
		if currentName == item.Name {
			fmt.Println("当前用户名称：", currentName, ", 体脂率:", item.FatRate, ", 排名:", index)
		}
	}
}

func (f *FatRateRank) UpdateUserInfo() {
	for person := range f.readPersonCh {
		for index, item := range f.Persons {
			if person.Name == item.Name {
				if item.FatRate < person.FatRate && item.FatRate+0.2 >= person.FatRate {
					item.FatRate = person.FatRate
					f.Persons[index] = item
				} else if item.FatRate > person.FatRate && item.FatRate-0.2 <= person.FatRate {
					item.FatRate = person.FatRate
					f.Persons[index] = item
				} else {
					fmt.Println("当前修改用户信息 { ", "用户名称：", item.Name, " ，用户原体脂率：", item.FatRate, ",用户修改体脂率：", person.FatRate, "} 不符合修改范围,实际修改范围应在[-0.2 ~ 0.2]之间")
					break
				}
			}
		}
	}
}
