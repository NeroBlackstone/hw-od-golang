package main

import "main/mytest"

// var total int = 1000
// var wg sync.WaitGroup

//	func main() {
//		for i := 0; i < 1000; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				total -= 1
//			}()
//		}
//		wg.Wait()
//		// 打印一下total的值
//		fmt.Println(total)
//	}
func main() {
	mytest.HJ5()
}
