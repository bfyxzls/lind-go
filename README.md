# lind-go
对go语言的学习

# go.mod
在Go语言中，`go.mod`文件用于管理项目的依赖关系和版本信息。一般来说，`go.mod`文件是由`go mod init`命令自动生成的，它会根据你的项目结构和依赖情况自动创建和更新。

一旦`go.mod`文件生成后，大部分情况下不需要手工维护。当你使用`go get`、`go build`或`go run`等命令时，Go工具会自动检查`go.mod`文件并下载所需的依赖项，同时更新`go.mod`和`go.sum`文件。

然而，在某些情况下，可能需要手工维护`go.mod`文件，比如手动添加或删除依赖项、指定特定的依赖版本等。总的来说，大部分情况下`go.mod`文件可以由Go工具自动维护，但在特殊情况下可能需要手动进行调整。

# 一 golang基础知识
Go（又称 Golang）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种计算机编程语言语言。
**设计初衷**
Go语言是谷歌推出的一种的编程语言，可以在不损失应用程序性能的情况下降低代码的复杂性。谷歌首席软件工程师罗布派克(Rob Pike)说：我们之所以开发Go，是因为过去10多年间软件开发的难度令人沮丧。派克表示，和今天的C++或C一样，Go是一种系统语言。他解释道，"使用它可以进行快速开发，同时它还是一个真正的编译语言，我们之所以现在将其开源，原因是我们认为它已经非常有用和强大。"
* 计算机硬件技术更新频繁，性能提高很快。目前主流的编程语言发展明显落后于硬件，不能合理利用多核多CPU的优势提升软件系统性能。
* 软件系统复杂度越来越高，维护成本越来越高，目前缺乏一个足够简洁高效的编程语言。
* 企业运行维护很多c/c++的项目，c/c++程序运行速度虽然很快，但是编译速度确很慢，同时还存在内存泄漏的一系列的困扰需要解决。

**应用领域**

![image-20230425134447756](D:\github\lind-go\assets\Sample 1.    HelloWorld.twinproj)

**数据类型**
* int :有符号的整数类型:具体占几个字节要看操作系统的分配:不过至少分配给32位。
* uint:非负整数类型:具体占几个字节要看操作系统的分配:不过至少分配给32位。
* int8:有符号的整数类型:占8位bit:1个字节。范围从负的2的8次方到正的2的8次方减1。
* int16:有符号的整数类型:占16位bit:2个字节。范围从负的2的16次方到正的2的16次方减1。
* int32:有符号的整数类型:占32位bit:4个字节。范围从负的2的32次方到正的2的32次方减1。
* int64:有符号的整数类型:占64位bit:8个字节。范围从负的2的64次方到正的2的64次方减1。
* uint8:无符号的正整数类型:占8位:从0到2的9次方减1.也就是0到255.
* uint16:无符号的正整数类型:占16位:从0到2的8次方减1.
* uint32:无符号的正整数类型:占32位:从0到2的32次方减1.
* uint64:无符号的正整数类型:占64位:从0到2的64次方减1.
* uintptr:无符号的储存指针位置的类型。也就是所谓的地址类型。
* rune :等于int32:这里是经常指文字符。
* byte:等于uint8:这里专门指字节符
* string:字符串:通常是一个切片类型:数组内部使用rune
* float32:浮点型:包括正负小数:IEEE-754 32位的集合
* float64:浮点型:包括正负小数:IEEE-754 64位的集合
* complex64:复数:实部和虚部是float32
* complex128:复数:实部和虚部都是float64
* error:错误类型,真实的类型是一个接口。
* bool:布尔类型

**格式化输出**
在Go语言的 fmt.Errorf 函数中，您可以使用格式化字符串和参数来构建错误消息。fmt.Errorf 的格式化参数采用类似于 Printf 函数的语法，其中使用占位符 % 来表示格式化参数，并根据需要提供相应的值。

以下是一些常见的格式化占位符：
```
    %s：字符串参数。
    %d：整数参数。
    %f：浮点数参数。
    %t：布尔值参数。
    %v：任意类型的参数。
```
您可以在格式化字符串中使用这些占位符，并在函数调用时提供相应的值。例如：
```go
import "fmt"

func main() {
name := "John"
age := 30
err := fmt.Errorf("Error: Name %s, Age %d", name, age)
fmt.Println(err)
}
```
上面的示例将输出 Error: Name John, Age 30。

