package main

import "fmt"

func main() {
        //常规写法
        for i := 0; i <10 ; i++ {
	fmt.Printf(" 遍历的值为：%d",i)
        }
        //变体写法
        //for condition ==while
        ///for{} ==for(;;)
        /**
        for key, value := range oldMap {
            newMap[key] = value
        }
        可以对lice、map、数组、字符串等进行迭代循环
         */
        var b int = 15
        var a int
        numbers := [6]int{1, 2, 3, 5}
        for i := 0; i <10 ; i++ {
	fmt.Println(i)
        }
        for a<b{
                a++
               fmt.Println("a的值为",a)
        }
        //遍历数组
        for i, x := range numbers {
      	fmt.Printf("-元素的下表为%d,元素的值为：%d",i,x)
        }
       //跳出循环
        for i := 1; i <= 10; i++ {
	if i > 5 {
	        break //i>5跳出循环
	}
	fmt.Printf("%d ", i)
        }
        fmt.Printf("\nline after for loop")

        //continue 语句
        for i := 1; i <= 10; i++ {
	if i%2 == 0 {
	        continue
	}
	fmt.Printf("lalla%d ", i)
        }
       //goto：可以无条件地转移到过程中指定的行。
        /* 定义局部变量 */
        var azc int = 10
        /* 循环 */
        LOOP: for azc < 20 {
	if azc == 15 {
	        /* 跳过迭代 */
	        azc = azc + 1
	        goto LOOP
	}
	fmt.Printf("azc的值为 : %d\n",azc)
	azc++
}






}
