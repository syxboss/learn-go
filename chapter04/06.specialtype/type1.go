package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore int = 100
	var mathScore2 Math = 200

	mathScore2 = mathScore
	fmt.Println(mathScore2)

	vg := &voteGate{students: []*student{
		&student{name: fmt.Sprintf("%d", 1)},
		&student{name: fmt.Sprintf("%d", 2)},
		&student{name: fmt.Sprintf("%d", 3)},
		&student{name: fmt.Sprintf("%d", 4)},
		&student{name: fmt.Sprintf("%d", 5)},
	}}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdXQ = &student{name: "小强"}
	var ldXQ Leader = Leader(*stdXQ) // 	var ldXQ leader = stdXQ   //
	ldXQ.Distribute()

	// 关联关系可以转换
	/*bytesTest1 := []byte{}
	var str1 = string(bytesTest1)
	fmt.Println(str1)*/
}

func getScoresOfStudent(name string) (Math, English, Chinese) {
	return 0, 0, 0
}

//=====================================================
type voteGate struct {
	students []*student
}

//func (g *voteGate) goRun()  (Leader *student) {}
func (g *voteGate) goRun() *Leader {
	for _, item := range g.students {
		randIdx := rand.Int()
		if randIdx%2 == 0 {
			item.voteA(g.students[randIdx%len(g.students)])
		} else {
			item.voteD(g.students[randIdx%len(g.students)])
		}
		//item.voteA(g.students[0])   // TODO 使用随机数代替
	}

	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 { // 判断是否大于0 ，因为如果没有学生，那么index就是默认值 -1
		return (*Leader)(g.students[maxScoreIndex])
	}
	return nil
}

// 使用嵌套兑现定义（继承）方式来定义班长
//type Leader struct {
//	student
//}

// 使用类型重定义
type Leader student

func (l *Leader) Distribute() {
	fmt.Println("发作业了")
}

// 指针无法定义,其他类型可以
type FooooTestRedefine []string //*int //func()

//Invalid receiver type 'FooooTestRedefine' ('FooooTestRedefine' is a pointer type)
func (f FooooTestRedefine) test11() {

}

type student struct {
	name     string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	std.agree++
}

func (std *student) voteD(target *student) {
	std.disagree++
}
