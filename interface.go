package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName string
	amount      int
}

type TimeAndMaterial struct {
	projectName string
	hrRate      int
	noOfhrs     int
}

func (obj *FixedBilling) calculate() int {
	return obj.amount
}

func (obj *FixedBilling) source() string {
	return obj.projectName
}

func (obj *TimeAndMaterial) calculate() int {
	return obj.hrRate * obj.noOfhrs
}

func (obj *TimeAndMaterial) source() string {
	return obj.projectName
}

// using Polymorphism for calculation based
// on the array of variables of interface type
func CalculateNetIncome(ic []Income) {
	var netincome int
	for _, income := range ic {
		fmt.Printf("Income for Project %s is %d \n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d \n ", netincome)
	return
}
