package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n",now)
	fmt.Printf("%t\n",now)
	fmt.Println(now)
	fmt.Println("location--",now.Location().String())
	year:=now.Year()
    fmt.Println("year:=",year)
	//自定义时间
	customTime := time.Date(2008, 7, 15, 13, 30,0,0, time.Local)
	fmt.Println(customTime)
	nowTime := time.Now()

	year, month, day := nowTime.Date()
	fmt.Println(year, month, day)		// 2019 November 01

	hour, min, sec := nowTime.Clock()
	fmt.Println(hour, min, sec)

	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Hour())

	fmt.Println(nowTime.YearDay())
	//时间睡眠
	time.Sleep(time.Second * 3)

	//时间戳
	timestamp1 := now.Unix() //时间戳
	timestamp2 := now.UnixNano()//纳秒时间戳

	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	//时间戳转换为时间
	tins:=time.Unix(timestamp1,0)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", tins.Year(), tins.Month(), tins.Day, tins.Hour(), tins.Minute(), tins.Second())


	//定时器
	//tick:=time.Tick(time.Second*1)
	//for i := range tick {
	//	fmt.Println(i)
	//}


	//时间格式化
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	//解析字符串时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("loc=",loc,"---",now.In(loc))
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))



}
