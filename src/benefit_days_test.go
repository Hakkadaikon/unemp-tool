package main

import (
	"testing"
)

func TestBenefitDays_Calc(t *testing.T) {
	tests := []struct {
		name           string
		age            int
		insuredPeriod  int
		reason         int
		expectedResult int
	}{
		// 会社都合の退職の場合
		{"会社都合退職, 29歳, 被保険者期間: 1年未満", 29, 0, 1, 90},
		{"会社都合退職, 30歳, 被保険者期間: 1年未満", 30, 0, 1, 90},
		{"会社都合退職, 34歳, 被保険者期間: 1年未満", 34, 0, 1, 90},
		{"会社都合退職, 35歳, 被保険者期間: 1年未満", 35, 0, 1, 90},
		{"会社都合退職, 44歳, 被保険者期間: 1年未満", 44, 0, 1, 90},
		{"会社都合退職, 45歳, 被保険者期間: 1年未満", 45, 0, 1, 90},
		{"会社都合退職, 59歳, 被保険者期間: 1年未満", 59, 0, 1, 90},
		{"会社都合退職, 60歳, 被保険者期間: 1年未満", 60, 0, 1, 90},

		{"会社都合退職, 29歳, 被保険者期間: 1年", 29, 1, 1, 90},
		{"会社都合退職, 30歳, 被保険者期間: 1年", 30, 1, 1, 120},
		{"会社都合退職, 34歳, 被保険者期間: 1年", 34, 1, 1, 120},
		{"会社都合退職, 35歳, 被保険者期間: 1年", 35, 1, 1, 150},
		{"会社都合退職, 44歳, 被保険者期間: 1年", 44, 1, 1, 150},
		{"会社都合退職, 45歳, 被保険者期間: 1年", 45, 1, 1, 180},
		{"会社都合退職, 59歳, 被保険者期間: 1年", 59, 1, 1, 180},
		{"会社都合退職, 60歳, 被保険者期間: 1年", 60, 1, 1, 150},

		{"会社都合退職, 29歳, 被保険者期間: 4年", 29, 4, 1, 90},
		{"会社都合退職, 30歳, 被保険者期間: 4年", 30, 4, 1, 120},
		{"会社都合退職, 34歳, 被保険者期間: 4年", 34, 4, 1, 120},
		{"会社都合退職, 35歳, 被保険者期間: 4年", 35, 4, 1, 150},
		{"会社都合退職, 44歳, 被保険者期間: 4年", 44, 4, 1, 150},
		{"会社都合退職, 45歳, 被保険者期間: 4年", 45, 4, 1, 180},
		{"会社都合退職, 59歳, 被保険者期間: 4年", 59, 4, 1, 180},
		{"会社都合退職, 60歳, 被保険者期間: 4年", 60, 4, 1, 150},

		{"会社都合退職, 29歳, 被保険者期間: 5年", 29, 5, 1, 120},
		{"会社都合退職, 30歳, 被保険者期間: 5年", 30, 5, 1, 180},
		{"会社都合退職, 34歳, 被保険者期間: 5年", 34, 5, 1, 180},
		{"会社都合退職, 35歳, 被保険者期間: 5年", 35, 5, 1, 180},
		{"会社都合退職, 44歳, 被保険者期間: 5年", 44, 5, 1, 180},
		{"会社都合退職, 45歳, 被保険者期間: 5年", 45, 5, 1, 240},
		{"会社都合退職, 59歳, 被保険者期間: 5年", 59, 5, 1, 240},
		{"会社都合退職, 60歳, 被保険者期間: 5年", 60, 5, 1, 180},

		{"会社都合退職, 29歳, 被保険者期間: 9年", 29, 9, 1, 120},
		{"会社都合退職, 30歳, 被保険者期間: 9年", 30, 9, 1, 180},
		{"会社都合退職, 34歳, 被保険者期間: 9年", 34, 9, 1, 180},
		{"会社都合退職, 35歳, 被保険者期間: 9年", 35, 9, 1, 180},
		{"会社都合退職, 44歳, 被保険者期間: 9年", 44, 9, 1, 180},
		{"会社都合退職, 45歳, 被保険者期間: 9年", 45, 9, 1, 240},
		{"会社都合退職, 59歳, 被保険者期間: 9年", 59, 9, 1, 240},
		{"会社都合退職, 60歳, 被保険者期間: 9年", 60, 9, 1, 180},

		{"会社都合退職, 29歳, 被保険者期間: 10年", 29, 10, 1, 180},
		{"会社都合退職, 30歳, 被保険者期間: 10年", 30, 10, 1, 210},
		{"会社都合退職, 34歳, 被保険者期間: 10年", 34, 10, 1, 210},
		{"会社都合退職, 35歳, 被保険者期間: 10年", 35, 10, 1, 240},
		{"会社都合退職, 44歳, 被保険者期間: 10年", 44, 10, 1, 240},
		{"会社都合退職, 45歳, 被保険者期間: 10年", 45, 10, 1, 270},
		{"会社都合退職, 59歳, 被保険者期間: 10年", 59, 10, 1, 270},
		{"会社都合退職, 60歳, 被保険者期間: 10年", 60, 10, 1, 210},

		{"会社都合退職, 29歳, 被保険者期間: 19年", 29, 19, 1, 0},
		{"会社都合退職, 30歳, 被保険者期間: 19年", 30, 19, 1, 0},
		{"会社都合退職, 34歳, 被保険者期間: 19年", 34, 19, 1, 210},
		{"会社都合退職, 35歳, 被保険者期間: 19年", 35, 19, 1, 240},
		{"会社都合退職, 44歳, 被保険者期間: 19年", 44, 19, 1, 240},
		{"会社都合退職, 45歳, 被保険者期間: 19年", 45, 19, 1, 270},
		{"会社都合退職, 59歳, 被保険者期間: 19年", 59, 19, 1, 270},
		{"会社都合退職, 60歳, 被保険者期間: 19年", 60, 19, 1, 210},

		{"会社都合退職, 29歳, 被保険者期間: 20年", 29, 20, 1, 0},
		{"会社都合退職, 30歳, 被保険者期間: 20年", 30, 20, 1, 0},
		{"会社都合退職, 34歳, 被保険者期間: 20年", 34, 20, 1, 0},
		{"会社都合退職, 35歳, 被保険者期間: 20年", 35, 20, 1, 270},
		{"会社都合退職, 44歳, 被保険者期間: 20年", 44, 20, 1, 270},
		{"会社都合退職, 45歳, 被保険者期間: 20年", 45, 20, 1, 330},
		{"会社都合退職, 59歳, 被保険者期間: 20年", 59, 20, 1, 330},
		{"会社都合退職, 60歳, 被保険者期間: 20年", 60, 20, 1, 240},

		{"自己都合退職, 40歳, 被保険者期間: 1年未満", 40, 0, 4, 90},
		{"自己都合退職, 40歳, 被保険者期間: 1年", 40, 1, 4, 90},
		{"自己都合退職, 40歳, 被保険者期間: 4年", 40, 4, 4, 90},
		{"自己都合退職, 40歳, 被保険者期間: 5年", 40, 5, 4, 90},
		{"自己都合退職, 40歳, 被保険者期間: 9年", 40, 9, 4, 90},
		{"自己都合退職, 40歳, 被保険者期間: 10年", 40, 10, 4, 120},
		{"自己都合退職, 40歳, 被保険者期間: 19年", 40, 19, 4, 120},
		{"自己都合退職, 40歳, 被保険者期間: 20年", 40, 20, 4, 150},

		// // 病気・怪我での退職の場合
		// {"病気・怪我での退職, 25歳, 被保険者期間: 3年", 25, 3, 3, 300},
		// {"病気・怪我での退職, 35歳, 被保険者期間: 7年", 35, 7, 3, 300},
		// {"病気・怪我での退職, 45歳, 被保険者期間: 12年", 45, 12, 3, 360},
		// {"病気・怪我での退職, 50歳, 被保険者期間: 20年", 50, 20, 3, 360},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			benefitDays := BenefitDays{}
			result := benefitDays.Calc(tt.age, tt.insuredPeriod, tt.reason)
			if result != tt.expectedResult {
				t.Errorf("Calc(%d, %d, %d) = %d; want %d",
					tt.age,
					tt.insuredPeriod,
					tt.reason,
					result,
					tt.expectedResult)
			}
		})
	}
}
