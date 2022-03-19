package main

import (
	"fmt"
	"sync"
)

/**
开启多个协程对共享内存产生竞争，单独使用等待组不能解决问题
 */

func main(){
 //waitDemo()
lockAndWaitDemo()

}

func lockAndWaitDemo()  {
	var mt sync.Mutex
	var wg sync.WaitGroup
	var money=10000
	for i := 0; i <10 ; i++ {
		wg.Add(1)
		go func(index int) {
			mt.Lock()
			fmt.Printf("协程 %d 抢到锁\n", index)
			for j := 0; j <100 ; j++ {
				money+=10
			}
			fmt.Printf("协程 %d 准备解锁\n", index)
			mt.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("最终的monet = ", money)		// 应该输出20000才正确
}
//没有加锁
func waitDemo(){

	var wg sync.WaitGroup
	var money=10000

	for i := 0; i <10 ; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j <100 ; j++ {
				money+=10
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("最终的结果为=",money)
}
