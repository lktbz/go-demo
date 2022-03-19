package main

import "fmt"

func main() {
/**
	6.1 延迟是什么?
  即延迟（defer）语句，延迟语句被用于执行一个函数调用，在这个函数之前，延迟语句返回。

  6.2 延迟函数
  你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

  如果有很多调用defer，那么defer是采用后进先出模式
  在离开所在的方法时，执行（报错的时候也会执行）
 */

	a := 1
	b := 2
	defer fmt.Println("延迟执行B的值为：",b)
	fmt.Println(a)
	b=3
	fmt.Println("修改b的值:",b)
	//通过比较延迟加载的和修改后的值，他俩并补相等
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)
}
func finished() {
	fmt.Println("Finished finding largest")
}
func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}