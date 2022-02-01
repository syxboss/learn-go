package main

func main() {
	dd := &downloadFromNetDisk{
		/*secret: func() string {
			return "1231234"
		},*/
		secret:   &mobileTokenDynamic{mobileNumber: "123456"},
		filePath: "接口异常.ppt",
	}
	dd.DownloadFile()
}

/*func dynamicSecret() string {
	return ""
}*/

type DynamicSecret interface {
	GetSecret() string
}

type mobileTokenDynamic struct {
	mobileNumber string
}

func (d *mobileTokenDynamic) GetSecret() string {
	return "something"
}