您还可以使用其他格式化标识符和选项来控制输出的格式，如精度、填充等。更多的格式化选项可以在标准库的 fmt 包文档中找到。

**基础组件分为以下几种**
* 引用类型
	* slice
	* interface
	* chan
	* map
* 非引用类型
	* array
	* func
	* struct

**声明包和引用包**
```go
package main

import (
"fmt"
"lind-go/common"
//自定义的本项目的包
_ "lind-go/common"
)
```
**赋值符号**
```
var a
b :=
```
其中var 这种方式不论是局部还是全局变量都可以使用，但是后者也就是:=只有局部变量可以使用。也就是只有函数内部才能使用。
并且，var后面的变量后面的类型是可以省略的，省略后，go会在编译过程中自动判断。所以如果不省略就是长这样 var a int 。
**数组的初始化**
```go
// 初始化的方式1
a := [6]string{}
// 初始化的方式2
var a [6]string

a[0] = "0"
a[1] = "1"
a[2] = "2"
a[3] = "3"
a[4] = "4"
a[5] = "5"
```
***二维数组初始化***
```go
a:= [][2]string{
		{"WWW-Authenticate", fmt.Sprintf("Basic realm=%s", realm)},
		{"client", fmt.Sprintf("client=%s", clientId)},
	}
// [][2]string表示一个二维数组，其中每个元素都是一个包含两个字符串的数组。换句话说，它是一个由多个长度为2的字符串数组组成的切片
// 这种类型在处理需要二维结构的数据时非常有用，比如表示一组键-值对、坐标点等。
```
# 二 golang 推荐的命名规范
很少见人总结一些命名规范，也可能是笔者孤陋寡闻， 作为一个两年的golang 开发者， 我根据很多知名的项目，如 moby, kubernetess 等总结了一些常见的命名规范。
命名规范可以使得代码更容易与阅读， 更少的出现错误。

**文件命名规范**
由于文件跟包无任何关系， 而又避免windows大小写的问题，所以推荐的明明规范如下：
* 文件名应一律使用小写
* 不同单词之间用下划线分割
* 命名应尽可能地见名知意

**常量命名规范**
常量明明用 camelcase来命名示例如下
```
const todayNews = "Hello"

// 如果超过了一个常量应该用括号的方法来组织

const (
systemName = "What"
sysVal = "dasdsada"
)
```
**变量命名规范**
与常量命名方式一样，变量也应该使用驼峰的命名方式, 但注意尽量不与包名一致或者以包名开头
```
var x string
x := new(string)
```
**函数命名规范**
由于Golang的特殊性（用大小写来控制函数的可见性），除特殊的性能测试与单元测试函数之外, 都应该遵循如下原则
* 使用驼峰命名
* 如果`包外不需要访问请用小写开头`的函数
* 如果需要`暴露出去给包外访问需要使用大写开头`的函数名称
一个典型的函数命名方法如下：
```go
// 注释一律使用双斜线， 对象暴露的方法
func (*fileDao) AddFile(file *model.File) bool {
result := db.NewRecord(*file)
if result {
db.Create(file)
}
return result
}

// 不需要给包外访问的函数如下
func removeCommaAndQuote(content string) string {
re, _ := regexp.Compile("[\\`\\,]+")
return strings.TrimSpace(re.ReplaceAllString(content, ""))
}
```
**接口命名规范**
接口命名也是要遵循驼峰方式命名， 可以用 type alias 来定义大写开头的type 给包外访问
```
type helloWorld interface {
func Hello();
}

type SayHello helloWorld
```
**Struct命名规范**
* 与接口命名规范类似

**receiver 命名规范**
golang 中存在receiver 的概念 receiver 名称应该尽量保持一致， 避免this, super，等其他语言的一些语义关键字如下
* 定义类型A
* 定义类型A的方法methodA，没有参数，没有返回值
* 定义类型A的方法 methodB， 参数int类型的size， 返回值是string类型， 在methodB中调用methodA
```
type A struct{}

