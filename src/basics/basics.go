package basics

import (
	"fmt"
	"time"
)

func Loops(){
	fmt.Println("Basics Demo...")
	forLoop()
}

func forLoop(){
	for i:=0 ; i<10; i++ {
		fmt.Println("Print i value ",i)
	}

	i :=10

	for ; i<= 100 ; {
		fmt.Println("Print i value in second loop ", i)
		i +=10
	}

	/* Multiple Initilization and incrementing */

	for no,i :=1,10; no <= 10 && i <= 100 ; no,i = no+1, i+1 {
		fmt.Println("Prininting both values",no,i)
	}

/*	for {
		fmt.Println("...it's infinite loop")
		time.Sleep(10)
	}*/

	states := []string{"AP","Chennai","Bangalore","Telangana"}
	for i:=0 ;i<len(states);i++ {
		fmt.Println("States Name ", i, states[i])
	}
	for i,state := range states {
		fmt.Println("Range of states ", i, state)
	}

	statesmap :=map[string]string{"India" : "Delhi", "Canada":"Ottawa","America":"washington"}

	for index,value := range statesmap {
		fmt.Println("Index - ", index, "value - " , value)
	}
	for key := range statesmap {
		fmt.Println("Only key - ", key)
	}

	time.Sleep(10)

	integers := make([]int, 10)
	for i = range integers {
		integers[i] = i
	}
	fmt.Println("Integers ", integers)


	name :="Alexa"
	for _, letter := range name {
		fmt.Printf("Character = %c",letter)
	}

	fmt.Println("Range of nested lists ")

	nestedLists := [][]int {
		[]int{1,2,3},
		[]int{4,5,6,10},
		[]int{9,5,6,10},
	}

	for _, firstList := range nestedLists {
		fmt.Println("\nFirst Lists", firstList)
		for _,secondList := range firstList {
			fmt.Print("Inside Lists \t", secondList,"\t")
		}
	}

}
