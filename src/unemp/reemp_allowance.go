package unemp 

type (
	ReempAllowance struct {
		numOfDayTotal  int
		numOfDayPassed int
		dailyAllowance int
	}
)

func (this ReempAllowance) judgeRate() float64 {
    numOfDayRemaining := this.calcNumOfDayRemaining()
    compDays := float64(numOfDayRemaining / 3)

    // 給付日数が総日数の3/1以上ある場合の給付率:0.7
    // そうでない場合の給付率:0.6
    if float64(numOfDayRemaining) > compDays {
        return 0.7
    }

    return 0.6
}

func (this ReempAllowance) calcNumOfDayRemaining() int {
    return this.numOfDayTotal - this.numOfDayPassed
}

func (this ReempAllowance) calcSimpleReempAllowance() float64 {
    // 給付率を意識しない単純計算した再就職手当を返却
    numOfDayRemaining := this.calcNumOfDayRemaining()
    return float64(numOfDayRemaining) * float64(this.dailyAllowance)
}

func (this ReempAllowance) CalcReempAllowance(
	numOfDayTotal  int,
	numOfDayPassed int,
	dailyAllowance int) int {
    this.numOfDayTotal  = numOfDayTotal 
    this.numOfDayPassed = numOfDayPassed 
    this.dailyAllowance = dailyAllowance

    rate := this.judgeRate()
    allowance := this.calcSimpleReempAllowance()

    return int(allowance * rate)
}
