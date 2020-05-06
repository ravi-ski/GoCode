package interfaces

import "fmt"

type totalMoney int

type incomeSources interface{
	salaryIncome() totalMoney
	otherIncome() totalMoney
	displayIncome()
}

func (money totalMoney) salaryIncome() totalMoney{

	fmt.Println("Salary Income ", money)
	return money + 1000
}

func (money totalMoney) otherIncome() totalMoney{

	fmt.Println("Other Income ", money)
	return money + 2000
}

func (money totalMoney) displayIncome(){

	fmt.Println("Display Income ", money)
}
func DemoInterface(){

	var money incomeSources
	money = totalMoney(0)
	money = money.salaryIncome()
	money = money.otherIncome()
	money.displayIncome()

}
