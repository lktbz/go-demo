package main

func main() {
	//传播错误。这意味着函数中某个子程序的失败，会变成该函数的失败.把流程中某个子函数的错误“传播”给主流程函数，并中断

	//res, err := subFunc(arg)
	//if err != nil{
	//	return nil, err
	//}
	//这样的错误返回也可以包装下
	//if err != nil {
	//	arg := "zs"
	//	return nil, fmt.Errorf(" %s : %v", arg,err)
	//}

	//重试

	/**
	如果错误的发生时偶然性的（类似TCP的部分错误处理），那么我们可以采用重试的策略，
	但是要注意重试的时间和次数，防止无限制的重试。在所有重试之后如果还是失败的话，再返回错误。
	 */

	// WaitForServer 尝试连接url参数对应的服务器
	// 它持续一分钟的重连，并采用指数级的等待时间增加
	// 如果所有重试都失败了，将返回错误
	//func WaitForServer(url string) error {
	//	const timeout = 1 * time.Minute
	//	deadline := time.Now().Add(timeout)
	//	for tries := 0; time.Now().Before(deadline); tries++ {
	//	_, err := http.Head(url)
	//	if err == nil {
	//	return nil // success
	//}
	//	log.Printf("server not responding (%s);retrying…", err)
	//	time.Sleep(time.Second << uint(tries)) // 指数递增
	//}
	//	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
	//}

	//第一时间结束
	//
	//这个策略一般用在 main 文件中，当主流程遇到错误时，直接退出结束程序。
	//
	//这个策略一般与 错误传播 策略合用，将子函数传播至主流程中，然后依照错误的严重性判断是否结束程序。
	//
	//// 在主程序中
	//if err := WaitForServer(url); err != nil {
	//	log.Fatalf("wrong %v\n",err)
	//	os.Exit(1)
	//}

	//只输出错误信息，不中断
	//
	//就是调用日志系统，常见于一些小问题，如图片丢失等。
	//
	//忽略错误
	//
	//dir, err := ioutil.TempDir("", "scratch")
	//if err != nil {
	//	return fmt.Errorf("failed to create temp dir: %v",err)
	//}
	//// ...use temp dir…
	//os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically


}

func subFunc(arg interface{}) (interface{}, interface{}) {
return "",'"'
}
