package calcator

func CalcBMI(tall float64, weight float64) (bmi float64) {
	if tall <= 0 {
		panic("身高不能为0或者负数")
	}
	bmi = weight / (tall * tall)
	return
}
