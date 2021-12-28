package calc

import (
	"fmt"
	"testing"
)

/**
测试计算体脂率的方法
*/
func TestCalcFatRate(t *testing.T) {

	inputBMI, inputAge, InputSex := 12.0, 32, "男"

	t.Logf("开始计算体脂率，输入 BMI：%f ，年龄：%d ， 性别：%s", inputBMI, inputAge, InputSex)
	fatRate, err := CalcFatRate(inputBMI, inputAge, InputSex)
	if err != nil {
		t.Fatalf("expecting no err , but got %v", err)
	}
	fmt.Println(fatRate)
}

func TestCalcFatRateBMI(t *testing.T) {

	inputBMI, inputAge, InputSex := -0.1, 32, "男"

	t.Logf("开始计算体脂率，输入 BMI：%f ，年龄：%d ， 性别：%s", inputBMI, inputAge, InputSex)
	fatRate, err := CalcFatRate(inputBMI, inputAge, InputSex)
	if err != nil {
		t.Fatalf("expecting no err , but got %v", err)
	}
	fmt.Println(fatRate)
}

func TestCalcFatRateAGE(t *testing.T) {

	inputBMI, inputAge, InputSex := 12.0, -1, "男" // 151 ,0 ,-1

	t.Logf("开始计算体脂率，输入 BMI：%f ，年龄：%d ， 性别：%s", inputBMI, inputAge, InputSex)
	fatRate, err := CalcFatRate(inputBMI, inputAge, InputSex)
	if err != nil {
		t.Fatalf("expecting no err , but got %v", err)
	}
	fmt.Println(fatRate)
}

func TestCalcFatRateSex(t *testing.T) {

	inputBMI, inputAge, InputSex := 12.0, 32, "123"

	t.Logf("开始计算体脂率，输入 BMI：%f ，年龄：%d ， 性别：%s", inputBMI, inputAge, InputSex)
	fatRate, err := CalcFatRate(inputBMI, inputAge, InputSex)
	if err != nil {
		t.Fatalf("expecting no err , but got %v", err)
	}
	fmt.Println(fatRate)
}
