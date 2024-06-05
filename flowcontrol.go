package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//和 for 一样，if 语句可以在条件表达式前执行一个简短语句。
//该语句声明的变量作用域仅在 if 之内。
//（在最后的 return 语句处使用 v 看看。）
//在 if 的简短语句中声明的变量同样可以在对应的任何 else 块中使用。
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {

    //for 循环
    sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

    //初始化语句和后置语句是可选的。
    sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)

    //for 是 Go 中的「while」 此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
    sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

    //如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
    //for {
    //}

    //Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的。
    fmt.Println(sqrt(2), sqrt(-4))

    fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)


    //Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只会运行选定的 case，而非之后所有的 case。 
        //在效果上，Go 的做法相当于这些语言中为每个 case 后面自动添加了所需的 break 语句。
        //在 Go 中，除非以 fallthrough 语句结束，否则分支会自动终止。 
    //Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。
    fmt.Print("Go 运行的系统环境：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

    //switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
    fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}

    //无条件的 switch 同 switch true 一样。
    //这种形式能将一长串 if-then-else 写得更加清晰。
    t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}

    //defer 推迟
    //defer 语句会将函数推迟到外层函数返回之后执行。
    //推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
    defer fmt.Println("world")
	fmt.Println("hello")

    //defer栈空间 推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。
    fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
