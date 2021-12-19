package main

import (
	"fmt"
	"github.com/spf13/cobra"
	//_ "github.com/spf13/cobra"
	learn_go_tool "learn.go.tools"
	"learn.go/chapter02/05.fatrate.refactorPackage/calc"
)

func main() {
	// 录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "fateRate",           // 名字
		Short: "体脂计算器，根据相关信息给出健康建议", //简称
		Long:  "该体脂计算器是基于BMI指数进行的。。。。。",
		Run: func(cmd *cobra.Command, args []string) { // 类似回调函数，执行的函数
			fmt.Println("name:", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			bmi := calc.CalcBMI(tall, weight)
			fateRate := calc.CalcFatRate(bmi, age, sex)
			fmt.Println(fateRate)
		},
	}
	max := learn_go_tool.Max(10, 11)
	fmt.Println(max)
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.ExecuteC()
	//panic: runtime error: index out of range [1] with length 1
	//arguments := os.Args // Args []string
	//fmt.Println(arguments)
	//name = arguments[1]
	//sex = arguments[2]
	//tall = arguments[3]
	//weight = arguments[4]
	//age = arguments[5]

	//fmt.Println("name:", name)
	//fmt.Println("sex:", sex)
	//fmt.Println("tall:", tall)
	//fmt.Println("weight:", weight)
	//fmt.Println("age:", age)
	// 计算

	// 评估结果
}
