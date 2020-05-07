package methods

import "fmt"

/* PART -1 : Functions & Methods on String Operations */
	/* Part - 1.1 : This section describes functions in golang based on call by value and call by reference */
	/* Part - 1.2 : Next section describes methods in golang based on call by value and call by reference */

/* PART -2 : Functions & Methods on Maps and Slices Operations */
	/* Part - 2.1 : This section describes functions in golang based on call by value and call by reference */
	/* Part - 2.2 : Next section describes methods in golang based on call by value and call by reference */


type empDetails struct {
	name string
	designation string
}

func DemoMethodsAndFunctions(){
	fmt.Println("How to use functions with pass by value:")
	callFunctionByValueWithStrings()
	callFunctionByPoitnerWithStrings()
	callMethodByValueWithStrings()
	callMethodByReferenceWithStrings()
}

func callFunctionByValueWithStrings(){
	fmt.Println("******   PART -1.1 ( Functions, call by value , operations on strings   **********************")

	empInfo := empDetails{"Ravi", "Software Engineer"}
	fmt.Println("Before call by value ", empInfo)
	changeEmpDesignation(empInfo)
	fmt.Println("After call by value - Did designation changed to Architect, NO", empInfo)
}
func changeEmpDesignation(details empDetails){

	fmt.Println("Emp designation is changing to Architect: ")
	details.designation = "Architect"
}


func callFunctionByPoitnerWithStrings() {

	fmt.Println("******    Functions, call by reference , operations on strings   ********************** ")

	empInfo := empDetails{"Ravi", "Software Engineer"}
	fmt.Println("Before call by reference", empInfo)
	changeEmpDesignationByPointer(&empInfo)
	fmt.Println("After call by reference- Did designation changed to Architect, Yes", empInfo)
}

func changeEmpDesignationByPointer(details *empDetails){

	fmt.Println("Emp designation is changing to Architect: ")
	details.designation = "Architect"
}

type custInterface  interface{
	//changeEmpDesignationByMethod(string)
}

type custDetails struct{
	name string
	designation string
}
func callMethodByValueWithStrings(){
	fmt.Println("******   PART -2.1 ( Methods , call by value , operations on strings   **********************")

	var customerInterface = custDetails{"Ravi", "Software Engineer"}
	fmt.Println("Before call by value ", customerInterface)
	customerInterface.changeEmpDesignationMethod("Architect")
	fmt.Println("After call by value - Did designation changed to Architect, NO", customerInterface)
}
func (e custDetails) changeEmpDesignationMethod(desig  string){

	fmt.Println("Emp designation is changing to Architect: ")
	e.designation = desig
}

func callMethodByReferenceWithStrings(){
	fmt.Println("******   PART -2.2 ( Methods , call by reference, operations on strings   **********************")

	var customerInterface = &custDetails{"Hari", "Software Engineer"}
	fmt.Println("Before call by value ", customerInterface)
	customerInterface.changeEmpDesignationByMethodByRef("Architect")
	fmt.Println("After call by value - Did designation changed to Architect, Yes", customerInterface)
}
func (e *custDetails) changeEmpDesignationByMethodByRef(desig  string){

	fmt.Println("Emp designation is changing to Architect: ")
	e.designation = desig
}
