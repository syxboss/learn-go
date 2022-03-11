package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	count := 1000
	scope := FatRateScope{
		min: 0.0,
		max: 0.4,
	}
	chanPerson := make(chan *Person, count)
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(num int) {
			defer wg.Done()
			PersonName := fmt.Sprintf("Person-%v", num)
			FatRate := rand.Float64() * (scope.max - scope.min)
			fmt.Println(FatRate)
			p := &Person{}
			// 注册
			p.register(PersonName, FatRate)
			jsonPerson := &PersonJson{
				Name:    PersonName,
				FatRate: FatRate,
			}
			// 将注册信息放入文件中，一人一行
			registerJsonToFile(p, jsonPerson)
			chanPerson <- p
		}(i)
	}
	wg.Wait()
	close(chanPerson)
	// 更新数据并排序
	persons := updateFatRate(chanPerson, scope)

	for index, newPerson := range persons {
		fmt.Println(newPerson.name, ",当前排名第：", index)
	}
	time.Sleep(3 * time.Second)
}

func registerJsonToFile(p *Person, person *PersonJson) {
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("生成json字符串出错：", err)
	}
	if str, err := os.Getwd(); err == nil {
		path := str
		if err := p.appendPersonIntoFile(data, path+"/homework/five/person.json"); err != nil {
			log.Println("写入文件出错：", err)
		}
	} else {
		log.Fatal("获取当前目录路径出错")
	}
}

// 排序
func updateFatRate(ch chan *Person, scope FatRateScope) (personRank Persons) {
	persons := Persons{}
	for person := range ch {
		// 更新相关数据，现在只改体脂率
		person.fatRate = rand.Float64() * (scope.max - scope.min)
		persons = append(persons, person)
	}
	//冒泡排序
	bubbleSort(persons)
	//快排
	//quickSort(persons, 0, 1000-1)
	return persons
}
