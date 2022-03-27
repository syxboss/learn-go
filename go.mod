module learn.go

go 1.17

replace (
	github.com/spf13/cobra => github.com/syxboss/cobra v1.3.0 //放在github上
	github.com/syxboss/awesomeProject => ./staging/src/github.com/syxboss/awesomeProject // 放在自己本地
	learn.go.tools => ../learn.go.tools
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/spf13/cobra v1.3.0
	//引入自己写的工具
	github.com/syxboss/awesomeProject v0.0.0-20211222055947-ae5cb5fa905d
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
	learn.go.tools v0.0.0-00010101000000-000000000000
)

require (
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//require learn.go.tools   // 识别不了