func (a *A) methodA() {
}
func (a *A) methodB(size int) string {
a.methodA()
}
```

**注释规范**
* 注释应一律使用双斜线

# 三 golang 方法接收者
* [定义]： golang的方法(Method)是一个带有receiver的函数Function，Receiver是一个特定的struct类型，当你将函数Function附加到该receiver， 这个方法Method就能获取该receiver的属性和其他方法。
* [面向对象]： golang方法Method允许你在类型上定义函数，是一个面向对象的行为代码， 这也有一些益处：同一个package可以有相同的方法名， 但是函数Function却不行。
```go
func (receiver receiver_type) some_func_name(arguments) return_values
```
从应用上讲，方法接受者分为值接收者，指针接收者，初级golang学者可能看过这两个接收者实际表现， 但是一直很混淆，很难记忆。

## 值类型方法接收者
* 值接受者： receiver是struct等值类型。
下面定义了值类型接受者Person, 尝试使用Person{}, &Person{}去调用接受者函数。
```
package main
import "fmt"

type Person struct {
 name  string
 age int
}

func (p Person) say() {
 fmt.Printf("I (%p) ma %s, %d years old \n",&p, p.name,p.age)
}

func (p Person) older(){  // 值类型方法接受者： 接受者是原类型值的副本
 p.age = p.age +1
 fmt.Printf("I (%p) am %s, %d years old\n", &p, p.name,p.age)
}

func main() {
  p1 := Person{name: "zhangsan", age: 20}
  p1.older()
  p1.say()
  fmt.Printf("I (%p) am  %s, %d years old\n",&p1, p1.name,p1.age)

  p2 := &Person{ name: "sili", age: 20}
  p2.older()   // 即使定义的是值类型接受者， 指针类型依旧可以使用，但我们传递进去的还是值类型的副本
  p2.say()
  fmt.Printf("I (%p) am %s, %d years old\n",p2, p2.name,p2.age)
}
```
尝试改变p1=Person{},p2=&Person{}的字段值:
```
I (0xc000098078) am zhangsan, 21 years old
I (0xc000098090) ma zhangsan, 20 years old
I (0xc000098060) am  zhangsan, 20 years old
I (0xc0000980c0) am sili, 21 years old
I (0xc0000980d8) ma sili, 20 years old
I (0xc0000980a8) am sili, 20 years old
```
p1=Person{} 未能修改原p1的字段值； p2=&Person{}也未能修改原p2的字段值。
* 通过Person{}值去调用函数， 传入函数的是原值的副本， 这里通过第一行和第三行的%p印证 (%p：输出地址值， 这两个非同一地址)。
* 即使定义的是值类型接收者，指针类型依旧可以调用函数， 但是传递进去的还是值类型的副本。
> 带来的效果是：对值类型接收者内的字段操作，并不影响原调用者。

## 指针类型接受者
方法接收者也可以定义在指针上，任何尝试对指针接收者的修改，会体现到调用者。
```
package main

import  "fmt"

type Person struct{
 name string
 age int
}

func  (p Person) say(){
 fmt.Printf("I (%p)  am %s, %d years old\n", &p, p.name,p.age)
}

func (p *Person) older(){   // 指针接受者，传递函数内部的是原类型值（指针）， 函数内的操作会体现到原指针指向的空间
 p.age = p.age +1
 fmt.Printf("I (%p)  am %s, %d years old\n", p, p.name,p.age)
}

