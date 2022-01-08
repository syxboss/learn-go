package main

import (
	"math"
	"sort"
)

// 榜单项
type RandItem struct {
	Name    string
	FatRate float64
}

//
type FatRateRank struct {
	items []RandItem
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	// Bug？？？？  预期得到的是：0
	//for i, item := range r.items {
	//	if item.Name == name {
	//		if item.FatRate >= minFatRate {
	//			item.FatRate = minFatRate // 此处无法修改值，因为item是复制出来的对象
	//		}
	//		r.items[i] = item // 将值放回去，这样才能真正的修改值
	//	}
	//}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate // 此处无法修改值，因为item是复制出来的对象
			}
			r.items[i] = item // 将值放回去，这样才能真正的修改值
			found = true
			break //
		}
	}
	if !found {
		r.items = append(r.items, RandItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRand(name string) (rank int, fatRate float64) {

	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})

	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}

	// 排序
	sort.Float64s(rankArr)

	// 寻找排行
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			return
		}
	}
	//rankArr := make([]float64, 0, len(personFatRate))
	//// 反转 k v
	//for nameItem, frItem := range personFatRate {
	//	fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
	//	rankArr = append(rankArr, frItem)
	//}
	//// 排序 sort
	//sort.Float64s(rankArr)
	//for i, frItem := range rankArr {
	//	_names := fatRate2PersonMap[frItem]
	//	for _, _name := range _names {
	//		if _name == name {
	//			rank = i + 1
	//			fatRate = frItem
	//			return
	//		}
	//	}
	//}
	return
}
