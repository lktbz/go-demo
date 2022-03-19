package main

import "fmt"

/**
go并不是一个纯面向对象的编程语言。在go中的面向对象，结构体替换了类。

Go并没有提供类class，但是它提供了结构体struct，方法method，可以在结构体上添加。提供了捆绑数据和方法的行为，这些数据和方法与类类似
 */

type Emps struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}
//方法
func (es Emps)LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", es.FirstName, es.LastName, (es.TotalLeaves - es.LeavesTaken))
}
func main() {
 	/**
	  构造函数写法

	e:=Emps{
 		FirstName:"sam",
 		LastName:"adof",
		TotalLeaves: 20,
		LeavesTaken: 30,
	}
	e.LeavesRemaining() 	 */

	/**
	  现在传入空构造会发生什么？
	 */
	e:=Emps{}
	e.LeavesRemaining()
	//  has 0 leaves remaining  显然不符合我们的期望
	/**
	通过运行结果可以知道，使用Employee的零值创建的变量是不可用的。它没有有效的名、姓，也没有有效的保留细节。在其他的OOP语言中，比如java，这个问题可以通过使用构造函数来解决。使用参数化构造函数可以创建一个有效的对象。

	go不支持构造函数。如果某个类型的零值不可用，则程序员的任务是不导出该类型以防止其他包的访问，并提供一个名为NewT(parameters)的函数，该函数初始化类型T和所需的值。在go中，它是一个命名一个函数的约定，它创建了一个T类型的值给NewT(parameters)。这就像一个构造函数。如果包只定义了一个类型，那么它的一个约定就是将这个函数命名为New(parameters)而不是NewT(parameters)
	 */
}
