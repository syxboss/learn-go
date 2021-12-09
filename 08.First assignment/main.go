package main

import "fmt"

/**
计算多个人的平均体脂
	实现完整的体脂计算器
	连续输入三人的姓名、性别、身高、体重、年龄信息
	计算每个人的BMI、体脂率
输出：
	每个人的姓名、BMI、体脂率、建议
	总人数、平均体脂率
*/

func main() {
	for {
		var count int
		fmt.Print("请输入需要计算的人数：")
		fmt.Scan(&count)
		// 初始化动态数组
		var bmis = make([]float64, count)
		var names = make([]string, count)
		var sexs = make([]string, count)
		var ages = make([]int, count)
		var fatRates = make([]float64, count)
		var sum float64
		// 赋值
		for i := 0; i < count; i++ {
			name, sex, age, fatRate, bmi := inputInfo(i)
			names[i] = name
			bmis[i] = bmi
			sexs[i] = sex
			ages[i] = age
			fatRates[i] = fatRate
			sum += fatRate
		}
		//总人数
		fmt.Println("结果如下：")
		fmt.Println("总人数为：", count)
		//评价
		for i := 0; i < count; i++ {
			Evaluation(names[i], sexs[i], ages[i], fatRates[i], bmis[i])
		}

		//平均体脂率
		fmt.Println("平均体脂率为：", sum/float64(count))

		var whetherContinue string
		fmt.Println("是否继续【y/n】？")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			return
		}
	}
}

/*
	输入操作
	return 需要的信息
*/
func inputInfo(i int) (string, string, int, float64, float64) {
	var name string
	fmt.Print("请输入第", i+1, "位的名字:")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("请输入第", i+1, "位体重【kg】:")
	fmt.Scanln(&weight) //取地址

	var tall float64
	fmt.Print("请输入第", i+1, "位身高【m】:")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("请输入第", i+1, "位年龄【岁】:")
	fmt.Scanln(&age)

	var sex string
	fmt.Print("请输入第", i+1, "位的性别【男、女】:")
	fmt.Scanln(&sex)

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}

	var bmi = weight / (tall * tall)

	//计算体脂率
	var fatRate = CalcFatRate(bmi, age, sexWeight)

	return name, sex, age, fatRate, bmi
}

/*
	计算体脂率
	return fatRate
*/
func CalcFatRate(bmi float64, age int, sexWeight int) float64 {
	return (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
}

/*
	评价
*/
func Evaluation(name string, sex string, age int, fatRate float64, bmi float64) {
	fmt.Println("您好！下面是您的体脂计算结果及建议：")
	fmt.Println("尊敬的用户：", name, ",bmi指数为：", bmi, "体脂率是：", fatRate)
	switch sex {
	case "男":
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
	case "女":
		// 编写女性的体脂率与体脂状态表
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
}
