package calc

import "testing"

/**
测试计算BMI
*/
func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0 // weight为0.0 / height为0.0
	expectedOutput := 1.0
	t.Logf("开始计算，输入：height：%f，weight：%f", inputHeight, inputWeight)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到：%f，error：%f", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err , but got %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expected %f ,but got %f", expectedOutput, actualOutput)
	}
}
