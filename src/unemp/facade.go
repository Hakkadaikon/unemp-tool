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

	var days Days
	var allowance Allowance

	console.Println("計算結果")
	console.Println("---------------------------------------------------")
	console.Println("給付日数:", days.Calc(age, insuredPeriod, reason))
	console.Println("失業手当 日当(円):", allowance.CalcDailyAllowance(age, totalWage))
	console.Println("失業手当 月額(円):", allowance.CalcMonthlyAllowance(age, totalWage))
	console.Println("---------------------------------------------------")
}
