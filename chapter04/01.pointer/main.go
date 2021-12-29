package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a) //3

	c := &a
	d := &c
	fmt.Println("&a:", &a)                         //0xc0000120a8
	fmt.Println("c:", c)                           //0xc0000120a8
	fmt.Println("&c:", &c)                         //0xc000006030
	fmt.Println("d:", d, ",*d:", *d, ",**d:", **d) //d: 0xc000006030 ,*d: 0xc0000120a8 ,**d: 3

	m := map[string]string{}
	mp1 := &m        // mp1的类型就是*map[string]string{}
	fmt.Println(mp1) //&map[]
	put(m)
	fmt.Println("*mp1=", *mp1) //*mp1= map[a:aaa]

	f1 := add
	f1(&a, &b)
	fmt.Println("fl ,add=", a) //5
	flp1 := &f1                // flp1 = *func(int,int)
	(*flp1)(&a, &b)
	fmt.Println("fl ,add=", a) //7

	{
		var nothing *int
		//*nothing = 3          // 这里是没有指向任何东西的
		fmt.Println("nothing", nothing) //invalid memory address or nil pointer dereference
	}
	{
		var nothingMap map[string]string
		//nothingMap["a"] = "BBB"
		fmt.Println("nothingMap", nothingMap) //panic: assignment to entry in nil map
	}
	{
		var nothingSlice []int
		//nothingSlice[0] =100
		nothingSlice = append(nothingSlice, 100)  //nothingSlice [100]
		fmt.Println("nothingSlice", nothingSlice) //panic: runtime error: index out of range [0] with length 0
	}
}

func put(m map[string]string) {
	m["a"] = "aaa"
}

func add(a *int, b *int) {
	*a = *a + *b
}
