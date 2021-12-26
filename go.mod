module learn.go

go 1.17

replace (
	github.com/spf13/cobra => github.com/syxboss/cobra v1.3.0 //放在github上
	github.com/syxboss/awesomeProject => ./staging/src/github.com/syxboss/awesomeProject // 放在自己本地
	learn.go.tools => ../learn.go.tools
)

require (
	github.com/spf13/cobra v1.3.0
	//引入自己写的工具
	github.com/syxboss/awesomeProject v0.0.0-20211222055947-ae5cb5fa905d
	learn.go.tools v0.0.0-00010101000000-000000000000
)

//require learn.go.tools   // 识别不了
