package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bfyxzls/lind-go/common"
	"github.com/bfyxzls/lind-go/generictype"
	"github.com/bfyxzls/lind-go/util"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 二分查找
func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		Mid := left + (right-left)>>1
		if nums[Mid] == target {
			return true
		} else if nums[Mid] > target {
			right = Mid - 1
		} else {
			left = Mid + 1
		}
	}
	return left < len(nums)
}

// 自定义类型
type Timestamps map[string]int64

func getTimestamps(t time.Time) *Timestamps {
	ts := Timestamps{}

	ye, mo, da := t.Year(), t.Month(), t.Day()
	ho, mi, se, lo := t.Hour(), t.Minute(), t.Second(), t.Location()

	ts["now"] = t.Unix()
	ts["second"] = time.Date(ye, mo, da, ho, mi, se, 0, lo).Unix()
	ts["minute"] = time.Date(ye, mo, da, ho, mi, 0, 0, lo).Unix()
	ts["hour"] = time.Date(ye, mo, da, ho, 0, 0, 0, lo).Unix()
	ts["day"] = time.Date(ye, mo, da, 0, 0, 0, 0, lo).Unix()
	ts["month"] = time.Date(ye, mo, 1, 0, 0, 0, 0, lo).Unix()
	ts["year"] = time.Date(ye, 1, 1, 0, 0, 0, 0, lo).Unix()

	return &ts
}

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

func testStringContain() {
	// 定义字符串数组
	strArr := []string{"apple", "banana", "cherry"}

	// 入参字符
	character := "apple?a=1"

	// 判断字符是否在数组中出现
	for _, str := range strArr {
		if strings.Contains(character, str) {
			fmt.Printf("字符 %s 包含数组中的元素 %s 中\n", character, str)
			return
		}
	}

	fmt.Printf("字符 %s 不存在于数组中的任何元素中\n", character)
}

func main() {
	testStringContain()

	lindRedisConfig := LindRedisConfig{RedisConfig{"lind"}, "info"}
	lindRedisConfig.printName()
	fmt.Println(lindRedisConfig.name)

	ipRange := "140.237.144.105-140.237.144.105"
	ips := strings.Split(ipRange, "-")
	startIP, _ := util.ConvertIPToScore(ips[0])
	endIP, _ := util.ConvertIPToScore(ips[1])

	fmt.Printf("startIp:%v,endIp:%v", startIP, endIP)
	now0 := time.Now()
	minuteAligned := now0.Truncate(time.Minute)
	timeStamp := strconv.FormatInt(minuteAligned.Unix(), 10)
	fmt.Println(timeStamp)

	// 泛型初始化与调用
	aCluster := generictype.NewClusterClient(generictype.ACluster{Title: "title"})
	aCluster.PrintMethod()

	bCluster := generictype.NewClusterClient(generictype.BCluster{Age: 40})
	bCluster.PrintMethod()

	// 定义新类型，遍历map类型
	ts := getTimestamps(time.Now())
	for k, v := range *ts {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}

	// int和[]byte的转换
	source := 10
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(source+1))
	ret := int64(binary.LittleEndian.Uint64(buf))
	fmt.Printf("reg=%v", ret)

	// 一个数组，每个元素是两个元素的小数组，可用于k/v存储
	a := [][2]string{
		{"WWW-Authenticate", fmt.Sprintf("Basic realm=%s", "baobao")},
		{"client", fmt.Sprintf("client=%s", "shop")},
	}
	fmt.Println(a[1][1])

	// url操作
	cp, _ := url.Parse("http://www.sina.com/home/index?type=3")
	fmt.Println(cp.Host)
	fmt.Println(cp.Path)
	fmt.Println(cp.RawQuery)
	fmt.Println(cp.Scheme)

	// 指针声明
	var aStart *int
	aVal := 10
	aStart = &aVal       //指针赋值，赋的是地址
	fmt.Println(aStart)  //输出地址
	fmt.Println(*aStart) //输出地址指向的值

	// 二分查找
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(binarySearch(nums, 10))
	fmt.Println(binarySearch(nums, 3))

	// 格式化输出
	fmt.Printf("auth:%s", common.Name)

	// 定义新对象
	mp3 := common.Mp3{"mp3"} // :=这个运算符可以使变量在不声明的情况下直接被赋值使用
	computer := common.Computer{"computer", nil}
	computer.Working(&mp3)
	//phone := Phone{"phone"} //没有实现usb的两个方法
	//computer.Working(&phone) //报错

	// 时间测试
	now := time.Now()
	// 返回日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	// 时间戳转换
	layout := "2006-01-02 15:04:05"
	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
	fmt.Println(t.Format(layout))
	const text = "My email is ccmouse@gmail.com"

	// 正则表达式
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	fmt.Println(match)

	// 直接初始化切片slices
	slices := []int{1, 2, 3}
	slices[1] = 22 //修改第2位元素值变成22
	fmt.Println(slices)
}
