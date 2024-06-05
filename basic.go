package main
//每个 Go 程序都由包构成。 程序从 main 包开始运行。


import (
    "fmt"
    "math/rand"
//    "math"
    "math/cmplx"
)

//此代码用圆括号将导入的包分成一组，这是“分组”形式的导入语句。 当然你也可以编写多个导入语句
//import "fmt"
//import "math"

//在导入一个包时，你只能引用其中已导出的名字。 任何「未导出」的名字在该包外均无法访问。eg: math.pi

//function 注意类型在变量名的后面
func add(x int, y int) int {
    return x + y;
}

//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
func add2(x, y int) int {
    return x + y;
}

//函数可以返回任意数量的返回值。
func swap(x, y string) (string, string) {
    return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
//返回值的命名应当能反应其含义，它可以作为文档使用。
//没有参数的 return 语句会直接返回已命名的返回值，也就是「裸」返回值。
func split(sum int) (x, y int) {
    x = sum * 4 /9
    y = sum - x
    return
}

//var 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。
var c, python, java bool
//变量声明可以包含初始值，每个变量对应一个。如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。
var i2, j2 int = 1, 2

//Go 的基本类型
//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte // uint8 的别名
//rune // int32 的别名
//     // 表示一个 Unicode 码位
//float32 float64
//complex64 complex128

//和导入语句一样，变量声明也可以「分组」成一个代码块。
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

//数值常量是高精度的 值。
//一个未指定类型的常量由上下文来决定其类型。
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func main() {
    fmt.Println("我最爱的数字是", rand.Intn(10))

//    fmt.Println(math.pi)

    fmt.Println(add(42, 13))

    fmt.Println(add2(42, 13))

    a, b := swap("hello", "world")

    fmt.Println(a, b)

    fmt.Println(split(17))

    var i int
    fmt.Println(i, c, python, java)

    var c2, python2, java2 = true, false, "no!"
    fmt.Println(i2, j2, c2, python2, java2)

    //短变量声明 在函数中，短赋值语句 := 可在隐式确定类型的 var 声明中使用。
    k := 3
    fmt.Println(k)

	fmt.Printf("类型：%T 值：%v\n", ToBe, ToBe)
	fmt.Printf("类型：%T 值：%v\n", MaxInt, MaxInt)
	fmt.Printf("类型：%T 值：%v\n", z, z)

    //没有明确初始化的变量声明会被赋予对应类型的 零值。
  	var i3 int
	var f3 float64
	var b3 bool
	var s3 string
	fmt.Printf("%v %v %v %q\n", i3, f3, b3, s3)

    //类型转换 表达式 T(v) 将值 v 转换为类型 T。
    //与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
    var i4 int = 42
    var f4 float64 = float64(i4)
    var u4 uint = uint(f4)
    fmt.Println(i4, f4, u4)
    //更加简短的形式：
    //i4 := 42
    //f4 := float64(i4)
    //u4 := uint(f4)

    //在声明一个变量而不指定其类型时（即使用不带类型的 := 语法 var = 表达式语法），变量的类型会通过右值推断出来。
    //当声明的右值确定了类型时，新变量的类型与其相同：
    //不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int、float64 或 complex128 了，这取决于常量的精度：
    v5 := 42
    fmt.Printf("v is of type %T\n",v5)

    //常量的声明与变量类似，只不过使用 const 关键字。
    //常量可以是字符、字符串、布尔值或数值。
    //常量不能用 := 语法声明。
    const Truth = true

    const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
