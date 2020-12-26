package entities

import (
	"testing"
)

func TestTrialDivision(t *testing.T) {
	test1 := TrialDivision(1)
	if test1 != 1 {
		t.Errorf("TrialDivision failed, expected %v got %v", 1, test1)
	} else {
		t.Logf("TrialDivision success, expected %v got %v", 1, test1)
	}

	test10 := TrialDivision(10)
	if test10 != 7 {
		t.Errorf("TrialDivision failed, expected %v got %v", 7, test10)
	} else {
		t.Logf("TrialDivision success, expected %v got %v", 7, test10)
	}

	test1000000 := TrialDivision(1000000)
	if test1000000 != 999983 {
		t.Errorf("TrialDivision failed, expected %v got %v", 999983, test1000000)
	} else {
		t.Logf("TrialDivision success, expected %v got %v", 999983, test1000000)
	}
}
