package main

import (
	"unemp-tool/front"
	"unemp-tool/unemp"
)

func main() {
	var param unemp.Parameter

	var console front.Console
	console.Println("失業手当を計算します")
	param.TotalWage = console.GetInt("過去6か月の賃金総額(円)")
	param.Age = console.GetInt("年齢")

	param.InsuredPeriod = console.GetInt(
		`被保険者期間(年)
		※ 1年未満の場合は0`)

	param.Reason = console.GetInt(
		`退職理由
		1: 勤め先の倒産や解雇など、会社都合の退職
		2: 有期雇用で更新を希望したがかなわず、退職
		3: 病気・怪我・妊娠・介護など致し方ない理由での退職
		4: 転職など自己都合での退職・懲戒解雇
		5: 定年退職(65歳未満)
		6: 65歳以上での退職`)
	param.NumOfDayPassed = console.GetInt("失業手当の期間経過日数")

	var facade unemp.Facade
	facade.Exec(param)
}
