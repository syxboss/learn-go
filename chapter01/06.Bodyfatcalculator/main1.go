package main

import (
	"fmt"
)

/**
Body fat calculator
BMI=体重（公斤）÷（身高×身高）（米）
体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）

要求：
1.计算出一个人的体脂
2.告诉他是偏瘦、标准、偏重、肥胖、严重肥胖
*/

// 版本2 能互动的体脂计算器
// 正则表达式
func main() {
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
	var bmi = weight / (tall * tall)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}

	var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100 // 体脂率
	fmt.Println("体脂率是：", fatRate)

	switch sex {
	case "男":
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
	case "女":
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
}
