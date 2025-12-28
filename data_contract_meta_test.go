package inference

import "testing"

func TestDataContractHash(t *testing.T) {
	if err := ValidateDataContract(); err != nil {
		t.Fatal(err)
	}
}
