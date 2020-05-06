package slicesMaps

import "fmt"

func DemoSlicesMaps(){
	/* Show to delete elements in slice */

	strSlice := []string{"Hi","How","are","you"}

	index :=2

	fmt.Println("Before Remove Slice =",strSlice)

	s := deleteSlice(strSlice,index)
	fmt.Println("After Remove Slice =",s)

	strAppSlice := appendSlices()
	fmt.Println("Appended Slice = ",strAppSlice)

}

func deleteSlice(strTempSlice []string,index int) []string{
	return append(strTempSlice[:index],strTempSlice[index+1:]...)
}

func appendSlices() []string{

	sliceOne :=[]string{"Hi","HOw","are","you"}
	sliceTwo :=[]string{"Hi","How","is","you"}

	return append(sliceOne,sliceTwo...)
}