package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// 人
type Person struct {
	Name    string  // 名字
	FatRate float64 // 体脂率
}

// 体脂的范围
type FatRateScope struct {
	min float64 // 最小值
	max float64 // 最大值
}

// 体脂排行榜
type FatRateRank struct {
	Persons []Person // 所有的人
	total   int      // 总人数
	//RankList []float64   // 体脂排序
	Scope *FatRateScope // 范围
	lock  sync.RWMutex  // 锁
}

func main() {
	scope := FatRateScope{
		min: 0.0,
		max: 0.4,
	}
	fateRateRank := &FatRateRank{
		total: 1000,
		Scope: &scope,
	}

	// 注册
	for i := 0; i < fateRateRank.total; i++ {
		go func(i int) {
			fateRateRank.Register(i)
		}(i)
		//for _, item := range fateRateRank.Persons {
		//	fmt.Println(item)
		//}
	}
	// 注册之后刷新当前排行
	fateRateRank.UpdateUserFatRate()
	//for _, item := range fateRateRank.Persons {
	//	fmt.Println(item)
	//}

	// 更新用户体脂信息刷新排行再从中获取用户排行
	for i := 0; i < fateRateRank.total; i++ {
		go func(i int) {
			fateRateRank.GetUserFatRateAndRank(i)
			fateRateRank.UpdateUserInfo(i)
			fateRateRank.GetUserFatRateAndRank(i)
		}(i)
	}
	time.Sleep(5 * time.Second)
}

// 用户注册
func (f *FatRateRank) Register(i int) {
	f.lock.Lock()
	defer f.lock.Unlock()
	PersonName := fmt.Sprintf("User-%v", i)
	FatRate := rand.Float64() * (f.Scope.max - f.Scope.min)
	f.Persons = append(f.Persons, Person{
		Name:    PersonName,
		FatRate: FatRate,
	})
}

//获得排行   每次都要遍历一遍，是否有优化方法？
func (f *FatRateRank) GetUserFatRateAndRank(i int) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	currentName := fmt.Sprintf("User-%v", i)
	for index, item := range f.Persons {
		if currentName == item.Name {
			fmt.Println("当前用户名称：", currentName, ", 体脂率:", item.FatRate, ", 排名:", index)
		}
	}
}

//更新用户排行
func (f *FatRateRank) UpdateUserFatRate() {
	f.lock.Lock()
	defer f.lock.Unlock()
	sort.Slice(f.Persons, func(i, j int) bool {
		return f.Persons[i].FatRate < f.Persons[j].FatRate
	})
}

// 更新用户体脂
func (f *FatRateRank) UpdateUserInfo(i int) {
	f.lock.Lock()
	defer f.lock.Unlock()
	currentName := fmt.Sprintf("User-%v", i)
	newFatRate := rand.Float64() * (f.Scope.max - f.Scope.min) // 随机获取最新体脂值
	for index, item := range f.Persons {
		if currentName == item.Name {
			if item.FatRate < newFatRate && item.FatRate+0.2 >= newFatRate {
				item.FatRate = newFatRate
				f.Persons[index] = item
			} else if item.FatRate > newFatRate && item.FatRate-0.2 <= newFatRate {
				item.FatRate = newFatRate
				f.Persons[index] = item
			} else {
				fmt.Println("当前修改用户信息 { ", "用户名称：", item.Name, " ，用户原体脂率：", item.FatRate, ",用户修改体脂率：", newFatRate, "} 不符合修改范围,实际修改范围应在[-0.2 ~ 0.2]之间")
				break
			}
			rank, FatRate := f.getRank(currentName)
			fmt.Println("更新用户信息如下：{", " 用户名称：", currentName, ", 体脂率:", FatRate, "， 用户排名:", rank, " }")
		}
	}
}

//获取排行
func (f *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	frs := map[float64]struct{}{}
	for _, item := range f.Persons {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return rank, fatRate
}
