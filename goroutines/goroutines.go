package goroutines

import (
	"fmt"
	"time"
)

func DemoGoRoutines(){
	fmt.Println("Demo on Go Routines ")
	callGoRoutines()
}

func callGoRoutines() {

	fmt.Println("Calling two asynchronous printers  threads")
	go startPrinter1()
	go startPrinter2()

	/* making part of the code snippet as asynchronous thread */
	go func () {
		for i :=0 ; i< 1000; i++ {
			fmt.Println("Inside main function ",i)
			time.Sleep(1)
		}
	}()

	fmt.Print("Main thread is waiting on 20sec sleep")
	time.Sleep(100)
}

func startPrinter1(){
	i := 0
	for ;i<1000; {
		fmt.Println("I am printer1>>>>>>>>>> ", i)
		time.Sleep(1)
		i = i+1
	}
}
func startPrinter2(){
	i := 0
	for ; i<10000; {
		fmt.Println("I am printer2<<<<<<<<<<<< ", i)
		time.Sleep(1)
		i = i+1
	}
}
