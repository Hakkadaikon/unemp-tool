package main

import (
	"testing"
)

func TestAllowance_CalcDailyAllowance(t *testing.T) {
	tests := []struct {
		name           string
		age            int
		totalWage      int
		expectedResult int
	}{
		{"Age 29, Wage 1271520", 29, 1271520, 7064},
		{"Age 29, Wage 1271880", 29, 1271880, 7065}, // 上限超え 7066 -> 7065
		{"Age 30, Wage 1411920", 30, 1411920, 7844},
		{"Age 30, Wage 1412280", 30, 1412280, 7845}, // 上限超え 7846 -> 7845
		{"Age 44, Wage 1411920", 44, 1411920, 7844},
		{"Age 44, Wage 1412280", 44, 1412280, 7845}, // 上限超え 7846 -> 7845
		{"Age 45, Wage 1554120", 45, 1554120, 8634},
		{"Age 45, Wage 1554480", 45, 1554480, 8635}, // 上限超え 8636 -> 8635
		{"Age 59, Wage 1554120", 59, 1554120, 8634},
		{"Age 59, Wage 1554480", 59, 1554480, 8635}, // 上限超え 8636 -> 8635
		{"Age 60, Wage 1335420", 60, 1335420, 7419},
		{"Age 60, Wage 1335780", 60, 1335780, 7420}, // 上限超え 7421 -> 7420
		{"Age 64, Wage 1335420", 64, 1335420, 7419},
		{"Age 64, Wage 1335780", 64, 1335780, 7420}, // 上限超え 7421 -> 7420
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Allowance{}
			result := a.CalcDailyAllowance(tt.age, tt.totalWage)
			if result != tt.expectedResult {
				t.Errorf("CalcDailyAllowance(%d, %d) = %d; want %d", tt.age, tt.totalWage, result, tt.expectedResult)
			}
		})
	}
}