func main() {
 p1 := Person{"zhangsan",20}
 p1.older()  // 虽然定义的是指针接受者，但是值类型依旧可以使用，但是会隐式传入指针值
 p1.say()
 fmt.Printf("I (%p)  am %s, %d years old\n", &p1, p1.name,p1.age)

 p2:= &Person{"sili",20}
 p2.older()
 p2.say()
 fmt.Printf("I (%p)  am %s, %d years old\n", p2, p2.name,p2.age)
}
```
尝试改变p1= Person{}, p2=&Person{}字段值
```
I (0xc000098060)  am zhangsan, 21 years old
I (0xc000098078)  am zhangsan, 21 years old
I (0xc000098060)  am zhangsan, 21 years old
I (0xc000098090)  am sili, 21 years old
I (0xc0000980a8)  am sili, 21 years old
I (0xc000098090)  am sili, 21 years old
```
p1=Person{} 成功修改字段值，p2=&Person{}也成功修改字段值。
* 通过p1也可以调用指针函数接收者， 但是实际会隐式传递指针值。
* 指针接收者，入参是原指针值，函数内的操作会体现到原调用者。
> 带来的效果： 任何对指针接收者的修改会体现到 原调用者。

## 什么时候使用指针接收者
* 需要对接受者的变更能体现到原调用者
* 当struct占用很大内存，最好使用指针接受者，否则每次调用接受者函数 都会形成struct的大副本

## golang方法的几种姿势
* 将接收者函数当扩展函数
```go
Person.say(p1)
(*Person).older(p2)
```
* golang 方法链条
```go
func (p Person) printName() Person{
  fmt.Printf("Name:%s", p.Name)
  return p
}
* Non_struct类型golang方法
```go
type myFloat float64
func (m myFloat) ceil() float64 {
   return  math.Ceil(float64(m))
}
```
# 四 golang 接口和结构体
Go语言中没有“类”的概念:也不支持像继承这种面向对象的概念。但是Go 语言的结构体与“类”都是复合结构体:而且Go 语言中结构体的组合方式比面向对象具有更高的扩展性和灵活性。
**结构体的定义和初始化**
```go
结构体定义:
type identifier struct {
field1 type1
field2 type2
}

结构体实例创建
s1 := new(Student) //第一种方式
s2 := Student{“james”, 35} //第二种方式
s3 := &Student { //第三种方式
Name: “LeBron”,
Age: 36,
}
```

**common文件夹定义常量**
>当前目录的common子目录下建立Info.go文件:内容如下:
```go
package common

const (
	Name  = "lind"
	Title = "go语言"
)

```
**接口定义**
```go
// 接口定义
type Usb interface {
	Start()
	Stop()
}
```
**结构体实现接口**
```go
// Mp3 将接口两个方法都实现了
type Mp3 struct {
	Name string
}

func (c *Mp3) Start() {
	fmt.Println(c.Name, "start working")
}
func (c *Mp3) Stop() {
	fmt.Println(c.Name, "stop work")
}
```
**定义一个使用接口的结构:它的方法参数是接口Usb的实例对象
```go
// 定义使用者的结构体
type Computer struct {
	Name string
}

// Working方法接收Usb接口对象:然后调用它的两个方法
func (c *Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}
```
**在main方法中使用它:程序启动时执行
```go
func main() {
	fmt.Println("auth:%s", common.Name)
	mp3 := Mp3{"mp3"}
	computer := Computer{"computer"}
	computer.Working(&mp3)
}
```
# 五 数据和切片的区别
在Go语言中，数组（array）和切片（slice）是两种不同的数据结构，它们有一些重要的区别。

### 数组（Array）

- 数组是一个固定长度的数据结构，定义时需要指定数组的长度。
- 数组的长度是其类型的一部分，因此`[5]int`和`[10]int`是两种不同的类型。
- 数组的长度是不可变的，一旦定义后无法改变。

```go
var arr1 [5]int // 定义一个包含5个整数的数组
```
### 切片（Slice）

- 切片是对数组的一个引用，它是一个动态长度的序列。
- 切片没有固定的长度，在创建时可以指定初始长度，但是可以动态增长。
- 切片是一个引用类型，它的底层是一个数组。

```go
var slice1 []int // 定义一个整数切片
```

### 区别总结

1. **长度：** 数组的长度是固定的，而切片的长度是动态可变的。
2. **引用：** 切片是对数组的引用，对切片的修改会影响到底层数组，而数组是一个完整的数据结构。
3. **传递方式：** 在函数参数传递中，数组是值传递，而切片是引用传递。
4. **使用场景：** 通常情况下，切片更加灵活，更适合在实际开发中使用。数组则更适合于长度固定且预先知道大小的情况。

# 六 对指针类型的理解
在Go语言中，指针是一种特殊的数据类型，它存储了一个变量的内存地址。通过指针，我们可以直接访问或修改变量的值，而不是对变量进行复制或传递。

