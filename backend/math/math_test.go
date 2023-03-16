package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 2)
	if result != 5 {
		t.Errorf("Expected: %d Recevied: %d", 5, result)
	}
}
