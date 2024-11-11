package unemp 

import (
    "unemp-tool/front"
)

type (
	Facade struct {}
)

func (this Facade) Exec() {
        var console front.Console
	console.Println("失業手当を計算します")

	totalWage := console.GetInt("過去6か月の賃金総額(円)")
	age := console.GetInt("年齢")

	insuredPeriod := console.GetInt(
			`被保険者期間(年)
                        ※ 1年未満の場合は0`)

	reason := console.GetInt(
		`退職理由
                1: 勤め先の倒産や解雇など、会社都合の退職
                2: 有期雇用で更新を希望したがかなわず、退職
                3: 病気・怪我・妊娠・介護など致し方ない理由での退職
                4: 転職など自己都合での退職・懲戒解雇
                5: 定年退職(65歳未満)
                6: 65歳以上での退職`)

        numOfDayPassed := console.GetInt("失業手当の期間経過日数")

	var days Days
	var unempAllowance UnempAllowance
        var reempAllowance ReempAllowance

        numOfDaysTotal          := days.Calc(age, insuredPeriod, reason)
        dailyUnempAllowance     := unempAllowance.CalcDailyAllowance(age, totalWage)
        monthlyUnempAllowance   := unempAllowance.CalcMonthlyAllowance(age, totalWage)
        unempAllowanceTotal     := dailyUnempAllowance * numOfDaysTotal
        reempAllowanceTotal     := reempAllowance.CalcReempAllowance(numOfDaysTotal, numOfDayPassed, dailyUnempAllowance)

	console.Println("計算結果")
	console.Println("---------------------------------------------------")
	console.Println("失業手当 給付日数 :", numOfDaysTotal)
	console.Println("失業手当 合計(円) :", unempAllowanceTotal)
	console.Println("失業手当 日当(円) :", dailyUnempAllowance)
	console.Println("失業手当 月額(円) :", monthlyUnempAllowance)
	console.Println("再就職手当(円)    :", reempAllowanceTotal)
	console.Println("---------------------------------------------------")
}