下面是对Go语言中指针类型的一些理解：

1. **声明和使用：** 在Go语言中，可以使用`*`操作符来声明指针类型，例如`var ptr *int`表示ptr是一个指向int类型的指针。要获取变量的地址，可以使用`&`操作符，比如`&variable`会返回variable的内存地址。

2. **传递参数：** 在函数参数列表中，可以使用指针类型作为参数，这样可以直接修改原始变量的值。这在需要修改函数外部变量的情况下非常有用。

3. **避免拷贝：** 通过传递指针而不是值，可以避免在函数调用时进行大量的内存拷贝，从而提高程序的性能。

4. **空指针：** 在Go语言中，指针的零值是nil，表示指针不指向任何有效的内存地址。

5. **安全性：** Go语言中的指针不能进行指针运算，这在一定程度上提高了程序的安全性。

6. 总的来说，在Go语言中，指针是一种强大的工具，可以用于管理内存、传递参数、避免拷贝等场景。然而，需要小心使用指针，以避免出现潜在的内存错误和问题。

# 七 方法参数指针传递的好处
> 相当于java和.net里的引用类型

在Go语言中，方法参数可以传递变量的地址（指针）或者直接传递变量值。使用指针作为方法参数有以下几个好处：

1. **修改原始变量**：通过传递变量的地址，方法可以直接修改原始变量的值。这意味着你可以在方法内部修改调用者的变量，而不是创建一个变量的副本。

2. **节省内存和时间**：当传递大的数据结构（比如大型数组、切片或结构体）时，传递地址可以避免复制整个数据结构，从而节省内存和时间。

3. **避免拷贝**：在Go语言中，将大的数据结构作为参数传递时，会发生一次浅拷贝。如果数据结构很大，这可能会影响性能。通过传递地址，可以避免这种拷贝。

4. **支持可选参数**：通过传递指针，可以实现方法的可选参数。通过检查指针是否为nil，方法可以知道是否需要处理某些参数。

尽管使用指针作为方法参数具有上述好处，但也需要谨慎使用。过度地使用指针可能会导致代码难以理解和维护，因此需要根据具体情况权衡使用指针和传值的利弊。

# 泛型类型
在Go语言中，这种定义使用了类型参数，是Go语言的泛型特性之一。这个定义表明 ClusterClient 是一个泛型类型，它接受一个类型参数 C，而 C 必须是 Cluster 类型或者实现了 Cluster 接口的类型。

通过这种方式，我们可以定义一个通用的 ClusterClient 类型，在实例化时指定具体的类型参数，从而使得 ClusterClient 可以适用于不同的 Cluster 类型。

举个例子，假设我们有两种不同的 Cluster 类型 ACluster 和 BCluster，它们都实现了 Cluster 接口，那么我们可以分别实例化两个不同类型的 ClusterClient：

```go
type ACluster struct {
    //...
}

type BCluster struct {
    //...
}

func (a ACluster) Method() {
    //...
}

func (b BCluster) Method() {
    //...
}

type Cluster interface {
    Method()
}

type ClusterClient[C Cluster] struct {
    cluster C
}

func main() {
    a := ClusterClient{cluster: ACluster{}}
    b := ClusterClient{cluster: BCluster{}}
}
```
通过这样的定义，我们可以很方便地创建不同类型的 ClusterClient 实例，而不需要为每种 Cluster 类型都单独定义一个 ClusterClient 类型。这样的泛型特性可以提高代码的复用性和灵活性。

# 组合代替继承
go中没有类，所以也就没有类的继承，如果希望使用面向对象中的继承特性，你可以通过组合来实现，即将结构的实例做为另一结构里的字段即可，如下代码实例：
```go
type RedisConfig struct {
	name string
}

func (c *RedisConfig) printName() {
	fmt.Println(c.name)
}

type LindRedisConfig struct {
	RedisConfig
	info string
}
func main() {
lindRedisConfig := LindRedisConfig{RedisConfig{"lind"}, "info"}
lindRedisConfig.printName() #可以直接使用RedisConfig中的方法printName，或者RedisConfig中的字段name
}
```