package methods

import "fmt"


	/* Part 1 shows, how to call functions using pass by value, part -2 shows how to call functions using pass by references.
	  Part 3 shows how to use methods in case of pass by value and part 4 shows in case of pass by references. */




type empDetails struct {
	name string
	designation string
}

func DemoMethodsAndFunctions(){
	fmt.Println("Demo to show how to use functions & methods with pass by value and pass by reference:")
	callFunctionByValueWithStrings()
	callFunctionByPoitnerWithStrings()
	callMethodByValueWithStrings()
	callMethodByReferenceWithStrings()
}

func callFunctionByValueWithStrings(){
	fmt.Println("******   PART -1 ( Functions, call by value , operations on strings   **********************")

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

	fmt.Println("******    Part -2 Functions, call by reference , operations on strings   ********************** ")

	empInfo := empDetails{"Ravi", "Software Engineer"}
	fmt.Println("Before call by reference", empInfo)
	changeEmpDesignationByPointer(&empInfo)
	fmt.Println("After call by reference- Did designation changed to Architect, Yes", empInfo)
}

func changeEmpDesignationByPointer(details *empDetails){

	fmt.Println("Emp designation is changing to Architect: ")
	details.designation = "Architect"
}


type custDetails struct{
	name string
	designation string
}
func callMethodByValueWithStrings(){
	fmt.Println("******   PART -3  ( Methods , call by value , operations on strings   **********************")

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
	fmt.Println("******   PART -4 ( Methods , call by reference, operations on strings   **********************")

	var customerInterface = &custDetails{"Hari", "Software Engineer"}
	fmt.Println("Before call by value ", customerInterface)
	customerInterface.changeEmpDesignationByMethodByRef("Architect")
	fmt.Println("After call by value - Did designation changed to Architect, Yes", customerInterface)
}
func (e *custDetails) changeEmpDesignationByMethodByRef(desig  string){

	fmt.Println("Emp designation is changing to Architect: ")
	e.designation = desig
}
