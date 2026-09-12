package sdkutil

import "testing"

func TestDataContractHash(t *testing.T) {
	if h, err := ValidateDataContract(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("hash: %s", h)
	}
}
