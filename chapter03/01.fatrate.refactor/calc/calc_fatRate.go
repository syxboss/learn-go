package calc

import gobmi "github.com/syxboss/awesomeProject" // replace掉之后是原来的路径，别忘了修改！

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	//gobmi.CalcFatRate(bmi,age,sex)  //gobmi "learn.go/staging/src/github.com/syxboss/awesomeProject"
	fatRate = gobmi.CalcFatRate(bmi, age, sex)

	return
}
