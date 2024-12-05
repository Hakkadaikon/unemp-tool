package unemp

import (
	"unemp-tool/exchange"
	"unemp-tool/front"
)

type (
	Facade struct{}
)

func (this *Facade) Exec(param Parameter) {
	totalWage := param.TotalWage
	age := param.Age
	insuredPeriod := param.InsuredPeriod
	reason := param.Reason
	numOfDayPassed := param.NumOfDayPassed

	var days Days
	var unempAllowance UnempAllowance
	var reempAllowance ReempAllowance

	numOfDaysTotal := days.Calc(age, insuredPeriod, reason)
	dailyUnempAllowance := unempAllowance.CalcDailyAllowance(age, totalWage)
	monthlyUnempAllowance := unempAllowance.CalcMonthlyAllowance(age, totalWage)
	unempAllowanceTotal := dailyUnempAllowance * numOfDaysTotal
	reempAllowanceTotal := reempAllowance.CalcReempAllowance(numOfDaysTotal, numOfDayPassed, dailyUnempAllowance)

	var bitcoin exchange.Bitcoin
	unempAllowanceTotalSatoshi, err := bitcoin.JpyToSatoshi(float64(unempAllowanceTotal))
	dailyUnempAllowanceSatoshi, err := bitcoin.JpyToSatoshi(float64(dailyUnempAllowance))
	monthlyUnempAllowanceSatoshi, err := bitcoin.JpyToSatoshi(float64(monthlyUnempAllowance))
	reempAllowanceTotalSatoshi, err := bitcoin.JpyToSatoshi(float64(reempAllowanceTotal))

	var console front.Console
	if err == nil {
		console.Println("計算結果")
		console.Println("---------------------------------------------------")
		console.Printf("失業手当 給付日数(合計) : %3d 日\n", numOfDaysTotal)
		console.Printf("失業手当 給付日数(残)   : %3d 日\n", numOfDaysTotal-numOfDayPassed)
		console.Printf("失業手当 合計           : %7d 円 / %10.0f satoshi\n", unempAllowanceTotal, unempAllowanceTotalSatoshi)
		console.Printf("失業手当 日当           : %7d 円 / %10.0f satoshi\n", dailyUnempAllowance, dailyUnempAllowanceSatoshi)
		console.Printf("失業手当 月額           : %7d 円 / %10.0f satoshi\n", monthlyUnempAllowance, monthlyUnempAllowanceSatoshi)
		console.Printf("再就職手当              : %7d 円 / %10.0f satoshi\n", reempAllowanceTotal, reempAllowanceTotalSatoshi)
		console.Println("---------------------------------------------------")
		console.Println("Buy bitcoin stashiで積み立てよう")
		console.Println("https://github.com/erechorse/stashi")
		return
	}

	console.Printf("err: %s\n", err)
	console.Println("計算結果")
	console.Println("---------------------------------------------------")
	console.Printf("失業手当 給付日数(合計) : %3d 日\n", numOfDaysTotal)
	console.Printf("失業手当 給付日数(残)   : %3d 日\n", numOfDaysTotal-numOfDayPassed)
	console.Printf("失業手当 合計           : %7d 円\n", unempAllowanceTotal)
	console.Printf("失業手当 日当           : %7d 円\n", dailyUnempAllowance)
	console.Printf("失業手当 月額           : %7d 円\n", monthlyUnempAllowance)
	console.Printf("再就職手当              : %7d 円\n", reempAllowanceTotal)
	console.Println("---------------------------------------------------")
}
