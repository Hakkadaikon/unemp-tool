package unemp

type (
	Days struct {
		age           int
		insuredPeriod int
		reason        int
	}
)

func (this Days) judgeInsuredPeriodKind() int {
	insuredPeriod := this.insuredPeriod
	if insuredPeriod <= 0 {
		return 0
	} else if insuredPeriod >= 1 && insuredPeriod < 5 {
		return 1
	} else if insuredPeriod >= 5 && insuredPeriod < 10 {
		return 2
	} else if insuredPeriod >= 10 && insuredPeriod < 20 {
		return 3
	}

	return 4
}

func (this Days) judgeAgeKind() int {
	age := this.age
	if age <= 29 {
		return 0
	} else if age >= 30 && age < 35 {
		return 1
	} else if age >= 35 && age < 45 {
		return 2
	} else if age >= 45 && age < 60 {
		return 3
	}

	return 4
}

func (this Days) selectDaysTable() [5][5]int {
	reason := this.reason
	// 会社都合
	var DaysTable1 = [5][5]int{
		{90, 90, 90, 90, 90},      // 1年未満
		{90, 120, 150, 180, 150},  // 1年以上5年未満
		{120, 180, 180, 240, 180}, // 5年以上10年未満
		{180, 210, 240, 270, 210}, // 10年以上20年未満
		{0, 240, 270, 330, 240},   // 20年以上
	}

	// 自己都合
	var DaysTable2 = [5][5]int{
		{90, 90, 90, 90, 90},      // 1年未満
		{90, 90, 90, 90, 90},      // 1年以上5年未満
		{90, 90, 90, 90, 90},      // 5年以上10年未満
		{120, 120, 120, 120, 120}, // 10年以上20年未満
		{150, 150, 150, 150, 150}, // 20年以上
	}

	// 病気・怪我
	var DaysTable3 = [5][5]int{
		{150, 150, 150, 150, 150}, // 1年未満
		{300, 300, 300, 360, 360}, // 1年以上5年未満
		{300, 300, 300, 360, 360}, // 5年以上10年未満
		{300, 300, 300, 360, 360}, // 10年以上20年未満
		{300, 300, 300, 360, 360}, // 20年以上
	}

	switch reason {
	// 1.勤め先の倒産や解雇など、会社都合の退職
	// 2.有期雇用で更新を希望したがかなわず、退職
	case 1, 2:
		return DaysTable1
		// 病気・怪我・妊娠・介護など致し方ない理由での退職
	case 3:
		return DaysTable3
	}

	// 4: 転職など自己都合での退職・懲戒解雇
	// 5: 定年退職(65歳未満)
	// 6: 65歳以上での退職
	return DaysTable2
}

func (this Days) Calc(age int, insuredPeriod int, reason int) int {
	// 雇用保険の被保険者期間が成立するのは、15歳から
	if (age - insuredPeriod) < 15 {
		return 0
	}

	this.age = age
	this.insuredPeriod = insuredPeriod
	this.reason = reason

	ageKind := this.judgeAgeKind()
	insuredPeriodKind := this.judgeInsuredPeriodKind()
	DaysTable := this.selectDaysTable()

	return DaysTable[insuredPeriodKind][ageKind]
}
