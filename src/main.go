package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
)

func getStdIn(desc string) int {
	var intVal int
	fmt.Println(desc)
	fmt.Print(":")

	fmt.Scan(&intVal)

	return intVal
}

func main() {
	totalWage := getStdIn("過去6か月の賃金総額(円)")
	age := getStdIn("年齢")

	insuredPeriod := getStdIn(heredoc.Doc(
		`被保険者期間(年)
    ※ 1年未満の場合は0`))

	reason := getStdIn(heredoc.Doc(
		`退職理由
    1: 勤め先の倒産や解雇など、会社都合の退職
    2: 有期雇用で更新を希望したがかなわず、退職
    3: 病気・怪我・妊娠・介護など致し方ない理由での退職
    4: 転職など自己都合での退職・懲戒解雇
    5: 定年退職(65歳未満)
    6: 65歳以上での退職`))

	var benefitDays BenefitDays
	var allowance Allowance

	fmt.Println("計算結果")
	fmt.Println("---------------------------------------------------")
	fmt.Println("給付日数:", benefitDays.Calc(age, insuredPeriod, reason))
	fmt.Println("失業手当 日当(円):", allowance.CalcDailyAllowance(age, totalWage))
	fmt.Println("失業手当 月額(円):", allowance.CalcMonthlyAllowance(age, totalWage))
	fmt.Println("---------------------------------------------------")
}
