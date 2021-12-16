package main

import (
	"fmt"
	"runtime/debug"
)

// 特殊函数init
// 自动调用，不能主动调用
// 可以定义多个，顺序执行，main函数之前
func init() {
	fmt.Println("我是init函数")
}

func init() {
	fmt.Println("我是init1函数")
}

func main() {
	for {
		mainFatRateBody()

		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning：catch error %v\n", re)
		debug.PrintStack() //打印调用栈
	}
}
func mainFatRateBody() {
	weight, tall, age, sex, sexWeight := getMaterialsFromImput()

	// fixme
	// 计算体脂
	fatRate := calcFatRate(weight, tall, age, sexWeight)
	if fatRate == 0 {
		panic("fat rate is not allowed to be active!")
	}

	// 得到健康提示
	getHealthinessSuggestions(sex, age, fatRate, getHealthinessSuggestionMale)
	getHealthinessSuggestions(sex, age, fatRate, getHealthinessSuggestionFamle)

	// 特殊变量
	check := getHealthinessSuggestions
	fmt.Println(check)
	var checkHealthinessFunc func(age int, fatRate float64) // 必须是函数，而且参数列表类型返回值必须一致 ,变量类型不可变
	fmt.Println(checkHealthinessFunc)
}

func getHealthinessSuggestions(sex string, age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate) // 工具
	switch sex {
	case "男":
		getHealthinessSuggestionMale(age, fatRate)
	case "女":
		getHealthinessSuggestionFamle(age, fatRate)
	}
}

// 函数作为产出
func generateCheckHealthinessForMale() func(age int, fatRate float64) {
	// 定制
	return func(age int, fatRate float64) {
		//返回定制函数内容
	}
}

func getHealthinessSuggestionFamle(age int, fatRate float64) {
	//TODO 编写女性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.2 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.2 && fatRate <= 0.27 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.27 && fatRate <= 0.34 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.34 && fatRate <= 0.39 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.21 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.21 && fatRate <= 0.28 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.28 && fatRate <= 0.35 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.35 && fatRate <= 0.40 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 60 {
		if fatRate <= 0.22 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.22 && fatRate <= 0.29 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.29 && fatRate <= 0.36 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.36 && fatRate <= 0.41 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else {
		fmt.Println("未成年人体脂改变迅速，不计算BMI体脂")
	}
}

func getHealthinessSuggestionMale(age int, fatRate float64) {
	//TODO 编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.11 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.11 && fatRate <= 0.17 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.17 && fatRate <= 0.22 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.22 && fatRate <= 0.27 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 60 {
		if fatRate <= 0.13 {
			fmt.Println("目前是:偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.13 && fatRate <= 0.19 {
			fmt.Println("目前是:标准。太棒了，要保持哦")
		} else if fatRate > 0.19 && fatRate <= 0.24 {
			fmt.Println("目前是:偏旁。吃晚饭多散散步，消化消化")
		} else if fatRate > 0.24 && fatRate <= 0.29 {
			fmt.Println("目前是:肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是:非常肥胖。健身游泳跑步，即刻开始")
		}
	} else {
		fmt.Println("未成年人体脂改变迅速，不计算BMI体脂")
	}
}

func calcFatRate(weight float64, tall float64, age int, sexWeight int) float64 {
	var bmi = weight / (tall * tall)
	var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100 // 体脂率
	fmt.Println("体脂率是：", fatRate)
	return fatRate
}

// refactor -> rename
func getMaterialsFromImput() (float64, float64, int, string, int) {
	// 录入各项
	var name string
	fmt.Print("请输入你的名字:")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("请输入体重【kg】:")
	fmt.Scanln(&weight) //取地址

	var tall float64
	fmt.Print("请输入身高【m】:")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("请输入年龄【岁】:")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("请输入sex【男、女】:")
	fmt.Scanln(&sex)

	var sexWeight int

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sex, sexWeight
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("是否录入下一个【y/n】？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
