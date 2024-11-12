package unemp

type (
	UnempAllowance struct {
		age       int
		totalWage int
	}
)

func (this UnempAllowance) judgeAgeKind() int {
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

func (this UnempAllowance) selectLimitDailyAllowance() int {
	var limitDailyAllowanceTable = [5]int{
		7065, 7845, 8635, 7420, 0}

	ageKind := this.judgeAgeKind()
	return limitDailyAllowanceTable[ageKind]
}

func (this UnempAllowance) calcDailyAllowanceFromTotalWage() int {
	return (this.totalWage / 180)
}

func (this UnempAllowance) validate(age int, totalWage int) bool {
	if age < 15 {
		return false
	}

	if totalWage <= 0 {
		return false
	}

	return true
}

func (this UnempAllowance) CalcDailyAllowance(age int, totalWage int) int {
	if !this.validate(age, totalWage) {
		return 0
	}

	this.age = age
	this.totalWage = totalWage

	calcAllowance := this.calcDailyAllowanceFromTotalWage()
	limitDailyAllowance := this.selectLimitDailyAllowance()

	if calcAllowance > limitDailyAllowance {
		return limitDailyAllowance
	}

	return calcAllowance
}

func (this UnempAllowance) CalcMonthlyAllowance(age int, totalWage int) int {
	return (this.CalcDailyAllowance(age, totalWage) * 28)
}
