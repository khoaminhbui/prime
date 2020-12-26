package usecases

import (
	"testing"

	"github.com/khoaminhbui/prime/prime-service/core/entities"
)

func TestFindLargestPrimeUnder(t *testing.T) {
	testInvalid := FindLargestPrimeUnder("a")
	if testInvalid.ErrorCode != entities.ERROR_CODE_INVALID_FORMAT {
		t.Errorf("FindLargestPrimeUnder failed, expected error code %v got %v", entities.ERROR_CODE_INVALID_FORMAT, testInvalid.ErrorCode)
	} else {
		t.Logf("FindLargestPrimeUnder passed, expected error code %v got %v", entities.ERROR_CODE_INVALID_FORMAT, testInvalid.ErrorCode)
	}

	testValid := FindLargestPrimeUnder("100")
	if testValid.ErrorCode != entities.ERROR_CODE_OK {
		t.Errorf("FindLargestPrimeUnder failed, expected error code %v got %v", entities.ERROR_CODE_OK, testValid.ErrorCode)
	} else {
		t.Logf("FindLargestPrimeUnder passed, expected error code %v got %v", entities.ERROR_CODE_OK, testValid.ErrorCode)
	}
}
