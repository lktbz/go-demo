package main

import "fmt"

/**
接口是编程的规范，他也可以作为函数的参数，以让函数更具备适用性。在下列示例中，有三个接口动物接口、飞翔接口、游泳接口，两个实现类鸟类与鱼类：

鸟类：实现了动物接口，飞翔接口
鱼类：实现了动物接口，游泳接口
*/
//通用接口动物，具有呼吸的功能
type Animal interface {
	Breath()
}

//飞行能力
type Flyer interface {
	Fly()
}

//游泳接口
type Swimer interface {
	Swim()
}

// 定义一个鸟类：其呼吸的方式是在陆地

type Bird struct {
	Name string
	Food string
	Kind string
}

func (b *Bird) Breath() {
	fmt.Println("鸟在陆地上呼吸")
}
func (b *Bird) Fly() {
	fmt.Printf("%s 在 飞", b.Name)
}

type Fish struct {
	Name string
	Kind string
}

func (f *Fish) Breath() {
	fmt.Println("鱼 在 水下 呼吸")
}
func (f *Fish) Fly() {
	fmt.Printf("%s 在游泳\n", f.Name)
}

/**
接口类型无法直接访问其具体实现类的成员，需要使用断言（type assertions），对接口的类型进行判断，类型断言格式：

t := i.(T)			//不安全写法：如果i没有完全实现T接口的方法，这个语句将会触发宕机
t, ok := i.(T)		// 安全写法：如果接口未实现接口，将会把ok掷为false，t掷为T类型的0值
i代表接口变量
T代表转换的目标类型
t代表转换后的变量
*/

func Display(a Animal)  {
	//直接调用接口中的方法
	a.Breath()
	// 调用实现类的成员：此时会报错
	ins,ok:=a.(*Bird)// 注意：这里必须是 *Bird类型，因为是*Bird实现了接口，不是Bird实现了接口
	if ok{
		fmt.Println("该鸟类的名字为：",ins.Name)
	}else{
		fmt.Println("该动物不是鸟类")
	}
}
/**
   接口调用的不同转换方式
 */
func Display1(a Animal) {
	instance, ok := a.(Flyer)			// 动物接口转换为了飞翔接口
	if ok {
		instance.Fly()
	} else {
		fmt.Println("该动物不会飞")
	}

}
func main() {
	/**
	  普通调用方式
	*/
	//b:=new(Bird)
	//b.Name="乌鸦"
	//b.Food="虫子"
	//b.Kind="鸟类"
	//b.Breath()
	//b.Fly()

	var b = &Bird{
		"斑鸠",
		"蚂蚱",
		"鸟类",
	}
	Display(b)
}
