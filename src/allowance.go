package main

type (
	Allowance struct {
		age       int
		totalWage int
	}
)

func (this Allowance) judgeAgeKind() int {
	age := this.age
	if age <= 29 {
		return 0
	} else if age >= 30 && age < 45 {
		return 1
	} else if age >= 45 && age < 60 {
		return 2
	} else if age >= 60 && age < 65 {
		return 3
	}

	return 4
}

func (this Allowance) selectLimitDailyAllowance() int {
	var limitDailyAllowanceTable = [5]int{
		7065, 7845, 8635, 7420, 0}

	ageKind := this.judgeAgeKind()
	return limitDailyAllowanceTable[ageKind]
}

func (this Allowance) calcDailyAllowanceFromTotalWage() int {
	return (this.totalWage / 180)
}

func (this Allowance) CalcDailyAllowance(age int, totalWage int) int {
	this.age = age
	this.totalWage = totalWage

	calcAllowance := this.calcDailyAllowanceFromTotalWage()
	limitDailyAllowance := this.selectLimitDailyAllowance()

	if calcAllowance > limitDailyAllowance {
		return limitDailyAllowance
	}

	return calcAllowance
}

func (this Allowance) CalcMonthlyAllowance(age int, totalWage int) int {
	return (this.CalcDailyAllowance(age, totalWage) * 28)
}
