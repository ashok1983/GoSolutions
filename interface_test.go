package main

import "testing"

func TestInterface(t *testing.T) {
	project1 := &FixedBilling{projectName: "proj-1", amount: 1000}
	project2 := &FixedBilling{projectName: "proj-2", amount: 5000}
	project3 := &TimeAndMaterial{projectName: "proj-3-tm", hrRate: 1000, noOfhrs: 8}

	income := []Income{project1, project2, project3}
	CalculateNetIncome(income)
	return
}
