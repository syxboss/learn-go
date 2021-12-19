module learn.go

go 1.15

replace learn.go.tools => ../learn.go.tools

require (
	github.com/spf13/cobra v1.3.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

//require learn.go.tools   // 识别不了
