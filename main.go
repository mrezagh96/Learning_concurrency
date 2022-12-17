package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")

	var wg sync.WaitGroup
	testint := 0

	for i := 0; i <= 20000; i++ {
		wg.Add(1)
		go increament(&testint, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func increament(pi *int, wg *sync.WaitGroup) {
	I := *pi
	fmt.Println(I)
	*pi++
	wg.Done()
}

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
