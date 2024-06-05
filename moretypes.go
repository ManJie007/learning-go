package main

import (
    "fmt"
    "strings"
    "math"
)

//一个 结构体（struct）就是一组 字段（field）。
type Vertex struct {
	X int
	Y int
}

//结构体字面量
//使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）。
//特殊的前缀 & 返回一个指向结构体的指针。
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予零值
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

var m2 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

//若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它。
var m3 = map[string]Vertex2{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	//Go 拥有指针。指针保存了值的内存地址。
	//类型 *T 是指向 T 类型值的指针，其零值为 nil。
	//& 操作符会生成一个指向其操作数的指针。
	//* 操作符表示指针指向的底层值。
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

    fmt.Println(Vertex{1, 2})

    //结构体字段可通过点号 . 来访问。
    v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

    //结构体指针
    //结构体字段可通过结构体指针来访问。
    //如果我们有一个指向结构体的指针 p 那么可以通过 (*p).X 来访问其字段 X。 不过这么写太啰嗦了，所以语言也允许我们使用隐式解引用，直接写 p.X 就可以。
    p2 := &v
	p2.X = 1e9
	fmt.Println(v)

    fmt.Println(v1, p, v2, v3)

    //类型 [n]T 表示一个数组，它拥有 n 个类型为 T 的值。
    //数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是个限制，不过没关系，Go 拥有更加方便的使用数组的方式。
    var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


    //切片
    //每个数组的大小都是固定的。而切片则为数组元素提供了动态大小的、灵活的视角。 在实践中，切片比数组更常用。
    //类型 []T 表示一个元素类型为 T 的切片。.
    //切片通过两个下标来界定，一个下界和一个上界，二者以冒号分隔：
    //a[low : high]
    //它会选出一个半闭半开区间，包括第一个元素，但排除最后一个元素。
    primes2 := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes2[1:4]
	fmt.Println(s)

    // 切片类似数组的引用
    //切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。
    //更改切片的元素会修改其底层数组中对应的元素。
    //和它共享底层数组的切片都会观测到这些修改。

    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a2 := names[0:2]
	b2 := names[1:3]
	fmt.Println(a2, b2)

	b2[0] = "XXX"
	fmt.Println(a2, b2)
	fmt.Println(names)

    //切片字面量
    //切片字面量类似于没有长度的数组字面量。
    q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

    //切片的默认行为
    //在进行切片时，你可以利用它的默认行为来忽略上下界。
    //切片下界的默认值为 0，上界则是该切片的长度。
    //对于数组
    //var a [10]int
    //来说，以下切片表达式和它是等价的：
    //a[0:10]
    //a[:10]
    //a[0:]
    //a[:]

    s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

    //切片的长度与容量
    //切片拥有 长度 和 容量。
    //切片的长度就是它所包含的元素个数。
    //切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
    //切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
    //你可以通过重新切片来扩展一个切片，给它提供足够的容量。 
    s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// 截取切片使其长度为 0
	s4 = s4[:0]
	printSlice(s4)

	// 扩展其长度
	s4 = s4[:4]
	printSlice(s4)

	// 舍弃前两个值
	s4 = s4[2:]
	printSlice(s4)


    //nil 切片
    //切片的零值是 nil。
    //nil 切片的长度和容量为 0 且没有底层数组。
    var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nil!")
	}

    //用 make 创建切片
    //切片可以用内置函数 make 来创建，这也是你创建动态数组的方式。
    //make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
    a3 := make([]int, 5)
	printSlice2("a3", a3)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

    //切片的切片
    //切片可以包含任何类型，当然也包括其他切片。
    // 创建一个井字棋（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
    
    //向切片追加元素
    //为切片追加新的元素是种常见的操作，为此 Go 提供了内置的 append 函数。
    //func append(s []T, vs ...T) []T
    //append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。
    //append 的结果是一个包含原切片所有元素加上新添加元素的切片。
    //当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。 返回的切片会指向这个新分配的数组。
    var s6 []int
	printSlice(s6)

	// 可在空切片上追加
	s6 = append(s6, 0)
	printSlice(s6)

	// 这个切片会按需增长
	s6 = append(s6, 1)
	printSlice(s6)

	// 可以一次性添加多个元素
	s6 = append(s6, 2, 3, 4)
	printSlice(s6)

    // range 遍历
    //for 循环的 range 形式可遍历切片或映射。
    //当使用 for 循环遍历切片时，每次迭代都会返回两个值。 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
    for i, v := range pow2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

    //可以将下标或值赋予 _ 来忽略它。
    //for i, _ := range pow
    //for _, value := range pow
    //若你只需要索引，忽略第二个变量即可。
    //for i := range pow
    pow3 := make([]int, 10)
	for i := range pow3 {
		pow3[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow3 {
		fmt.Printf("%d\n", value)
	}

    //map 映射
    //map 映射将键映射到值。
    //映射的零值为 nil 。nil 映射既没有键，也不能添加键。
    //make 函数会返回给定类型的映射，并将其初始化备用。
    m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m2)

	fmt.Println(m3)

    /**
    修改映射
    在映射 m 中插入或修改元素：
    m[key] = elem
    获取元素：
    elem = m[key]
    删除元素：
    delete(m, key)
    通过双赋值检测某个键是否存在：
    elem, ok = m[key]
    若 key 在 m 中，ok 为 true ；否则，ok 为 false。
    若 key 不在映射中，则 elem 是该映射元素类型的零值。
    注：若 elem 或 ok 还未声明，你可以使用短变量声明：
    elem, ok := m[key]
    */
    m4 := make(map[string]int)

	m4["答案"] = 42
	fmt.Println("值：", m4["答案"])

	m4["答案"] = 48
	fmt.Println("值：", m4["答案"])

	delete(m4, "答案")
	fmt.Println("值：", m4["答案"])

	v2, ok := m4["答案"]
	fmt.Println("值：", v2, "是否存在？", ok)

    //函数值
    //函数也是值。它们可以像其他值一样传递。
    //函数值可以用作函数的参数或返回值。
    hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

    //函数闭包
    //Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。 该函数可以访问并赋予其引用的变量值，换句话说，该函数被“绑定”到了这些变量。
    //例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
    pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}


}
