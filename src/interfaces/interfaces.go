package interfaces

import "fmt"

type custInterface  interface{
	changeEmpDesignationByValue(string)
	changeEmpDesignationByRef(string)
}

type custDetails struct{
	name string
	designation string
}

func DemoInterface(){
	fmt.Println("Demo on Interface ")
	callInterfaceByValueWithString()
	callInterfaceByRefWithString()
}
func callInterfaceByValueWithString(){
	fmt.Println("******   PART 1 ( Inteface Methods , call by value , operations on strings   **********************")

	var customerInterface custInterface = &custDetails{"Ravi", "Software Engineer"}
	fmt.Println("Before call by value ", customerInterface)
	customerInterface.changeEmpDesignationByValue("Architect")
	fmt.Println("After call by value - Did designation changed to Architect, NO", customerInterface)
}
func (e custDetails) changeEmpDesignationByValue(desig  string){

	fmt.Println("Emp designation is changing to Architect: ")
	e.designation = desig
}

func callInterfaceByRefWithString(){
	fmt.Println("******   PART -2  Interface Methods , call by reference, operations on strings   **********************")

	var customerInterface custInterface= &custDetails{"Hari", "Software Engineer"}
	fmt.Println("Before call by value ", customerInterface)
	customerInterface.changeEmpDesignationByRef("Architect")
	fmt.Println("After call by value - Did designation changed to Architect, Yes", customerInterface)
}
func (e *custDetails) changeEmpDesignationByRef(desig  string){

	fmt.Println("Emp designation is changing to Architect: ")
	e.designation = desig
}

