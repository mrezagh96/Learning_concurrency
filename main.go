package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go methode(ch, &wg)

	majicnumber := <-ch
	wg.Wait()
	fmt.Println(majicnumber)
}
func methode(ch chan int, wg *sync.WaitGroup) {

	ch <- 1111
	fmt.Println("finish1")
	fmt.Println("finish2")
	fmt.Println("finish3")
	fmt.Println("finish4")
	fmt.Println("finish5")
	fmt.Println("finish6")
	fmt.Println("finish7")
	fmt.Println("finish8")
	wg.Done()
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// type Counter struct {
// 	sync.Mutex
// 	value int
// }
//
// func main() {
// 	fmt.Println("start")
//
// 	var wg sync.WaitGroup
// 	counter := Counter{}
//
// 	for i := 0; i < 20000; i++ {
// 		wg.Add(1)
// 		go increament(&counter, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(counter.value)
// 	fmt.Println("end")
// }
//
// func increament(counter *Counter, wg *sync.WaitGroup) {
// 	counter.Lock()
// 	I := counter.value
// 	counter.value = I + 1
// 	wg.Done()
// 	counter.Unlock()
// }

//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//func main() {
//
//	var wg sync.WaitGroup
//
//	function0 := func() {
//		fmt.Println("1- this is a test for Github and git structures checking")
//		fmt.Println("2- hi here is  mohammadreza")
//		wg.Done()
//	}
//	wg.Add(1)
//	go function0()
//	wg.Wait()
//
//	function1(&wg)
//	wg.Wait()
//
//}
//
//func function1(Wg *sync.WaitGroup) {
//	Wg.Add(1)
//	go aaa(Wg)
//}
//
//func aaa(Wg *sync.WaitGroup) {
//	fmt.Println("3- say hello to everybody ")
//	fmt.Println("4- this is an edit an other device ")
//	Wg.Done()
//}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//
//func main() {
//	fmt.Println("start")
//	message("Toplearn.com")
//	wg.Wait()
//	fmt.Println("End")
//}
//
//func message(text string) {
//	wg.Add(10)
//	for i := 1; i <= 10; i++ {
//		go printMessage(text, i)
//	}
//}
//
//func printMessage(text string, index int) {
//	if index == 2 {
//		time.Sleep(2 * time.Second)
//	}
//	fmt.Println(text, index)
//	wg.Done()
//}
