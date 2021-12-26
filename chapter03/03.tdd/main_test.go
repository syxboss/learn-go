package main

import "testing"

func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
			//t.Fatal() // 报错直接退出
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期 王强的体脂是0.32 ，但是得到的是：%f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期 王强 第一，但是得到的是：%d", randOfWQ)
			//t.Fatal() // 报错直接退出
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期 王强的体脂是0.32 ，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
			//t.Fatal() // 报错直接退出
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静的体脂是0.28 ，但是得到的是：%f", fatRateOfLJ)
		}
	}
}
func TestCase2(t *testing.T) {

}
